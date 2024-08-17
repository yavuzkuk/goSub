package header

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/fatih/color"
)

func RequestHeader(url string) {
	fmt.Println("-----------------------------" + color.BlueString("HTTP Request Header") + "-----------------------------")

	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	requestHeader, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Fatal(err)
	}

	parca := strings.Split(string(requestHeader), "\r\n")

	for _, v := range parca {
		fmt.Println(v)
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Request error --> ", err)
	}

	ResponseHeader(resp)
}

func ResponseHeader(resp *http.Response) {
	fmt.Println("-----------------------------" + color.BlueString("HTTP Response Header") + "-----------------------------")

	fmt.Println(resp.Status)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(resp.Header.Get("Request URL"))

	for k, v := range resp.Header {
		fmt.Println(k, " ---- ", v)
	}

}
