//Run with cron (or whatever you want to use), for example every 5 or 10 mins.
//Change the access token to the one that digitalocean gave you (Personal access token)
//In order to this to work you have to have created first the domain that you want to use
//(it won't create it, it only updates it)
package main

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/digitalocean/godo"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
        var accessToken string = "your personal access token"
        var domain string = "your domain"
        changeDnsIp(accessToken, domain)
}

func getOwnIp() string {
	resp, _ := http.Get("http://www.telize.com/ip")
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	return strings.Trim(string(data), "\n")
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
