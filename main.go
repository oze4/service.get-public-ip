package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", getIPHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func getIPHandler(w http.ResponseWriter, r *http.Request) {
	n := net{}
	n.GetPublicIP()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(net{PublicIP: n.PublicIP})
}
