package main

import "fmt"
import "log"
import "math/rand"
import "time"
//import "strings"
import planet "many_moons/planet"
//import ship "many_moons/ship"
import civilization "many_moons/civilization"
import ui "github.com/gizak/termui/v3"
import "github.com/gizak/termui/v3/widgets"

func main() {
	// set the current year to 0
	currentyear := 5000
	// for randos
	rand.Seed(time.Now().UTC().UnixNano())
	// array of civilizations 
	civilizations := make([]*civilization.Civilization, 0)
	// array of planets
	planets := make([]*planet.Planet, 0)	
	// selected planet 
	selectedplanet := 0 

	//All colors
	//termui.ColorCyan
	//termui.ColorGreen
	//termui.ColorYellow
	//termui.ColorWhite	
	//termui.ColorBlue
	//termui.ColorMagenta
	//termui.ColorRed
	
	// set up civilizations
	// 	       {civilization, color, atk, def, nav, gov, tec, res, shipsavail, maxshipsavail, shiptimer, maxshiptimer, colonizationtime, adsli, ademsli, eisli}
	c0 := civilization.New("Balanced",   "cyan",    2, 2, 2, 2, 2, 2, 0, 1, 0, 30, 30, 1, 2, -1)
	c1 := civilization.New("Warlike",    "red",     3, 2, 2, 2, 1, 1, 1, 1, 0, 30, 30, 0, 0, 0)
	c2 := civilization.New("Defensive",  "magenta", 2, 3, 2, 2, 1, 1, 1, 1, 0, 30, 30, 0, 0, 0)
	c3 := civilization.New("Explorer",   "green",   1, 1, 3, 2, 2, 2, 1, 1, 0, 30, 30, 0, 0, 0)
	//c4 := civilization.New("Autocracy",  "blue",    2, 1, 2, 3, 1, 2, 1, 1, 30, 30, 0, 0, 0)
	//c5 := civilization.New("Technology", "yellow",  1, 2, 2, 1, 3, 2, 1, 1, 30, 30, 0, 0, 0)				

	//add them to master list
	civilizations = append(civilizations, &c0, &c1, &c2, &c3) //&c4, &c5)	

	//generate all planets, add them to master list
	planets = GenerateSpace(planets)

	//select first planet
	//selected = 0

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	
	//render all planets - this looks horrible
	planet0box := widgets.NewParagraph()
	planet0box.Text = fmt.Sprintf("\n  %s  \n", planets[0].Appearance())
	planet0box.Border = false
	planet0box.SetRect((planets[0].Xcoord()-4), (planets[0].Ycoord()-2), (planets[0].Xcoord()+4), (planets[0].Ycoord()+2))
	planet0box.TextStyle.Fg = ui.ColorWhite

	planet1box := widgets.NewParagraph()
	planet1box.Text = fmt.Sprintf("\n  %s  \n", planets[1].Appearance())
	planet1box.Border = false
	planet1box.SetRect((planets[1].Xcoord()-4), (planets[1].Ycoord()-2), (planets[1].Xcoord()+4), (planets[1].Ycoord()+2))
	planet1box.TextStyle.Fg = ui.ColorWhite

	planet2box := widgets.NewParagraph()
	planet2box.Text = fmt.Sprintf("\n  %s  \n", planets[2].Appearance())
	planet2box.Border = false
	planet2box.SetRect((planets[2].Xcoord()-4), (planets[2].Ycoord()-2), (planets[2].Xcoord()+4), (planets[2].Ycoord()+2))
	planet2box.TextStyle.Fg = ui.ColorWhite

	planet3box := widgets.NewParagraph()
	planet3box.Text = fmt.Sprintf("\n  %s  \n", planets[3].Appearance())
	planet3box.Border = false
	planet3box.SetRect((planets[3].Xcoord()-4), (planets[3].Ycoord()-2), (planets[3].Xcoord()+4), (planets[3].Ycoord()+2))
	planet3box.TextStyle.Fg = ui.ColorWhite

	planet4box := widgets.NewParagraph()
	planet4box.Text = fmt.Sprintf("\n  %s  \n", planets[4].Appearance())
	planet4box.Border = false
	planet4box.SetRect((planets[4].Xcoord()-4), (planets[4].Ycoord()-2), (planets[4].Xcoord()+4), (planets[4].Ycoord()+2))
	planet4box.TextStyle.Fg = ui.ColorWhite

	planet5box := widgets.NewParagraph()
	planet5box.Text = fmt.Sprintf("\n  %s  \n", planets[5].Appearance())
	planet5box.Border = false
	planet5box.SetRect((planets[5].Xcoord()-4), (planets[5].Ycoord()-2), (planets[5].Xcoord()+4), (planets[5].Ycoord()+2))
	planet5box.TextStyle.Fg = ui.ColorWhite

	planet6box := widgets.NewParagraph()
	planet6box.Text = fmt.Sprintf("\n  %s  \n", planets[6].Appearance())
	planet6box.Border = false
	planet6box.SetRect((planets[6].Xcoord()-4), (planets[6].Ycoord()-2), (planets[6].Xcoord()+4), (planets[6].Ycoord()+2))
	planet6box.TextStyle.Fg = ui.ColorWhite

	planet7box := widgets.NewParagraph()
	planet7box.Text = fmt.Sprintf("\n  %s  \n", planets[7].Appearance())
	planet7box.Border = false
	planet7box.SetRect((planets[7].Xcoord()-4), (planets[7].Ycoord()-2), (planets[7].Xcoord()+4), (planets[7].Ycoord()+2))
	planet7box.TextStyle.Fg = ui.ColorWhite

	planet8box := widgets.NewParagraph()
	planet8box.Text = fmt.Sprintf("\n  %s  \n", planets[8].Appearance())
	planet8box.Border = false
	planet8box.SetRect((planets[8].Xcoord()-4), (planets[8].Ycoord()-2), (planets[8].Xcoord()+4), (planets[8].Ycoord()+2))
	planet8box.TextStyle.Fg = ui.ColorWhite
	
	planet9box := widgets.NewParagraph()
	planet9box.Text = fmt.Sprintf("\n  %s  \n", planets[0].Appearance())
	planet9box.Border = false
	planet9box.SetRect((planets[9].Xcoord()-4), (planets[9].Ycoord()-2), (planets[9].Xcoord()+4), (planets[9].Ycoord()+2))
	planet9box.TextStyle.Fg = ui.ColorWhite
	
	planet10box := widgets.NewParagraph()
	planet10box.Text = fmt.Sprintf("\n  %s  \n", planets[10].Appearance())
	planet10box.Border = false
	planet10box.SetRect((planets[10].Xcoord()-4), (planets[10].Ycoord()-2), (planets[10].Xcoord()+4), (planets[10].Ycoord()+2))
	planet10box.TextStyle.Fg = ui.ColorWhite

	planet11box := widgets.NewParagraph()
	planet11box.Text = fmt.Sprintf("\n  %s  \n", planets[11].Appearance())
	planet11box.Border = false
	planet11box.SetRect((planets[11].Xcoord()-4), (planets[11].Ycoord()-2), (planets[11].Xcoord()+4), (planets[11].Ycoord()+2))
	planet11box.TextStyle.Fg = ui.ColorWhite

	planet12box := widgets.NewParagraph()
	planet12box.Text = fmt.Sprintf("\n  %s  \n", planets[12].Appearance())
	planet12box.Border = false
	planet12box.SetRect((planets[12].Xcoord()-4), (planets[12].Ycoord()-2), (planets[12].Xcoord()+4), (planets[12].Ycoord()+2))
	planet12box.TextStyle.Fg = ui.ColorWhite

	planet13box := widgets.NewParagraph()
	planet13box.Text = fmt.Sprintf("\n  %s  \n", planets[13].Appearance())
	planet13box.Border = false
	planet13box.SetRect((planets[13].Xcoord()-4), (planets[13].Ycoord()-2), (planets[13].Xcoord()+4), (planets[13].Ycoord()+2))
	planet13box.TextStyle.Fg = ui.ColorWhite

	planet14box := widgets.NewParagraph()
	planet14box.Text = fmt.Sprintf("\n  %s  \n", planets[14].Appearance())
	planet14box.Border = false
	planet14box.SetRect((planets[14].Xcoord()-4), (planets[14].Ycoord()-2), (planets[14].Xcoord()+4), (planets[14].Ycoord()+2))
	planet14box.TextStyle.Fg = ui.ColorWhite
	
	planet15box := widgets.NewParagraph()
	planet15box.Text = fmt.Sprintf("\n  %s  \n", planets[15].Appearance())
	planet15box.Border = false
	planet15box.SetRect((planets[15].Xcoord()-4), (planets[15].Ycoord()-2), (planets[15].Xcoord()+4), (planets[15].Ycoord()+2))
	planet15box.TextStyle.Fg = ui.ColorWhite

	//render map and stat panes
	planetmap := widgets.NewParagraph()
	planetmap.Title = CurrentYearText(currentyear)
	planetmap.Text = fmt.Sprintf("%d", len(planets))
	planetmap.SetRect(0, 0, 78, 34)
	planetmap.BorderStyle.Fg = ui.ColorWhite

	civstatspane := widgets.NewParagraph()
	civstatspane.Title = "Civilization Stats"
	civstatspane.Text = CivilizationStatsText(civilizations)
	civstatspane.SetRect(78, 0, 116, 9)
	civstatspane.BorderStyle.Fg = ui.ColorBlue

	sliderspane := widgets.NewParagraph()
	sliderspane.Title = "Policy Sliders"
	sliderspane.Text = PlayerSlidersText(civilizations)
	sliderspane.SetRect(78, 9, 116, 14)
	sliderspane.BorderStyle.Fg = ui.ColorCyan

	playercivstatspane := widgets.NewParagraph()
	playercivstatspane.Title = fmt.Sprintf("%s", civilizations[0].Name())
	playercivstatspane.Text = PlayerCivilizationStatsText(civilizations)
	playercivstatspane.SetRect(78, 14, 116, 19)
	playercivstatspane.BorderStyle.Fg = ui.ColorCyan

	selectplanetinfopane := widgets.NewParagraph()
	selectplanetinfopane.Title = "Planet Info"
	selectplanetinfopane.Text = SelectedPlanetText(planets, selectedplanet)
	selectplanetinfopane.SetRect(78, 19, 116, 34)
	selectplanetinfopane.BorderStyle.Fg = ui.ColorBlue

	ui.Render(planetmap, civstatspane, sliderspane, playercivstatspane, selectplanetinfopane, planet0box, planet1box, planet2box, planet3box, planet4box, 		planet5box, planet6box, planet7box, planet8box, planet9box, planet10box, planet11box, planet12box, planet13box, planet14box, planet15box)

	tickerCount := 1
	tickerCount++
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-ticker:
			currentyear = currentyear + 1
			EveryCivilizationCountdown(civilizations)
			
			//refactor this
			planet0box := widgets.NewParagraph()
			planet0box.Text = fmt.Sprintf("\n  %s  \n", planets[0].Appearance())
			planet0box.Border = false
			planet0box.SetRect((planets[0].Xcoord()-4), (planets[0].Ycoord()-2), (planets[0].Xcoord()+4), (planets[0].Ycoord()+2))
			planet0box.TextStyle.Fg = ui.ColorWhite

			planet1box := widgets.NewParagraph()
			planet1box.Text = fmt.Sprintf("\n  %s  \n", planets[1].Appearance())
			planet1box.Border = false
			planet1box.SetRect((planets[1].Xcoord()-4), (planets[1].Ycoord()-2), (planets[1].Xcoord()+4), (planets[1].Ycoord()+2))
			planet1box.TextStyle.Fg = ui.ColorWhite

			planet2box := widgets.NewParagraph()	
			planet2box.Text = fmt.Sprintf("\n  %s  \n", planets[2].Appearance())
			planet2box.Border = false
			planet2box.SetRect((planets[2].Xcoord()-4), (planets[2].Ycoord()-2), (planets[2].Xcoord()+4), (planets[2].Ycoord()+2))
			planet2box.TextStyle.Fg = ui.ColorWhite

			planet3box := widgets.NewParagraph()
			planet3box.Text = fmt.Sprintf("\n  %s  \n", planets[3].Appearance())
			planet3box.Border = false
			planet3box.SetRect((planets[3].Xcoord()-4), (planets[3].Ycoord()-2), (planets[3].Xcoord()+4), (planets[3].Ycoord()+2))
			planet3box.TextStyle.Fg = ui.ColorWhite
		
			planet4box := widgets.NewParagraph()
			planet4box.Text = fmt.Sprintf("\n  %s  \n", planets[4].Appearance())
			planet4box.Border = false
			planet4box.SetRect((planets[4].Xcoord()-4), (planets[4].Ycoord()-2), (planets[4].Xcoord()+4), (planets[4].Ycoord()+2))
			planet4box.TextStyle.Fg = ui.ColorWhite

			planet5box := widgets.NewParagraph()
			planet5box.Text = fmt.Sprintf("\n  %s  \n", planets[5].Appearance())
			planet5box.Border = false
			planet5box.SetRect((planets[5].Xcoord()-4), (planets[5].Ycoord()-2), (planets[5].Xcoord()+4), (planets[5].Ycoord()+2))
			planet5box.TextStyle.Fg = ui.ColorWhite

			planet6box := widgets.NewParagraph()
			planet6box.Text = fmt.Sprintf("\n  %s  \n", planets[6].Appearance())
			planet6box.Border = false
			planet6box.SetRect((planets[6].Xcoord()-4), (planets[6].Ycoord()-2), (planets[6].Xcoord()+4), (planets[6].Ycoord()+2))
			planet6box.TextStyle.Fg = ui.ColorWhite

			planet7box := widgets.NewParagraph()
			planet7box.Text = fmt.Sprintf("\n  %s  \n", planets[7].Appearance())
			planet7box.Border = false
			planet7box.SetRect((planets[7].Xcoord()-4), (planets[7].Ycoord()-2), (planets[7].Xcoord()+4), (planets[7].Ycoord()+2))
			planet7box.TextStyle.Fg = ui.ColorWhite

			planet8box := widgets.NewParagraph()
			planet8box.Text = fmt.Sprintf("\n  %s  \n", planets[8].Appearance())
			planet8box.Border = false
			planet8box.SetRect((planets[8].Xcoord()-4), (planets[8].Ycoord()-2), (planets[8].Xcoord()+4), (planets[8].Ycoord()+2))
			planet8box.TextStyle.Fg = ui.ColorWhite
	
			planet9box := widgets.NewParagraph()
			planet9box.Text = fmt.Sprintf("\n  %s  \n", planets[0].Appearance())
			planet9box.Border = false
			planet9box.SetRect((planets[9].Xcoord()-4), (planets[9].Ycoord()-2), (planets[9].Xcoord()+4), (planets[9].Ycoord()+2))
			planet9box.TextStyle.Fg = ui.ColorWhite
	
			planet10box := widgets.NewParagraph()
			planet10box.Text = fmt.Sprintf("\n  %s  \n", planets[10].Appearance())
			planet10box.Border = false
			planet10box.SetRect((planets[10].Xcoord()-4), (planets[10].Ycoord()-2), (planets[10].Xcoord()+4), (planets[10].Ycoord()+2))
			planet10box.TextStyle.Fg = ui.ColorWhite

			planet11box := widgets.NewParagraph()
			planet11box.Text = fmt.Sprintf("\n  %s  \n", planets[11].Appearance())
			planet11box.Border = false
			planet11box.SetRect((planets[11].Xcoord()-4), (planets[11].Ycoord()-2), (planets[11].Xcoord()+4), (planets[11].Ycoord()+2))
			planet11box.TextStyle.Fg = ui.ColorWhite

			planet12box := widgets.NewParagraph()
			planet12box.Text = fmt.Sprintf("\n  %s  \n", planets[12].Appearance())
			planet12box.Border = false
			planet12box.SetRect((planets[12].Xcoord()-4), (planets[12].Ycoord()-2), (planets[12].Xcoord()+4), (planets[12].Ycoord()+2))
			planet12box.TextStyle.Fg = ui.ColorWhite

			planet13box := widgets.NewParagraph()
			planet13box.Text = fmt.Sprintf("\n  %s  \n", planets[13].Appearance())
			planet13box.Border = false
			planet13box.SetRect((planets[13].Xcoord()-4), (planets[13].Ycoord()-2), (planets[13].Xcoord()+4), (planets[13].Ycoord()+2))
			planet13box.TextStyle.Fg = ui.ColorWhite

			planet14box := widgets.NewParagraph()
			planet14box.Text = fmt.Sprintf("\n  %s  \n", planets[14].Appearance())
			planet14box.Border = false
			planet14box.SetRect((planets[14].Xcoord()-4), (planets[14].Ycoord()-2), (planets[14].Xcoord()+4), (planets[14].Ycoord()+2))
			planet14box.TextStyle.Fg = ui.ColorWhite
	
			planet15box := widgets.NewParagraph()
			planet15box.Text = fmt.Sprintf("\n  %s  \n", planets[15].Appearance())	
			planet15box.Border = false
			planet15box.SetRect((planets[15].Xcoord()-4), (planets[15].Ycoord()-2), (planets[15].Xcoord()+4), (planets[15].Ycoord()+2))
			planet15box.TextStyle.Fg = ui.ColorWhite

			planetmap := widgets.NewParagraph()
			planetmap.Title = CurrentYearText(currentyear)
			planetmap.Text = ""
			planetmap.SetRect(0, 0, 78, 34)
			planetmap.BorderStyle.Fg = ui.ColorWhite

			civstatspane := widgets.NewParagraph()
			civstatspane.Title = "Civilization Stats"
			civstatspane.Text = CivilizationStatsText(civilizations)
			civstatspane.SetRect(78, 0, 116, 9)
			civstatspane.BorderStyle.Fg = ui.ColorBlue

			sliderspane := widgets.NewParagraph()
			sliderspane.Title = "Policy Sliders"
			sliderspane.Text = PlayerSlidersText(civilizations)
			sliderspane.SetRect(78, 9, 116, 14)
			sliderspane.BorderStyle.Fg = ui.ColorCyan

			playercivstatspane := widgets.NewParagraph()
			playercivstatspane.Title = fmt.Sprintf("%s", civilizations[0].Name())
			playercivstatspane.Text = PlayerCivilizationStatsText(civilizations)
			playercivstatspane.SetRect(78, 14, 116, 19)
			playercivstatspane.BorderStyle.Fg = ui.ColorCyan

			selectplanetinfopane := widgets.NewParagraph()
			selectplanetinfopane.Title = "Planet Info"
			selectplanetinfopane.Text = SelectedPlanetText(planets, selectedplanet)
			selectplanetinfopane.SetRect(78, 19, 116, 34)
			selectplanetinfopane.BorderStyle.Fg = ui.ColorBlue
			//	

			ui.Render(planetmap, civstatspane, sliderspane, playercivstatspane, selectplanetinfopane, planet0box, planet1box, planet2box, 				planet3box, planet4box, planet5box, planet6box, planet7box, planet8box, planet9box, planet10box, planet11box, planet12box, 				planet13box, planet14box, planet15box)
		}
	}
}

