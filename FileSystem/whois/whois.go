package whois

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/fatih/color"
)

func Whois(url string) {
	fmt.Println("-----------------------------" + color.BlueString("Whois") + "-----------------------------")

	parts := strings.Split(url, ".")
	if len(parts) < 2 {
		fmt.Errorf("Geçersiz url adı: %s", url)
	}
	tld := parts[len(parts)-1]
	var whoisServer string
	switch tld {
	case "com":
		whoisServer = "whois.verisign-grs.com:43"
	case "net":
		whoisServer = "whois.verisign-grs.com:43"
	case "org":
		whoisServer = "whois.publicinterestregistry.net:43"
	case "gov":
		whoisServer = "whois.dotgov.gov:43"
	case "edu":
		whoisServer = "whois.educause.edu:43"
	case "tr":
		whoisServer = "whois.nic.tr:43"
	default:
		fmt.Errorf("Desteklenmeyen url uzantısı: .%s", tld)
	}

	conn, err := net.Dial("tcp", whoisServer)
	if err != nil {
		fmt.Printf("Bağlantı hatası: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Fprintf(conn, url+"\r\n")
	desc := false

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {

		bodyString := scanner.Text()

		splitScanner := strings.SplitN(bodyString, ">>>", 2)
		for _, v := range splitScanner {
			if strings.Contains(v, "For more information on Whois status codes, please visit") {
				desc = true
			} else {
				if !desc {
					// gizlilik için saklanan kısımlar ekrana yazdırılmicak
					if !strings.Contains(v, "REDACTED FOR PRIVACY") {
						fmt.Println(v)
					}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Cevap okuma hatası: %v\n", err)
	}
}
