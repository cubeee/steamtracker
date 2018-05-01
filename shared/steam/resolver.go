package steam

import (
	"encoding/json"
	"fmt"

	"github.com/cubeee/steamtracker/shared/config"
	"github.com/cubeee/steamtracker/shared/model"
	"github.com/cubeee/steamtracker/shared/steam/service"
	"github.com/spf13/viper"
	"gopkg.in/resty.v1"
)

type Resolver struct {
}

type ProfileResponse struct {
	Response struct {
		Players *[]model.Player `json:"players"`
	} `json:"response"`
}

type OwnedGamesResponse struct {
	Response struct {
		GameCount int           `json:"game_count"`
		Games     *[]model.Game `json:"games"`
	} `json:"response"`
}

func (resolver *Resolver) GetProfile(identifiers []string) (*[]model.Player, error) {
	s := service.ProfileService{}
	response := &ProfileResponse{}
	if err := resolver.getResponse(s, response, identifiers); err != nil {
		return nil, err
	}
	return response.Response.Players, nil
}

func (resolver Resolver) GetGames(identifier string) (int, *[]model.Game, error) {
	s := service.OwnedGamesService{}
	response := OwnedGamesResponse{}
	if err := resolver.getResponse(s, &response, []string{identifier}); err != nil {
		return -1, nil, err
	}
	return response.Response.GameCount, response.Response.Games, nil
}

func (resolver Resolver) getResponse(service service.Service, out interface{}, args []string) error {
	url := resolver.getUrl(service, args)
	resp, err := resty.R().Get(url) // TODO: register api key use, maybe check limits
	if err != nil {
		return err
	}

	if err := json.Unmarshal(resp.Body(), out); err != nil {
		return err
	}
	return nil
}

func (Resolver) getUrl(service service.Service, params []string) string {
	apiKey := config.GetString("api_key")
	if apiKey == "" {
		panic("No Steam api key set")
	}
	return fmt.Sprintf("https://api.steampowered.com/%s/%s/%s/?key=%s%s",
		service.Interface(), service.CallName(), service.Version(), apiKey, service.Parameters(params))
}