func CivilizationStatsText(civilizations []*civilization.Civilization) string {
	civilizationstatstext := "\n          ATK DEF NAV GOV TEC RES"
	for _, c := range civilizations {
		civilizationstatstext = civilizationstatstext + fmt.Sprintf("\n%s:%s%d   %d   %d   %d   %d   %d", c.Name(), SpacePadding(len(c.Name())), c.Attack(), c.Defense(), c.Navigation(), c.Government(), c.Technology(), c.Resources())
	}
	return civilizationstatstext
}

func PlayerSlidersText(civilizations []*civilization.Civilization) string {
	atkdefval := civilizations[0].Atkdefslider()
	autdemval := civilizations[0].Autdemslider()
	envindval := civilizations[0].Envindslider()
	playersliderstext := ""
	playersliderstext = playersliderstext + fmt.Sprintf("DEF %s ATK %s", SliderValueToString(atkdefval), SliderValueToMessage(atkdefval, "DEF", "ATK"))
	playersliderstext = playersliderstext + fmt.Sprintf("\nAUT %s DEM %s", SliderValueToString(autdemval), SliderValueToMessage(autdemval, "GOV", "NAV"))
	playersliderstext = playersliderstext + fmt.Sprintf("\nENV %s IND %s", SliderValueToString(envindval), SliderValueToMessage(envindval, "RES", "TEC"))
	
	return playersliderstext
}

