package fleet

import z "github.com/Oudwins/zog"

type ShipMaster struct {
	// id
	Id int32

	// 図鑑No
	AlbumId int32

	// 名称
	Name string

	// 読み
	Yomi string

	// 艦種id
	Type int32

	// 艦型 〇〇型的なやつのid
	Type2 int32

	// 装備スロット数
	SlotCount int32

	// 装備搭載数
	Slots []int32

	// 改造段階 0で無印
	Version int32

	// 最終改造状態か否か
	IsFinal bool

	// 未改造時ID
	OriginalId int32

	// デフォルト射程
	Range int32

	// 耐久
	Hp int32

	// ケッコン後耐久
	Hp2 int32

	// 最大耐久(改修の限界)
	MaxHp int32

	// 火力
	Fire int32

	// 雷装
	Torpedo int32

	// 火力+雷装
	Night int32

	// 対空
	AntiAir int32

	// 装甲
	Armor int32

	// 運初期値
	Luck int32

	// 運最大値
	MaxLuck int32

	// 索敵初期値
	MinScout int32

	// 索敵最大値
	MaxScout int32

	// 対潜初期値
	MinAsw int32

	// 対潜最大値
	MaxAsw int32

	// 回避初期値
	MinAvoid int32

	// 回避最大値
	MaxAvoid int32

	// 速力
	Speed int32

	// 改装直前のid
	BeforeId int32

	// 次改装Lv
	NextLv int32

	// 母港マスタソート順
	Sort int32

	// 搭載燃料
	Fuel int32

	// 搭載弾薬
	Ammo int32

	// 改装設計図消費数
	Blueprints int32

	// 戦闘詳報消費数
	ActionReports int32

	// カタパルト消費数
	Catapults int32

	// 空母か(CV CVB CVL)
	IsCV bool

	// 戦艦か(BB FBB BBV BBB)
	IsBB bool
}

var ShipMasterSchema = z.Struct(z.Shape{
	"id":            z.Int32(),
	"albumId":       z.Int32(),
	"name":          z.String(),
	"yomi":          z.String(),
	"type":          z.Int32(),
	"type2":         z.Int32(),
	"slotCount":     z.Int32(),
	"slots":         z.Slice(z.Int32()),
	"version":       z.Int32(),
	"isFinal":       z.Bool(),
	"originalId":    z.Int32(),
	"range":         z.Int32(),
	"hp":            z.Int32(),
	"hp2":           z.Int32(),
	"maxHp":         z.Int32(),
	"fire":          z.Int32(),
	"torpedo":       z.Int32(),
	"night":         z.Int32(),
	"antiAir":       z.Int32(),
	"armor":         z.Int32(),
	"luck":          z.Int32(),
	"maxLuck":       z.Int32(),
	"minScout":      z.Int32(),
	"maxScout":      z.Int32(),
	"minAsw":        z.Int32(),
	"maxAsw":        z.Int32(),
	"minAvoid":      z.Int32(),
	"maxAvoid":      z.Int32(),
	"speed":         z.Int32(),
	"beforeId":      z.Int32(),
	"nextLv":        z.Int32(),
	"sort":          z.Int32(),
	"fuel":          z.Int32(),
	"ammo":          z.Int32(),
	"blueprints":    z.Int32(),
	"actionReports": z.Int32(),
	"catapults":     z.Int32(),
	"isCV":          z.Bool(),
	"isBB":          z.Bool(),
})
