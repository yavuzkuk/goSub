package cmd

import (
	filesystem "Cyrops/FileSystem"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gokuk",
	Short: "This tool created with Golang. With this tool, you can scan the website you provide as parameters.",
	Run: func(cmd *cobra.Command, args []string) {
		if DNSType != "" {
			DNSType = strings.ToUpper(DNSType)
		}
		// if Robots {
		// 	filesystem.Robots(Url)
		// }
		// fmt.Println(Whois)
		// if Whois {
		// 	filesystem.ServerInfo(Url)
		// 	filesystem.DNSRecord(Url, DNSType)
		// 	filesystem.SPFRecord(Url)
		// 	filesystem.GetIp(Url)
		// }

		// filesystem.BruteForceFile(Url, DirectoryWordlist, RequestNumber, FilterStatusCode)
		filesystem.SubDomainSearch(Url, SubdomainWordlist)
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
var SubdomainWordlist string
var Robots bool
var DNSType string
var RequestNumber int
var FilterStatusCode string
var Whois bool

func init() {
	rootCmd.Flags().StringVarP(&Url, "url", "u", "", "You need to specify URL required")
	rootCmd.Flags().StringVarP(&DirectoryWordlist, "wordlist", "w", "wordlist/seclistWebContent.txt", "You can specify Directory Wordlist")
	rootCmd.Flags().StringVarP(&SubdomainWordlist, "subdomain-wordlist", "s", "wordlist/seclistSubdomains5000.txt", "You can specify Subdomain Wordlist")
	rootCmd.Flags().BoolVarP(&Robots, "robots", "r", true, "With default value the tool check the robots.txt file")
	rootCmd.Flags().BoolVar(&Robots, "no-robots", false, "Disable the robots.txt feature")
	rootCmd.Flags().StringVarP(&DNSType, "DNS Record Type", "d", "A-AAAA-NS-MX-TXT", "A Record: IPv4 address\nAAAA Record:Ipv6 address\nMX Record:Mail record\nNS Record:Name server record\nTXT Record:Domain info text")
	rootCmd.Flags().IntVarP(&RequestNumber, "count", "c", 10, "Request count")
	rootCmd.Flags().StringVarP(&FilterStatusCode, "Filter HTTP Status Code", "f", "200,404", "You can filter HTTP Statsus Code with -f parameter")

	rootCmd.Flags().BoolVarP(&Whois, "whois", "", false, "With default value the tool check the robots.txt file")

	err := rootCmd.MarkFlagRequired("url")

	if err != nil {
		log.Fatalln(err)
	}

}
