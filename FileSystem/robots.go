package filesystem

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fatih/color"
)

func Robots(url string) {

	newUrl := IsHTTPS(url) + "/robots.txt"

	resp, err := http.Get(newUrl)

	if err != nil {
		fmt.Println("Robots.txt error -->", err)
	}
	fmt.Println("-----------------------------" + color.BlueString("robots.txt") + "-----------------------------")
	if resp.StatusCode == 200 {
		fmt.Printf("%s website have robots.txt file ---> ", newUrl)
		color.Green(strconv.Itoa(resp.StatusCode))
	} else {
		fmt.Printf("%s website don't have robots.txt file ---> ", newUrl)
		color.Red(strconv.Itoa(resp.StatusCode))
	}
	fmt.Println("--------------------------------------------------------------------")
}
