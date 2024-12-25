package main

import (
	"encoding/json"
	"fmt"
)



type course struct{
	Name string `json:"coursename"`
	Price int
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"`
}
func main(){
	fmt.Println("Json learning")
	//EncodeJson()
	DecodeJson()
}


func EncodeJson(){
	courses := []course{
		{"ReactJs ",299,"youtube.com","123abc", []string{"web-dev","js"}},
		{"Mern ",399,"youtube.com","123bc", []string{"web-dev","js"}},
		{"Angular ",599,"youtube.com","123cdcabc", nil},
	}
	//package this as data
	//  finalJson,err := json.Marshal(courses)
	finalJson,err := json.MarshalIndent(courses,"","\t")
	 checkNilErr(err)
	 fmt.Printf("%s\n",finalJson)
}

func DecodeJson(){
	jsonDataFromWeb := []byte(`
	 {
                "coursename": "ReactJs ",
                "Price": 299,
                "website": "youtube.com",
                "tags": ["web-dev","js"]
     }
	`)

	var devcourse course

	checkvalid := json.Valid(jsonDataFromWeb)
	if checkvalid {
		fmt.Println("Valid data")
		json.Unmarshal(jsonDataFromWeb, &devcourse)
		fmt.Printf("%#v\n",devcourse)
	}else{
		fmt.Println("UNVALID JSON")
	}
}

func checkNilErr(err error){
	if err!=nil{
		panic(err)
	}
}