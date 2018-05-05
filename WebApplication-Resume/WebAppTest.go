package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
)

func main(){
	urlGetBase := "http://localhost:8080/api/getResume/"
	fmt.Println("URL: >",urlGetBase)

	urlPostBase := "http://localhost:8080/api/uploadResumeDetails"

	values := map[string]string{
		"Name": "Clemence",
		"CurrentJobTitle": "Student",
		"CurrentJobDescription": "Study",
		"CurrentJobCompany": "SUTD",
	}

	values1 := map[string]string{
		"Name": "NotClemence",
		"CurrentJobTitle": "NotStudent",
		"CurrentJobDescription": "NotStudy",
		"CurrentJobCompany": "NotSUTD",
	}

	jsonValue, _ := json.Marshal(values)
	jsonValue1, _ := json.Marshal(values1)

	pResp, pErr := http.Post(urlPostBase, "application/json", bytes.NewBuffer(jsonValue))
	if pErr!=nil{
		panic(pErr)
	}
	fmt.Println(pResp)

	pResp2, pErr := http.Post(urlPostBase, "application/json", bytes.NewBuffer(jsonValue1))
	if pErr!=nil{
		panic(pErr)
	}
	fmt.Println(pResp2)

	// Get request for 0
	resp,err := http.Get(fmt.Sprintf("%s%d",urlGetBase,0))
	resp1,err1:= http.Get(fmt.Sprintf("%s%d",urlGetBase,1))
	if err!=nil{
		panic(err)
	}
	if err1!=nil{
		panic(err1)
	}

	fmt.Println(resp)
	fmt.Println(resp1)


}
