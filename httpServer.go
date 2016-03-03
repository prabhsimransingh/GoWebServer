package main

import (
	"fmt"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p Page) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	fmt.Fprint(w, "<h1>Sample GO Web server!</h1>")
}

func main() {
	var p Page
	err := http.ListenAndServe("localhost:8000",p)
	checkError(err)

}


func checkError(err error) {
	if err != nil {
		//Application will stop and exit the program
		panic(err)
	}

}
