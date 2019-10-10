package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Accounts for public
type Accounts struct {
	ID   string
	Pass string
}

// Teststruct for public :
type Teststruct struct {
	Test    string
	TestTwo string
	Arr     []Accounts
}

func testHandle(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t Teststruct
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	fmt.Println(t.Test)
	fmt.Println(t.TestTwo)
	fmt.Println(t.Arr[0].ID)
	fmt.Println(t.Arr[0].Pass)
	fmt.Println(t.Arr[1].ID)
	fmt.Println(t.Arr[1].Pass)

}

func main() {

	fmt.Println("Server Start...")
	http.HandleFunc("/", testHandle)
	http.ListenAndServe(":9999", nil)

}
