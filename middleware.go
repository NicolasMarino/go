package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

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
	restoSoftData.Total = Data.Total
	var datosRS []models.ItemsRestoSoft

	for x := 0; x < len(Data.Items); x++ {
		var nuevosRS models.ItemsRestoSoft
		nuevosRS.Name = Data.Items[x].Name
		nuevosRS.Price = Data.Items[x].Price
		nuevosRS.Quantity = Data.Items[x].Quantity
		datosRS = append(datosRS, nuevosRS)
	}
	restoSoftData.Items = datosRS

	restoSoftData.Customer.Name = Data.Customer.Name
	restoSoftData.Customer.Location.Latitude = strings.Split(Data.Address.Coordinates, ",")[0]
	restoSoftData.Customer.Location.Longitude = strings.Split(Data.Address.Coordinates, ",")[1]
	restoSoftData.Business.Name = Data.Restaurant.Name

	restoSoftDataJSON, _ := json.MarshalIndent(restoSoftData, "", "    ")

	fmt.Print(string(restoSoftDataJSON))

}

func postXResto(order models.Data) {
	//Para ver cada orden
	//fmt.Println(order)

	//fmt.Println(responseObject.Datos)
	/*var datos models.RestoSoft

	for i := 0; i < len(responseObject.Datos); i++ {

		if responseObject.Datos[i].Integration == "RestoSoft" {
			//fmt.Println("es RestoSoft")
		} else if responseObject.Datos[i].Integration == "XResto" {
			//fmt.Println("es XResto")
		}
		for j := 0; j < len(responseObject.Datos[i].Items); j++ {
			responseObject.Datos[i].Items[j].ID = 0
			responseObject.Datos[i].Items[j].Options = []models.Option{}
		}

		Data := responseObject.Datos[i]
		datos.Date = Data.RegisteredDate
		datos.Notes = Data.Notes
		datos.Total = Data.Total
		var datosRS []models.ItemsRs

		for x := 0; x < len(Data.Items); x++ {
			var nuevosRS models.ItemsRs
			nuevosRS.Name = Data.Items[x].Name
			nuevosRS.Price = Data.Items[x].Price
			nuevosRS.Quantity = Data.Items[x].Quantity
			datosRS = append(datosRS, nuevosRS)
		}
		//datos.Items = []models.ItemsRs{}
		datos.Items = datosRS

		datos.Customer.Name = Data.Customer.Name
		datos.Customer.Adress.Coordinates = Data.Address.Coordinates
		datos.Business.Name = Data.Restaurant.Name

		//fmt.Print(datos)

		//dataJson, _ := json.MarshalIndent(datos, "", " ")

		//fmt.Print(string(dataJson))
	}
	*/
}
