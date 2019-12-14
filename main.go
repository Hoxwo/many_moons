package main

import "fmt"
import "time"
import "math/rand"
import "strings"
import planet "many_moons/planet"
import ship "many_moons/ship"
import civilization "many_moons/civilization"
import ui "github.com/gizak/termui/v3"

func main() {
	// set the current year to 0
	currentyear := 0
	//for randos
	rand.Seed(time.Now().UTC().UnixNano())
	// array of civilizations 
	civilizations := make([]*civilization.Civilization, 0)
	// array of planets
	coinPrices := make(map[string]float64)
	//index of selected planet
	selected := 0
	// current state
	state := 1 // 1 - watching, 2 - buying, 3 - selling
	//last player quantity input
	lastPlayerInput := 0

	//All colors
	//termui.ColorCyan
	//termui.ColorGreen
	//termui.ColorYellow
	//termui.ColorWhite	
	//termui.ColorBlue
	//termui.ColorMagenta
	//termui.ColorRed
	
	// set up civilizations
	// 	       {civilization, color, atk, def, nav, gov, tec, res, shipsavail, maxshipsavail, shiptimer, colonizationtime, adsli, ademsli, eisli}
	c0 := civilization.New("Balanced",   "cyan",    2, 2, 2, 2, 2, 2, 1, 1, 30, 30, 0, 0, 0)
	c1 := civilization.New("Warlike",    "red",     3, 2, 2, 2, 1, 1, 1, 1, 30, 30, 0, 0, 0)
	c2 := civilization.New("Defensive",  "magenta", 2, 3, 2, 2, 1, 1, 1, 1, 30, 30, 0, 0, 0)
	c3 := civilization.New("Explorer",   "green",   1, 1, 3, 2, 2, 2, 1, 1, 30, 30, 0, 0, 0)
	c4 := civilization.New("Autocracy",  "blue",    2, 1, 2, 3, 1, 2, 1, 1, 30, 30, 0, 0, 0)
	c5 := civilization.New("Technology", "yellow",  1, 2, 2, 1, 3, 2, 1, 1, 30, 30, 0, 0, 0)				

	//add them to master list
	civilizations = append(civilizations, &c0, &c1, &c2, &c3, &c4, &c5)	


	//select first planet
	selected = 0

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	
	planetmap := ui.NewCanvas()
	planetmap.SetRect(0, 78, 0, 34)
	planetmap.SetLine(image.Pt(0, 0), image.Pt(10, 20), ui.ColorWhite)

	datapane := ui.NewCanvas()
	datapane.SetRect(78, 126, 0, 34)
	datapane.SetLine(image.Pt(0, 0), image.Pt(10, 20), ui.ColorWhite)


	ui.Render(planetmap, datapane)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
