package filesystem

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

func SubDomainSearch(url string) {
	fmt.Println("-----------------------------" + color.BlueString("Subdomain") + "-----------------------------")

	var GoogleSub = map[string]string{}
	var BingSub = map[string]string{}
	var YahooSub = map[string]string{}
	var SSLSub = map[string]string{}

	var totalSub = map[string]string{}

	newUrl := SplitUrl(url)

	SSLSub = SSLScraping(newUrl)
	BingSub = BingDork(newUrl)
	YahooSub = Yahoo(newUrl)
	GoogleSub = GoogleDork(newUrl)

	for k, _ := range GoogleSub {
		totalSub[k] = ""
	}

	for k, _ := range BingSub {
		totalSub[k] = ""
	}

	for k, _ := range YahooSub {
		totalSub[k] = ""
	}

	for k, _ := range SSLSub {
		totalSub[k] = ""
	}

	for k, _ := range BingSub {
		totalSub[k] = ""
	}

	for k, _ := range totalSub {
		fmt.Println(k)
	}
	fmt.Println(len(totalSub))

	// DNSDumpster(newUrl)

}

func DNSDumpster(url string) {

	domain := "https://www.virustotal.com/ui/domains/omu.edu.tr/subdomains?relationships=resolutions"

	randomNumber := rand.Intn(5)
	var userAgent = [5]string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:91.0) Gecko/20100101 Firefox/91.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 12_3) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Safari/605.1.15",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.1823.51 Safari/537.36 Edg/114.0.1823.51",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 OPR/80.0.4170.63",
	}

	client := http.Client{}

	request, err := http.NewRequest("GET", domain, nil)

	request.Header.Add("User-agent", userAgent[randomNumber])
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	request.Header.Set("Accept-Language", "en-US,en;q=0.5")
	request.Header.Set("Referer", "https://www.google.com/")
	request.Header.Add("x-recaptcha-response", "true")

	response, err := client.Do(request)

	if err != nil {
		fmt.Println("Response error --> ", err)
	}

	if err != nil {
		fmt.Println("Body error --> ", err)
	}

	fmt.Println(response)
}

func SSLScraping(url string) map[string]string {

	target := "https://crt.sh/?q=" + url
	subdomains := map[string]string{}

	res, err := http.Get(target)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("HTTP isteği başarısız: %d %s", res.StatusCode, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	html := string(body)

	regexpPattern2 := "<TD>([a-z . -]*)*</TD>"
	re := regexp.MustCompile(regexpPattern2)

	var subdomain = map[string]string{}

	match := re.FindAllStringSubmatch(html, -1)
	for _, v := range match {
		pattern := regexp.MustCompile("</?TD>")
		newText := pattern.ReplaceAllString(v[0], "")
		subdomain[newText] = ""
	}

	for k, _ := range subdomain {
		subdomains[k] = ""
	}

	return subdomains
}

func BingDork(url string) map[string]string {

	var domain string = "https://www.bing.com/search?q=insite:" + url + "&first="
	var subdomains = map[string]string{}
	for i := 0; i < 101; i += 10 {

		targetUrl := domain + strconv.Itoa(i)

		response, err := http.Get(targetUrl)

		if err != nil {
			fmt.Println("Error --> ", err)
		}

		doc, err := goquery.NewDocumentFromReader(response.Body)

		if err != nil {
			fmt.Println("Document error --> ", err)
		}

		doc.Find("h2 a").Each(func(i int, s *goquery.Selection) {
			href, exists := s.Attr("href")

			if exists && strings.Contains(href, url) {
				newUrl := SplitUrl(href)
				subdomains[newUrl] = ""
			}
		})
	}

	return subdomains
}

func GoogleDork(url string) map[string]string {

	randomNumber := rand.Intn(5)
	var userAgent = [5]string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:91.0) Gecko/20100101 Firefox/91.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 12_3) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Safari/605.1.15",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.1823.51 Safari/537.36 Edg/114.0.1823.51",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 OPR/80.0.4170.63",
	}

	mainDomain := "https://www.google.com/search?q=inurl:" + url + "&start="

	subdomains := map[string]string{}

	for i := 0; i < 150; i += 10 {

		targetUrl := mainDomain + strconv.Itoa(i)

		client := http.Client{}

		request, err := http.NewRequest("GET", targetUrl, nil)

		if err != nil {
			fmt.Println("Request error --> ", err)
		}

		request.Header.Add("User-agent", userAgent[randomNumber])
		request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
		request.Header.Set("Accept-Language", "en-US,en;q=0.5")
		request.Header.Set("Referer", "https://www.google.com/")

		response, err := client.Do(request)

		if err != nil {
			fmt.Println("Client request error --> ", err)
		}

		if response.StatusCode == 200 {

			doc, _ := goquery.NewDocumentFromReader(response.Body)

			doc.Find("cite").Each(func(i int, s *goquery.Selection) {
				href := s.Text()

				if strings.Contains(href, url) {
					links := SplitUrl(strings.Split(href, " ")[0])

					subdomains[links] = ""
				}
			})
		}
	}
	return subdomains
}

func Yahoo(url string) map[string]string {

	var subdomains = map[string]string{}
	mainUrl := "https://search.yahoo.com/search?p=intitle:" + url + "&b="

	for i := 1; i < 252; i += 10 {

		newUrl := mainUrl + strconv.Itoa(i)
		response, err := http.Get(newUrl)

		if err != nil {
			fmt.Println("Error --> ", err)
		}

		doc, _ := goquery.NewDocumentFromReader(response.Body)

		doc.Find("div h3 a").Each(func(i int, s *goquery.Selection) {
			link, exists := s.Attr("href")

			if exists && strings.Contains(link, url) {
				newLink := strings.Split(link, " ")[0]
				newLink = SplitUrl(newLink)

				subdomains[newLink] = ""
			}

		})
	}
	return subdomains
}
