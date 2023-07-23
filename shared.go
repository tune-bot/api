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

var colourInRotation color = black

func rotateColor() color {
	colourInRotation = ((colourInRotation + 1) % white) + red
	return colourInRotation
}

func printLnColor(msg string, colour color) {
	fmt.Printf("\033[%dm%s\033[0m\n", colour, msg)
}

func successResponse(data []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	printLnColor("200 - OK", rotateColor())
	w.Write(data)
}

func errorResponse(err error, w http.ResponseWriter) {
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest) // TODO use more descriptive http codes
		printLnColor("400 - Bad Request: "+err.Error(), rotateColor())
		w.Write([]byte("{\"error\":\"" + strings.Replace(err.Error(), "\"", "\\\"", -1) + "\"}"))
	}
}
