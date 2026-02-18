package aerialcombat

import (
	"bytes"
	"testing"

	z "github.com/Oudwins/zog"
	"github.com/Oudwins/zog/parsers/zjson"
	"github.com/stretchr/testify/assert"
)

func TestAntiAirCutIn(t *testing.T) {
	t.Parallel()

	var aaci AntiAirCutIn

	const raw = `{"id":5,"rateCorr":1.5,"fixCorrA":2,"fixCorrB":3,"rate":0.49504950495049505}`

	errList := AntiAirCutInSchema.Parse(zjson.Decode(bytes.NewReader([]byte(raw))), &aaci)
	if errList != nil {
		t.Fatal(z.Issues.Prettify(errList))
	}

	assert.Equal(t, int32(5), aaci.Id)
	assert.Equal(t, float64(1.5), aaci.RateCorr)
	assert.Equal(t, int32(2), aaci.FixCorrA)
	assert.Equal(t, int32(3), aaci.FixCorrB)
	assert.Equal(t, float64(0.49504950495049505), aaci.Rate)
}
