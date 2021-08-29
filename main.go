package main

import (
	"idealista-flats/config"
	"idealista-flats/idealista"
	"idealista-flats/storage"
	"idealista-flats/telegram"
	"log"
)

func main() {
	cfg := config.LoadConfigs()
	t := telegram.New(cfg.TelegramBot)
	client := idealista.New(cfg.IdealistaID, cfg.IdealistaSecret)
	results := client.GetProperties(cfg.Search)
	for _, property := range results.Properties {
		if !idealista.IsValidBasicProperty(property, *cfg) {
			continue
		}
		if storage.InsertDeal(property) {
			propertyDetails := client.GetProperty(property.PropertyCode)
			if idealista.IsValidDetailedProperty(propertyDetails, *cfg) {
				err := t.SendMessage(cfg.TelegramChannel, property.URL)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}
