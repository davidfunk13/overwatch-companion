package helpers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func SearchBattletags(b string) []byte{

	url := "https://playoverwatch.com/en-us/search/account-by-name/"+b
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	
	res, err := client.Do(req)
	
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	return body
}