package fleet

import (
	z "github.com/Oudwins/zog"
	"github.com/kafe523/kancolatex/internal/classes/consts/fleettype"
)

type FleetInfo struct {
	// 艦隊一覧
	Fleets []Fleet

	// 連合艦隊？
	IsUnion bool

	// 艦隊タイプ
	FleetType fleettype.FleetType

	// 司令部レベル
	AdmiralLevel int32

	// 計算を行う艦隊番号
	MainFleetIndex int32

	// 連合艦隊 自動生成
	UnionFleet *Fleet

	// 計算済みフラグ
	Calculated bool

	// 履歴に追加しなくてもいいフラグ
	IgnoreHistory bool
}

var FleetInfoSchema = z.Struct(z.Shape{
	"fleets":        z.Slice(FleetSchema),
	"isUnion":       z.Bool(),
	"fleetType":     z.IntLike[fleettype.FleetType]().OneOf(fleettype.FleetTypeValues()),
	"admiralLevel":  z.Int32(),
	"unionFleet":    z.Ptr(FleetSchema),
	"calculated":    z.Bool(),
	"ignoreHistory": z.Bool(),
})
