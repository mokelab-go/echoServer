package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func readBody(req *http.Request) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	return buf.String()
}

func handler(w http.ResponseWriter, r *http.Request) {
	body := map[string]interface{}{
		"method": r.Method,
		"url":    r.URL.Path,
	}
	reqHeader := make(map[string]string)
	for key, val := range r.Header {
		reqHeader[key] = val[0]
	}
	body["header"] = reqHeader
	switch r.Method {
	case "POST":
		body["body"] = readBody(r)
		break
	case "PUT":
		body["body"] = readBody(r)
		break
	}
	data, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	h := w.Header()
	h.Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(data)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