func SpacePadding(stringLength int) string { 
	padding := "  "
	for i := 0;  i <= (8 - stringLength); i++ {
                padding = padding + " "
        } 
 	return padding
}

func SliderValueToString(slidervalue int) string {
	if(slidervalue == -2) {
        	return "<|====>"  	
	} else if(slidervalue == -1) {
        	return "<=|===>"  	
	} else if(slidervalue == 0) {
        	return "<==|==>"  	
	} else if(slidervalue == 1) {
        	return "<===|=>"  	
	} else if(slidervalue == 2) {
        	return "<====|>"  	
	} else {
		return "ERROR"
	}
}

func SliderValueToMessage(slidervalue int, firsttrait string, secondtrait string) string {
	if(slidervalue == -2) {
        	return fmt.Sprintf("[+2 to %s](fg:blue) [-2 to %s](fg:red)", firsttrait, secondtrait)  	
	} else if(slidervalue == -1) {
        	return fmt.Sprintf("[+1 to %s](fg:blue) [-1 to %s](fg:red)", firsttrait, secondtrait)	
	} else if(slidervalue == 0) {
        	return fmt.Sprintf("[+0 to %s](fg:white) [+0 to %s](fg:white)", firsttrait, secondtrait)	
	} else if(slidervalue == 1) {
        	return fmt.Sprintf("[-1 to %s](fg:red) [+1 to %s](fg:blue)", firsttrait, secondtrait)
	} else if(slidervalue == 2) {
        	return fmt.Sprintf("[-2 to %s](fg:red) [+2 to %s](fg:blue)", firsttrait, secondtrait)	
	} else {
		return "ERROR"
	}
}

