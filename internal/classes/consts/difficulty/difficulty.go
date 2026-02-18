package difficulty

//go:generate go run github.com/dmarkham/enumer -json -text -typederrors -type=Difficulty
type Difficulty int32

const (
	HARD   Difficulty = 0
	MEDIUM Difficulty = 1
	EASY   Difficulty = 2
	CASUAL Difficulty = 3
)
