package supporttype

//go:generate go run github.com/dmarkham/enumer -json -text -typederrors -type=SupportType
type SupportType int32

const (
	SHELLING           SupportType = 0
	AIRSTRIKE          SupportType = 1
	ANTI_SUBMARINE     SupportType = 2
	LONG_RANGE_TORPEDO SupportType = 3
	NOT_FOUND_DD       SupportType = 4 //?
	NONE               SupportType = 5
)
