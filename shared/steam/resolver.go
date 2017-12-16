package steam

import (
	"encoding/json"
	"fmt"

	"github.com/cubeee/steamtracker/shared/model"
	"github.com/cubeee/steamtracker/shared/steam/service"
	"github.com/spf13/viper"
	"gopkg.in/resty.v1"
)

type Resolver struct {
}

type ProfileResponse struct {
	Response struct {
		Players []model.Player `json:"players"`
	} `json:"response"`
}

func (resolver Resolver) GetProfile(identifiers []string) ([]model.Player, error) {
	s := service.ProfileService{}
	url := resolver.getUrl(s, identifiers) // TODO: register api key use, maybe check limits

	resp, err := resty.R().Get(url)
	if err != nil {
		return nil, err
	}

	response := ProfileResponse{}
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		return nil, err
	}
	return response.Response.Players, nil
}

func (Resolver) getUrl(service service.Service, params []string) string {
	apiKey := viper.GetString("api_key")
	if apiKey == "" {
		panic("No Steam api key set")
	}
	return fmt.Sprintf("https://api.steampowered.com/%s/%s/%s/?key=%s%s",
		service.Interface(), service.CallName(), service.Version(), apiKey, service.Parameters(params))
}
