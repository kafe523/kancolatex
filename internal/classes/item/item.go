package item

import (
	z "github.com/Oudwins/zog"
)

type Item struct {
	Data ItemMaster

	// 熟練度 0~120
	Level int32

	// 改修値
	Remodel int32

	// 搭載数
	FullSlot int32

	// 改修値による火力値増分(昼)
	BonusFire float64

	// 改修値による火力値増分(夜)
	BonusNightFire float64

	// 改修値による火力値増分(遠征)
	BonusExpeditionFire float64

	// 改修値による雷装値増分
	BonusTorpedo float64

	// 改修値による爆装値増分
	BonusBomber float64

	// 改修値による対空値増分
	BonusAntiAir float64

	// 改修値による対空値増分(遠征)
	BonusExpeditionAntiAir float64

	// 改修値による装甲値増分
	BonusArmor float64

	// 改修値による命中値増分
	BonusAccuracy float64

	// 改修値による対潜値増分
	BonusAsw float64

	// 改修値による対空値増分(遠征)
	BonusExpeditionAsw float64

	// 改修値による索敵値増分
	BonusScout float64

	// 改修値による対空値増分(遠征)
	BonusExpeditionScout float64

	// 熟練度による制空値増分
	BonusAirPower float64

	// 装備索敵 計算で利用 (装備の素の索敵値 + 改修係数 × √★) × 装備係数
	ItemScout float64

	// 装備加重対空値
	AntiAirWeight float64

	// 装備防空ボーナス 誤差考慮のため100倍された値
	AntiAirBonus float64

	// 昼砲撃戦火力(通常)
	DayBattleFirePower float64

	// 昼砲撃戦火力(空母)
	AircraftDayBattleFirePower float64

	// 夜戦火力 装備ボーナス分は後付け(艦娘が必要なため)
	NightBattleFirePower float64

	// 改修効果込み実火力値 装備ボーナス分は後付け(艦娘が必要なため)
	ActualFire float64

	// 制空値計算時に適用される実対空値
	ActualAntiAir float64

	// 改修効果込み実雷装値 装備ボーナス分は後付け(艦娘が必要なため)
	ActualTorpedo float64

	// 改修効果込み実爆装値
	ActualBomber float64

	// 改修効果込み実対潜値 装備ボーナス分は後付け(艦娘が必要なため)
	ActualAsw float64

	// 改修効果込み実装甲値 装備ボーナス分は後付け(艦娘が必要なため)
	ActualArmor float64

	// 改修効果込み実命中値 装備ボーナスは後付け(艦娘が必要なため)
	ActualAccuracy float64

	// 改修効果込み実索敵値 装備ボーナス分は後付け(艦娘が必要なため)
	ActualScout float64

	// 改修効果込み実回避 装備ボーナス分は後付け(艦娘が必要なため)
	ActualAvoid float64

	// 実射程値 装備ボーナス分後付け(艦娘が必要なため)
	ActualRange float64

	// 制空値計算時に適用される実対空値(防空時)
	ActualDefenseAntiAir float64

	// 装備制空値
	FullAirPower int32

	// 装備制空値(防空時)
	FullDefenseAirPower int32

	// 航空支援制空値
	SupportAirPower int32

	// 対潜支援制空値
	SupportAswAirPower int32

	// 触接選択率 [確保時, 優勢時, 劣勢時]
	ContactSelectRates []float64

	// 輸送量
	Tp float64

	// 輸送量(戦車)
	Tp2 float64

	// 輸送量(戦車)
	Tp3 float64

	// 基地1出撃燃料
	Fuel int32

	// 基地1出撃弾薬
	Ammo int32

	// 鋼材
	Steel int32

	// 基地配備時消費ボーキ
	Bauxite int32

	// 偵察機補正 -出撃時
	ReconCorr float64

	// 偵察機補正 -防空時
	ReconCorrDefense float64

	// 最大搭載数から1までの制空値を計算し終えた配列 機体のみ有効 計算用
	// private
	CalculatedAirPower []int32

	// 最大搭載数から1までの防空制空値を計算し終えた配列 機体のみ有効 計算用
	// private
	CalculatedDefenseAirPower []int32

	// 航空戦雷装ボーナス 艦娘インスタンス化時限定変更
	AttackerTorpedoBonus int32

	// 航空戦雷装ボーナス(熟練甲板要員＋航空整備員による) 艦娘インスタンス化時限定変更
	CrewTorpedoBonus int32

	// 航空戦爆装ボーナス(熟練甲板要員＋航空整備員による) 艦娘インスタンス化時限定変更
	CrewBomberBonus int32

	// 現在搭載数における制空値 計算用
	AirPower int32

	// 現在搭載数における制空値 計算用
	DefenseAirPower int32

	// 現在搭載数 計算用
	Slot int32

	// 搭載数推移 表示用
	SlotHistories []int32

	// 戦闘後搭載数 表示用
	SlotResult int32

	// 全滅率 表示用
	DeathRate int32

	// 戦闘後搭載数の最小 夜襲計算用
	MinSlot int32

	// 戦闘後搭載数の最大 夜襲計算用
	MaxSlot int32

	// 第2艦隊搭載機かどうかフラグ 計算用
	IsEscortItem bool

	// 本隊航空戦に参加するかどうかフラグ 計算用
	DisabledItem bool

	// スロット分布のための記録が必要
	NeedRecord bool

	// 補給前搭載数分布
	Dist []int32

	// 棒立ち率特定のための親識別用index => FleetクラスのallPlaneに突っ込む際、装備者がわからなくなるため特定したい
	ParentIndex int32

	// 所持なしなのに配備されてる
	NoStock bool
}

