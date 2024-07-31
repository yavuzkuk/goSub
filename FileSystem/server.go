package filesystem

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

func ServerInfo(url string) {
	newUrl := HTTPS(url)

	fmt.Println("-----------------------------" + color.BlueString("Server Type") + "-----------------------------")

	resp, err := http.Get(newUrl)

	if err != nil {
		fmt.Println("Server info error -->", err)
	}

	if resp.Header.Get("Server") == "" {
		fmt.Println(color.RedString("Server type unknow"))
	} else {
		fmt.Printf(color.GreenString("Server type %s\n", resp.Header.Get("Server")))
	}
}
