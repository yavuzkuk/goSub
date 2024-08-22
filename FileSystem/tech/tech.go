package tech

import (
	filesystem "Cyrops/FileSystem"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
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

	// ServerName(url)
	// DetectChat(url)
	ServerInfo(url)
	XpoweredBy(url)
	CDN(url)
	DetectCMS(url)
	DetectAnalytics(url)
	JSDetect(url)
	Icons(url)
	OtherDetect(url)
	CSSDetect(url)
}

func CSSDetect(url string) {

	cssLibraries := map[string]string{"Bootstrap": "bootstrap", "Bulma": "bulma", "Tailwind CSS": "tailwind", "Foundation": "foundation", "Materialize": "materialize", "Semantic UI": "semantic", "PureCSS": "pure", "UIKit": "uikit", "Skeleton": "skeleton", "Milligram": "milligram", "Spectre": "spectre", "Tachyons": "tachyons", "Material Design Lite": "material", "Normalize": "normalize", "Base CSS": "base", "MUI": "MuiCont"}

	useCss := map[string]string{}
	newUrl := filesystem.HTTPS(url)

	response, err := http.Get(newUrl)

	if err != nil {
		fmt.Println("Response errorr --> ", err)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		fmt.Println("Document error --> ", err)
	}

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		source, exists := s.Attr("src")

		if exists {
			for k, v := range cssLibraries {
				if strings.Contains(source, v) {
					useCss[k] = ""
				}
			}
		}
	})

	doc.Find("link").Each(func(i int, s *goquery.Selection) {
		source, exists := s.Attr("href")

		if exists {
			for k, v := range cssLibraries {
				if strings.Contains(source, v) {
					useCss[k] = ""
				}
			}
		}

	})

	doc.Find("div").Each(func(i int, s *goquery.Selection) {
		source, exists := s.Attr("class")

		if exists {
			for k, v := range cssLibraries {
				if strings.Contains(source, v) {
					useCss[k] = ""
				}
			}
		}

	})

	if len(useCss) > 0 {
		fmt.Println("************************" + color.MagentaString("CSS") + "************************")
		for k, _ := range useCss {
			fmt.Println(k)
		}
	}

}

func OtherDetect(url string) {

	otherDetect := map[string]string{}

	newUrl := filesystem.HTTPS(url)

	response, err := http.Get(newUrl)

	if err != nil {
		fmt.Println("Response error --> ", err)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		fmt.Println("Document error --> ", err)
	}

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		property, exists := s.Attr("property")

		if exists && strings.Contains(property, "og:") {
			otherDetect["Open Graph"] = ""
		}
	})

	if len(otherDetect) != 0 {
		fmt.Println("************************" + color.MagentaString("Other") + "************************")
		for k, _ := range otherDetect {
			fmt.Println(k)
		}
	}
}

// func DetectChat(url string) {

// 	chat := map[string]string{}
// 	chatSystems := map[string]string{"Intercom": "intercom", "Zendesk Chat": "zendesk", "Drift": "drift", "Tawk.to": "tawk", "LiveChat": "livechat", "Crisp": "crisp", "Olark": "olark", "HubSpot": "hubspot", "Freshchat": "freshchat", "JivoChat": "jivosite", "Pure Chat": "purechat", "SnapEngage": "snapengage", "Userlike": "userlike", "Smartsupp": "smartsupp", "Comm100": "comm100", "Zoho SalesIQ": "zoho", "Tidio": "tidio", "LivePerson": "liveperson", "HelpCrunch": "helpcrunch"}

// 	newUrl := filesystem.HTTPS(url)

// 	response, err := http.Get(newUrl)

// 	if err != nil {
// 		fmt.Println("Response error --> ", err)
// 	}

// 	doc, err := goquery.NewDocumentFromReader(response.Body)

// 	if err != nil {
// 		fmt.Println("Document error --> ", err)
// 	}

// 	doc.Find("script").Each(func(i int, s *goquery.Selection) {
// 		source, exists := s.Attr("src")

// 		if exists {
// 			for k, v := range chatSystems {
// 				if strings.Contains(source, v) {
// 					chat[k] = ""
// 				}
// 			}
// 		}
// 	})

// 	if len(chat) > 0 {
// 		fmt.Println("************************" + color.MagentaString("Fonts") + "************************")
// 		for k, _ := range chat {
// 			fmt.Println(k)
// 		}
// 	}