var ItemSchema = z.Struct(z.Shape{
	"data":                       ItemMasterSchema,
	"level":                      z.Int32(),
	"remodel":                    z.Int32(),
	"fullSlot":                   z.Int32(),
	"bonusFire":                  z.Float64(),
	"bonusNightFire":             z.Float64(),
	"bonusExpeditionFire":        z.Float64(),
	"bonusTorpedo":               z.Float64(),
	"bonusBomber":                z.Float64(),
	"bonusAntiAir":               z.Float64(),
	"bonusExpeditionAntiAir":     z.Float64(),
	"bonusArmor":                 z.Float64(),
	"bonusAccuracy":              z.Float64(),
	"bonusAsw":                   z.Float64(),
	"bonusExpeditionAsw":         z.Float64(),
	"bonusScout":                 z.Float64(),
	"bonusExpeditionScout":       z.Float64(),
	"bonusAirPower":              z.Float64(),
	"itemScout":                  z.Float64(),
	"antiAirWeight":              z.Float64(),
	"antiAirBonus":               z.Float64(),
	"dayBattleFirePower":         z.Float64(),
	"aircraftDayBattleFirePower": z.Float64(),
	"nightBattleFirePower":       z.Float64(),
	"actualFire":                 z.Float64(),
	"actualAntiAir":              z.Float64(),
	"actualTorpedo":              z.Float64(),
	"actualBomber":               z.Float64(),
	"actualAsw":                  z.Float64(),
	"actualArmor":                z.Float64(),
	"actualAccuracy":             z.Float64(),
	"actualScout":                z.Float64(),
	"actualAvoid":                z.Float64(),
	"actualRange":                z.Float64(),
	"actualDefenseAntiAir":       z.Float64(),
	"fullAirPower":               z.Int32(),
	"fullDefenseAirPower":        z.Int32(),
	"supportAirPower":            z.Int32(),
	"supportAswAirPower":         z.Int32(),
	"contactSelectRates":         z.Slice(z.Float64()),
	"tp":                         z.Float64(),
	"tp2":                        z.Float64(),
	"tp3":                        z.Float64(),
	"fuel":                       z.Int32(),
	"ammo":                       z.Int32(),
	"steel":                      z.Int32(),
	"bauxite":                    z.Int32(),
	"reconCorr":                  z.Float64(),
	"reconCorrDefense":           z.Float64(),
	"calculatedAirPower":         z.Slice(z.Int32()),
	"calculatedDefenseAirPower":  z.Slice(z.Int32()),
	"attackerTorpedoBonus":       z.Int32().Default(0),
	"crewTorpedoBonus":           z.Int32().Default(0),
	"crewBomberBonus":            z.Int32().Default(0),
	"airPower":                   z.Int32(),
	"defenseAirPower":            z.Int32(),
	"slot":                       z.Int32(),
	"slotHistories":              z.Slice(z.Int32()),
	"slotResult":                 z.Int32().Default(0),
	"deathRate":                  z.Int32().Default(0),
	"minSlot":                    z.Int32().Default(0),
	"maxSlot":                    z.Int32().Default(0),
	"isEscortItem":               z.Bool().Default(false),
	"disabledItem":               z.Bool().Default(false),
	"needRecord":                 z.Bool().Default(false),
	"dist":                       z.Slice(z.Int32()).Default([]int32{}),
	"parentIndex":                z.Int32().Default(-1),
	"noStock":                    z.Bool(),
})
