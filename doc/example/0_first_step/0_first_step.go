package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"math/rand"
	"net"
	"net/http"
	"time"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	UserAgent   = "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_1) AppleWebKit" +
		"/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"
)

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	//httpProxy := &url.URL{
	//	Scheme: "HTTP",
	//	Host: "192.168.124.220:1080",
	//}

	c := colly.NewCollector()

	//HTTP configuration
	c.WithTransport(&http.Transport{
		//Proxy: http.ProxyURL(httpProxy),
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", UserAgent)
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://google.com/")
}
