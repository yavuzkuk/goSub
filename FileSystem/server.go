package filesystem

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

func ServerInfo(url string) {
	newUrl := IsHTTPS(url)

	fmt.Println("-----------------------------" + color.BlueString("Server Type") + "-----------------------------")

	resp, err := http.Get(newUrl)

	if err != nil {
		fmt.Println("Server info error -->", err)
	}

	fmt.Println(color.GreenString(resp.Header.Get("Server")))

	fmt.Println("---------------------------------------------------------------------")

}