func CurrentYearText(currentyear int) string {
	return fmt.Sprintf("Current year: %d", currentyear)
}

func EveryCivilizationCountdown(civilizations []*civilization.Civilization) {
	for _, c := range civilizations {
		if(c.Shiptimer() < c.Maxshiptimer()) { 
			c.SetShiptimer(c.Shiptimer() + 1)
		} else {
			c.SetShiptimer(0)
			if(c.Shipsavailable() < c.Maxshipsavailable()) {
				c.SetShipsavailable(c.Shipsavailable() + 1)			
			}
		}
	}
}

func PlayerCivilizationStatsText(civilizations []*civilization.Civilization) string {
	shipsavailable := civilizations[0].Shipsavailable()
	maxshipsavailable := civilizations[0].Maxshipsavailable()
	shiptimer := civilizations[0].Shiptimer()
	maxshiptimer := civilizations[0].Maxshiptimer()
	playercivstatstext := ""
	playercivstatstext = playercivstatstext + fmt.Sprintf("%d/%d ships available", shipsavailable, maxshipsavailable)
	playercivstatstext = playercivstatstext + fmt.Sprintf("\n%d years until next ship ready", (maxshiptimer - shiptimer))

	return playercivstatstext
}

func SelectedPlanetText(planets []*planet.Planet, selectedplanet int) string {
	name := planets[selectedplanet].Name()
	planettype := planets[selectedplanet].Planettype()
	occupied := planets[selectedplanet].Occupied()
	resources := planets[selectedplanet].Resources()
	xcoord := planets[selectedplanet].Xcoord()
	ycoord := planets[selectedplanet].Ycoord()
	selectedPlanetText := ""
	selectedPlanetText = selectedPlanetText + fmt.Sprintf("Planet name: %s", name)
	selectedPlanetText = selectedPlanetText + fmt.Sprintf("\nPlanet type: %s", planettype)
	selectedPlanetText = selectedPlanetText + fmt.Sprintf("\nPlanet occupied by: %s", occupied)
	selectedPlanetText = selectedPlanetText + fmt.Sprintf("\nPlanet resource value: %d", resources)
	selectedPlanetText = selectedPlanetText + fmt.Sprintf("\nPlanet coordinates: (%d,%d)", xcoord, ycoord)
	selectedPlanetText = selectedPlanetText + fmt.Sprintf("\nDistance from your base planet: 0") //TO DO
	if(planettype == "Ice Giant") {
		selectedPlanetText = selectedPlanetText + fmt.Sprintf("\n[NAV or TEC level 4 required](fg:red)")
		selectedPlanetText = selectedPlanetText + fmt.Sprintf("\n[to colonize this planet](fg:red)")
	} else if(planettype == "Gas Giant") {
		selectedPlanetText = selectedPlanetText + fmt.Sprintf("\n[NAV or TEC level 7 required](fg:red)")
		selectedPlanetText = selectedPlanetText + fmt.Sprintf("\n[to colonize this planet](fg:red)")	
	} else {
		selectedPlanetText = selectedPlanetText + fmt.Sprintf("\n")
		selectedPlanetText = selectedPlanetText + fmt.Sprintf("\n")	
	}
	selectedPlanetText = selectedPlanetText + fmt.Sprintf("\n")
	selectedPlanetText = selectedPlanetText + fmt.Sprintf("\n")
	selectedPlanetText = selectedPlanetText + fmt.Sprintf("\n")
	selectedPlanetText = selectedPlanetText + fmt.Sprintf("\n[<, >] to select planet")
	selectedPlanetText = selectedPlanetText + fmt.Sprintf("\n[s] to send ship")

	return selectedPlanetText
}

