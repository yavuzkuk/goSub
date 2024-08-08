package filesystem

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/fatih/color"
	wappalyzer "github.com/projectdiscovery/wappalyzergo"
)

func Tech(url string) {
	fmt.Println("-----------------------------" + color.BlueString("Technologies") + "-----------------------------")

	newUrl := HTTPS(url)
	resp, err := http.DefaultClient.Get(newUrl)
	if err != nil {
		log.Fatal(err)
	}
	data, _ := io.ReadAll(resp.Body)

	wappalyzerClient, err := wappalyzer.New()
	fingerprints := wappalyzerClient.Fingerprint(resp.Header, data)

	for k, _ := range fingerprints {
		fmt.Println(color.GreenString(k))
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
