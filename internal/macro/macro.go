package macro

import (
	"bytes"
	"errors"
	"fmt"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"

	z "github.com/Oudwins/zog"
	"github.com/Oudwins/zog/parsers/zjson"
	"github.com/kafe523/kancolatex/internal/classes"
	"github.com/kafe523/kancolatex/internal/classes/calcmanager"
	"github.com/kafe523/kancolatex/internal/classes/fleet"
)

var FleetLabelLookup = map[string]int{
	"A": 0, "B": 1, "C": 2, "D": 3,
	// "U": // union fleet have special treatment
}

var ShipLabelLookup = map[string]int{
	// Single / JTF 1-6
	"A": 0, "B": 1, "C": 2, "D": 3, "E": 4, "F": 5,
	// Single 7 (Vanguard Formation)
	"G": 6,
	// JTF 7-12
	"U": 6, "V": 7, "W": 8, "X": 9, "Y": 10, "Z": 11,
}

var ItemLabelLookup = map[string]int{
	"A": 0, "B": 1, "C": 2, "D": 3, "E": 4,
	"X": 99, // exItem is standalone object
}

type MacroContext struct {
	CalcManager    calcmanager.CalcManager
	TranslateItems map[string]string
	TranslateShips map[string]string
	PreCalcMap     map[string]string
	AccessPatMap   map[string]string
}

func (mctx *MacroContext) Populate(noroRaw *[]byte) error {
	if noroRaw != nil {
		errList := calcmanager.CalcManagerSchema.Parse(zjson.Decode(bytes.NewReader(*noroRaw)), &mctx.CalcManager)
		if errList != nil {
			return errors.New(z.Issues.Prettify(errList))
		}
	}

	mctx.TranslateItems = map[string]string{}
	mctx.TranslateShips = map[string]string{}

	mctx.precalcInit()
	mctx.accessStandardPatternInit()

	return nil
}

func (mctx *MacroContext) precalcInit() {
	mctx.PreCalcMap = map[string]string{}

	var unitScoutScore = make([]float64, 4)
	if mctx.CalcManager.FleetInfo.UnionFleet != nil {
		unitScoutScore = mctx.CalcManager.FleetInfo.UnionFleet.GetUnionScoutScore(mctx.CalcManager.FleetInfo.AdmiralLevel, 4)
	}

	mctx.PreCalcMap["fleetUlosA"] = fmt.Sprintf("%v", math.Floor(unitScoutScore[0]*100)/100)
	mctx.PreCalcMap["fleetUlosB"] = fmt.Sprintf("%v", math.Floor(unitScoutScore[1]*100)/100)
	mctx.PreCalcMap["fleetUlosC"] = fmt.Sprintf("%v", math.Floor(unitScoutScore[2]*100)/100)
	mctx.PreCalcMap["fleetUlosD"] = fmt.Sprintf("%v", math.Floor(unitScoutScore[3]*100)/100)

	fleets := slices.Sorted(maps.Keys(FleetLabelLookup))
	ships := slices.Sorted(maps.Keys(ShipLabelLookup))
	items := slices.Sorted(maps.Keys(ItemLabelLookup))

	for fleetIdx, f := range mctx.CalcManager.FleetInfo.Fleets {
		scoutScore := fleet.GetScoutScore(f.Ships, mctx.CalcManager.FleetInfo.AdmiralLevel, 4)
		mctx.PreCalcMap["fleet"+fleets[fleetIdx]+"losA"] = fmt.Sprintf("%v", math.Floor(scoutScore[0]*100)/100)
		mctx.PreCalcMap["fleet"+fleets[fleetIdx]+"losB"] = fmt.Sprintf("%v", math.Floor(scoutScore[1]*100)/100)
		mctx.PreCalcMap["fleet"+fleets[fleetIdx]+"losC"] = fmt.Sprintf("%v", math.Floor(scoutScore[2]*100)/100)
		mctx.PreCalcMap["fleet"+fleets[fleetIdx]+"losD"] = fmt.Sprintf("%v", math.Floor(scoutScore[3]*100)/100)

		for shipIdx, s := range mctx.CalcManager.FleetInfo.Fleets[fleetIdx].Ships {
			for itemIdx, i := range mctx.CalcManager.FleetInfo.Fleets[fleetIdx].Ships[shipIdx].Items {
				mctx.PreCalcMap["fleet"+fleets[fleetIdx]+"ship"+ships[shipIdx]+"item"+items[itemIdx]+"levelAlt"] = strconv.Itoa(int(classes.GetProfLevel(i.Level)))
				mctx.PreCalcMap["fleet"+fleets[fleetIdx]+"ship"+ships[shipIdx]+"item"+items[itemIdx]+"equipped"] = func() string {
					if i.Data.Id > 0 {
						return "1"
					}
					return "0"
				}()
			}
			mctx.PreCalcMap["fleet"+fleets[fleetIdx]+"ship"+ships[shipIdx]+"item"+"X"+"levelAlt"] = strconv.Itoa(int(classes.GetProfLevel(s.ExItem.Level)))
			mctx.PreCalcMap["fleet"+fleets[fleetIdx]+"ship"+ships[shipIdx]+"item"+"X"+"equipped"] = func() string {
				if s.ExItem.Data.Id > 0 {
					return "1"
				}
				return "0"
			}()
		}
	}
}

func (mctx *MacroContext) GetMacroPattern(name string) string {
	result, ok := mctx.AccessPatMap[strings.TrimSpace(name)]
	if !ok {
		return ""
	}
	return result
}

