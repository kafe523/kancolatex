package aerialcombat

import z "github.com/Oudwins/zog"

type ShootDownStatus struct {
	AntiAirWeightList []int32
	RateDownList      []float64
	FixDownList       []int32
	MinimumDownList   []int32
}

var ShootDownStatusSchema = z.Struct(z.Shape{
	"antiAirWeightList": z.Slice(z.Int32()).Required(),
	"rateDownList":      z.Slice(z.Float64()).Required(),
	"fixDownList":       z.Slice(z.Int32()).Required(),
	"minimumDownList":   z.Slice(z.Int32()).Required(),
})

type ShootDownInfo struct {
	// 対空射撃回避区分毎の撃墜性能(固定 割合 最低保証)セット
	ShootDownStatusList []ShootDownStatus

	// 発動可能ランダム域ボーダー
	Border int32

	// 対空砲火可能艦数
	MaxRange int32
}

var ShootDownInfoSchema = z.Struct(z.Shape{
	"shootDownStatusList": z.Slice(ShootDownStatusSchema).Required(),
	"border":              z.Int32().Required(),
	"maxRange":            z.Int32().Required(),
})
