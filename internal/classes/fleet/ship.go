package fleet

import (
	z "github.com/Oudwins/zog"
	"github.com/kafe523/kancolatex/internal/classes/aerialcombat"
	"github.com/kafe523/kancolatex/internal/classes/item"
	"github.com/kafe523/kancolatex/internal/classes/item/itembonus"
)

type ShipDisplayStatus struct {
	HP        int32
	FirePower int32
	Armor     int32
	Torpedo   int32
	Avoid     int32
	AntiAir   int32
	Asw       int32
	LoS       int32
	Luck      int32
	Range     int32
	Accuracy  int32
	Bomber    int32
}

var ShipDisplayStatusSchema = z.Struct(z.Shape{
	"HP":        z.Int32(),
	"firePower": z.Int32(),
	"armor":     z.Int32(),
	"torpedo":   z.Int32(),
	"avoid":     z.Int32(),
	"antiAir":   z.Int32(),
	"asw":       z.Int32(),
	"LoS":       z.Int32(),
	"luck":      z.Int32(),
	"range":     z.Int32(),
	"accuracy":  z.Int32(),
	"bomber":    z.Int32(),
})

type Ship struct {
	// 艦娘マスタ情報
	Data ShipMaster

	// 装備一覧
	Items []item.Item

	// 補強増設枠
	ExItem item.Item

	// 練度
	Level int32

	// 表示ステータス (装備 + 装備ボーナス込み)
	DisplayStatus ShipDisplayStatus

	// 装備ボーナス合計 まとめ
	ItemBonusStatus itembonus.ItemBonusStatus

	// 装備フィットボーナスすべて
	ItemBonuses []itembonus.ItemBonusStatus

	// 耐久
	Hp int32

	// 計算で適用する運
	Luck int32

	// 砲戦火力基礎値(連合とかで変わるので、とりあえずの基本値)
	BaseDayBattleFirePower int32

	// 支援射撃火力
	SupportFirePower int32

	// 夜戦火力
	NightBattleFirePower int32

	// 計算で適用する対空
	AntiAir int32

	// 計算で適用する装甲
	ActualArmor int32

	// 素の索敵値
	Scout int32

	// 全装備による索敵値
	ItemsScout float64

	// 回避値
	Avoid int32

	// 通常命中
	Accuracy int32

	// 支援命中
	SupportAccuracy int32

	// 装備「なし」対潜値
	Asw int32

	// 改修分対潜値
	ImproveAsw int32

	// 装備による対潜上昇値
	ItemAsw int32

	// 先制対潜可
	EnabledTSBK bool

	// 輸送量
	Tp int32

	// 輸送量(戦車)
	Tp2 int32

	// 輸送量(戦車)
	Tp3 int32

	// 速力
	Speed int32

	// 消費燃料
	Fuel int32

	// 消費弾薬
	Ammo int32

	// 出撃海域
	Area int32

	// 所持数登録ユニークid
	UniqueId int32

	// 噴進弾幕率
	HunshinRate int32

	// 有効無効
	IsActive bool

	// 装備が空 計算対象として数えない
	IsEmpty bool

	// 防空ボーナス
	AntiAirBonus int32

	// 随伴艦フラグ
	IsEscort bool

	// 噴式機ありなし
	HasJet bool

	// 制空値(搭載数満タン)
	FullAirPower int32

	// 航空支援制空値
	SupportAirPower int32

	// 対潜支援制空値
	SupportAswAirPower int32

	// 対潜支援参加可能
	EnabledASWSupport bool

	// 発動可能対空CI
	AntiAirCutIn []aerialcombat.AntiAirCutIn

	// 特殊高角砲所持数
	SpecialKokakuCount int32

	// 高角砲所持数
	KokakuCount int32

	// 特殊機銃所持数
	SpecialKijuCount int32

	// 機銃所持数
	KijuCount int32

	// 対空電探所持数
	AntiAirRadarCount int32

	// 水上電探所持数
	SurfaceRadarCount int32

	// 高射装置所持数
	KoshaCount int32

	// 水偵/水爆の裝備索敵值 * int(sqrt(水偵/水爆の機數)
	SumSPRos int32

	// 夜偵発動率
	NightContactRate int32

	// 夜襲 発動可能判定
	EnabledAircraftNightAttack bool

	// 夜襲 熟練甲板要員火力ボーナス
	NightAttackCrewFireBonus int32

	// 夜襲 熟練甲板要員爆装ボーナス
	NightAttackCrewBomberBonus int32

	// 補強増設空いてる？
	ReleaseExpand bool

	// 所持なしなのに配備されてる
	NoStock bool

	// 装備置き場用プロパティ
	IsTray bool

	// 海色リボン 白たすき
	SpEffectItemId int32

	// 先制対潜不足対潜値
	MissingAsw int32

	// 先制対潜可になるまでの残りLevel
	NeedTSBKLevel int32

	// 固定撃墜 画面表示用
	FixDown int32

	// 割合撃墜 画面表示用
	RateDown float64

	// 棒立ち率
	AllPlaneDeathRate int32
}

