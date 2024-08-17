package filesystem

import (
	"fmt"
	"net"
	"strings"

	"github.com/fatih/color"
)

func DNSRecord(url string, dnstypes string) {
	seperatedDns := strings.Split(dnstypes, "-")
	fmt.Println("-----------------------------" + color.BlueString("DNS Record Type") + "-----------------------------")

	for _, v := range seperatedDns {
		if v == "MX" {
			MxRecord(url)
		} else if v == "NS" {
			NsRecord(url)
		} else if v == "A" {
			ARecord(url)
		} else if v == "AAAA" {
			AAAARecord(url)
		} else if v == "TXT" {
			TXTRecord(url)
		} else {
			fmt.Printf("Wrong arguments --> %s ", color.RedString(v))
		}
	}
}

func NsRecord(url string) {
	newUrl := SplitUrl(url)

	ns, err := net.LookupNS(newUrl)

	if err != nil {
		fmt.Println("MX connect error -->", err)
	}
	if len(ns) > 0 {
		for _, v := range ns {
			fmt.Println("Name Server ---> ", color.GreenString(v.Host))
		}
	} else {
		fmt.Println("Name Server ---> ", color.RedString("Not found"))
	}
}

func MxRecord(url string) {

	newUrl := SplitUrl(url)

	mx, err := net.LookupMX(newUrl)

	if err != nil {
		fmt.Println("MX error --> ", color.RedString(err.Error()))
	}

	if len(mx) > 0 {
		for _, v := range mx {
			fmt.Println("MX error ---> ", color.GreenString(v.Host))
		}
	}
}

func TXTRecord(url string) {
	newUrl := SplitUrl(url)

	txt, err := net.LookupTXT(newUrl)

	if err != nil {
		fmt.Println("TXT error --> ", color.RedString(err.Error()))
	}

	if len(txt) > 0 {
		fmt.Println("TXT Record  ---> ", color.GreenString(txt[0]))
	}
}

func ARecord(url string) {

	newUrl := SplitUrl(url)

	ipaddress, err := net.LookupIP(newUrl)

	if err != nil {
		fmt.Println("MX connect error -->", err)
	}

	for _, v := range ipaddress {
		if len(v) <= 5 {
			fmt.Println("A (IPv4): ---> ", color.GreenString(v.String()))
		}
	}
}

func AAAARecord(url string) {

	newUrl := SplitUrl(url)

	ipaddress, err := net.LookupIP(newUrl)

	if err != nil {
		fmt.Println("MX connect error -->", err)
	}

	if len(ipaddress) >= 2 {
		fmt.Println("AAAA (IPv6)  ---> ", color.GreenString(ipaddress[0].String()))
	} else {
		fmt.Println("AAAA (IPv6)  ---> ", color.RedString("Not Found"))
	}
}

func SPFRecord(url string) {
	fmt.Println("-----------------------------" + color.BlueString("SPF Record") + "-----------------------------")

	newUrl := SplitUrl(url)

	resp, err := net.LookupTXT(newUrl)

	if err != nil {
		fmt.Println("SPF error --> ", color.RedString(err.Error()))
	}

	for _, v := range resp {
		contains := strings.Contains(v, "spf")
		if contains {
			fmt.Println("SPF kaydı var --> ", color.GreenString(v))
		}
	}
}
