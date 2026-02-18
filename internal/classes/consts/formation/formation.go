package formation

//go:generate go run github.com/dmarkham/enumer -json -text -typederrors -type=Formation
type Formation int32

const (
	LINE_AHEAD   Formation = 1
	DOUBLE_LINE  Formation = 2
	DIAMOND      Formation = 3
	ECHELON      Formation = 4
	LINE_ABREAST Formation = 5
	VANGUARD     Formation = 6
	FORMATION1   Formation = 11 //15
	FORMATION2   Formation = 12
	FORMATION3   Formation = 13
	FORMATION4   Formation = 14 //11
)

// Somehow noro6 JTF formation have different number.
