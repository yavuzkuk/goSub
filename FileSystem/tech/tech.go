package tech

import (
	filesystem "Cyrops/FileSystem"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/fatih/color"
)

func HTTPGet(url string) *http.Response {
	newUrl := filesystem.HTTPS(url)
	resp, err := http.Get(newUrl)

	if err != nil {
		fmt.Println("ERROR --> ", err)
	}

	return resp
}

func Tech(url string) {
	fmt.Println("-----------------------------" + color.BlueString("Technologies") + "-----------------------------")

	// for k, v := range res.Header {
	// 	fmt.Println(k, " --> ", v)
	// }

	ServerName(url)
	XpoweredBy(url)
	CDN(url)
	DetectCMS(url)
	OSDetection(url)
	DetectAnalytics(url)
	JSDetect(url)
}

func JSDetect(url string) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var jsFrameworks = [...]string{"bootstrap", "sweetalert", "jquery", "owl"}

	var requests []string

	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if ev, ok := ev.(*network.EventRequestWillBeSent); ok {
			requests = append(requests, ev.Request.URL)
		}
	})

	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(2*time.Second),
	); err != nil {
		log.Fatal(err)
	}

	values := make(map[string]string)

	for _, request := range requests {
		if strings.Contains(request, ".js") {
			// splitFile := strings.Split(request, "/")
			for _, v := range jsFrameworks {
				if strings.Contains(request, v) {
					values[v] = ""
				}
			}
		}
	}

	for k, _ := range values {
		strings.ToUpper(k)
	}
}

func DetectAnalytics(url string) {
	resp := HTTPGet(url)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Sayfa içeriği okunamadı: %v\n", err)
		fmt.Println("Tespit edilemedi")
	}

	content := string(body)
	var analytics []string

	gaPattern := regexp.MustCompile(`https?://www\.googletagmanager\.com`)
	if gaPattern.MatchString(content) {
		analytics = append(analytics, "Google Analytics")
	}

	fbPixelPattern := regexp.MustCompile(`https?://connect\.facebook\.net/.*/fbevents\.js`)
	if fbPixelPattern.MatchString(content) {
		analytics = append(analytics, "Facebook Pixel")
	}

	matomoPattern := regexp.MustCompile(`https?://.*\.matomo\.org/matomo\.js|piwik\.js`)
	if matomoPattern.MatchString(content) {
		analytics = append(analytics, "Matomo (Piwik)")
	}

	yandexMetricaPattern := regexp.MustCompile(`https?://mc\.yandex\.ru/metrika/watch\.js`)
	if yandexMetricaPattern.MatchString(content) {
		analytics = append(analytics, "Yandex Metrica")
	}

	hotjarPattern := regexp.MustCompile(`https?://static\.hotjar\.com/c/hotjar-.*\.js`)
	if hotjarPattern.MatchString(content) {
		analytics = append(analytics, "Hotjar")
	}

	crazyEggPattern := regexp.MustCompile(`https?://script\.crazyegg\.com/pages/scripts/\d+/\d+\.js`)
	if crazyEggPattern.MatchString(content) {
		analytics = append(analytics, "Crazy Egg")
	}

	if len(analytics) == 0 {
		analytics = append(analytics, "Analitik firması tespit edilemedi")
	} else {
		fmt.Print("Analytics: ")
		for _, v := range analytics {
			fmt.Print(color.GreenString(v), "    ")
		}
	}
}

func OSDetection(url string) {
	resp := HTTPGet(url)

	serverHeader := resp.Header.Get("Server")

	if strings.Contains(serverHeader, "Unix") || strings.Contains(serverHeader, "Linux") {
		fmt.Println("OS: ", color.GreenString("Linux/Unix"))
	} else if strings.Contains(serverHeader, "Win32") || strings.Contains(serverHeader, "Win64") || strings.Contains(serverHeader, "Windows") {
		fmt.Println("OS: ", color.GreenString("Windows"))
	} else if strings.Contains(serverHeader, "Darwin") {
		fmt.Println("OS: ", color.GreenString("macOS"))
	} else {
		fmt.Println(color.RedString("OS not detection"))
	}

}

