package aircalcresult

import (
	z "github.com/Oudwins/zog"
	"github.com/kafe523/kancolatex/internal/classes/consts/airstate"
)

type AirCalcResult struct {
	// 制空状態（複数ある場合は一番確率が高いやつ）
	AirState struct {
		Text  string
		Value airstate.AirState
		Color string
	}

	// 制空状態テキスト
	AirStateText string

	// 制空ゲージバーの長さ
	AirStateBarWidth int32

	// 各制空状態の割合
	Rates []int32

	// 各制空状態の割合(支援)
	SupportRates []int32

	// 計算処理ループ中の表示戦闘での制空値記録用
	LoopSumAirPower int32

	// 計算処理ループ中の表示戦闘での敵制空値記録用
	LoopSumEnemyAirPower int32

	// 計算処理ループ中の表示戦闘での敵制空値記録用
	LoopSumEnemySupportAirPower int32

	// 表示戦闘での平均制空値
	AvgAirPower int32

	// 表示戦闘での敵制空値平均
	AvgEnemyAirPower int32

	// 表示戦闘での敵制空値平均
	AvgEnemySupportAirPower int32

	// 撃ち落とされた艦載機合計 ボーキ算出に必要
	AvgDownSlot int32

	// 使われた鋼材合計
	AvgUsedSteels int32

	// 敵制空値が不明
	IsUnknownEnemyAirPower bool
}

var AirCalcResultSchema = z.Struct(z.Shape{
	"airState": z.Struct(z.Shape{
		"text":  z.String().Default("不発"),
		"value": z.IntLike[airstate.AirState]().OneOf(airstate.AirStateValues()).Default(airstate.NONE),
		"color": z.String().Default("secondary"),
	}),
	"airStateText":                z.String().Default(""),
	"airStateBarWidth":            z.Int32().Default(0),
	"rates":                       z.Slice(z.Int32()).Len(6).Default([]int32{0, 0, 0, 0, 0, 0}),
	"supportRates":                z.Slice(z.Int32()).Len(6).Default([]int32{0, 0, 0, 0, 0, 0}),
	"loopSumAirPower":             z.Int32().Default(0),
	"loopSumEnemyAirPower":        z.Int32().Default(0),
	"loopSumEnemySupportAirPower": z.Int32().Default(0),
	"avgAirPower":                 z.Int32().Default(0),
	"avgEnemyAirPower":            z.Int32().Default(0),
	"avgEnemySupportAirPower":     z.Int32().Default(0),
	"avgDownSlot":                 z.Int32().Default(0),
	"avgUsedSteels":               z.Int32().Default(0),
	"isUnknownEnemyAirPower":      z.Bool().Default(false),
})