// }

func Icons(url string) {
	iconsMap := map[string]string{"Font Awesome": "fontawesome", "Material Icons": "materialicons", "Ionicons": "ionicons", "Bootstrap Icons": "bootstrap-icons", "Google Font": "fonts.googleapis"}
	useFont := map[string]string{}

	newUrl := filesystem.HTTPS(url)

	response, err := http.Get(newUrl)

	if err != nil {
		fmt.Println("Response error --> ", err)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		fmt.Println("Document error --> ", err)
	}

	doc.Find("link").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")

		if exists {
			for k, v := range iconsMap {
				if strings.Contains(href, v) {
					useFont[k] = ""
				}
			}
		}
	})

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		source, exists := s.Attr("src")

		if exists {
			for k, v := range iconsMap {
				if strings.Contains(source, v) {
					useFont[k] = ""
				}
			}
		}
	})

	if len(useFont) != 0 {
		fmt.Println("************************" + color.MagentaString("Fonts") + "************************")

		for k, _ := range useFont {
			fmt.Println(k)
		}
	}

}

func JSDetect(url string) {
	newUrl := filesystem.HTTPS(url)

	jsLib := map[string]string{}

	jsVersion := []string{"prettyPhoto", "jquery-migrate", "isotope", "swiper", "core-js", "lightbox", "clipboard", "slick", "aos", "owl", "fancybox", "jquery", "react", "vue", "angular", "lodash", "underscore", "moment", "axios", "chart", "three", "leaflet", "anime", "popper", "swiper", "select2", "owl.carousel", "gsap", "handlebars", "mustache", "backbone", "knockout", "redux", "socket.io", "leaflet", "highcharts", "semantic-ui", "fullcalendar", "sweetalert", "toastr"}

	response, err := http.Get(newUrl)

	if err != nil {
		fmt.Println("Response error --> ", err)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		fmt.Println("Document error --> ", err)
	}

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		source, exists := s.Attr("src")

		if exists {
			for _, v := range jsVersion {
				if strings.Contains(source, v) {
					jsLib[v] = ""
				}
			}
		}
	})

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		id, exists := s.Attr("id")

		if exists {
			for _, v := range jsVersion {
				if strings.Contains(id, v) {
					jsLib[v] = ""
				}
			}
		}
	})

	doc.Find("div").Each(func(i int, s *goquery.Selection) {
		class, exists := s.Attr("class")

		if exists {
			for _, v := range jsVersion {
				if strings.Contains(class, v) {
					jsLib[v] = ""
				}
			}
		}
	})

	if len(jsLib) != 0 {
		fmt.Println("************************" + color.MagentaString("JS") + "************************")
		for k, _ := range jsLib {
			fmt.Println(k)
		}
	}
}

func DetectAnalytics(url string) {

	analytics := map[string]string{}

	analyticsSystems := map[string]string{"Google Tag Manager": "googletagmanager.com", "Google Analytics": "google-analytics.com", "Google Tag Services": "https://www.googletagservices.com", "Google Ad Services": "https://www.googleadservices.com", "Facebook Pixel": "connect.facebook.net", "Segment": "https://cdn.segment.com", "Hotjar": "https://www.hotjar.com", "Matomo": "https://cdn.matomo.cloud", "Google Analytics SSL": "https://ssl.google-analytics.com", "Twitter Analytics": "https://analytics.twitter.com", "Amplitude": "https://cdn.amplitude.com", "Mixpanel": "https://cdn.mixpanel.com", "FullStory": "https://www.fullstory.com", "Sentry": "https://cdn.sentry.io", "Heap Analytics": "https://www.heap.io", "Twitter Ads": "https://static.ads-twitter.com", "Cloudflare Analytics": "https://www.cloudflare.com/cdn-cgi/trace", "Intercom": "https://cdn.intercom.io", "Monsterinsights": "monsterinsights", "Fathom": "fathom"}
	newUrl := filesystem.HTTPS(url)

	response, err := http.Get(newUrl)

	if err != nil {
		fmt.Println("Request error --> ", err)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		fmt.Println("Document error --> ", err)
	}

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		source, exists := s.Attr("src")

		if exists {
			for k, v := range analyticsSystems {
				if strings.Contains(source, v) {
					analytics[k] = ""
				}
			}
		}
	})

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		source := s.Text()

		for k, v := range analyticsSystems {
			if strings.Contains(source, v) {
				analytics[k] = ""
			}
		}

	})

	doc.Find("link").Each(func(i int, s *goquery.Selection) {
		source, exists := s.Attr("href")

		if exists {
			for k, v := range analyticsSystems {
				if strings.Contains(source, v) {
					analytics[k] = ""
				}
			}
		}
	})

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		id, exists := s.Attr("id")

		if exists {
			for k, v := range analyticsSystems {
				if strings.Contains(id, v) {
					analytics[k] = ""
				}
			}
		}
	})

	if len(analytics) != 0 {
		fmt.Println("************************" + color.MagentaString("Analytics") + "************************")
		for k, _ := range analytics {

			fmt.Println(k)
		}
	}
}

