package ssl

import (
	filesystem "Cyrops/FileSystem"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/fatih/color"
)

func SSL(url string) {
	fmt.Println("-----------------------------" + color.BlueString("SSL") + "-----------------------------")

	newUrl := filesystem.SplitUrl(url)

	conn, err := tls.Dial("tcp", newUrl+":443", nil)
	if err != nil {
		fmt.Println("TLS error --> ", err)
	}
	defer conn.Close()

	certs := conn.ConnectionState().PeerCertificates

	if len(certs) == 0 {
		fmt.Println("Sertifika bulunamadı.")
		return
	}

	for _, v := range certs {
		fmt.Printf("Common Name (CN): %s\n", v.Subject.CommonName)
		fmt.Printf("Issuer Name: %s\n", v.Issuer.CommonName)
		fmt.Printf("Not Before: %s\n", v.NotBefore.Format(time.RFC3339))
		fmt.Printf("Not After: %s\n", v.NotAfter.Format(time.RFC3339))

		currentTime := time.Now()
		if currentTime.Before(v.NotBefore) {
			fmt.Println(color.YellowString("Sertifika henüz geçerli değil."))
		} else if currentTime.After(v.NotAfter) {
			fmt.Println(color.RedString("Sertifika geçerli."))
		} else {
			fmt.Println(color.GreenString("Sertifika geçerli."))
		}
	}
}
