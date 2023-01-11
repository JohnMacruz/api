package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {

	GetMethod()
	PostMethod()
}

func GetMethod() {

	resp, err := http.Get("https://reqres.in/api/users/2")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	fmt.Println("status code", resp.StatusCode)
	fmt.Println("The status code text we got is:", http.StatusText(resp.StatusCode))
	fmt.Println("Get method success")

}

// func PostMethod() {
// 	data := url.Values{
// 		"name": {"morpheus"},
// 		"job":  {"leader"},
// 	}

// 	resp, err := http.PostForm("https://reqres.in/api/users", data)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var res map[string]interface{}

// 	json.NewDecoder(resp.Body).Decode(&res)

// 	fmt.Println(res["form"])

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(int(resp.StatusCode))
// 	fmt.Println(string(body))
// 	fmt.Println("Post method success")
// }

func PostMethod() {

	datas := url.Values{
		"name": {"macruz"},
		"job":  {"leaders"},
	}
	resp, err := http.PostForm("https://reqres.in/api/users", datas)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	fmt.Println("status code", resp.StatusCode)
	fmt.Println("The status code text we got is:", http.StatusText(resp.StatusCode))
	fmt.Println("Post method success")

}
