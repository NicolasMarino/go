package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/nicolasmarino/api/models"
)

func main() {
	datos := getOrdersPY()

	for i := 0; i < len(datos.Datos); i++ {
		if datos.Datos[i].Integration == "RestoSoft" {
			postRestoSoft(datos.Datos[i])
			fmt.Println("---------")
		} else if datos.Datos[i].Integration == "XResto" {
			postXResto(datos.Datos[i])
		}
	}

}

func getOrdersPY() models.Orders {
	response, err := http.Get("http://vmrdr.mocklab.io/pedidosya/v1/orders")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject models.Orders
	json.Unmarshal(responseData, &responseObject)

	return responseObject
}

func postRestoSoft(order models.Data) {
	//Para ver cada orden
	//fmt.Println(order) Para ver cada orden.
	var restoSoftData models.RestoSoft

	Data := order
	restoSoftData.Date = strings.Split(Data.RegisteredDate.String(), " ")[0]
	restoSoftData.Notes = Data.Notes
	restoSoftData.Total = Data.Subtotal
	var datosRS []models.ItemsRestoSoft

	for x := 0; x < len(Data.Items); x++ {
		var nuevosRS models.ItemsRestoSoft
		nuevosRS.Name = Data.Items[x].Name
		nuevosRS.Price = Data.Items[x].Price
		nuevosRS.Quantity = Data.Items[x].Quantity
		datosRS = append(datosRS, nuevosRS)
		for xi := 0; xi < len(Data.Items[x].Options); xi++ {
			nuevosRS.Name = Data.Items[x].Options[xi].Name
			nuevosRS.Quantity = Data.Items[x].Options[xi].Quantity
			nuevosRS.Price = 0
			datosRS = append(datosRS, nuevosRS)
		}
	}

	restoSoftData.Items = datosRS
	restoSoftData.Customer.Name = Data.Customer.GetFullName()
	restoSoftData.Customer.Location.Longitude = strings.Split(Data.Address.Coordinates, ",")[0]
	restoSoftData.Customer.Location.Latitude = strings.Split(Data.Address.Coordinates, ",")[1]
	restoSoftData.Business.Name = Data.Restaurant.Name

	restoSoftDataJSON, _ := json.MarshalIndent(restoSoftData, "", "    ")

	url := "http://vmrdr.mocklab.io/restosoft/v1/orders"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(restoSoftDataJSON))
	req.Header.Set("Authorization", "restosoft-test-developer")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}

func postXResto(order models.Data) {
	//Para ver cada orden
	//fmt.Println(order)
	var orderXResto models.Order
	Data := order

	orderXResto.Customer.Name = Data.Customer.GetFullName()

	orderXResto.Customer.Coordinates = Data.Address.Coordinates

	//fmt.Print(Data.Restaurant.Name)
	orderXResto.Status = Data.State

	orderXResto.Business.Name = Data.Restaurant.Name

	orderXResto.Date.Year = strings.Split(strings.Split(Data.RegisteredDate.String(), " ")[0], "-")[0]
	orderXResto.Date.Month = strings.Split(strings.Split(Data.RegisteredDate.String(), " ")[0], "-")[1]
	orderXResto.Date.Day = strings.Split(strings.Split(Data.RegisteredDate.String(), " ")[0], "-")[2]

	orderXResto.Notes = Data.Notes
	orderXResto.Total = Data.Subtotal

	var datosXResto []models.ItemsRestoSoft
	for x := 0; x < len(Data.Items); x++ {
		var nuevosRS models.ItemsRestoSoft
		nuevosRS.Name = Data.Items[x].Name
		nuevosRS.Price = Data.Items[x].Price
		nuevosRS.Quantity = Data.Items[x].Quantity
		datosXResto = append(datosXResto, nuevosRS)
		for xi := 0; xi < len(Data.Items[x].Options); xi++ {
			nuevosRS.Name = Data.Items[x].Options[xi].Name
			nuevosRS.Quantity = Data.Items[x].Options[xi].Quantity
			nuevosRS.Price = 0
			datosXResto = append(datosXResto, nuevosRS)
		}
	}

	orderXResto.Items = datosXResto

	dataXML, _ := xml.MarshalIndent(orderXResto, "", "    ")
	fmt.Print(string(dataXML))

	url := "http://vmrdr.mocklab.io/xresto/v2/orders"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(dataXML))
	req.Header.Set("Authorization", "xresto-test-developer")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}

func structToLowerFirstMap(in interface{}) map[string]interface{} {
	v := reflect.ValueOf(in)
	vType := v.Type()

	result := make(map[string]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		name := vType.Field(i).Name
		result[lowerFirst(name)] = v.Field(i).Interface()
	}

	return result

}

func lowerFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

/*
func structToMap(i interface{}) (values url.Values) {
	values = url.Values{}
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		// You ca use tags here...
		// tag := typ.Field(i).Tag.Get("tagname")
		// Convert each type into a string for the url.Values string map
		var v string
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
		}
		values.Set(typ.Field(i).Name, v)
	}
	return
}*/
