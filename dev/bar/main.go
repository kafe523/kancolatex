package main

import (
	"fmt"
	"os"

	"github.com/kafe523/kancolatex/internal/access"
	"github.com/kafe523/kancolatex/internal/macro"
)

func main() {
	c, err := os.ReadFile("sample.json")
	if err != nil {
		panic(err)
	}

	mctx := macro.MacroContext{}
	err = mctx.Populate(&c)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", mctx.CalcManager.FleetInfo.Fleets[0].Ships[0].Data.Name)
	fmt.Printf("%#v\n", access.Access(&mctx, access.PatternTokenizer("CalcManager.FleetInfo.Fleets.0.Ships.0.Data.Name")))

	// fmt.Printf("%#v\n", access.Access(&mctx, access.PatternTokenizer("CalcManager.FleetInfo.Fleets.0.Ships.0.Items.0.Level")))

	fmt.Printf("%#v\n", mctx.PreCalcMap["fleetAshipAitemAlevelAlt"])
	fmt.Printf("%#v\n", access.Access(&mctx, access.PatternTokenizer("PreCalcMap.fleetAshipAitemAlevelAlt")))
}
