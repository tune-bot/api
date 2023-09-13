package main

import (
	"net/http"
	"strings"

	"github.com/tune-bot/core"
)

func successResponse(data []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	core.PrintSuccess("200 - OK")
	w.Write(data)
}

func errorResponse(err error, w http.ResponseWriter) {
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest) // TODO use more descriptive http codes
		core.PrintError("400 - Bad Request: " + err.Error())
		w.Write([]byte("{\"error\":\"" + strings.Replace(err.Error(), "\"", "\\\"", -1) + "\"}"))
	}
}
