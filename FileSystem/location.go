package filesystem

import (
	"fmt"
	"log"
	"net"

	"github.com/fatih/color"
	"github.com/oschwald/geoip2-golang"
)

func GetLocation(url string) {
	fmt.Println("-----------------------------" + color.BlueString("Server Location") + "-----------------------------")

	ip := JustIp(url)

	db, err := geoip2.Open("Geo/GeoIP2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ipaddrs := net.ParseIP(ip)
	record, err := db.City(ipaddrs)
	if err != nil {
		log.Fatal(err)
	}
	if record.City.Names["en"] == "" {
		fmt.Println(color.RedString("City can't find"))
	} else {
		fmt.Println(record.City.Names["en"])
	}

	if record.Country.Names["en"] == "" {
		fmt.Println(color.RedString("Country can't find"))
	} else {
		fmt.Println(record.Country.Names["en"])
	}
}
