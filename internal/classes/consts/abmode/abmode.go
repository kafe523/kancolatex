package abmode

// 基地札種類
//
//go:generate go run github.com/dmarkham/enumer -json -text -typederrors -type=ABMode
type ABMode int32

const (
	WAIT    ABMode = 0
	BATTLE  ABMode = 1
	DEFENSE ABMode = 2
)