func GenerateSpace(planets []*planet.Planet) []*planet.Planet {
	numberplanets := 15 //random(16, 20)
	xcoordsused := make([]int, 0)
	ycoordsused := make([]int, 0) 
	for i := 0;  i <= numberplanets; i++ {
		planettype := random(1,5) // 1 - small rocky planet, 2 - large rocky planet, 3 - ice giant, 4 - gas giant
		
		//don't duplicate any coordinates
		xcoord := 0
		for xcoord == 0 || contains(xcoordsused, xcoord) {
			xcoord = random(3,74)
			if(contains(xcoordsused, xcoord) == false) {
				xcoordsused = append(xcoordsused, xcoord)
				break
			}		
		}
		
		ycoord := 0
		for ycoord == 0 || contains(ycoordsused, ycoord) {
			ycoord = random(3, 31)
			if(contains(ycoordsused, ycoord) == false) {
				ycoordsused = append(ycoordsused, ycoord)
				break
			}		
		}
	
		if(planettype == 1) {
			p0 := planet.New("ST", "o",   1, "uncolonized", "Small Terrestrial", xcoord, ycoord)
			planets = append(planets, &p0)
		} else if(planettype == 2) {
			p0 := planet.New("LT", "O",   2, "uncolonized", "Large Terrestrial", xcoord, ycoord)
			planets = append(planets, &p0)		
		} else if(planettype == 3) {
			p0 := planet.New("IG", "(o)",   2, "uncolonized", "Ice Giant", xcoord, ycoord)
			planets = append(planets, &p0)
		} else {
			p0 := planet.New("GG", "(O)",   3, "uncolonized", "Gas Giant", xcoord, ycoord)
			planets = append(planets, &p0)
		}
	}
	
        return planets
}

//6 sided die = random(1,7)
func random(min, max int) int {
    return rand.Intn(max - min) + min
}

func isThreeAway(numtocheck int, othernum int) bool {
	if((numtocheck > othernum && (numtocheck - 3) > othernum) || (numtocheck < othernum && (othernum - 3) < numtocheck)) {
		return true        
	} else {
		return false
	}
}

func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

