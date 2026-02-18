package firecap

// 火力キャップ
//go:generate go run github.com/dmarkham/enumer -json -text -typederrors -type=Cap
type Cap int32

const (
	AS      Cap = 170
	BATTLE  Cap = 220
	SUPPORT Cap = 170
	NIGHT   Cap = 360
	LBAS    Cap = 220
)
