package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	IdealistaID           string
	IdealistaSecret       string
	TelegramBot           string
	TelegramChannel       string
	Search                string
	NotValidSentences     []string
	MinimumPictures       int
	NotValidNeighborhoods []string
	ShowOnlyAgency        bool
	NotValidFloors        []string
	NotValidAgencies      []string
}

func LoadConfigs() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return &Config{
		IdealistaID:           viper.GetString("idealista.client_id"),
		IdealistaSecret:       viper.GetString("idealista.client_secret"),
		TelegramBot:           viper.GetString("telegram.api_key"),
		TelegramChannel:       viper.GetString("telegram.channel_id"),
		Search:                viper.GetString("search"),
		NotValidSentences:     viper.GetStringSlice("filters.not_valid_sentences"),
		MinimumPictures:       viper.GetInt("filters.minimum_pictures"),
		NotValidNeighborhoods: viper.GetStringSlice("filters.not_valid_neighborhoods"),
		NotValidFloors:        viper.GetStringSlice("filters.not_valid_floors"),
		ShowOnlyAgency:        viper.GetBool("filters.show_only_agencies"),
		NotValidAgencies:      viper.GetStringSlice("filters.not_valid_agencies"),
	}
}
