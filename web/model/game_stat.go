package model

import (
	sharedModel "github.com/cubeee/steamtracker/shared/model"
)

type GameStatistic struct {
	MinutesPlayed int64
	Game          sharedModel.Game
}
