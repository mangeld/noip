//Run with cron (or whatever you want to use), for example every 5 or 10 mins.
//Change the access token to the one that digitalocean gave you (Personal access token)
//In order to this to work you have to have created first the domain that you want to use
//(it won't create it, it only updates it)
package main

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/digitalocean/godo"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
)

type MyIp struct {
    Ip string `json:"ip"`
    Hostname string `json:""`
    City string `json:""`
    Region string `json:""`
    Country string `json:""`
    Loc string `json:""`
    Org string `json:""`
}

func main() {
        var accessToken string = "your token"
        var domain string = "Domain to update"
        changeDnsIp(accessToken, domain)
}

func getOwnIp() string {
	resp, err := http.Get("http://ipinfo.io/json")
	if err != nil {
	    fmt.Printf(err.Error())
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	response := MyIp{}
	err = json.Unmarshal(data, &response)
	if err != nil {
	    fmt.Printf(err.Error())
	}
	return response.Ip
}

func changeDnsIp(accessToken string, domainName string) {

	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: accessToken},
	}

	client := godo.NewClient(t.Client())
	listOps := godo.ListOptions{Page: 1, PerPage: 50}

	records, _, _ := client.Domains.Records(domainName, &listOps)
	var ipRecord = godo.DomainRecord{}
	for _, r := range records {
		if r.Type == "A" {
			ipRecord = r
		}
	}

	editRequest := godo.DomainRecordEditRequest{Data: getOwnIp()}
	client.Domains.EditRecord(domainName, ipRecord.ID, &editRequest)
}
