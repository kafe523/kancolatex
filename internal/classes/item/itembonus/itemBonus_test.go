package itembonus

import (
	"bytes"
	"testing"

	z "github.com/Oudwins/zog"
	"github.com/Oudwins/zog/parsers/zjson"
	"github.com/stretchr/testify/assert"
)

var testcaseItemBonusStatusSchema = map[string]func(*testing.T){
	"empty": func(t *testing.T) {
		t.Parallel()

		var ibs ItemBonusStatus

		raw := "{}"

		errList := ItemBonusStatusSchema.Parse(zjson.Decode(bytes.NewReader([]byte(raw))), &ibs)
		if errList != nil {
			t.Fatal(z.Issues.Prettify(errList))
		}

		assert.Equal(t, int32(0), ibs.FirePower)
		assert.Equal(t, int32(0), ibs.Torpedo)
		assert.Equal(t, int32(0), ibs.AntiAir)
		assert.Equal(t, int32(0), ibs.Armor)
		assert.Equal(t, int32(0), ibs.Asw)
		assert.Equal(t, int32(0), ibs.Scout)
		assert.Equal(t, int32(0), ibs.Avoid)
		assert.Equal(t, int32(0), ibs.Accuracy)
		assert.Equal(t, int32(0), ibs.Bomber)
		assert.Equal(t, int32(0), ibs.Range)
		assert.Equal(t, int32(0), ibs.FromTypeId)
	},
	"all": func(t *testing.T) {
		t.Parallel()

		var ibs ItemBonusStatus

		raw := `{
			"firePower":  1,
			"torpedo":    2,
			"antiAir":    3,
			"armor":      4,
			"asw":        5,
			"scout":      6,
			"avoid":      7,
			"accuracy":   8,
			"bomber":     9,
			"range":      8,
			"fromTypeId": 7
		}`

		errList := ItemBonusStatusSchema.Parse(zjson.Decode(bytes.NewReader([]byte(raw))), &ibs)
		if errList != nil {
			t.Fatal(z.Issues.Prettify(errList))
		}

		assert.Equal(t, int32(1), ibs.FirePower)
		assert.Equal(t, int32(2), ibs.Torpedo)
		assert.Equal(t, int32(3), ibs.AntiAir)
		assert.Equal(t, int32(4), ibs.Armor)
		assert.Equal(t, int32(5), ibs.Asw)
		assert.Equal(t, int32(6), ibs.Scout)
		assert.Equal(t, int32(7), ibs.Avoid)
		assert.Equal(t, int32(8), ibs.Accuracy)
		assert.Equal(t, int32(9), ibs.Bomber)
		assert.Equal(t, int32(8), ibs.Range)
		assert.Equal(t, int32(7), ibs.FromTypeId)
	},
}

func TestItemBonusStatusSchema(t *testing.T) {
	for tn, tc := range testcaseItemBonusStatusSchema {
		t.Run(tn, tc)
	}
}
