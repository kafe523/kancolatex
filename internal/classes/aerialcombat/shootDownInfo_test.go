package aerialcombat

import (
	"bytes"
	"testing"

	z "github.com/Oudwins/zog"
	"github.com/Oudwins/zog/parsers/zjson"
	"github.com/stretchr/testify/assert"
)

func TestShootDownStatusSchema(t *testing.T) {
	t.Parallel()
	var sds ShootDownStatus

	const raw = `{"antiAirWeightList":[44],"fixDownList":[9],"rateDownList":[0.22],"minimumDownList":[1]}`

	errList := ShootDownStatusSchema.Parse(zjson.Decode(bytes.NewReader([]byte(raw))), &sds)
	if errList != nil {
		t.Fatal(z.Issues.Prettify(errList))
	}
}

func TestShootDownInfoSchema(t *testing.T) {
	t.Parallel()
	var sdi ShootDownInfo

	const raw = `{"shootDownStatusList":[{"antiAirWeightList":[44],"fixDownList":[9],"rateDownList":[0.22],"minimumDownList":[1]},{"antiAirWeightList":[26],"fixDownList":[5],"rateDownList":[0.13],"minimumDownList":[1]},{"antiAirWeightList":[26],"fixDownList":[5],"rateDownList":[0.13],"minimumDownList":[0]},{"antiAirWeightList":[22],"fixDownList":[4],"rateDownList":[0.11],"minimumDownList":[0]},{"antiAirWeightList":[22],"fixDownList":[4],"rateDownList":[0.11],"minimumDownList":[0]},{"antiAirWeightList":[17],"fixDownList":[3],"rateDownList":[0.085],"minimumDownList":[0]},{"antiAirWeightList":[44],"fixDownList":[9],"rateDownList":[0.22],"minimumDownList":[1]}],"border":1,"maxRange":1}`

	errList := ShootDownInfoSchema.Parse(zjson.Decode(bytes.NewReader([]byte(raw))), &sdi)
	if errList != nil {
		t.Fatal(z.Issues.Prettify(errList))
	}

	// Skip type check
	// assert.Equal(t, nil, sdi.ShootDownStatusList)

	assert.Equal(t, int32(1), sdi.Border)
	assert.Equal(t, int32(1), sdi.MaxRange)
}
