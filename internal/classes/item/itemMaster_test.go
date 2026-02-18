package item

import (
	"bytes"
	"testing"

	z "github.com/Oudwins/zog"
	"github.com/Oudwins/zog/parsers/zjson"
	"github.com/stretchr/testify/assert"
)

var testItemMasterSchemaCase = map[string]func(*testing.T){
	"流星改(一航戦/熟練)": func(t *testing.T) {
		t.Parallel()
		var im ItemMaster

		const raw = `{"id":343,"apiTypeId":8,"iconTypeId":8,"name":"流星改(一航戦/熟練)","abbr":"","fire":0,"torpedo":15,"bomber":0,"antiAir":3,"armor":0,"asw":7,"antiBomber":0,"accuracy":2,"interception":0,"avoid":0,"scout":6,"range":0,"radius":6,"cost":9,"canRemodel":true,"avoidId":1,"grow":3,"sortieAntiAir":3,"defenseAntiAir":3,"isSpecial":true,"isPlane":true,"isFighter":false,"isAttacker":true,"isAswPlane":false,"isAswBomber1":false,"isAswBomber2":false,"isAutoGyro":false,"isABAttacker":false,"isJet":false,"isHeavyJet":false,"isBakusen":false,"isRocket":false,"isRecon":false,"isShinzan":false,"enabledAttackLandBase":false,"isStrictDepthCharge":false,"isTorpedoAttacker":true,"isNightAircraftItem":false,"isSPPlane":false,"isLateModelTorpedo":false,"isEnemyItem":false,"bonuses":[],"airbaseMaxSlot":18}`

		errList := ItemMasterSchema.Parse(zjson.Decode(bytes.NewReader([]byte(raw))), &im)
		if errList != nil {
			t.Fatal(z.Issues.Prettify(errList))
		}

		assert.Equal(t, int32(343), im.Id)
		assert.Equal(t, int32(8), im.ApiTypeId)
		assert.Equal(t, int32(8), im.IconTypeId)
		assert.Equal(t, "流星改(一航戦/熟練)", im.Name)
		assert.Equal(t, "", im.Abbr)
		assert.Equal(t, int32(0), im.Fire)
		assert.Equal(t, int32(15), im.Torpedo)
		assert.Equal(t, int32(0), im.Bomber)
		assert.Equal(t, int32(3), im.AntiAir)
		assert.Equal(t, int32(0), im.Armor)
		assert.Equal(t, int32(7), im.Asw)
		assert.Equal(t, int32(0), im.AntiBomber)
		assert.Equal(t, int32(2), im.Accuracy)
		assert.Equal(t, int32(0), im.Interception)
		assert.Equal(t, int32(0), im.Avoid)
		assert.Equal(t, int32(6), im.Scout)
		assert.Equal(t, int32(0), im.Range)
		assert.Equal(t, int32(6), im.Radius)
		assert.Equal(t, int32(9), im.Cost)
		assert.Equal(t, true, im.CanRemodel)
		assert.Equal(t, int32(1), im.AvoidId)
		assert.Equal(t, int32(3), im.Grow)
		assert.Equal(t, float64(3), im.SortieAntiAir)
		assert.Equal(t, int32(3), im.DefenseAntiAir)
		assert.Equal(t, true, im.IsSpecial)
		assert.Equal(t, true, im.IsPlane)
		assert.Equal(t, false, im.IsFighter)
		assert.Equal(t, true, im.IsAttacker)
		assert.Equal(t, false, im.IsAswPlane)
		assert.Equal(t, false, im.IsAswBomber1)
		assert.Equal(t, false, im.IsAswBomber2)
		assert.Equal(t, false, im.IsAutoGyro)
		assert.Equal(t, false, im.IsABAttacker)
		assert.Equal(t, false, im.IsJet)
		assert.Equal(t, false, im.IsHeavyJet)
		assert.Equal(t, false, im.IsBakusen)
		assert.Equal(t, false, im.IsRocket)
		assert.Equal(t, false, im.IsRecon)
		assert.Equal(t, false, im.IsShinzan)
		assert.Equal(t, false, im.EnabledAttackLandBase)
		assert.Equal(t, false, im.IsStrictDepthCharge)
		assert.Equal(t, true, im.IsTorpedoAttacker)
		assert.Equal(t, false, im.IsNightAircraftItem)
		assert.Equal(t, false, im.IsSPPlane)
		assert.Equal(t, false, im.IsLateModelTorpedo)
		assert.Equal(t, false, im.IsEnemyItem)
		assert.Equal(t, []struct {
			Key        string
			Text       []string
			IsOnlyAB   bool
			IsOnlyShip bool
		}{}, im.Bonuses)
		assert.Equal(t, int32(18), im.AirbaseMaxSlot)
	},
}

func TestItemMasterSchema(t *testing.T) {
	for tn, tc := range testItemMasterSchemaCase {
		t.Run(tn, tc)
	}
}
