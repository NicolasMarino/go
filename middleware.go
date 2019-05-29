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
	"strconv"
	"strings"

	"github.com/nicolasmarino/api/models"
)

func main() {
	datos := getOrdersPY()

	for i := 0; i < len(datos.Datos); i++ {

		if datos.Datos[i].Integration == "RestoSoft" {
			logs := postRestoSoft(datos.Datos[i])
			saveLog(logs)
		} else if datos.Datos[i].Integration == "XResto" {
			logs := postXResto(datos.Datos[i])
			saveLog(logs)
		}

	}

}

func saveLog(response string) {
	f, err := os.OpenFile("ordenesEnviadas.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write([]byte(string(response)))
	if err != nil {
		log.Fatal(err)
	}

	f.Close()
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

func postRestoSoft(order models.Data) string {
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
	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//fmt.Println(string(body))
	responseString := "En Pedidos Ya: " + strconv.FormatInt(order.ID, 10) + ", en RestoSoft: " + string(body) + ", estado:" + string(resp.Status) + "." + "\n"

	return responseString
}

func postXResto(order models.Data) string {
	//Para ver cada orden
	//fmt.Println(order)
	var orderXResto models.Order
	Data := order

	orderXResto.Customer.Name = Data.Customer.GetFullName()
	orderXResto.Customer.Coordinates = Data.Address.Coordinates
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

	dataXML, _ := xml.MarshalIndent(orderXResto, "", "  ")

	url := "http://vmrdr.mocklab.io/xresto/v2/orders"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(dataXML))
	req.Header.Set("Authorization", "xresto-test-developer")
	req.Header.Set("Content-Type", "text/xml")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//fmt.Println(string(body))
	responseString := "En Pedidos Ya: " + strconv.FormatInt(order.ID, 10) + ", en XResto: " + string(body) + ", estado:" + string(resp.Status) + "." + "\n"

	return responseString
}
