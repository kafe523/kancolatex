package fleet

import (
	"bytes"
	"testing"

	z "github.com/Oudwins/zog"
	"github.com/Oudwins/zog/parsers/zjson"
	"github.com/stretchr/testify/assert"
)

func TestShipDisplayStatus(t *testing.T) {
	t.Parallel()
	var sds ShipDisplayStatus

	const raw = `{"HP":81,"firePower":66,"armor":81,"torpedo":23,"avoid":76,"antiAir":100,"asw":11,"LoS":99,"luck":20,"range":2,"accuracy":2,"bomber":11}`

	errList := ShipDisplayStatusSchema.Parse(zjson.Decode(bytes.NewReader([]byte(raw))), &sds)
	if errList != nil {
		t.Fatal(z.Issues.Prettify(errList))
	}

	assert.Equal(t, int32(81), sds.HP)
	assert.Equal(t, int32(66), sds.FirePower)
	assert.Equal(t, int32(81), sds.Armor)
	assert.Equal(t, int32(23), sds.Torpedo)
	assert.Equal(t, int32(76), sds.Avoid)
	assert.Equal(t, int32(100), sds.AntiAir)
	assert.Equal(t, int32(11), sds.Asw)
	assert.Equal(t, int32(99), sds.LoS)
	assert.Equal(t, int32(20), sds.Luck)
	assert.Equal(t, int32(2), sds.Range)
	assert.Equal(t, int32(2), sds.Accuracy)
	assert.Equal(t, int32(11), sds.Bomber)
}

