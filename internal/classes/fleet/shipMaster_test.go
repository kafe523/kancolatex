package fleet

import (
	"bytes"
	"testing"

	z "github.com/Oudwins/zog"
	"github.com/Oudwins/zog/parsers/zjson"
	"github.com/stretchr/testify/assert"
)

var testShipMasterSchemaCase = map[string]func(*testing.T){
	"赤城改二": func(t *testing.T) {
		t.Parallel()
		var sm ShipMaster

		const raw = `{"id":594,"albumId":404,"name":"赤城改二","yomi":"あかぎ","type":11,"type2":14,"slotCount":5,"slots":[21,21,32,12,4],"version":2,"isFinal":true,"originalId":83,"range":2,"hp":81,"hp2":89,"maxHp":91,"fire":60,"torpedo":0,"night":60,"antiAir":85,"armor":81,"luck":20,"maxLuck":77,"minScout":51,"maxScout":91,"minAsw":0,"maxAsw":0,"minAvoid":33,"maxAvoid":73,"speed":10,"beforeId":277,"nextLv":92,"sort":2016,"fuel":95,"ammo":90,"blueprints":0,"actionReports":0,"catapults":0,"isCV":true,"isBB":false}`

		errList := ShipMasterSchema.Parse(zjson.Decode(bytes.NewReader([]byte(raw))), &sm)
		if errList != nil {
			t.Fatal(z.Issues.Prettify(errList))
		}

		assert.Equal(t, int32(594), sm.Id)
		assert.Equal(t, int32(404), sm.AlbumId)
		assert.Equal(t, "赤城改二", sm.Name)
		assert.Equal(t, "あかぎ", sm.Yomi)
		assert.Equal(t, int32(11), sm.Type)
		assert.Equal(t, int32(14), sm.Type2)
		assert.Equal(t, int32(5), sm.SlotCount)

		// Skip type check
		// assert.Equal(t, nil, sm.Slots)

		assert.Equal(t, int32(2), sm.Version)
		assert.Equal(t, true, sm.IsFinal)
		assert.Equal(t, int32(83), sm.OriginalId)
		assert.Equal(t, int32(2), sm.Range)
		assert.Equal(t, int32(81), sm.Hp)
		assert.Equal(t, int32(89), sm.Hp2)
		assert.Equal(t, int32(91), sm.MaxHp)
		assert.Equal(t, int32(60), sm.Fire)
		assert.Equal(t, int32(0), sm.Torpedo)
		assert.Equal(t, int32(60), sm.Night)
		assert.Equal(t, int32(85), sm.AntiAir)
		assert.Equal(t, int32(81), sm.Armor)
		assert.Equal(t, int32(20), sm.Luck)
		assert.Equal(t, int32(77), sm.MaxLuck)
		assert.Equal(t, int32(51), sm.MinScout)
		assert.Equal(t, int32(91), sm.MaxScout)
		assert.Equal(t, int32(0), sm.MinAsw)
		assert.Equal(t, int32(0), sm.MaxAsw)
		assert.Equal(t, int32(33), sm.MinAvoid)
		assert.Equal(t, int32(73), sm.MaxAvoid)
		assert.Equal(t, int32(10), sm.Speed)
		assert.Equal(t, int32(277), sm.BeforeId)
		assert.Equal(t, int32(92), sm.NextLv)
		assert.Equal(t, int32(2016), sm.Sort)
		assert.Equal(t, int32(95), sm.Fuel)
		assert.Equal(t, int32(90), sm.Ammo)
		assert.Equal(t, int32(0), sm.Blueprints)
		assert.Equal(t, int32(0), sm.ActionReports)
		assert.Equal(t, int32(0), sm.Catapults)
		assert.Equal(t, true, sm.IsCV)
		assert.Equal(t, false, sm.IsBB)
	},
}

func TestShipMasterSchema(t *testing.T) {
	for tn, tc := range testShipMasterSchemaCase {
		t.Run(tn, tc)
	}
}
