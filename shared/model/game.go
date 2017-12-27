package model

type Game struct {
	AppId    int64  `json:"appid" gorm:"primary_key:pk_game_id,column:app_id"`
	Name     string `json:"name" gorm:"column:name"`
	Playtime int64  `json:"playtime_forever" gorm:"-"`
	Icon     string `json:"img_icon_url" gorm:"column:icon_url"`
	Logo     string `json:"img_logo_url" gorm:"column:logo_url"`
}

func (Game) TableName() string {
	return "game"
}
