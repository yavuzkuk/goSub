package filesystem

import (
	"fmt"
	"net"

	"github.com/fatih/color"
)

func DNSRecord(url string) {
	fmt.Println("-----------------------------" + color.BlueString("DNS Record Type") + "-----------------------------")
	MxRecord(url)
	NxRecord(url)
	AsRecord(url)
	fmt.Println("-----------------------------------------------------------------------")
}

func NxRecord(url string) {
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

func AsRecord(url string) {

	newUrl := SplitUrl(url)

	ipaddress, err := net.LookupIP(newUrl)

	if err != nil {
		fmt.Println("MX connect error -->", err)
	}

	if len(ipaddress) >= 2 {
		fmt.Println("IPv4  ---> ", color.GreenString(ipaddress[1].String()))
		fmt.Println("IPv6  ---> ", color.GreenString(ipaddress[0].String()))
	} else {
		fmt.Println("IPv4  ---> ", color.GreenString(ipaddress[0].String()))
	}
}
