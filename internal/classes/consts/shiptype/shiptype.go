package shiptype

// 艦種
//
//go:generate go run github.com/dmarkham/enumer -json -text -typederrors -type=ShipType
type ShipType int32

const (
	DE   ShipType = 1
	DD   ShipType = 2
	CL   ShipType = 3
	CLT  ShipType = 4
	CA   ShipType = 5
	CAV  ShipType = 6
	CVL  ShipType = 7
	FBB  ShipType = 8
	BB   ShipType = 9
	BBV  ShipType = 10
	CV   ShipType = 11
	BBB  ShipType = 12
	SS   ShipType = 13
	SSV  ShipType = 14
	AO_2 ShipType = 15
	AV   ShipType = 16
	LHA  ShipType = 17
	CVB  ShipType = 18
	AR   ShipType = 19
	AS   ShipType = 20
	CT   ShipType = 21
	AO   ShipType = 22
)
