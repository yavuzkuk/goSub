/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	filesystem "Cyrops/FileSystem"
	"Cyrops/cmd"
)

func main() {
	cmd.Execute()

	// _ := strings.Split(cmd.DNSType, "-")
	// fmt.Println("Main fonksiyon tarafı bu Url: ", cmd.Url, "\n")
	// fmt.Println("Main fonksiyon tarafı bu Wordlist: ", cmd.WordList, "\n")
	// fmt.Println("Main fonksiyon tarafı bu Robots.txt: ", cmd.Robots, "\n")
	// fmt.Println("Main fonksiyon tarafı bu DNS: ", dnsTypes, "\n")

	// filesystem.Robots(cmd.Url)

	// filesystem.BruteForceFile(cmd.Url, cmd.WordList, cmd.RequestNumber, cmd.FilterStatusCode)
	// filesystem.SubDomainSearch(cmd.Url, cmd.WordList)

	// filesystem.GetIp(cmd.Url)

	filesystem.ServerInfo(cmd.Url)

}
