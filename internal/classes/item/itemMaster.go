package item

import z "github.com/Oudwins/zog"

type ItemMaster struct {
	// id
	Id int32

	// 種別
	ApiTypeId int32

	// アイコン用種別
	IconTypeId int32

	// 名称
	Name string

	// 略称
	Abbr string

	// 火力
	Fire int32

	// 雷装
	Torpedo int32

	// 爆装
	Bomber int32

	// 対空
	AntiAir int32

	// 装甲
	Armor int32

	// 対潜
	Asw int32

	// 対爆
	AntiBomber int32

	// 命中
	Accuracy int32

	// 迎撃
	Interception int32

	// 回避
	Avoid int32

	// 索敵
	Scout int32

	// 射程
	Range int32

	// 半径
	Radius int32

	// コスト
	Cost int32

	// 改修可否
	CanRemodel bool

	// 射撃回避id
	AvoidId int32

	// 熟練度成長定数
	Grow int32

	// 出撃対空
	SortieAntiAir float64

	// 防空対空
	DefenseAntiAir int32

	// 特殊機銃とか高角砲とかそういうの 一覧で緑色になる
	IsSpecial bool

	// 航空機フラグ
	IsPlane bool

	// 戦闘機フラグ
	IsFighter bool

	// 攻撃機フラグ
	IsAttacker bool

	// 対潜哨戒機フラグ
	IsAswPlane bool

	// 対潜哨戒機爆装ありフラグ 爆装4以上 (一式戦 隼II型改(20戦隊) / 熟練)
	IsAswBomber1 bool

	// 対潜哨戒機爆装ありフラグ2 爆装1以上 爆装4未満 (三式指揮連絡機改二)
	IsAswBomber2 bool

	// オートジャイロフラグ
	IsAutoGyro bool

	// 基地攻撃機フラグ
	IsABAttacker bool

	// 噴式機フラグ
	IsJet bool

	// 重噴式フラグ (Ho229 等)
	IsHeavyJet bool

	// カテゴリ(爆戦)フラグ
	IsBakusen bool

	// ロケット戦闘機フラグ
	IsRocket bool

	// 偵察機フラグ
	IsRecon bool

	// 大型陸上機フラグ
	IsShinzan bool

	// 対地可能フラグ
	EnabledAttackLandBase bool

	// 爆雷(狭義)フラグ
	IsStrictDepthCharge bool

	// 雷装による攻撃を行うかどうか => 現行では火力計算で0.8倍1.5倍が発生するのに使う
	IsTorpedoAttacker bool

	// 夜間航空機
	IsNightAircraftItem bool

	// 水上機かどうか
	IsSPPlane bool

	// 潜水か後期型魚雷かどうか
	IsLateModelTorpedo bool

	// 敵装備かどうか
	IsEnemyItem bool

	// 特効情報
	Bonuses []struct {
		Key        string
		Text       []string
		IsOnlyAB   bool
		IsOnlyShip bool
	}

	// 基地配備時の最大搭載数
	AirbaseMaxSlot int32
}

var ItemMasterSchema = z.Struct(z.Shape{
	"id":                    z.Int32(),
	"apiTypeId":             z.Int32(),
	"iconTypeId":            z.Int32(),
	"name":                  z.String(),
	"abbr":                  z.String(),
	"fire":                  z.Int32(),
	"torpedo":               z.Int32(),
	"bomber":                z.Int32(),
	"antiAir":               z.Int32(),
	"armor":                 z.Int32(),
	"asw":                   z.Int32(),
	"antiBomber":            z.Int32(),
	"accuracy":              z.Int32(),
	"interception":          z.Int32(),
	"avoid":                 z.Int32(),
	"scout":                 z.Int32(),
	"range":                 z.Int32(),
	"radius":                z.Int32(),
	"cost":                  z.Int32(),
	"canRemodel":            z.Bool(),
	"avoidId":               z.Int32(),
	"grow":                  z.Int32(),
	"sortieAntiAir":         z.Float64(),
	"defenseAntiAir":        z.Int32(),
	"isSpecial":             z.Bool(),
	"isPlane":               z.Bool(),
	"isFighter":             z.Bool(),
	"isAttacker":            z.Bool(),
	"isAswPlane":            z.Bool(),
	"isAswBomber1":          z.Bool(),
	"isAswBomber2":          z.Bool(),
	"isAutoGyro":            z.Bool(),
	"isABAttacker":          z.Bool(),
	"isJet":                 z.Bool(),
	"isHeavyJet":            z.Bool(),
	"isBakusen":             z.Bool(),
	"isRocket":              z.Bool(),
	"isRecon":               z.Bool(),
	"isShinzan":             z.Bool(),
	"enabledAttackLandBase": z.Bool(),
	"isStrictDepthCharge":   z.Bool(),
	"isTorpedoAttacker":     z.Bool(),
	"isNightAircraftItem":   z.Bool(),
	"isSPPlane":             z.Bool(),
	"isLateModelTorpedo":    z.Bool(),
	"isEnemyItem":           z.Bool(),
	"bonuses":               z.Slice(z.Struct(z.Shape{"key": z.String(), "text": z.Slice(z.String()), "isOnlyAB": z.Bool(), "isOnlyShip": z.Bool()})),
	"airbaseMaxSlot":        z.Int32(),
})
