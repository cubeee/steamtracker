package service

import "strings"

type ProfileService struct {
	Service
}

func (ProfileService) Interface() string {
	return "ISteamUser"
}

func (ProfileService) CallName() string {
	return "GetPlayerSummaries"
}

func (ProfileService) Version() string {
	return "v0002"
}

func (ProfileService) Parameters(identifiers []string) string {
	joinedIdentifiers := strings.Join(identifiers, ",")
	return "&steamids=" + joinedIdentifiers
}
