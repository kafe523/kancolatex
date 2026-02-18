package calcmanager

import (
	z "github.com/Oudwins/zog"
	"github.com/kafe523/kancolatex/internal/classes/fleet"
)

type CalcManager struct {
	// 防空計算モードか否か 基地航空隊欄で制御
	IsDefense bool

	AirbaseInfo map[string]any // AirbaseInfo

	BattleInfo map[string]any // BattleInfo

	FleetInfo fleet.FleetInfo

	MainBattle int32
}

var CalcManagerSchema = z.Struct(z.Shape{
	"isDefense": z.Bool(),
	// "airbaseInfo": AirbaseInfoSchema,
	// "battleInfo": BattleInfoSchema,
	"fleetInfo":  fleet.FleetInfoSchema,
	"mainBattle": z.Int32(),
})
