package repos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type myIp struct {
	Ip       string `json:"ip"`
	Hostname string `json:""`
	City     string `json:""`
	Region   string `json:""`
	Country  string `json:""`
	Loc      string `json:""`
	Org      string `json:""`
}

type MyIpRepo struct {
	client http.Client
}

func BuildMyIpRepo() MyIpRepo {
	return MyIpRepo{client: http.Client{}}
}

func (o *MyIpRepo) GetIp() string {
	resp, err := o.client.Get("http://ipinfo.io/json")
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	response := myIp{}
	if err = json.Unmarshal(data, &response); err != nil {
		fmt.Printf(err.Error())
	}
	return response.Ip
}
