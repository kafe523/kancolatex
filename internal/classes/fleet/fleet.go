package fleet

import (
	"math"
	"slices"

	z "github.com/Oudwins/zog"
	"github.com/kafe523/kancolatex/internal/classes/aerialcombat"
	"github.com/kafe523/kancolatex/internal/classes/aircalcresult"
	"github.com/kafe523/kancolatex/internal/classes/consts/formation"
	"github.com/kafe523/kancolatex/internal/classes/consts/supporttype"
	"github.com/kafe523/kancolatex/internal/classes/item"
)

type Fleet struct {
	// この艦隊を構成する艦娘一覧
	Ships []Ship

	// 陣形
	Formation formation.Formation

	// 連合艦隊フラグ
	IsUnion bool

	// 艦隊制空値
	FullAirPower int32

	// 支援タイプ Const.SUPPORT_TYPE参照
	SupportTypes []supporttype.SupportType

	// 航空支援制空値
	SupportAirPower int32

	// 対潜支援制空値
	SupportAswAirPower int32

	// 対潜支援実行可能であるかどうか
	EnabledAswSupport bool

	// 輸送量
	Tp int32

	// 輸送量(戦車)主力
	MainTP2 int32

	// 輸送量(戦車)随伴
	EscortTP2 int32

	// 輸送量(戦車)
	Tp2 int32

	// 輸送量(戦車)
	Tp3 int32

	// 航空戦が可能な機体ありなし(主力 & 随伴)
	HasPlane bool

	// 航空戦が可能な機体ありなし(主力のみ)
	HasMainPlane bool

	// 噴式機ありなし
	HasJet bool

	// 艦隊防空値(ブラウザ版表示値)
	FleetAntiAir float64

	// この艦隊の全艦載機装備 搭載数1以上 計算用
	AllPlanes []item.Item

	// stage2 撃墜テーブル
	ShootDownList []aerialcombat.ShootDownInfo

	// stage2 撃墜テーブル(空襲)
	ShootDownListAirRaid []aerialcombat.ShootDownInfo

	// stage2 撃墜テーブル(どちらかが連合)
	UnionShootDownList []aerialcombat.ShootDownInfo

	// stage2 撃墜テーブル(どちらかが連合かつ空襲)
	UnionShootDownListAirRaid []aerialcombat.ShootDownInfo

	// 発動可能対空CI全種
	AllAntiAirCutIn []aerialcombat.AntiAirCutIn

	// 艦隊索敵補正
	FleetRosCorr int32

	// 夜偵発動率
	NightContactRate int32

	// 艦隊区分(速力)
	FleetSpeed string

	// 現在搭載数における制空値 計算用
	AirPower int32

	// 計算結果の各制空状態の割合
	Results []aircalcresult.AirCalcResult

	// 表示戦闘の各制空状態の割合
	MainResult aircalcresult.AirCalcResult

	// この艦隊の随伴艦隊制空値 連合艦隊専用
	EscortAirPower int32
}

var FleetSchema = z.Struct(z.Shape{
	"ships":                     z.Slice(ShipSchema),
	"formation":                 z.IntLike[formation.Formation]().OneOf(formation.FormationValues()),
	"isUnion":                   z.Bool(),
	"fullAirPower":              z.Int32(),
	"supportTypes":              z.Slice(z.IntLike[supporttype.SupportType]().OneOf(supporttype.SupportTypeValues())),
	"supportAirPower":           z.Int32(),
	"supportAswAirPower":        z.Int32(),
	"enabledAswSupport":         z.Bool(),
	"tp":                        z.Int32(),
	"mainTP2":                   z.Int32(),
	"escortTP2":                 z.Int32(),
	"tp2":                       z.Int32(),
	"tp3":                       z.Int32(),
	"hasPlane":                  z.Bool(),
	"hasMainPlane":              z.Bool(),
	"hasJet":                    z.Bool(),
	"fleetAntiAir":              z.Float64(),
	"allPlanes":                 z.Slice(item.ItemSchema),
	"shootDownList":             z.Slice(aerialcombat.ShootDownInfoSchema),
	"shootDownListAirRaid":      z.Slice(aerialcombat.ShootDownInfoSchema),
	"unionShootDownList":        z.Slice(aerialcombat.ShootDownInfoSchema),
	"unionShootDownListAirRaid": z.Slice(aerialcombat.ShootDownInfoSchema),
	"allAntiAirCutIn":           z.Slice(aerialcombat.AntiAirCutInSchema),
	"fleetRosCorr":              z.Int32(),
	"nightContactRate":          z.Int32(),
	"fleetSpeed":                z.String(),
	"airPower":                  z.Int32(),
	"results":                   z.Slice(aircalcresult.AirCalcResultSchema).Default([]aircalcresult.AirCalcResult{aircalcresult.AirCalcResult{}}),
	"mainResult":                aircalcresult.AirCalcResultSchema,
	"escortAirPower":            z.Int32().Default(0),
})

func GetScoutScore(argShips []Ship, admiralLevel, cCount int32) []float64 {
	// Σ(√艦娘の素の索敵値) + Σ{(装備の素の索敵値 + 改修係数×√★)×装備係数}×分岐点係数 - ⌈艦隊司令部Lv.×司令部補正係数⌉ + 2×(6 - 分岐点に到達した際の隻数)
	var scoutScore []float64
	var block3 float64 = float64(admiralLevel) * 0.4
	var ships []Ship = slices.Collect(func(yield func(Ship) bool) {
		for _, s := range argShips {
			if !s.IsActive || s.IsEmpty {
				continue
			}

			if !yield(s) {
				return
			}
		}
	})

	// 分岐点係数
	for i := int32(1); i <= cCount; i++ {
		var block1 float64
		var block2 float64

		for j := 0; j < len(ships); j++ {
			ship := ships[j]
			// Σ(√艦娘の素の索敵値 + 装備ボーナス)
			block1 += math.Sqrt(float64(ship.Scout) + float64(ship.ItemBonusStatus.Scout))
			// Σ{(装備の素の索敵値 + 改修係数×√★)×装備係数}×分岐点係数
			block2 += ship.ItemsScout * float64(i)
		}
		scoutScore = append(scoutScore, block1+block2-math.Ceil(block3)+float64(2*(6-len(ships))))
	}

	return scoutScore
}

func (this *Fleet) GetUnionScoutScore(admiralLevel, cCount int32) []float64 {
	var mainShips []Ship = slices.Collect(func(yield func(Ship) bool) {
		for _, v := range this.Ships {
			if !v.IsEscort {
				if !yield(v) {
					return
				}
			}
		}
	})
	var mainScouts []float64 = GetScoutScore(mainShips, admiralLevel, cCount)

	var escortShips []Ship = slices.Collect(func(yield func(Ship) bool) {
		for _, v := range this.Ships {
			if v.IsEscort {
				if !yield(v) {
					return
				}
			}
		}

	})

	var escortScouts []float64 = GetScoutScore(escortShips, admiralLevel, cCount)

	var result []float64
	for i := 1; i <= int(cCount); i++ {
		result = append(result, mainScouts[i]+escortScouts[i])
	}

	return result
}
