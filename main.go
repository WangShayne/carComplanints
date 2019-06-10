package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	fmt.Println(1111)
	url := "http://www.12365auto.com/zlts/0-0-0-0-0-0_0-0-1.shtml"

	resp,err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Println("error code is :",resp.StatusCode)
		return
	}

	data,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Println(err)
	}
	fmt.Printf("%s\n",data)
}

