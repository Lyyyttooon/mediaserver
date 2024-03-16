package main

import "github.com/Lyyyttooon/mediaserver/protocol/http"

func main() {
	conn, err := http.New("http://127.0.0.1:18080/ping")
	if err != nil {
		panic(err)
	}
	conn.Get()
}
