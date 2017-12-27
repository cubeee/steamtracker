package service

type OwnedGamesService struct {
	Service
}

func (OwnedGamesService) Interface() string {
	return "IPlayerService"
}

func (OwnedGamesService) CallName() string {
	return "GetOwnedGames"
}

func (OwnedGamesService) Version() string {
	return "v0001"
}

func (OwnedGamesService) Parameters(params []string) string {
	return "&steamid=" + params[0] + "&format=json&include_appinfo=1&include_played_free_games=1"
}
