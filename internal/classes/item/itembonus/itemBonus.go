package itembonus

import z "github.com/Oudwins/zog"

type ItemBonusStatus struct {
	FirePower  int32
	Torpedo    int32
	AntiAir    int32
	Armor      int32
	Asw        int32
	Scout      int32
	Avoid      int32
	Accuracy   int32
	Bomber     int32
	Range      int32
	FromTypeId int32
}

var ItemBonusStatusSchema = z.Struct(z.Shape{
	"firePower":  z.Int32(),
	"torpedo":    z.Int32(),
	"antiAir":    z.Int32(),
	"armor":      z.Int32(),
	"asw":        z.Int32(),
	"scout":      z.Int32(),
	"avoid":      z.Int32(),
	"accuracy":   z.Int32(),
	"bomber":     z.Int32(),
	"range":      z.Int32(),
	"fromTypeId": z.Int32(),
})

type Bonus struct {
	// 上昇するステータス
	Bonus ItemBonusStatus
	// 艦型(未改造id)
	ShipBase []int32
	// 艦型(白露型 など)
	ShipClass []int32
	// 国籍(艦型と同じ)
	ShipCountry []int32
	// 艦種(駆逐 軽巡...)
	ShipType []int32
	// 艦種娘固有id
	ShipId []int32
	// 必須の組み合わせ装備種別
	RequiresType []int32
	// 必須の組み合わせ水上電探有無
	RequiresSR int32
	// 必須の組み合わせ対空電探有無
	RequiresAR int32
	// 必須の組み合わせ命中+8以上電探有無
	RequiresAccR int32
	// 必須の組み合わせ装備id
	RequiresId []int32
	// 必須の組み合わせ装備id個数 requiresIdに一致する装備の最低限必要数
	RequiresIdNum int32
	// 必須の組み合わせ装備idの改修値 requiresIdに一致する装備の最低限必要改修★数
	RequiresIdLevel int32
	// 必須の組み合わせ装備id
	RequiresId2 []int32
	// 必須の組み合わせ装備idの改修値 requiresIdに一致する装備の最低限必要改修★数
	RequiresIdLevel2 int32
	// 該当装備の最低限必要改修★数
	Remodel int32
	/*
	 * これがなければ累積可能
	 * 存在する場合、上記の条件を全て満たしている装備の最低限の個数 4連装酸素魚雷辺りがわかりやすい
	 * 2024/06現在、この値が「ない≒累積可能」場合は、require系の条件は一切存在しない
	 */
	Num int32
}

type Bonuses struct {
	Types   []int32
	Ids     []int32
	Bonuses []Bonus
}