var ShipSchema = z.Struct(z.Shape{
	"data":                       ShipMasterSchema,
	"items":                      z.Slice(item.ItemSchema),
	"exItem":                     item.ItemSchema,
	"level":                      z.Int32(),
	"displayStatus":              ShipDisplayStatusSchema,
	"itemBonusStatus":            itembonus.ItemBonusStatusSchema,
	"itemBonuses":                z.Slice(itembonus.ItemBonusStatusSchema),
	"hp":                         z.Int32(),
	"luck":                       z.Int32(),
	"baseDayBattleFirePower":     z.Int32(),
	"supportFirePower":           z.Int32(),
	"nightBattleFirePower":       z.Int32(),
	"antiAir":                    z.Int32(),
	"actualArmor":                z.Int32(),
	"scout":                      z.Int32(),
	"itemsScout":                 z.Float64(),
	"avoid":                      z.Int32(),
	"accuracy":                   z.Int32(),
	"supportAccuracy":            z.Int32(),
	"asw":                        z.Int32(),
	"improveAsw":                 z.Int32(),
	"itemAsw":                    z.Int32(),
	"enabledTSBK":                z.Bool(),
	"tp":                         z.Int32(),
	"tp2":                        z.Int32(),
	"tp3":                        z.Int32(),
	"speed":                      z.Int32(),
	"fuel":                       z.Int32(),
	"ammo":                       z.Int32(),
	"area":                       z.Int32(),
	"uniqueId":                   z.Int32(),
	"hunshinRate":                z.Int32(),
	"isActive":                   z.Bool(),
	"isEmpty":                    z.Bool(),
	"antiAirBonus":               z.Int32(),
	"isEscort":                   z.Bool(),
	"hasJet":                     z.Bool(),
	"fullAirPower":               z.Int32(),
	"supportAirPower":            z.Int32(),
	"supportAswAirPower":         z.Int32(),
	"enabledASWSupport":          z.Bool(),
	"antiAirCutIn":               z.Slice(aerialcombat.AntiAirCutInSchema),
	"specialKokakuCount":         z.Int32(),
	"kokakuCount":                z.Int32(),
	"specialKijuCount":           z.Int32(),
	"kijuCount":                  z.Int32(),
	"antiAirRadarCount":          z.Int32(),
	"surfaceRadarCount":          z.Int32(),
	"koshaCount":                 z.Int32(),
	"sumSPRos":                   z.Int32(),
	"nightContactRate":           z.Int32(),
	"enabledAircraftNightAttack": z.Bool(),
	"nightAttackCrewFireBonus":   z.Int32(),
	"nightAttackCrewBomberBonus": z.Int32(),
	"releaseExpand":              z.Bool(),
	"noStock":                    z.Bool(),
	"isTray":                     z.Bool(),
	"spEffectItemId":             z.Int32(),
	"missingAsw":                 z.Int32().Default(0),
	"needTSBKLevel":              z.Int32().Default(0),
	"fixDown":                    z.Int32().Default(0),
	"rateDown":                   z.Float64().Default(0),
	"allPlaneDeathRate":          z.Int32().Default(0),
})
