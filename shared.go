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

var successColourInRotation color = cyan

func rotateSuccessColor() color {
	switch successColourInRotation {
	case green:
		successColourInRotation = blue
	case blue:
		successColourInRotation = cyan
	case cyan:
		successColourInRotation = green
	}
	return successColourInRotation
}

var errorColourInRotation color = magenta

func rotateErrorColor() color {
	switch errorColourInRotation {
	case red:
		errorColourInRotation = yellow
	case yellow:
		errorColourInRotation = magenta
	case magenta:
		errorColourInRotation = red
	}
	return errorColourInRotation
}

func printLnColor(msg string, colour color) {
	fmt.Printf("\033[%dm%s\033[0m\n", colour, msg)
}

func successResponse(data []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	printLnColor("200 - OK", rotateSuccessColor())
	w.Write(data)
}

func errorResponse(err error, w http.ResponseWriter) {
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest) // TODO use more descriptive http codes
		printLnColor("400 - Bad Request: "+err.Error(), rotateErrorColor())
		w.Write([]byte("{\"error\":\"" + strings.Replace(err.Error(), "\"", "\\\"", -1) + "\"}"))
	}
}
