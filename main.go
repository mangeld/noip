//Run with cron (or whatever you want to use), for example every 5 or 10 mins.
//Change the access token to the one that digitalocean gave you (Personal access token)
//In order to this to work you have to have created first the domain that you want to use
//(it won't create it, it only updates it)
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mangeld/noip/repos"

	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

type Config struct {
	AccessToken string
	Domain      string
}

func (c *Config) Init() {
	c.AccessToken = os.Getenv("ACCESS_TOKEN")
	c.Domain = os.Getenv("DOMAIN")
}

func (t *Config) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{AccessToken: t.AccessToken}
	return token, nil
}

type MyIp struct {
	Ip       string `json:"ip"`
	Hostname string `json:""`
	City     string `json:""`
	Region   string `json:""`
	Country  string `json:""`
	Loc      string `json:""`
	Org      string `json:""`
}

type TokenSource struct {
	AccessToken string
}

func main() {
	config := Config{}
	config.Init()
	fmt.Printf("Using token: %v... for domain: %v\n", config.AccessToken[:8], config.Domain)
	changeDnsIp(&config)
}

func getOwnIp() string {
	repo := repos.BuildMyIpRepo()
	return repo.GetIp()
}

func changeDnsIp(config *Config) {
	oauth_client := oauth2.NewClient(context.Background(), config)
	client := godo.NewClient(oauth_client)
	listOps := godo.ListOptions{Page: 1, PerPage: 50}

	records, _, _ := client.Domains.Records(context.Background(), config.Domain, &listOps)
	var ipRecord = godo.DomainRecord{}
	for _, r := range records {
		if r.Type == "A" {
			ipRecord = r
		}
	}

	editRequest := godo.DomainRecordEditRequest{Data: getOwnIp()}
	fmt.Printf("Updating record %v to new ip: %v\n", config.Domain, editRequest.Data)
	_, _, err := client.Domains.EditRecord(context.Background(), config.Domain, ipRecord.ID, &editRequest)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