var testShipSchemaCase = map[string]func(*testing.T){
	"赤城改二": func(t *testing.T) {
		t.Parallel()
		var s Ship

		const raw = `{"data":{"id":594,"albumId":404,"name":"赤城改二","yomi":"あかぎ","type":11,"type2":14,"slotCount":5,"slots":[21,21,32,12,4],"version":2,"isFinal":true,"originalId":83,"range":2,"hp":81,"hp2":89,"maxHp":91,"fire":60,"torpedo":0,"night":60,"antiAir":85,"armor":81,"luck":20,"maxLuck":77,"minScout":51,"maxScout":91,"minAsw":0,"maxAsw":0,"minAvoid":33,"maxAvoid":73,"speed":10,"beforeId":277,"nextLv":92,"sort":2016,"fuel":95,"ammo":90,"blueprints":0,"actionReports":0,"catapults":0,"isCV":true,"isBB":false},"items":[{"data":{"id":343,"apiTypeId":8,"iconTypeId":8,"name":"流星改(一航戦/熟練)","abbr":"","fire":0,"torpedo":15,"bomber":0,"antiAir":3,"armor":0,"asw":7,"antiBomber":0,"accuracy":2,"interception":0,"avoid":0,"scout":6,"range":0,"radius":6,"cost":9,"canRemodel":true,"avoidId":1,"grow":3,"sortieAntiAir":3,"defenseAntiAir":3,"isSpecial":true,"isPlane":true,"isFighter":false,"isAttacker":true,"isAswPlane":false,"isAswBomber1":false,"isAswBomber2":false,"isAutoGyro":false,"isABAttacker":false,"isJet":false,"isHeavyJet":false,"isBakusen":false,"isRocket":false,"isRecon":false,"isShinzan":false,"enabledAttackLandBase":false,"isStrictDepthCharge":false,"isTorpedoAttacker":true,"isNightAircraftItem":false,"isSPPlane":false,"isLateModelTorpedo":false,"isEnemyItem":false,"bonuses":[],"airbaseMaxSlot":18},"level":120,"remodel":2,"fullSlot":21,"bonusFire":0.4,"bonusNightFire":0,"bonusExpeditionFire":0,"bonusTorpedo":0.4,"bonusBomber":0,"bonusAntiAir":0,"bonusExpeditionAntiAir":0,"bonusArmor":0,"bonusAccuracy":0,"bonusAsw":0.4,"bonusExpeditionAsw":0,"bonusScout":0,"bonusExpeditionScout":0,"bonusAirPower":3.4641016151377544,"itemScout":4.800000000000001,"antiAirWeight":0,"antiAirBonus":60,"dayBattleFirePower":0.4,"aircraftDayBattleFirePower":23.1,"nightBattleFirePower":15,"actualFire":0.4,"actualAntiAir":3,"actualTorpedo":15.4,"actualBomber":0,"actualAsw":7.4,"actualArmor":0,"actualAccuracy":2,"actualScout":6,"actualAvoid":0,"actualRange":0,"actualDefenseAntiAir":3,"fullAirPower":17,"fullDefenseAirPower":17,"supportAirPower":13,"supportAswAirPower":13,"contactSelectRates":[0.42857142857142855,0.375,0.3333333333333333],"tp":0,"tp2":0,"tp3":0,"fuel":21,"ammo":13,"steel":0,"bauxite":162,"reconCorr":1,"reconCorrDefense":1,"calculatedAirPower":[3,6,7,8,9,10,10,11,11,12,12,13,13,14,14,15,15,15,16,16,16,17],"calculatedDefenseAirPower":[3,6,7,8,9,10,10,11,11,12,12,13,13,14,14,15,15,15,16,16,16,17],"attackerTorpedoBonus":0,"crewTorpedoBonus":0,"crewBomberBonus":0,"airPower":17,"defenseAirPower":17,"slot":21,"slotHistories":[],"slotResult":0,"deathRate":0,"minSlot":21,"maxSlot":0,"isEscortItem":false,"disabledItem":false,"needRecord":false,"dist":[],"parentIndex":0,"noStock":false},{"data":{"id":476,"apiTypeId":7,"iconTypeId":7,"name":"F4U-7","abbr":"","fire":3,"torpedo":8,"bomber":11,"antiAir":10,"armor":0,"asw":4,"antiBomber":0,"accuracy":0,"interception":0,"avoid":2,"scout":2,"range":0,"radius":6,"cost":11,"canRemodel":false,"avoidId":1,"grow":6,"sortieAntiAir":10,"defenseAntiAir":10,"isSpecial":false,"isPlane":true,"isFighter":false,"isAttacker":true,"isAswPlane":false,"isAswBomber1":false,"isAswBomber2":false,"isAutoGyro":false,"isABAttacker":false,"isJet":false,"isHeavyJet":false,"isBakusen":false,"isRocket":false,"isRecon":false,"isShinzan":false,"enabledAttackLandBase":false,"isStrictDepthCharge":false,"isTorpedoAttacker":false,"isNightAircraftItem":false,"isSPPlane":false,"isLateModelTorpedo":false,"isEnemyItem":false,"bonuses":[{"key":"59-1","text":["B"],"isOnlyAB":false,"isOnlyShip":false},{"key":"59-1","text":["3"],"isOnlyAB":false,"isOnlyShip":false},{"key":"61-1","text":["A1"],"isOnlyAB":false,"isOnlyShip":false},{"key":"61-1","text":["C2"],"isOnlyAB":false,"isOnlyShip":false}],"airbaseMaxSlot":18},"level":120,"remodel":0,"fullSlot":21,"bonusFire":0,"bonusNightFire":0,"bonusExpeditionFire":0,"bonusTorpedo":0,"bonusBomber":0,"bonusAntiAir":0,"bonusExpeditionAntiAir":0,"bonusArmor":0,"bonusAccuracy":0,"bonusAsw":0,"bonusExpeditionAsw":0,"bonusScout":0,"bonusExpeditionScout":0,"bonusAirPower":3.4641016151377544,"itemScout":1.2,"antiAirWeight":0,"antiAirBonus":200,"dayBattleFirePower":3,"aircraftDayBattleFirePower":37.5,"nightBattleFirePower":11,"actualFire":3,"actualAntiAir":10,"actualTorpedo":8,"actualBomber":11,"actualAsw":4,"actualArmor":0,"actualAccuracy":0,"actualScout":2,"actualAvoid":2,"actualRange":0,"actualDefenseAntiAir":10,"fullAirPower":49,"fullDefenseAirPower":49,"supportAirPower":45,"supportAswAirPower":45,"contactSelectRates":[0,0,0],"tp":0,"tp2":0,"tp3":0,"fuel":21,"ammo":13,"steel":0,"bauxite":198,"reconCorr":1,"reconCorrDefense":1,"calculatedAirPower":[3,13,17,20,23,25,27,29,31,33,35,36,38,39,40,42,43,44,45,47,48,49],"calculatedDefenseAirPower":[3,13,17,20,23,25,27,29,31,33,35,36,38,39,40,42,43,44,45,47,48,49],"attackerTorpedoBonus":0,"crewTorpedoBonus":0,"crewBomberBonus":0,"airPower":49,"defenseAirPower":49,"slot":21,"slotHistories":[],"slotResult":0,"deathRate":0,"minSlot":21,"maxSlot":0,"isEscortItem":false,"disabledItem":false,"needRecord":false,"dist":[],"parentIndex":0,"noStock":false},{"data":{"id":0,"apiTypeId":0,"iconTypeId":0,"name":"","abbr":"","fire":0,"torpedo":0,"bomber":0,"antiAir":0,"armor":0,"asw":0,"antiBomber":0,"accuracy":0,"interception":0,"avoid":0,"scout":0,"range":0,"radius":0,"cost":0,"canRemodel":false,"avoidId":0,"grow":0,"sortieAntiAir":0,"defenseAntiAir":0,"isSpecial":false,"isPlane":false,"isFighter":false,"isAttacker":false,"isAswPlane":false,"isAswBomber1":false,"isAswBomber2":false,"isAutoGyro":false,"isABAttacker":false,"isJet":false,"isHeavyJet":false,"isBakusen":false,"isRocket":false,"isRecon":false,"isShinzan":false,"enabledAttackLandBase":false,"isStrictDepthCharge":false,"isTorpedoAttacker":false,"isNightAircraftItem":false,"isSPPlane":false,"isLateModelTorpedo":false,"isEnemyItem":false,"bonuses":[],"airbaseMaxSlot":18},"level":0,"remodel":0,"fullSlot":32,"bonusFire":0,"bonusNightFire":0,"bonusExpeditionFire":0,"bonusTorpedo":0,"bonusBomber":0,"bonusAntiAir":0,"bonusExpeditionAntiAir":0,"bonusArmor":0,"bonusAccuracy":0,"bonusAsw":0,"bonusExpeditionAsw":0,"bonusScout":0,"bonusExpeditionScout":0,"bonusAirPower":0,"itemScout":0,"antiAirWeight":0,"antiAirBonus":0,"dayBattleFirePower":0,"aircraftDayBattleFirePower":0,"nightBattleFirePower":0,"actualFire":0,"actualAntiAir":0,"actualTorpedo":0,"actualBomber":0,"actualAsw":0,"actualArmor":0,"actualAccuracy":0,"actualScout":0,"actualAvoid":0,"actualRange":0,"actualDefenseAntiAir":0,"fullAirPower":0,"fullDefenseAirPower":0,"supportAirPower":0,"supportAswAirPower":0,"contactSelectRates":[0,0,0],"tp":0,"tp2":0,"tp3":0,"fuel":0,"ammo":0,"steel":0,"bauxite":0,"reconCorr":1,"reconCorrDefense":1,"calculatedAirPower":[],"calculatedDefenseAirPower":[],"attackerTorpedoBonus":0,"crewTorpedoBonus":0,"crewBomberBonus":0,"airPower":0,"defenseAirPower":0,"slot":32,"slotHistories":[],"slotResult":0,"deathRate":0,"minSlot":32,"maxSlot":32,"isEscortItem":false,"disabledItem":false,"needRecord":false,"dist":[],"parentIndex":-1,"noStock":false},{"data":{"id":0,"apiTypeId":0,"iconTypeId":0,"name":"","abbr":"","fire":0,"torpedo":0,"bomber":0,"antiAir":0,"armor":0,"asw":0,"antiBomber":0,"accuracy":0,"interception":0,"avoid":0,"scout":0,"range":0,"radius":0,"cost":0,"canRemodel":false,"avoidId":0,"grow":0,"sortieAntiAir":0,"defenseAntiAir":0,"isSpecial":false,"isPlane":false,"isFighter":false,"isAttacker":false,"isAswPlane":false,"isAswBomber1":false,"isAswBomber2":false,"isAutoGyro":false,"isABAttacker":false,"isJet":false,"isHeavyJet":false,"isBakusen":false,"isRocket":false,"isRecon":false,"isShinzan":false,"enabledAttackLandBase":false,"isStrictDepthCharge":false,"isTorpedoAttacker":false,"isNightAircraftItem":false,"isSPPlane":false,"isLateModelTorpedo":false,"isEnemyItem":false,"bonuses":[],"airbaseMaxSlot":18},"level":0,"remodel":0,"fullSlot":12,"bonusFire":0,"bonusNightFire":0,"bonusExpeditionFire":0,"bonusTorpedo":0,"bonusBomber":0,"bonusAntiAir":0,"bonusExpeditionAntiAir":0,"bonusArmor":0,"bonusAccuracy":0,"bonusAsw":0,"bonusExpeditionAsw":0,"bonusScout":0,"bonusExpeditionScout":0,"bonusAirPower":0,"itemScout":0,"antiAirWeight":0,"antiAirBonus":0,"dayBattleFirePower":0,"aircraftDayBattleFirePower":0,"nightBattleFirePower":0,"actualFire":0,"actualAntiAir":0,"actualTorpedo":0,"actualBomber":0,"actualAsw":0,"actualArmor":0,"actualAccuracy":0,"actualScout":0,"actualAvoid":0,"actualRange":0,"actualDefenseAntiAir":0,"fullAirPower":0,"fullDefenseAirPower":0,"supportAirPower":0,"supportAswAirPower":0,"contactSelectRates":[0,0,0],"tp":0,"tp2":0,"tp3":0,"fuel":0,"ammo":0,"steel":0,"bauxite":0,"reconCorr":1,"reconCorrDefense":1,"calculatedAirPower":[],"calculatedDefenseAirPower":[],"attackerTorpedoBonus":0,"crewTorpedoBonus":0,"crewBomberBonus":0,"airPower":0,"defenseAirPower":0,"slot":12,"slotHistories":[],"slotResult":0,"deathRate":0,"minSlot":12,"maxSlot":12,"isEscortItem":false,"disabledItem":false,"needRecord":false,"dist":[],"parentIndex":-1,"noStock":false},{"data":{"id":0,"apiTypeId":0,"iconTypeId":0,"name":"","abbr":"","fire":0,"torpedo":0,"bomber":0,"antiAir":0,"armor":0,"asw":0,"antiBomber":0,"accuracy":0,"interception":0,"avoid":0,"scout":0,"range":0,"radius":0,"cost":0,"canRemodel":false,"avoidId":0,"grow":0,"sortieAntiAir":0,"defenseAntiAir":0,"isSpecial":false,"isPlane":false,"isFighter":false,"isAttacker":false,"isAswPlane":false,"isAswBomber1":false,"isAswBomber2":false,"isAutoGyro":false,"isABAttacker":false,"isJet":false,"isHeavyJet":false,"isBakusen":false,"isRocket":false,"isRecon":false,"isShinzan":false,"enabledAttackLandBase":false,"isStrictDepthCharge":false,"isTorpedoAttacker":false,"isNightAircraftItem":false,"isSPPlane":false,"isLateModelTorpedo":false,"isEnemyItem":false,"bonuses":[],"airbaseMaxSlot":18},"level":0,"remodel":0,"fullSlot":4,"bonusFire":0,"bonusNightFire":0,"bonusExpeditionFire":0,"bonusTorpedo":0,"bonusBomber":0,"bonusAntiAir":0,"bonusExpeditionAntiAir":0,"bonusArmor":0,"bonusAccuracy":0,"bonusAsw":0,"bonusExpeditionAsw":0,"bonusScout":0,"bonusExpeditionScout":0,"bonusAirPower":0,"itemScout":0,"antiAirWeight":0,"antiAirBonus":0,"dayBattleFirePower":0,"aircraftDayBattleFirePower":0,"nightBattleFirePower":0,"actualFire":0,"actualAntiAir":0,"actualTorpedo":0,"actualBomber":0,"actualAsw":0,"actualArmor":0,"actualAccuracy":0,"actualScout":0,"actualAvoid":0,"actualRange":0,"actualDefenseAntiAir":0,"fullAirPower":0,"fullDefenseAirPower":0,"supportAirPower":0,"supportAswAirPower":0,"contactSelectRates":[0,0,0],"tp":0,"tp2":0,"tp3":0,"fuel":0,"ammo":0,"steel":0,"bauxite":0,"reconCorr":1,"reconCorrDefense":1,"calculatedAirPower":[],"calculatedDefenseAirPower":[],"attackerTorpedoBonus":0,"crewTorpedoBonus":0,"crewBomberBonus":0,"airPower":0,"defenseAirPower":0,"slot":4,"slotHistories":[],"slotResult":0,"deathRate":0,"minSlot":4,"maxSlot":4,"isEscortItem":false,"disabledItem":false,"needRecord":false,"dist":[],"parentIndex":-1,"noStock":false}],"exItem":{"data":{"id":0,"apiTypeId":0,"iconTypeId":0,"name":"","abbr":"","fire":0,"torpedo":0,"bomber":0,"antiAir":0,"armor":0,"asw":0,"antiBomber":0,"accuracy":0,"interception":0,"avoid":0,"scout":0,"range":0,"radius":0,"cost":0,"canRemodel":false,"avoidId":0,"grow":0,"sortieAntiAir":0,"defenseAntiAir":0,"isSpecial":false,"isPlane":false,"isFighter":false,"isAttacker":false,"isAswPlane":false,"isAswBomber1":false,"isAswBomber2":false,"isAutoGyro":false,"isABAttacker":false,"isJet":false,"isHeavyJet":false,"isBakusen":false,"isRocket":false,"isRecon":false,"isShinzan":false,"enabledAttackLandBase":false,"isStrictDepthCharge":false,"isTorpedoAttacker":false,"isNightAircraftItem":false,"isSPPlane":false,"isLateModelTorpedo":false,"isEnemyItem":false,"bonuses":[],"airbaseMaxSlot":18},"level":0,"remodel":0,"fullSlot":0,"bonusFire":0,"bonusNightFire":0,"bonusExpeditionFire":0,"bonusTorpedo":0,"bonusBomber":0,"bonusAntiAir":0,"bonusExpeditionAntiAir":0,"bonusArmor":0,"bonusAccuracy":0,"bonusAsw":0,"bonusExpeditionAsw":0,"bonusScout":0,"bonusExpeditionScout":0,"bonusAirPower":0,"itemScout":0,"antiAirWeight":0,"antiAirBonus":0,"dayBattleFirePower":0,"aircraftDayBattleFirePower":0,"nightBattleFirePower":0,"actualFire":0,"actualAntiAir":0,"actualTorpedo":0,"actualBomber":0,"actualAsw":0,"actualArmor":0,"actualAccuracy":0,"actualScout":0,"actualAvoid":0,"actualRange":0,"actualDefenseAntiAir":0,"fullAirPower":0,"fullDefenseAirPower":0,"supportAirPower":0,"supportAswAirPower":0,"contactSelectRates":[0,0,0],"tp":0,"tp2":0,"tp3":0,"fuel":0,"ammo":0,"steel":0,"bauxite":0,"reconCorr":1,"reconCorrDefense":1,"calculatedAirPower":[],"calculatedDefenseAirPower":[],"attackerTorpedoBonus":0,"crewTorpedoBonus":0,"crewBomberBonus":0,"airPower":0,"defenseAirPower":0,"slot":0,"slotHistories":[],"slotResult":0,"deathRate":0,"minSlot":0,"maxSlot":0,"isEscortItem":false,"disabledItem":false,"needRecord":false,"dist":[],"parentIndex":-1,"noStock":false},"level":99,"displayStatus":{"HP":81,"firePower":66,"armor":81,"torpedo":23,"avoid":76,"antiAir":100,"asw":11,"LoS":99,"luck":20,"range":2,"accuracy":2,"bomber":11},"itemBonusStatus":{"firePower":3,"torpedo":0,"antiAir":2,"armor":0,"asw":0,"scout":0,"avoid":1,"accuracy":0,"bomber":0,"range":0},"itemBonuses":[{"firePower":3,"antiAir":2,"avoid":1,"fromTypeId":8}],"hp":81,"luck":20,"baseDayBattleFirePower":210,"supportFirePower":208,"nightBattleFirePower":89,"antiAir":85,"actualArmor":81,"scout":91,"itemsScout":6.000000000000001,"avoid":73,"accuracy":2,"supportAccuracy":2,"asw":0,"improveAsw":0,"itemAsw":11,"enabledTSBK":false,"tp":0,"tp2":0,"tp3":0,"speed":10,"fuel":95,"ammo":90,"area":0,"uniqueId":40,"hunshinRate":0,"isActive":true,"isEmpty":false,"antiAirBonus":2,"isEscort":false,"hasJet":false,"fullAirPower":66,"supportAirPower":58,"supportAswAirPower":58,"enabledASWSupport":false,"antiAirCutIn":[],"specialKokakuCount":0,"kokakuCount":0,"specialKijuCount":0,"kijuCount":0,"antiAirRadarCount":0,"surfaceRadarCount":0,"koshaCount":0,"sumSPRos":0,"nightContactRate":0,"enabledAircraftNightAttack":false,"nightAttackCrewFireBonus":0,"nightAttackCrewBomberBonus":0,"releaseExpand":false,"noStock":false,"isTray":false,"spEffectItemId":0,"missingAsw":0,"needTSBKLevel":0,"fixDown":9,"rateDown":0.22,"allPlaneDeathRate":0}`

		errList := ShipSchema.Parse(zjson.Decode(bytes.NewReader([]byte(raw))), &s)
		if errList != nil {
			t.Fatal(z.Issues.Prettify(errList))
		}

		// Skip type check
		// assert.Equal(t, nil, s.Data)
		// assert.Equal(t, nil, s.items)
		// assert.Equal(t, nil, s.exItem)

		assert.Equal(t, int32(99), s.Level)

		// Skip type check
		// assert.Equal(t, nil, s.DisplayStatus)
		// assert.Equal(t, nil, s.ItemBonusStatus)
		// assert.Equal(t, nil, s.ItemBonuses)

		assert.Equal(t, int32(81), s.Hp)
		assert.Equal(t, int32(20), s.Luck)
		assert.Equal(t, int32(210), s.BaseDayBattleFirePower)
		assert.Equal(t, int32(208), s.SupportFirePower)
		assert.Equal(t, int32(89), s.NightBattleFirePower)
		assert.Equal(t, int32(85), s.AntiAir)
		assert.Equal(t, int32(81), s.ActualArmor)
		assert.Equal(t, int32(91), s.Scout)
		assert.Equal(t, float64(6.000000000000001), s.ItemsScout)
		assert.Equal(t, int32(73), s.Avoid)
		assert.Equal(t, int32(2), s.Accuracy)
		assert.Equal(t, int32(2), s.SupportAccuracy)
		assert.Equal(t, int32(0), s.Asw)
		assert.Equal(t, int32(0), s.ImproveAsw)
		assert.Equal(t, int32(11), s.ItemAsw)
		assert.Equal(t, false, s.EnabledTSBK)
		assert.Equal(t, int32(0), s.Tp)
		assert.Equal(t, int32(0), s.Tp2)
		assert.Equal(t, int32(0), s.Tp3)
		assert.Equal(t, int32(10), s.Speed)
		assert.Equal(t, int32(95), s.Fuel)
		assert.Equal(t, int32(90), s.Ammo)
		assert.Equal(t, int32(0), s.Area)
		assert.Equal(t, int32(40), s.UniqueId)
		assert.Equal(t, int32(0), s.HunshinRate)
		assert.Equal(t, true, s.IsActive)
		assert.Equal(t, false, s.IsEmpty)
		assert.Equal(t, int32(2), s.AntiAirBonus)
		assert.Equal(t, false, s.IsEscort)
		assert.Equal(t, false, s.HasJet)
		assert.Equal(t, int32(66), s.FullAirPower)
		assert.Equal(t, int32(58), s.SupportAirPower)
		assert.Equal(t, int32(58), s.SupportAswAirPower)
		assert.Equal(t, false, s.EnabledASWSupport)

		// Skip type check
		// assert.Equal(t, nil, s.AntiAirCutIn)

		assert.Equal(t, int32(0), s.SpecialKokakuCount)
		assert.Equal(t, int32(0), s.KokakuCount)
		assert.Equal(t, int32(0), s.SpecialKijuCount)
		assert.Equal(t, int32(0), s.KijuCount)
		assert.Equal(t, int32(0), s.AntiAirRadarCount)
		assert.Equal(t, int32(0), s.SurfaceRadarCount)
		assert.Equal(t, int32(0), s.KoshaCount)
		assert.Equal(t, int32(0), s.SumSPRos)
		assert.Equal(t, int32(0), s.NightContactRate)
		assert.Equal(t, false, s.EnabledAircraftNightAttack)
		assert.Equal(t, int32(0), s.NightAttackCrewFireBonus)
		assert.Equal(t, int32(0), s.NightAttackCrewBomberBonus)
		assert.Equal(t, false, s.ReleaseExpand)
		assert.Equal(t, false, s.NoStock)
		assert.Equal(t, false, s.IsTray)
		assert.Equal(t, int32(0), s.SpEffectItemId)
		assert.Equal(t, int32(0), s.MissingAsw)
		assert.Equal(t, int32(0), s.NeedTSBKLevel)
		assert.Equal(t, int32(9), s.FixDown)
		assert.Equal(t, float64(0.22), s.RateDown)
		assert.Equal(t, int32(0), s.AllPlaneDeathRate)
	},
}

func TestShipSchema(t *testing.T) {
	for tn, tc := range testShipSchemaCase {
		t.Run(tn, tc)
	}
}
