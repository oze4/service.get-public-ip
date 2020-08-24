package main

import (
	"net/http"
	"strings"
	"io/ioutil"
)

type net struct {
	PublicIP string `json:"publicIp"`
}

func (n *net) GetPublicIP() {
	p, e := n.getPublicIP()
	if e != nil {
		panic(e.Error())
	}
	n.PublicIP = strings.TrimSpace(string(*p))
}

func (n *net) getPublicIP() (*string, error) {
	r, e := http.Get("https://icanhazip.com/")
	if e != nil {
		return nil, e
	}
	defer r.Body.Close()

	b, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return nil, e
	}

	sb := string(b)
	return &sb, nil
}