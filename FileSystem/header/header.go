package header

import (
	filesystem "Cyrops/FileSystem"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/fatih/color"
)

func RequestHeader(url string) {
	fmt.Println("-----------------------------" + color.BlueString("HTTP Request Header") + "-----------------------------")

	randomNumber := rand.Intn(5)
	var userAgent = [5]string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:91.0) Gecko/20100101 Firefox/91.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 12_3) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Safari/605.1.15",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.1823.51 Safari/537.36 Edg/114.0.1823.51",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 OPR/80.0.4170.63",
	}

	client := http.Client{}

	newUrl := filesystem.HTTPS(url)
	request, err := http.NewRequest("GET", newUrl, nil)

	request.Header.Add("User-agent", userAgent[randomNumber])
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	request.Header.Set("Accept-Language", "tr-TR,tr;q=0.9,en-US;q=0.8,en;q=0.7")
	request.Header.Set("Referer", "https://www.google.com/")
	request.Header.Set("Connection", "close")

	resp, err := client.Do(request)

	if err != nil {
		fmt.Println("Do error --> ", err)
	}

	for k, v := range request.Header {
		fmt.Println(k, " ----- ", v[0])
	}

	responseHeader(resp)
}

func responseHeader(resp *http.Response) {
	fmt.Println("-----------------------------" + color.BlueString("HTTP Response Header") + "-----------------------------")

	fmt.Println(resp.Status)
	fmt.Println("Request URL --- ", resp.Request.URL)
	fmt.Println("Content-Type --- ", resp.Header.Get("Content-Type"))
	fmt.Println("Server --- ", resp.Header.Get("Server"))

	if resp.Header.Get("Date") != "" {
		fmt.Println("Date --- ", resp.Header.Get("Date"))
	}

	if resp.Header.Get("X-Powered-By") != "" {
		fmt.Println("X-Powered-By --- ", resp.Header.Get("X-Powered-By"))
	}

	if resp.Header.Get("Cache-Control") != "" {
		fmt.Println("Cache-Control --- ", resp.Header.Get("Cache-Control"))
	}

	if resp.Header.Get("Content-Length") != "" {
		fmt.Println("Content-Length --- ", resp.Header.Get("Content-Length"))
	}

	if resp.Header.Get("X-Frame-Options") != "" {
		fmt.Println("X-Frame-Options --- ", resp.Header.Get("X-Frame-Options"))
	}

	if resp.Header.Get("Cache-Control") != "" {
		fmt.Println("Cache-Control --- ", resp.Header.Get("Cache-Control"))
	}
}
