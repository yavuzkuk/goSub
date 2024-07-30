/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Cyrops",
	Short: "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {
		if DNSType != "" {
			DNSType = strings.ToUpper(DNSType)
			i := 0
			var dnsTypes = [5]string{"A", "AAAA", "MX", "NS", "A-AAAA-MX-NX"}
			for _, v := range dnsTypes {
				if v == DNSType {
					i++
					DNSType = v
				}
			}
			if i == 0 {
				DNSType = "A-AAAA-MX-NX"
			}
		}
	},
	Version: "1",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var Url string
var WordList string
var Robots bool
var DNSType string
var RequestNumber int
var FilterStatusCode string

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.Cyrops.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.Flags().StringVarP(&Url, "url", "u", "", "You need to specify URL")
	rootCmd.Flags().StringVarP(&WordList, "wordlist", "w", "wordlist/seclistWebContent.txt", "You can specify Wordlist")
	rootCmd.Flags().BoolVarP(&Robots, "robots.txt", "r", true, "With default value the tool check the robots.txt file")
	rootCmd.Flags().StringVarP(&DNSType, "DNS Record Type", "d", "A-AAAA-MX-NX", "A Record: IPv4 address\nAAAA Record:Ipv6 address\nMX Record:Mail record\nNS Record:Name server record")
	rootCmd.Flags().IntVarP(&RequestNumber, "count", "c", 10, "Request count")
	rootCmd.Flags().StringVarP(&FilterStatusCode, "Filter HTTP Status Code", "f", "200,404", "You can filter HTTP Statsus Code with -f parameter")

	err := rootCmd.MarkFlagRequired("url")

	if err != nil {
		log.Fatalln(err)
	}

}
