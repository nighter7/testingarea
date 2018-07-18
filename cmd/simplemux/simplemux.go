package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

var x *viper.Viper

func init() {
	x = viper.New()
	x.SetDefault("port", ":8081")
}

func TestHandler(arg1 http.ResponseWriter, arg2 *http.Request) {
	arg1.Header().Set("TestHeader", "hi")
	fmt.Fprint(arg1, "hola")
}

func TestHandler2(arg1 http.ResponseWriter, arg2 *http.Request) {
	fmt.Fprint(arg1, "hola2")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/test", TestHandler).Methods("GET")
	r.HandleFunc("/test2", TestHandler2)

	log.Fatal(http.ListenAndServe(x.GetString("port"), r))
}
