package main

import (
	"fmt"
	"net/http"
	"strings"
)

type color int

var black color = 30
var red color = 31
var green color = 32
var yellow color = 33
var blue color = 34
var magenta color = 35
var cyan color = 36
var white color = 37
var reset color = 0

func printLnColor(msg string, colour color) {
	fmt.Println("\\0%d%s\\u00\n", colour, msg)
}

func successResponse(data []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	printLnColor("200 - OK", green)
	w.Write(data)
}

func errorResponse(err error, w http.ResponseWriter) {
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest) // TODO use more descriptive http codes
		fmt.Println("400 - Bad Request")
		fmt.Println(err.Error())
		w.Write([]byte("{\"error\":\"" + strings.Replace(err.Error(), "\"", "\\\"", -1) + "\"}"))
	}
}
