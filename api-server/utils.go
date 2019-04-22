package main

import (
	"net/http"
	"encoding/json"
)

func Respond(w http.ResponseWriter, body interface{}, headers map[string]string) {
	data, _ := json.Marshal(map[string]interface{}{
		"data": body,
	})

	header := w.Header()
	header.Add("Content-Type", "application/json")
	for k, v := range headers {
		header.Add(k, v)
	}
	w.WriteHeader(200)
	w.Write(data)
}
