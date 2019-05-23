package repos

import (
	"context"

	"golang.org/x/oauth2"

	"github.com/digitalocean/godo"
)

type DigitaloceanRepo struct {
	client godo.Client
}

type oauth2Token struct {
	oauth2.TokenSource
}

func BuildDigitaloceanRepo(token string) DigitaloceanRepo {
	return DigitaloceanRepo{client: godo.NewClient(oauth2.NewClient(context.Background()))}
}

func (self *DigitaloceanRepo) SaveARecord(ip string, record string) {

}
