package main

import (
	"fmt"
	"net/url"
)

const myurl string ="https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj4r3fef"

func main(){
	fmt.Println("Learning about Urls")
	fmt.Println(myurl)

	//parsing
	result,err :=url.Parse(myurl)
	checkNilErr(err)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("the type of query params : %T\n",qparams)

	for _,val := range qparams {
		fmt.Println("param is: ", val)
	}
    
	partsofUrl := &url.URL{
		Scheme: "https",
		Host:"lco.dev",
		RawPath: "user=hitesh",
	}

	anotherURl:= partsofUrl.String()
	fmt.Println(anotherURl)
}

func checkNilErr(err error){
	if err!=nil{
		panic(err)
	}
}