func ServerInfo(url string) {

	WebServers := []string{"Apache", "Nginx", "IIS", "LiteSpeed", "Caddy", "Tomcat", "OpenResty", "Gunicorn", "Node.js", "lighttpd", "uWSGI", "Jetty", "Resin", "Oracle HTTP Server", "Zeus", "XAMPP", "WampServer", "Mongoose", "Tornado", "WEBrick", "GlassFish", "JBoss", "WildFly", "Varnish", "Tengine", "H2O", "Microsoft Azure Web Apps", "Google App Engine", "Heroku", "Apache Traffic Server"}

	operatingSystems := []string{"Windows", "macOS", "Linux", "FreeBSD", "OpenBSD", "NetBSD", "Solaris", "AIX", "HP-UX", "CentOS", "Debian", "Ubuntu", "Fedora", "Red Hat Enterprise Linux (RHEL)", "Arch Linux", "Alpine Linux", "Slackware", "openSUSE", "SUSE Linux Enterprise Server (SLES)", "Kali Linux", "Android", "iOS", "Chrome OS", "Windows Server", "VMware ESXi", "QNX", "Zorin OS", "Raspberry Pi OS", "Amazon Linux"}

	found := map[string]string{}

	resp := HTTPGet(url)

	serverHeader := resp.Header.Get("Server")

	if serverHeader != "" {
		splitString := strings.Split(serverHeader, " ")

		for _, v := range splitString {
			for _, w := range WebServers {
				if strings.Contains(v, w) {
					found["Web Server"] = w
				}
			}

			for _, o := range operatingSystems {
				if strings.Contains(v, o) {
					found["OS"] = o
				}
			}

		}
	}

	if len(found) > 0 {
		fmt.Println("************************" + color.MagentaString("Server Info") + "************************")
		for k, v := range found {
			fmt.Println(k, " -- ", v)
		}
	}

}

// func ServerName(url string) {

// 	resp := HTTPGet(url)
// 	server := resp.Header.Get("Server")

// 	if server == "" {
// 		fmt.Println(color.RedString("Server type not found."))
// 	} else {
// 		fmt.Println("************************" + color.MagentaString("Server") + "************************")

// 		fmt.Println(server)
// 	}
// }

func XpoweredBy(url string) {

	poweredBy := map[string]string{}

	extens := map[string]string{"PHP": ".php", "Python": ".py", "Ruby": ".rb", "Node.js": ".js", "Java": ".jsp", "ASP.NET": ".aspx", "Perl": ".pl", "ColdFusion": ".cfm", "Go": ".go", "Scala": ".scala", "C#": ".cs", "Elixir": ".ex", "Erlang": ".erl", "Rust": ".rs", "Kotlin": ".kt", "Next.js": "_next"}

	response := HTTPGet(url)

	xpowered := response.Header.Get("x-powered-by")

	if xpowered == "" {
		doc, err := goquery.NewDocumentFromReader(response.Body)

		if err != nil {
			fmt.Println("Document error --> ", err)
		}

		doc.Find("link").Each(func(i int, s *goquery.Selection) {
			href, exists := s.Attr("href")

			if exists {
				for k, v := range extens {
					expression := `^http(s)?://([^/]+/)+[^/]+` + v + `(\?.*)?$`

					reg := regexp.MustCompile(expression)

					result := reg.FindString(href)
					if strings.Contains(result, v) {
						poweredBy[k] = ""
					}
				}
			}
		})

		// doc.Find("script").Each(func(i int, s *goquery.Selection) {
		// 	href, exists := s.Attr("src")

		// 	if exists {
		// 		for k, v := range extens {
		// 			if strings.Contains(href, v) {
		// 				poweredBy[k] = ""
		// 			}

		// 		}
		// 	}
		// })

		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			href, exists := s.Attr("href")

			if exists && strings.Contains(href, url) {
				for k, v := range extens {
					expression := `^http(s)?://([^/]+/)+[^/]+` + v + `(\?.*)?$`

					reg := regexp.MustCompile(expression)

					result := reg.FindString(href)
					if strings.Contains(result, v) {
						poweredBy[k] = ""
					}
				}
			}
		})

		if len(poweredBy) > 0 {
			fmt.Println("************************" + color.MagentaString("X-Powered-By") + "************************")

			for k, _ := range poweredBy {
				fmt.Println(k)
			}
		}

	} else {
		fmt.Println("************************" + color.MagentaString("X-Powered-By") + "************************")
		fmt.Println(xpowered)
	}

}

