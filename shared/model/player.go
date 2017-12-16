package model

import (
	"time"
)

type Player struct {
	Id           int64  `orm:"auto;index"`
	Identifier   string `json:"steamid"`
	Name         string `json:"personaname"`
	CreationTime time.Time
	LastUpdated  time.Time
	Avatar       string `json:"avatar"`
	AvatarMedium string `json:"avatarmedium"`
	AvatarFull   string `json:"avatarfull"`
	CountryCode  string `json:"loccountrycode"`
	GameCount    int
}

func (Player) TableName() string {
	return "player"
}
