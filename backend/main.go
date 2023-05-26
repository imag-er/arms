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

func getall_procfunc(w http.ResponseWriter, r *http.Request) {
	raw_json, err := ioutil.ReadFile("item.json")

	if err != nil {
		fmt.Println("error:", err)
	}
	var items []Item

	err = json.Unmarshal(raw_json, &items)
	if err != nil {
		fmt.Println("error:", err)
	}

	for _, item := range items {
		fmt.Println(item)
	}

	w.Header().Set("Content-Type", "application/json") // 指定内容类型
	w.Header().Set("Access-Control-Allow-Origin", "*") // 允许跨域访问

	length, err := w.Write(raw_json)

	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("sent ", length, " bytes")
	}

	fmt.Println("-----------------")
}
func add_procfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got ADD request ")
	fmt.Fprint(w, r.Body)
}
func delete_procfunc(w http.ResponseWriter, r *http.Request) {
	delete_id := r.FormValue("id")
	fmt.Println("got DELETE request ", delete_id)

	w.Header().Set("Content-Type", "application/json") // 指定内容类型
	w.Header().Set("Access-Control-Allow-Origin", "*") // 允许跨域访问

	fmt.Fprint(w, delete_id)

	fmt.Println("-----------------")

}
func main() {
	http.HandleFunc("/data/add", add_procfunc)
	http.HandleFunc("/data/getall", getall_procfunc)
	http.HandleFunc("/data/delete", delete_procfunc)

	fmt.Println("server opened on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("open server failed", err)
	}

}