func CDN(url string) {

	cdnMap := map[string]string{
		"Cloudflare":        "cloudflare",
		"Akamai":            "akamai",
		"Amazon CloudFront": "cloudfront",
		"Google Cloud CDN":  "googleapis",
		"Microsoft Azure":   "azureedge",
		"Fastly":            "fastly",
		"StackPath":         "stackpath",
		"KeyCDN":            "keycdn",
		"CDN77":             "cdn77",
		"Alibaba Cloud CDN": "alicdn",
		"jsDelivr":          "jsdelivr",
		"Bootstrap CDN":     "bootstrapcdn",
		"cdnjs":             "cdnjs",
		"MaxCDN":            "maxcdn",
	}

	useCDN := map[string]string{}

	newUrl := filesystem.HTTPS(url)
	resp, err := http.Get(newUrl)

	if err != nil {
		fmt.Println("HTTP get error --> ", err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		fmt.Println("Document error --> ", err)
	}

	doc.Find("link").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")

		if exists && strings.Contains(href, "cdn") {
			for k, v := range cdnMap {
				if strings.Contains(href, v) {
					useCDN[k] = v
				}
			}
		}
	})

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("src")

		if exists && strings.Contains(href, "cdn") {
			for k, v := range cdnMap {
				if strings.Contains(href, v) {
					useCDN[k] = v
				}
			}
		}
	})

	if len(useCDN) > 0 {
		fmt.Println("************************" + color.MagentaString("CDN") + "************************")
		for k, _ := range useCDN {
			fmt.Println(k)
		}
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

	usedCMS := map[string]string{}

	cmsMap := map[string]string{"WordPress": "wordpress", "Joomla": "joomla", "Drupal": "drupal", "Magento": "magento", "Shopify": "shopify", "Wix": "wix", "Squarespace": "squarespace", "Blogger": "blogger", "TYPO3": "typo3", "Ghost": "ghost", "Concrete5": "concrete5", "Grav": "grav", "SilverStripe": "silverstripe", "MODX": "modx", "HubSpot CMS": "hubspot cms", "ExpressionEngine": "expressionengine", "Craft CMS": "craft cms", "Sitecore": "sitecore", "Umbraco": "umbraco", "Weebly": "weebly"}

	newUrl := filesystem.HTTPS(url)

	response, err := http.Get(newUrl)

	if err != nil {
		fmt.Println("Response error --> ", err)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		fmt.Println("Document error --> ", err)
	}

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		content, exists := s.Attr("content")
		content = strings.ToLower(content)
		if exists {
			for k, v := range cmsMap {
				if strings.Contains(content, v) {
					usedCMS[k] = ""
				}
			}
		}
	})

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		content, exists := s.Attr("src")
		if exists {
			for k, v := range cmsMap {
				if strings.Contains(content, v) {
					usedCMS[k] = ""
				}
			}
		}
	})

	doc.Find("link").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			for k, v := range cmsMap {
				if strings.Contains(href, v) {
					usedCMS[k] = ""
				}
			}
		}
	})

	if len(usedCMS) != 0 {
		fmt.Println("************************" + color.MagentaString("CMS") + "************************")
		for k, _ := range usedCMS {
			fmt.Println(k)
		}
	}
}
