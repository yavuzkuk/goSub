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
		} else if v == "NX" {
			NsRecord(url)
		} else if v == "A" {
			ARecord(url)
		} else if v == "AAAA" {
			AAAARecord(url)
		}
	}
	// MxRecord(url)
	// NsRecord(url)
	// AsRecord(url)
}

func NsRecord(url string) {
	newUrl := SplitUrl(url)

	ns, err := net.LookupNS(newUrl)

	if err != nil {
		fmt.Println("MX connect error -->", err)
	}

	for _, v := range ns {
		fmt.Println("Name Server ---> ", color.GreenString(v.Host))
	}
}

func MxRecord(url string) {

	newUrl := SplitUrl(url)

	mx, err := net.LookupMX(newUrl)

	if err != nil {
		fmt.Println("MX connect error -->", err)
	}

	for _, v := range mx {
		fmt.Println("Mail Server ---> ", color.GreenString(v.Host))
	}
}

func ARecord(url string) {

	newUrl := SplitUrl(url)

	ipaddress, err := net.LookupIP(newUrl)

	if err != nil {
		fmt.Println("MX connect error -->", err)
	}

	if len(ipaddress) >= 2 {
		fmt.Println("IPv4  ---> ", color.GreenString(ipaddress[1].String()))
	} else {
		fmt.Println("IPv4  ---> ", color.GreenString(ipaddress[0].String()))
	}
}

func AAAARecord(url string) {

	newUrl := SplitUrl(url)

	ipaddress, err := net.LookupIP(newUrl)

	if err != nil {
		fmt.Println("MX connect error -->", err)
	}

	if len(ipaddress) >= 2 {
		fmt.Println("IPv6  ---> ", color.GreenString(ipaddress[0].String()))
	} else {
		fmt.Println("IPv6  ---> ", color.RedString("Not Found"))
	}
}
