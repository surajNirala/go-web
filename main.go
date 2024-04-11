package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/surajNirala/web/handlers"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func airth(data string) func(int, int) int {
	// fmt.Println("type :", data)
	if data == "add" {
		return add
	} else {
		return sub
	}
}

func arrSort() map[string]string {
	var arr = make(map[string]string)
	arr["1"] = "test1"
	arr["2"] = "test2"
	arr["5"] = "test5"
	arr["10"] = "test10"
	arr["9"] = "test9"
	arr["3"] = "test3"
	arr["4"] = "test4"
	return arr
	// fmt.Println("Before Sorting: ", arr)
}

func ascOrder(arrData []int) {
	sort.Slice(arrData, func(i, j int) bool {
		return arrData[i] < arrData[j]
	})
	fmt.Println("Sorted sortKey (ascending):", arrData)
}

func descOrder(arrData []int) {
	sort.Slice(arrData, func(i int, j int) bool {
		return arrData[i] > arrData[j]
	})
	fmt.Println("Sorted sortKey (descending):", arrData)
}

func getKeyFromArr(arrData map[string]string) []int {
	// keys := []int{}
	keys := make([]int, 0, len(arrData))
	for key, _ := range arrData {
		intKey, err := strconv.Atoi(key)
		if err != nil {
			panic(err) // Handle error if conversion fails
		}
		keys = append(keys, intKey)
	}
	return keys
}

func getValue() {
	data := make(map[string]string)
	data["1"] = "s1"
	data["2"] = "s2"
	data["4"] = "s4"
	data["3"] = "s3"
	fmt.Println("data", data)
	key := make([]string, 0)
	for index, _ := range data {
		key = append(key, index)
	}
	fmt.Println("key := ", key)
}

type constant string
type unconstant string

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home Page!")
	fmt.Println("Welcome to Home Page!")
}

func main() {

	router := mux.NewRouter()
	// Define your API routes
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/get-user", handlers.GetUser).Methods("GET")

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", router))

	// getValue()
	// var a constant = "10"
	// a = "20"
	// var a1 unconstant = "srj"
	// var a1 unconstant = "srjss"
	// fmt.Println("a => ", a)
	// fmt.Println("a1 => ", a1)
	// arrData := arrSort()
	// sortKey := getKeyFromArr(arrData)
	// fmt.Println("Key-Value Pair", arrData)
	// fmt.Println("Keys", sortKey)
	// sort.Strings(sortKey)
	// fmt.Println("Sorted-Keys", sortKey)
	// ascOrder(sortKey)
	// fmt.Println("ss")

	// res := airth("add")
	// fmt.Printf("TYpe %T", res)
	// fmt.Println("Res :=", res(2, 1))
}
