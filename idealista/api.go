package idealista

import (
	"context"
	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2/clientcredentials"
	"log"
	"time"
)

type Idealista struct {
	client *resty.Client
}

func New(clientID string, clientSecret string) *Idealista {
	conf := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     "https://api.idealista.com/oauth/token",
	}
	hc := conf.Client(context.Background())
	client := resty.New().SetTimeout(1 * time.Minute).
	SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	client.SetTransport(hc.Transport)
	return &Idealista{
		client: client,
	}
}

func (c Idealista) GetProperties(search string) *Results {
	var properties *Results
	_, err := c.client.R().
		SetBody(search).
		SetResult(&properties).
		Post("https://api.idealista.com/3.5/es/search")
	if err != nil {
		log.Println(err)
	}
	return properties
}

func (c Idealista) GetProperty(propertyId string) PropertyDetails {
	var property PropertyDetails
	_, err := c.client.R().
		SetResult(&property).
		Get("https://api.idealista.com/3/es/detail/" + propertyId)
	if err != nil {
		log.Println(err)
	}
	return property
}
