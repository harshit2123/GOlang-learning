package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main(){
	fmt.Println("Learning about web servers")
	//PerformGetRequest()
//	PerformPostRequest()
PerformPostFormRequest()
}

func PerformPostRequest(){
      const myurl = "http://localhost:3000/post"

	  //fake json payload

	  requestBody := strings.NewReader(`
	        {
		   "coursename": "yo",
			"price": 23,
			"couseteacher":"harshitsir"}
	  `)

   response, err := http.Post(myurl,"application/json",requestBody)
   checkNilErr(err)
   defer response.Body.Close()
   
   content,_ :=  ioutil.ReadAll(response.Body)
   fmt.Println(string(content))

}


func PerformGetRequest(){
	const myurl = "http://localhost:3000/get"

	response,err :=http.Get(myurl)
	checkNilErr(err)
	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Lentg: ", response.ContentLength)

	var responseString strings.Builder
	content,_ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Bytecount is: ",byteCount)
	fmt.Println(responseString.String())

}

func PerformPostFormRequest(){
	const myurl = "http://localhost:3000/postform"
	//form data
	data := url.Values{}

	data.Add("firstname","harshit")
	data.Add("lastname","yadav")
	data.Add("email","harshit@go.dev")

	response, err := http.PostForm(myurl,data)
	checkNilErr(err)

	defer response.Body.Close()

	// response,err :=http.Get(myurl)
	// checkNilErr(err)
	// defer response.Body.Close()

	content,_ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func checkNilErr(err error){
	if err!=nil{
		panic(err)
	}
}