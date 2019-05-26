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

    var responseObject models.Response
    json.Unmarshal(responseData, &responseObject)
	fmt.Println("hola")
    fmt.Println(len(responseObject.Datos))


    for i := 0; i < len(responseObject.Datos); i++ {
		fmt.Println(responseObject.Datos[i].ID)
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
		fmt.Print("items:")
		for j := 0; j < len(responseObject.Datos[i].Items); j++ {
			fmt.Println(responseObject.Datos[i].Items[j].Options)
		}


		//fmt.Println(string(jsonMap))

		fmt.Println("-----")
    }

}