package filesystem

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

func Folders(url string) {

	var newUrl string
	if strings.Contains(url, "http") || strings.Contains(url, "https") {
		newUrl = url
	} else {
		newUrl = HTTPS(url)
	}
	BodyContent(newUrl)
}

func BodyContent(url string) {

	client := http.Client{}

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Request error --> ", err)
	}

	request.Header.Add("User-agent", userAgent[rand.Intn(len(userAgent))])

	response, err := client.Do(request)

	if err != nil {
		fmt.Println("Response error --> ", err)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		fmt.Println("Document error --> ", err)
	}

	// fmt.Println("-----------------------------" + color.BlueString("Link scan") + "-----------------------------")

	// doc.Find("link").Each(func(i int, s *goquery.Selection) {
	// 	source, exists := s.Attr("href")
	// 	rel, exists2 := s.Attr("rel")

	// 	exists3 := strings.Contains(source, "cdn")
	// 	exists4 := strings.Contains(source, "api")

	// 	if exists && exists2 && !exists3 && !exists4 && rel == "stylesheet" {
	// 		fmt.Println(source)
	// 	}
	// })

	// fmt.Println("-----------------------------" + color.BlueString("Script scan") + "-----------------------------")

	// doc.Find("script").Each(func(i int, s *goquery.Selection) {
	// 	source, exists := s.Attr("src")

	// 	contains := strings.ContainsFunc(source, func(r rune) bool {
	// 		extens := []string{"cdn", "cse", "code", "google"}

	// 		for _, v := range extens {
	// 			if strings.Contains(source, v) {
	// 				return true
	// 			}
	// 		}
	// 		return false
	// 	})

	// 	if exists && !contains {
	// 		fmt.Println(source)
	// 	}
	// })
	fmt.Println("-----------------------------" + color.BlueString("Image scan") + "-----------------------------")

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		source, exists := s.Attr("src")

		isCnd := strings.Contains(source, "cdn")
		base64C := strings.Contains(source, "base64")

		if base64C {

			result := strings.Split(source, "base64,")[1]
			fmt.Println(result)

		}

		if exists && !isCnd {
			fmt.Println(source)
		}
	})
}
