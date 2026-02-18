package aerialcombat

import z "github.com/Oudwins/zog"

type AntiAirCutIn struct {
	// カットイン種別id
	Id int32

	// 割合撃墜ボーナス
	RateCorr float64

	// 固定撃墜ボーナスA
	FixCorrA int32

	// 固定撃墜ボーナスB
	FixCorrB int32

	// 発動率
	Rate float64
}

var AntiAirCutInSchema = z.Struct(z.Shape{
	"id":       z.Int32().Required(),
	"rateCorr": z.Float64().Required(),
	"fixCorrA": z.Int32().Required(),
	"fixCorrB": z.Int32().Required(),
	"rate":     z.Float64().Required(),
})
