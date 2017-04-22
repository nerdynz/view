package view

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// JSON
func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Template
func Template(w http.ResponseWriter, f func(interface{}, *bytes.Buffer), status int, data interface{}) {
	buffer := new(bytes.Buffer)
	f(data, buffer)
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/html")
	w.Write(buffer.Bytes())
}
