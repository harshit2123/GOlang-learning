package main

import (
	"fmt"
)

func main(){
	fmt.Println("Learning about maps now")

	lang := make(map[string]string)

	lang["JS"] = "javascript"
	lang["RB"] = "RUBY"
	lang["PY"] = "Python"


	fmt.Println("List of all lang: ", lang["JS"])
	delete(lang,"RB")
	fmt.Println("LIST after deleteion: ", lang)

}