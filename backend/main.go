package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Item struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Owner string `json:"owner"`
	Time  string `json:"time"`
}

func get_procfunc(w http.ResponseWriter, r *http.Request) {

	fmt.Println("received:", r.Body)
	
	file, err := ioutil.ReadFile("item.json")

	if err != nil {
		fmt.Println("error:", err)
	}
	var items []Item

	err = json.Unmarshal(file, &items)
	if err != nil {
		fmt.Println("error:", err)
	}

	for _, item := range items {
		fmt.Println(item)
	}
	fmt.Print("111")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
func add_procfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got request ", r.Body)
	fmt.Fprint(w, r.Body)
}

func main() {
	http.HandleFunc("/data/add", add_procfunc)
	http.HandleFunc("/data/get", get_procfunc)

	fmt.Println("server opened on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("open server failed", err)
	}

}
