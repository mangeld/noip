package main

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/digitalocean/godo"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	var accessToken string = "your personal token string from digitalocean"
	changeDnsIp(accessToken, "the domain which A (ip) record you want to change, ex: home.mydomain.com")

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