func (mctx *MacroContext) accessStandardPatternInit() {
	mctx.AccessPatMap = map[string]string{}

	mctx.AccessPatMap["fleetUfullAirPower"] = "CalcManager.FleetInfo.UnionFleet.FullAirPower"
	mctx.AccessPatMap["fleetUlosA"] = "PreCalcMap.fleetAlosA"
	mctx.AccessPatMap["fleetUlosB"] = "PreCalcMap.fleetAlosB"
	mctx.AccessPatMap["fleetUlosC"] = "PreCalcMap.fleetAlosC"
	mctx.AccessPatMap["fleetUlosD"] = "PreCalcMap.fleetAlosD"
	mctx.AccessPatMap["fleetUspeedKanji"] = "CalcManager.FleetInfo.UnionFleet.FleetSpeed"
	mctx.AccessPatMap["fleetUspeedEn"] = "PreCalcMap.fleetUspeedEn"

	for fleetName, fleetIndex := range FleetLabelLookup {
		// fleetA
		fleetPrefix := "fleet" + fleetName
		// CalcManager.FleetInfo.Fleets.0.
		fleetBase := "CalcManager.FleetInfo.Fleets." + strconv.Itoa(fleetIndex) + "."

		mctx.AccessPatMap[fleetPrefix+"fullAirPower"] = fleetBase + "FullAirPower"
		mctx.AccessPatMap[fleetPrefix+"losA"] = "PreCalcMap." + fleetPrefix + "losA"
		mctx.AccessPatMap[fleetPrefix+"losB"] = "PreCalcMap." + fleetPrefix + "losB"
		mctx.AccessPatMap[fleetPrefix+"losC"] = "PreCalcMap." + fleetPrefix + "losC"
		mctx.AccessPatMap[fleetPrefix+"losD"] = "PreCalcMap." + fleetPrefix + "losD"
		mctx.AccessPatMap[fleetPrefix+"speedKanji"] = fleetBase + "FleetSpeed"
		mctx.AccessPatMap[fleetPrefix+"speedEn"] = "PreCalcMap." + "fleet" + fleetName + "speed"

		for shipName, shipIndex := range ShipLabelLookup {
			// fleetAShipA
			shipPrefix := fleetPrefix + "ship" + shipName
			// CalcManager.FleetInfo.Fleets.0.Ships.0.
			shipBase := fleetBase + "Ships." + strconv.Itoa(shipIndex) + "."

			mctx.AccessPatMap[shipPrefix+"nameJp"] = shipBase + "Data.Name"
			mctx.AccessPatMap[shipPrefix+"nameEn"] = "PreCalcMap." + shipPrefix + "nameEn"
			mctx.AccessPatMap[shipPrefix+"Level"] = shipBase + "Level"
			mctx.AccessPatMap[shipPrefix+"fullAirPower"] = shipBase + "fullAirPower"
			mctx.AccessPatMap[shipPrefix+"displayStatusHp"] = shipBase + "DisplayStatus.HP"
			mctx.AccessPatMap[shipPrefix+"displayStatusFirePower"] = shipBase + "DisplayStatus.FirePower"
			mctx.AccessPatMap[shipPrefix+"displayStatusArmor"] = shipBase + "DisplayStatus.Armor"
			mctx.AccessPatMap[shipPrefix+"displayStatusTorpedo"] = shipBase + "DisplayStatus.Torpedo"
			mctx.AccessPatMap[shipPrefix+"displayStatusAvoid"] = shipBase + "DisplayStatus.Avoid"
			mctx.AccessPatMap[shipPrefix+"displayStatusAsw"] = shipBase + "DisplayStatus.Asw"
			mctx.AccessPatMap[shipPrefix+"displayStatusLos"] = shipBase + "DisplayStatus.Los"
			mctx.AccessPatMap[shipPrefix+"displayStatusLuck"] = shipBase + "DisplayStatus.Luck"
			mctx.AccessPatMap[shipPrefix+"displayStatusRange"] = shipBase + "DisplayStatus.Range"
			mctx.AccessPatMap[shipPrefix+"displayStatusAccuracy"] = shipBase + "DisplayStatus.Accuracy"

			for itemName, itemIndex := range ItemLabelLookup {
				// fleetAshipAitemA
				itemPrefix := shipPrefix + "item" + itemName
				// CalcManager.FleetInfo.Fleets.0.Ships.0.Items.0.
				itemBase := shipBase + "Items." + strconv.Itoa(itemIndex) + "."
				if itemName == "X" {
					itemBase = shipBase + "ExItem."
				}

				mctx.AccessPatMap[itemPrefix+"nameJp"] = itemBase + "Data.Name"
				mctx.AccessPatMap[itemPrefix+"nameEn"] = "PreCalcMap." + itemPrefix + "nameEn"
				mctx.AccessPatMap[itemPrefix+"levelAlt"] = "PreCalcMap." + itemPrefix + "levelAlt"
				mctx.AccessPatMap[itemPrefix+"level"] = itemBase + "level"
				mctx.AccessPatMap[itemPrefix+"id"] = itemBase + "Data.Id"
				mctx.AccessPatMap[itemPrefix+"typeid"] = itemBase + "Data.ApiTypeId"
				mctx.AccessPatMap[itemPrefix+"apitypeid"] = itemBase + "Data.ApiTypeId"
				mctx.AccessPatMap[itemPrefix+"iconid"] = itemBase + "Data.IconId"
				mctx.AccessPatMap[itemPrefix+"icontypeid"] = itemBase + "Data.IconId"
				mctx.AccessPatMap[itemPrefix+"equipped"] = "PreCalcMap." + itemPrefix + "equipped"

			}
		}
	}
}

func (mctx *MacroContext) AccessKeyStr() string {
	return strings.Join(slices.Sorted(maps.Keys(mctx.AccessPatMap)), ",")
}
