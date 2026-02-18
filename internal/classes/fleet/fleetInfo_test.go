package fleet

import (
	"bytes"
	"testing"

	z "github.com/Oudwins/zog"
	"github.com/Oudwins/zog/parsers/zjson"
	"github.com/kafe523/kancolatex/internal/classes/consts/fleettype"
	"github.com/stretchr/testify/assert"
)

func TestFleetInfoSchema(t *testing.T) {
	t.Parallel()
	var fi FleetInfo

	const raw = `{"fleetInfo":{},"isUnion":false,"fleetType":0,"admiralLevel":120,"mainFleetIndex":0,"calculated":false,"ignoreHistory":false}`

	errList := FleetInfoSchema.Parse(zjson.Decode(bytes.NewReader([]byte(raw))), &fi)
	if errList != nil {
		t.Fatal(z.Issues.Prettify(errList))
	}

	// Skip Check in fleet_test
	// assert.Equal(t, nil, fi.Fleets)

	assert.Equal(t, false, fi.IsUnion)
	assert.Equal(t, fleettype.SINGLE, fi.FleetType)
	assert.Equal(t, int32(120), fi.AdmiralLevel)
	assert.Equal(t, int32(0), fi.MainFleetIndex)
	assert.Nil(t, fi.UnionFleet)
	assert.Equal(t, false, fi.Calculated)
	assert.Equal(t, false, fi.IgnoreHistory)

}
