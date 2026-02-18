package aircalcresult

import (
	"bytes"
	"testing"

	z "github.com/Oudwins/zog"
	"github.com/Oudwins/zog/parsers/zjson"
	"github.com/kafe523/kancolatex/internal/classes/consts"
	"github.com/stretchr/testify/assert"
)

func TestAirCalcResultSchema(t *testing.T) {
	t.Parallel()
	var acr AirCalcResult

	const raw = `{"airState":{"text":"不発","value":5,"color":"secondary"},"airStateText":"","airStateBarWidth":0,"rates":[0,0,0,0,0,0],"supportRates":[0,0,0,0,0,0],"loopSumAirPower":0,"loopSumEnemyAirPower":0,"loopSumEnemySupportAirPower":0,"avgAirPower":0,"avgEnemyAirPower":0,"avgEnemySupportAirPower":0,"avgDownSlot":0,"avgUsedSteels":0,"isUnknownEnemyAirPower":false}`

	errList := AirCalcResultSchema.Parse(zjson.Decode(bytes.NewReader([]byte(raw))), &acr)
	if errList != nil {
		t.Fatal(z.Issues.Prettify(errList))
	}

	assert.Equal(t, consts.AIR_STATUS[len(consts.AIR_STATUS)-1].Text, acr.AirState.Text)
	assert.Equal(t, consts.AIR_STATUS[len(consts.AIR_STATUS)-1].Value, acr.AirState.Value)
	assert.Equal(t, consts.AIR_STATUS[len(consts.AIR_STATUS)-1].Color, acr.AirState.Color)

	assert.Equal(t, "", acr.AirStateText)
	assert.Equal(t, int32(0), acr.AirStateBarWidth)

	// Skip type check
	// assert.Equal(t, nil, acr.Rates)
	assert.Equal(t, 6, len(acr.Rates))

	// Skip type check
	// assert.Equal(t, nil, acr.SupportRates)
	assert.Equal(t, 6, len(acr.SupportRates))

	assert.Equal(t, int32(0), acr.LoopSumAirPower)
	assert.Equal(t, int32(0), acr.LoopSumEnemyAirPower)
	assert.Equal(t, int32(0), acr.LoopSumEnemySupportAirPower)
	assert.Equal(t, int32(0), acr.AvgAirPower)
	assert.Equal(t, int32(0), acr.AvgEnemyAirPower)
	assert.Equal(t, int32(0), acr.AvgEnemySupportAirPower)
	assert.Equal(t, int32(0), acr.AvgDownSlot)
	assert.Equal(t, int32(0), acr.AvgUsedSteels)
	assert.Equal(t, false, acr.IsUnknownEnemyAirPower)
}
