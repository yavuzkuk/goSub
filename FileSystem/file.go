package filesystem

import (
	"Cyrops/wordlist"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Response struct {
	Subdomains []string `json:"subdomains"`
}

type IPInfo struct {
	IPAddress     string  `json:"ip"`
	City          string  `json:"city"`
	Region        string  `json:"region"`
	Country       string  `json:"country_name"`
	PostalCode    string  `json:"postal"`
	EuropeanUnion bool    `json:"in_eu"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	TimeZone      string  `json:"timezone"`
	CallingCode   string  `json:"country_code"`
	Currency      string  `json:"currency"`
	Languages     string  `json:"languages"`
	ASN           string  `json:"asn"`
	Organization  string  `json:"org"`
}

func parseStatusCodes(stringStatusCode string) []int {
	newStatusCode := strings.Split(stringStatusCode, ",")

	var integerStatusCodes []int

	for _, v := range newStatusCode {
		num, err := strconv.Atoi(v)

		if err != nil {
			fmt.Println("Integer response error -> ", err)
		}

		integerStatusCodes = append(integerStatusCodes, num)
	}

	return integerStatusCodes
}

func HTTPS(url string) string {
	isHttps := strings.Contains(url, "https://")

	if !isHttps {
		url = "https://" + url
		return url
	}
	return url
}

func SplitUrl(url string) string {

	containsHttps := strings.Contains(url, "https://")
	containsHttp := strings.Contains(url, "http://")

	var newUrl []string
	var currentUrl string = url
	if containsHttps {
		newUrl = strings.Split(url, "/")
		currentUrl = newUrl[2]
	}

	if containsHttp {
		newUrl = strings.Split(url, "/")
		currentUrl = newUrl[2]
	}

	containsWWW := strings.Contains(currentUrl, "www")

	if containsWWW {
		newUrl = strings.Split(currentUrl, "www.")

		currentUrl = newUrl[1]
	}

	return currentUrl
}

func BruteForceFile(url string, wordlistPath string, requestCount int, stringStatusCode string) {
	fmt.Println("-----------------------------" + color.BlueString("Directory scan") + "-----------------------------")
	integerStatusCodes := parseStatusCodes(stringStatusCode)

	url = HTTPS(url)

	wordArray := wordlist.ReadWordlistFile(wordlistPath)
	var totalSubDomain = 0
	var counter int = 0
	var requestCounter int = 0
	for i := 0; i < len(wordArray); i++ {
		counter++
		newUrl := url + "/" + wordArray[i]

		if counter == requestCount {
			time.Sleep(10 * time.Second)
			counter = 0
		}

		resp, err := http.Get(newUrl)
		requestCounter++

		if err != nil {
			fmt.Println("HTTP request error ", err)
		}

		for _, v := range integerStatusCodes {
			if v == resp.StatusCode {
				if resp.StatusCode == 200 {
					totalSubDomain++
					fmt.Printf("URL: %-70s ---- %d/%d ---> ", newUrl, requestCounter, len(wordArray))
					color.Green(strconv.Itoa(resp.StatusCode))
				} else if resp.StatusCode == 404 {
					fmt.Printf("URL: %-70s ----> %d/%d --->", newUrl, requestCounter, len(wordArray))
					color.Red(strconv.Itoa(resp.StatusCode))
				} else {
					fmt.Printf("URL: %-70s ----> %d/%d --->", newUrl, requestCounter, len(wordArray))
					color.Cyan(strconv.Itoa(resp.StatusCode))
				}
			}
		}
	}

	if totalSubDomain == 0 {
		fmt.Println(color.RedString("No directory found with the given URL and wordlist."))
	}
}

func GetIp(url string) string {
	fmt.Println("-----------------------------" + color.BlueString("IP Info") + "-----------------------------")
	isHttp := strings.Contains(url, "/")

	var newUrl string = url

	if isHttp {
		urlsArray := strings.Split(url, "/")
		newUrl = urlsArray[2]
	}

	ipaddress, err := net.LookupIP(newUrl)
	var IPv4 string
	if err != nil {
		fmt.Println("Ip address error ->", err)
	}

	if ipaddress != nil {
		for _, v := range ipaddress {
			if len(v) > 14 {
				fmt.Println("--- IPv6 -->", color.GreenString(v.String()))
			} else {
				fmt.Println("--- IPv4 -->", color.GreenString(v.String()))
			}
		}
	} else {
		fmt.Println("Ip address error")
	}
	return IPv4
}

func JustIp(url string) string {
	isHttp := strings.Contains(url, "/")

	var newUrl string = url

	if isHttp {
		urlsArray := strings.Split(url, "/")
		newUrl = urlsArray[2]
	}

	ipaddress, err := net.LookupIP(newUrl)
	var IPv4 string
	if err != nil {
		fmt.Println("Ip address error ->", err)
	}

	if ipaddress != nil {
		for _, v := range ipaddress {
			if len(v) > 14 {
				// fmt.Println("Domain -->", url, " --- IPv6 -->", color.GreenString(v.String()))
			} else {
				IPv4 = v.String()
				// fmt.Println("Domain -->", url, " --- IPv4 -->", color.GreenString(v.String()))
			}
		}
	}
	return IPv4
}
