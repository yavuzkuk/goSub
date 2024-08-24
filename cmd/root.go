package cmd

import (
	filesystem "Cyrops/FileSystem"
	"Cyrops/FileSystem/header"
	"Cyrops/FileSystem/ssl"
	"Cyrops/FileSystem/tech"
	"Cyrops/FileSystem/whois"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mbndr/figlet4go"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goSUB",
	Short: "This tool created with Golang. With this tool, you can scan the website you provide as parameters.",
	Run: func(cmd *cobra.Command, args []string) {
		if DNSType != "" {
			DNSType = strings.ToUpper(DNSType)
		}

		if !Banner {
			ascii := figlet4go.NewAsciiRender()

			options := figlet4go.NewRenderOptions()
			options.FontColor = []figlet4go.Color{
				figlet4go.ColorRed,
			}

			renderStr, _ := ascii.RenderOpts("gosub", options)
			fmt.Print(renderStr)
		}

		if All {
			filesystem.Robots(Url)
			filesystem.ServerInfo(Url)
			filesystem.DNSRecord(Url, DNSType)
			filesystem.SPFRecord(Url)
			filesystem.GetIp(Url)
			whois.Whois(Url)
			filesystem.BruteForceFile(Url, DirectoryWordlist, RequestNumber, FilterStatusCode)
			filesystem.SubDomainSearch(Url)
			header.RequestHeader(Url)
			header.RequestHeader(Url)
			filesystem.WebArchive(Url)
			ssl.SSL(Url)
			tech.Tech(Url)
			filesystem.GetLocation(Url)
		}

		if Robots {
			filesystem.Robots(Url)
		}

		if DNS {
			filesystem.ServerInfo(Url)
			filesystem.DNSRecord(Url, DNSType)
			filesystem.SPFRecord(Url)
			filesystem.GetIp(Url)
		}

		if Whois {

			whois.Whois(Url)
		}

		if Brute {
			filesystem.BruteForceFile(Url, DirectoryWordlist, RequestNumber, FilterStatusCode)
		}

		if Sub {
			filesystem.SubDomainSearch(Url)
		}

		if Request {
			header.RequestHeader(Url)
			header.RequestHeader(Url)
		}

		// filesystem.WebArchive(Url)

		// filesystem.Folders(Url)

		if Directory {
			filesystem.WebArchive(Url)
		}

		if SSL {
			ssl.SSL(Url)
		}

		if Tech {
			tech.Tech(Url)
		}

		if Location {
			filesystem.GetLocation(Url)
		}
	},
	Version: "1",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var Url string
var DirectoryWordlist string
var Robots bool
var DNSType string
var RequestNumber int
var FilterStatusCode string
var Whois bool
var DNS bool
var SSL bool
var Location bool
var Tech bool
var Request bool
var Directory bool
var Sub bool
var Brute bool
var All bool
var Banner bool

func init() {

	rootCmd.Flags().StringVarP(&Url, "url", "u", "", "You need to specify URL required")
	rootCmd.Flags().StringVarP(&DirectoryWordlist, "wordlist", "w", "wordlist/seclistWebContent.txt", "You can specify Directory Wordlist")
	rootCmd.Flags().BoolVarP(&Robots, "robots", "r", false, "With default value the tool check the robots.txt file")
	rootCmd.Flags().StringVarP(&DNSType, "DNS Record Type", "d", "A-AAAA-NS-MX-TXT", "A Record: IPv4 address\nAAAA Record:Ipv6 address\nMX Record:Mail record\nNS Record:Name server record\nTXT Record:Domain info text")
	rootCmd.Flags().IntVarP(&RequestNumber, "count", "c", 10, "Request count")
	rootCmd.Flags().StringVarP(&FilterStatusCode, "Filter HTTP Status Code", "f", "200,404", "You can filter HTTP Statsus Code with -f parameter")
	rootCmd.Flags().BoolVarP(&DNS, "DNS", "", false, "With default value the tool check the DNS record")
	rootCmd.Flags().BoolVarP(&SSL, "ssl", "", false, "Check SSL certification")
	rootCmd.Flags().BoolVarP(&Location, "location", "l", false, "Enable location")
	rootCmd.Flags().BoolVarP(&Tech, "tech", "t", false, "Enable Technologies search")
	rootCmd.Flags().BoolVarP(&Whois, "whois", "", false, "Whois")
	rootCmd.Flags().BoolVarP(&Request, "rr", "", false, "Request & response header")
	rootCmd.Flags().BoolVarP(&Directory, "dir", "", false, "All directory")
	rootCmd.Flags().BoolVarP(&Sub, "sub", "", false, "Subdomains scan")
	rootCmd.Flags().BoolVarP(&Brute, "brute", "b", false, "Brute force")
	rootCmd.Flags().BoolVarP(&All, "all", "", false, "All parameters")
	rootCmd.Flags().BoolVarP(&Banner, "banner", "x", false, "Hide banner")

	err := rootCmd.MarkFlagRequired("url")

	if err != nil {
		log.Fatalln(err)
	}

}
