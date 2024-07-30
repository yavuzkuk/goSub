package filesystem

import (
	"Cyrops/wordlist"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

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

func IsHTTPS(url string) string {
	isHttps := strings.Contains(url, "https://")

	if !isHttps {
		url = "https://" + url
		return url
	}
	return url
}

func BruteForceFile(url string, wordlistPath string, requestCount int, stringStatusCode string) {

	integerStatusCodes := parseStatusCodes(stringStatusCode)

	url = IsHTTPS(url)

	wordArray := wordlist.ReadWordlistFile(wordlistPath)
	var counter int = 0
	for i := 0; i < len(wordArray); i++ {
		counter++
		newUrl := url + "/" + wordArray[i]

		if counter == requestCount {
			time.Sleep(5 * time.Second)
			counter = 0
		}

		resp, err := http.Get(newUrl)

		if err != nil {
			fmt.Println("HTTP request error ", err)
		}

		for _, v := range integerStatusCodes {
			if v == resp.StatusCode {
				if resp.StatusCode == 200 {
					fmt.Printf("URL: %-70s ----> ", newUrl)
					color.Green(strconv.Itoa(resp.StatusCode))
				} else if resp.StatusCode == 404 {
					fmt.Printf("URL: %-70s ----> ", newUrl)
					color.Red(strconv.Itoa(resp.StatusCode))
				} else {
					fmt.Printf("URL: %-70s ----> ", newUrl)
					color.Cyan(strconv.Itoa(resp.StatusCode))
				}
			}
		}
	}
}

func SubDomainSearch(url string, wordlistPath string) {

	var counter int

	newUrl := IsHTTPS(url)

	wordlist := wordlist.ReadWordlistFile(wordlistPath)

	for i := 0; i < len(wordlist); i++ {
		counter++
		newUrls := strings.Split(newUrl, "/")

		subdomain := "https://" + wordlist[i] + "." + newUrls[2]

		resp, err := http.Get(subdomain)

		if err != nil && resp == nil {
			continue
		}
		if resp.StatusCode == 200 {
			fmt.Printf("%-45s %d/%d -----> ", subdomain, counter, len(wordlist))
			color.Green(strconv.Itoa(resp.StatusCode))
		}
	}
}

func GetIp(url string) {

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
	if len(ipaddress) >= 2 {
		fmt.Print("Domain: ", newUrl, " ----> IPv4 ")
		color.Green(ipaddress[1].String())
		fmt.Print("Domain: ", newUrl, " ----> IPv6 ")
		color.Green(ipaddress[0].String())
		IPv4 = ipaddress[1].String()
	} else {
		fmt.Print("Domain: ", newUrl, " ----> IPv4 ")
		color.Green(ipaddress[0].String())
		IPv4 = ipaddress[0].String()
	}

	GetLocation(IPv4)
}

func GetLocation(ip string) {
	apiEndpoint := "https://ipapi.co/" + ip + "/json"
	resp, err := http.Get(apiEndpoint)

	if err != nil {
		fmt.Println("Location error ->", err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var location IPInfo
	err = json.Unmarshal(body, &location)

	if err != nil {
		fmt.Println("JSON Unmarshall error ->", err)
	}

	fmt.Println(location.CallingCode + "/" + location.Country + "---" + location.Region + "---" + location.City)
	fmt.Println(location.Organization)

}
