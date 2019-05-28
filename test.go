package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/nicolasmarino/api/models"
)

type User struct {
	Name string
}

func main() {
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
	fmt.Println(len(responseObject.Datos))
	//fmt.Println(responseObject.Datos)
	var datos models.RestoSoft

	for i := 0; i < len(responseObject.Datos); i++ {
		/*fmt.Println(responseObject.Datos[i].ID)
		fmt.Println(responseObject.Datos[i].State)
		fmt.Println(responseObject.Datos[i].Pickup)
		fmt.Println(responseObject.Datos[i].Notes)
		fmt.Println(responseObject.Datos[i].RegisteredDate)
		fmt.Println(responseObject.Datos[i].Integration)
		fmt.Println(responseObject.Datos[i].Customer)
		fmt.Println(responseObject.Datos[i].Address)
		fmt.Println(responseObject.Datos[i].Restaurant)
		fmt.Println(responseObject.Datos[i].Total)
		fmt.Println(responseObject.Datos[i].Shipping)
		fmt.Println(responseObject.Datos[i].Subtotal)
		fmt.Print("items:")*/

		//opciones := models.Option{}*/

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

		/*for x := 0; x < len(Data.Items); x++ {
			datos.Items[x].Name = Data.Items[i].Name
			datos.Items[x].Quantity = Data.Items[i].Quantity
			datos.Items[x].Price = Data.Items[i].Price
		}*/
		//fmt.Print(datos)
		dataJson, _ := json.MarshalIndent(datos, "", " ")

		fmt.Print(string(dataJson))

	}
}
