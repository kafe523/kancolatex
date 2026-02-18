package item

import (
	"bytes"
	"testing"

	z "github.com/Oudwins/zog"
	"github.com/Oudwins/zog/parsers/zjson"
	"github.com/stretchr/testify/assert"
)

var testItemSchemaCase = map[string]func(*testing.T){
	"流星改(一航戦/熟練)": func(t *testing.T) {
		t.Parallel()

		var im Item

		const raw = `{"data":{"id":343,"apiTypeId":8,"iconTypeId":8,"name":"流星改(一航戦/熟練)","abbr":"","fire":0,"torpedo":15,"bomber":0,"antiAir":3,"armor":0,"asw":7,"antiBomber":0,"accuracy":2,"interception":0,"avoid":0,"scout":6,"range":0,"radius":6,"cost":9,"canRemodel":true,"avoidId":1,"grow":3,"sortieAntiAir":3,"defenseAntiAir":3,"isSpecial":true,"isPlane":true,"isFighter":false,"isAttacker":true,"isAswPlane":false,"isAswBomber1":false,"isAswBomber2":false,"isAutoGyro":false,"isABAttacker":false,"isJet":false,"isHeavyJet":false,"isBakusen":false,"isRocket":false,"isRecon":false,"isShinzan":false,"enabledAttackLandBase":false,"isStrictDepthCharge":false,"isTorpedoAttacker":true,"isNightAircraftItem":false,"isSPPlane":false,"isLateModelTorpedo":false,"isEnemyItem":false,"bonuses":[],"airbaseMaxSlot":18},"level":120,"remodel":2,"fullSlot":21,"bonusFire":0.4,"bonusNightFire":0,"bonusExpeditionFire":0,"bonusTorpedo":0.4,"bonusBomber":0,"bonusAntiAir":0,"bonusExpeditionAntiAir":0,"bonusArmor":0,"bonusAccuracy":0,"bonusAsw":0.4,"bonusExpeditionAsw":0,"bonusScout":0,"bonusExpeditionScout":0,"bonusAirPower":3.4641016151377544,"itemScout":4.800000000000001,"antiAirWeight":0,"antiAirBonus":60,"dayBattleFirePower":0.4,"aircraftDayBattleFirePower":23.1,"nightBattleFirePower":15,"actualFire":0.4,"actualAntiAir":3,"actualTorpedo":15.4,"actualBomber":0,"actualAsw":7.4,"actualArmor":0,"actualAccuracy":2,"actualScout":6,"actualAvoid":0,"actualRange":0,"actualDefenseAntiAir":3,"fullAirPower":17,"fullDefenseAirPower":17,"supportAirPower":13,"supportAswAirPower":13,"contactSelectRates":[0.42857142857142855,0.375,0.3333333333333333],"tp":0,"tp2":0,"tp3":0,"fuel":21,"ammo":13,"steel":0,"bauxite":162,"reconCorr":1,"reconCorrDefense":1,"calculatedAirPower":[3,6,7,8,9,10,10,11,11,12,12,13,13,14,14,15,15,15,16,16,16,17],"calculatedDefenseAirPower":[3,6,7,8,9,10,10,11,11,12,12,13,13,14,14,15,15,15,16,16,16,17],"attackerTorpedoBonus":0,"crewTorpedoBonus":0,"crewBomberBonus":0,"airPower":17,"defenseAirPower":17,"slot":21,"slotHistories":[],"slotResult":0,"deathRate":0,"minSlot":21,"maxSlot":0,"isEscortItem":false,"disabledItem":false,"needRecord":false,"dist":[],"parentIndex":0,"noStock":false}`

		errList := ItemSchema.Parse(zjson.Decode(bytes.NewReader([]byte(raw))), &im)
		if errList != nil {
			t.Fatal(z.Issues.Prettify(errList))
		}

		// Skip data, checked in itemMaster
		// assert.Equal(t, nil, im.Data)
		assert.Equal(t, int32(120), im.Level)
		assert.Equal(t, int32(2), im.Remodel)
		assert.Equal(t, int32(21), im.FullSlot)
		assert.Equal(t, float64(0.4), im.BonusFire)
		assert.Equal(t, float64(0), im.BonusNightFire)
		assert.Equal(t, float64(0), im.BonusExpeditionFire)
		assert.Equal(t, float64(0.4), im.BonusTorpedo)
		assert.Equal(t, float64(0), im.BonusBomber)
		assert.Equal(t, float64(0), im.BonusAntiAir)
		assert.Equal(t, float64(0), im.BonusExpeditionAntiAir)
		assert.Equal(t, float64(0.4), im.BonusAsw)
		assert.Equal(t, float64(0), im.BonusExpeditionAsw)
		assert.Equal(t, float64(0), im.BonusScout)
		assert.Equal(t, float64(0), im.BonusExpeditionScout)
		assert.Equal(t, float64(3.4641016151377544), im.BonusAirPower)
		assert.Equal(t, float64(4.800000000000001), im.ItemScout)
		assert.Equal(t, float64(0), im.AntiAirWeight)
		assert.Equal(t, float64(60), im.AntiAirBonus)
		assert.Equal(t, float64(0.4), im.DayBattleFirePower)
		assert.Equal(t, float64(23.1), im.AircraftDayBattleFirePower)
		assert.Equal(t, float64(15), im.NightBattleFirePower)
		assert.Equal(t, float64(0.4), im.ActualFire)
		assert.Equal(t, float64(15.4), im.ActualTorpedo)
		assert.Equal(t, float64(0), im.ActualBomber)
		assert.Equal(t, float64(7.4), im.ActualAsw)
		assert.Equal(t, float64(0), im.ActualArmor)
		assert.Equal(t, float64(2), im.ActualAccuracy)
		assert.Equal(t, float64(6), im.ActualScout)
		assert.Equal(t, float64(0), im.ActualAvoid)
		assert.Equal(t, float64(0), im.ActualRange)
		assert.Equal(t, float64(3), im.ActualDefenseAntiAir)
		assert.Equal(t, int32(17), im.FullAirPower)
		assert.Equal(t, int32(17), im.FullDefenseAirPower)
		assert.Equal(t, int32(13), im.SupportAirPower)
		assert.Equal(t, int32(13), im.SupportAswAirPower)

		// Skip type check
		// assert.Equal(t, nil, im.ContactSelectRates)

		assert.Equal(t, float64(0), im.Tp)
		assert.Equal(t, float64(0), im.Tp2)
		assert.Equal(t, float64(0), im.Tp3)
		assert.Equal(t, int32(21), im.Fuel)
		assert.Equal(t, int32(13), im.Ammo)
		assert.Equal(t, int32(0), im.Steel)
		assert.Equal(t, int32(162), im.Bauxite)

		assert.Equal(t, float64(1), im.ReconCorr)
		assert.Equal(t, float64(1), im.ReconCorrDefense)

		// Skip type check
		// assert.Equal(t, nil, im.CalculatedAirPower)
		// assert.Equal(t, nil, im.CalculatedDefenseAirPower)

		assert.Equal(t, int32(0), im.AttackerTorpedoBonus)
		assert.Equal(t, int32(0), im.CrewTorpedoBonus)
		assert.Equal(t, int32(0), im.CrewBomberBonus)
		assert.Equal(t, int32(17), im.AirPower)
		assert.Equal(t, int32(17), im.DefenseAirPower)
		assert.Equal(t, int32(21), im.Slot)

		// Skip type check
		// assert.Equal(t, nil, im.SlotHistories)

		assert.Equal(t, int32(0), im.SlotResult)
		assert.Equal(t, int32(0), im.DeathRate)
		assert.Equal(t, int32(21), im.MinSlot)
		assert.Equal(t, int32(0), im.MaxSlot)
		assert.Equal(t, false, im.IsEscortItem)
		assert.Equal(t, false, im.DisabledItem)
		assert.Equal(t, false, im.NeedRecord)

		// Skip type check
		// assert.Equal(t, nil, im.Dist)
		assert.Equal(t, int32(0), im.ParentIndex)
		assert.Equal(t, false, im.NoStock)
	},
}

func TestItemSchema(t *testing.T) {
	for tn, tc := range testItemSchemaCase {
		t.Run(tn, tc)
	}
}