func ServerName(url string) {

	resp := HTTPGet(url)
	server := resp.Header.Get("Server")

	if server == "" {
		fmt.Println(color.RedString("Server type not found."))
	} else {
		fmt.Printf("Server: %s \n", color.GreenString(server))
	}
}

func XpoweredBy(url string) {

	resp := HTTPGet(url)

	extens := []string{".js", ".net", ".php", ".node", ".java", ".ruby", ".django", ".laravel", ".symfony", ".cakephp", ".codeigniter", ".flask", ".fastapi", ".spring", ".gin", ".echo"}
	xpoweredBy := resp.Header.Get("X-Powered-By")

	if xpoweredBy == "" {
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("Respond body error ")
		}

		for _, v := range extens {
			if strings.Contains(string(body), v) {
				fmt.Printf("Programming language: %s \n", color.GreenString(v))
			}
		}
	} else {
		fmt.Printf("Programming language: %s \n", color.GreenString(xpoweredBy))
	}
}

func CDN(url string) {

	newUrl := filesystem.HTTPS(url)
	resp, err := http.Get(newUrl)

	if err != nil {
		fmt.Println("HTTP get error --> ", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
	}

	cdnPattern := `https://cdn\.[a-zA-Z0-9-]+\.[a-z]+[^\s]*`

	re := regexp.MustCompile(cdnPattern)
	matches := re.FindAllString(string(body), -1)

	cdnList := []string{}
	if len(matches) > 0 {
		fmt.Printf("CDN: ")
		for _, match := range matches {
			var isContain int = sort.SearchStrings(cdnList, match)

			if isContain == 0 {
				splitUrl := strings.Split(match, "/")[2]
				cdnList = append(cdnList, splitUrl)
				fmt.Println(color.GreenString(splitUrl), "    ")
			}
		}
	} else {
		fmt.Println(color.RedString("CDN not found"))
	}
}

func Ssl(url string) {

	fmt.Println("-----------------------------" + color.BlueString("SSL") + "-----------------------------")

	newUrl := url + ":443"
	conn, err := tls.Dial("tcp", newUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cert := conn.ConnectionState().PeerCertificates[0]
	fmt.Printf("Subject: %s\n", cert.Subject.CommonName)

	fmt.Printf("Issuer: %s\n", cert.Issuer.CommonName)

	fmt.Printf("Expires: %s\n", cert.NotAfter.Format("2 January 2006"))

	fmt.Printf("Renewed: %s\n", cert.NotBefore.Format("2 January 2006"))

	serialNumber := cert.SerialNumber.String()
	fmt.Printf("Serial Num: %s\n", serialNumber)
}

func DetectCMS(url string) {
	resp := HTTPGet(url)

	defer resp.Body.Close()

	headers := resp.Header

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Resp body error --> ", err)
	}

	respBody := string(body)

	if strings.Contains(respBody, "WordPress") || strings.Contains(headers.Get("Set-Cookie"), "wordpress") {
		fmt.Println("WordPress")
	} else if strings.Contains(respBody, "Joomla") || strings.Contains(headers.Get("Set-Cookie"), "joomla") {
		fmt.Println("Joomla")
	} else if strings.Contains(respBody, "Drupal") || strings.Contains(headers.Get("Set-Cookie"), "drupal") {
		fmt.Println("Drupal")
	} else if strings.Contains(respBody, "Wix") {
		fmt.Println("Wix")
	} else if strings.Contains(respBody, "Squarespace") {
		fmt.Println("Squarespace")
	} else if strings.Contains(respBody, "Shopify") {
		fmt.Println("Shopify")
	} else {
		fmt.Println(color.RedString("CMS not specified"))
	}
}
