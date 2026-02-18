package classes

import "github.com/kafe523/kancolatex/internal/classes/consts"

func GetProfLevel(level int32) int32 {
	for i := int32(len(consts.PROF_LEVEL_BORDER)) - 2; i >= 0; i -= 1 {
		if level >= consts.PROF_LEVEL_BORDER[i] {
			return i
		}
	}

	return 0
}
