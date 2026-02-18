package fleettype

//go:generate go run github.com/dmarkham/enumer -json -text -typederrors -type=FleetType
type FleetType int32

const (
	SINGLE FleetType = 0
	CTF    FleetType = 1
	SFT    FleetType = 2
	TCF    FleetType = 3
)
