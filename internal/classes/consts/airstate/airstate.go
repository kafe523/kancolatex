package airstate

//go:generate go run github.com/dmarkham/enumer -json -text -typederrors -type=AirState
type AirState int32

const (
	KAKUHO   AirState = 0
	YUSEI    AirState = 1
	KINKO    AirState = 2
	RESSEI   AirState = 3
	SOSHITSU AirState = 4
	NONE     AirState = 5
)
