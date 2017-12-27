package model

import "time"

type GameSnapshot struct {
	Id            int64 `orm:"auto;index"`
	PlayerId      int64 `column:"player_id"`
	GameId        int64 `column:"game_id"`
	MinutesPlayed int64 `column:"minutes_played"`
	Date          time.Time
}

func (GameSnapshot) TableName() string {
	return "game_snapshot"
}

func NewSnapshot(playerId int64, gameId int64, minutesPlayed int64) *GameSnapshot {
	return &GameSnapshot{PlayerId: playerId, GameId: gameId, MinutesPlayed: minutesPlayed, Date: time.Now().UTC()}
}
