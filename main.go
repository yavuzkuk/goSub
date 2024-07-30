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

	// fmt.Println(cmd.Robots)
	// if cmd.Robots {
	// 	filesystem.Robots(cmd.Url)
	// }

	// filesystem.BruteForceFile(cmd.Url, cmd.WordList, cmd.RequestNumber, cmd.FilterStatusCode)
	// filesystem.SubDomainSearch(cmd.Url, cmd.WordList)

	// filesystem.GetIp(cmd.Url)

	// filesystem.ServerInfo(cmd.Url)

	filesystem.DNSRecord(cmd.Url, cmd.DNSType)

}
