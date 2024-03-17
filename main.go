package main

import "github.com/Lyyyttooon/mediaserver/protocol/http"

func main() {
	conn, err := http.New("https://movie.douban.com/j/search_tags?type=movie&source=index")
	if err != nil {
		panic(err)
	}
	conn.Get()
}
