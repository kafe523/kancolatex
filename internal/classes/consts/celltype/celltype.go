package celltype

//go:generate go run github.com/dmarkham/enumer -json -text -typederrors -type=CellType
type CellType int32

const (
	NORMAL              CellType = 1
	GRAND               CellType = 2
	AIR_RAID            CellType = 3
	NIGHT               CellType = 4
	HIGH_AIR_RAID       CellType = 5
	AERIAL_COMBAT       CellType = 6
	SUPER_HIGH_AIR_RAID CellType = 7
	AIR_SUPPORTED_ASW   CellType = 8
	RADAR               CellType = 9
)
