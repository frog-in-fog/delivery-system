package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RenderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	log.Println("RESP: ", string(js))
	w.Write(js)
}
