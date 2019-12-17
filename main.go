package main

import "fmt"
import "log"
import "math/rand"
import "time"
import "image"
import planet "many_moons/planet"
//import ship "many_moons/ship"
import civilization "many_moons/civilization"
import ui "github.com/gizak/termui/v3"
import "github.com/gizak/termui/v3/widgets"

func main() {
	// set to the starting year
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
	
	//render map and stat panes
	planetmapframe := widgets.NewParagraph()
	planetmapframe.Title = CurrentYearText(currentyear)
	planetmapframe.SetRect(0, 0, 78, 34)
	planetmapframe.BorderStyle.Fg = ui.ColorWhite

	//canvas-style planet map
	planetmap := ui.NewCanvas()
	planetmap.SetRect(1, 1, 76, 33)
	for i := 0; i < len(planets); i++ {
		pt1 := image.Pt(((planets[i].Xcoord()*2)+0), ((planets[i].Ycoord()*4)+0))
		pt2 := image.Pt(((planets[i].Xcoord()*2)+1), ((planets[i].Ycoord()*4)+0))
		pt3 := image.Pt(((planets[i].Xcoord()*2)+0), ((planets[i].Ycoord()*4)+1))
		pt4 := image.Pt(((planets[i].Xcoord()*2)+1), ((planets[i].Ycoord()*4)+1))
		planetmap.SetPoint(pt1, ui.ColorWhite)
		planetmap.SetPoint(pt2, ui.ColorWhite)
		planetmap.SetPoint(pt3, ui.ColorWhite)
		planetmap.SetPoint(pt4, ui.ColorWhite)
	}	
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

	ui.Render(planetmapframe, planetmap, civstatspane, sliderspane, playercivstatspane, selectplanetinfopane)

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
			
			//render map and stat panes
			planetmapframe := widgets.NewParagraph()
			planetmapframe.Title = CurrentYearText(currentyear)
			planetmapframe.SetRect(0, 0, 78, 34)
			planetmapframe.BorderStyle.Fg = ui.ColorWhite

			//canvas-style planet map
			planetmap := ui.NewCanvas()
			planetmap.SetRect(2, 1, 75, 33)
			for i := 0; i < len(planets); i++ {
				pt1 := image.Pt(((planets[i].Xcoord()*2)+0), ((planets[i].Ycoord()*4)+0))
				pt2 := image.Pt(((planets[i].Xcoord()*2)+1), ((planets[i].Ycoord()*4)+0))
				pt3 := image.Pt(((planets[i].Xcoord()*2)+0), ((planets[i].Ycoord()*4)+1))
				pt4 := image.Pt(((planets[i].Xcoord()*2)+1), ((planets[i].Ycoord()*4)+1))
				planetmap.SetPoint(pt1, ui.ColorWhite)
				planetmap.SetPoint(pt2, ui.ColorWhite)
				planetmap.SetPoint(pt3, ui.ColorWhite)
				planetmap.SetPoint(pt4, ui.ColorWhite)
			}	
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

			ui.Render(planetmapframe, planetmap, civstatspane, sliderspane, playercivstatspane, selectplanetinfopane)
		}
	}
}

func CivilizationStatsText(civilizations []*civilization.Civilization) string {
	civilizationstatstext := "\n          ATK DEF NAV GOV TEC RES"
	for _, c := range civilizations {
		if(c.Color() == "cyan"){
		civilizationstatstext = civilizationstatstext + fmt.Sprintf("[\n%s:%s%d   %d   %d   %d   %d   %d](fg:cyan)", c.Name(), 										     SpacePadding(len(c.Name())), c.Attack(), c.Defense(), 
									     c.Navigation(), c.Government(), c.Technology(), c.Resources())
		}else if(c.Color() == "red"){
		civilizationstatstext = civilizationstatstext + fmt.Sprintf("[\n%s:%s%d   %d   %d   %d   %d   %d](fg:red)", c.Name(), 										     SpacePadding(len(c.Name())), c.Attack(), c.Defense(), 
									     c.Navigation(), c.Government(), c.Technology(), c.Resources())
		}else if(c.Color() == "magenta"){
		civilizationstatstext = civilizationstatstext + fmt.Sprintf("[\n%s:%s%d   %d   %d   %d   %d   %d](fg:magenta)", c.Name(), 										     SpacePadding(len(c.Name())), c.Attack(), c.Defense(), 
									     c.Navigation(), c.Government(), c.Technology(), c.Resources())
		}else if(c.Color() == "green"){
		civilizationstatstext = civilizationstatstext + fmt.Sprintf("[\n%s:%s%d   %d   %d   %d   %d   %d](fg:green)", c.Name(), 									     SpacePadding(len(c.Name())), c.Attack(), c.Defense(), 
									     c.Navigation(), c.Government(), c.Technology(), c.Resources())
		}

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
		for xcoord == 0 || tooCloseX(xcoordsused, xcoord) {
			xcoord = random(3,74)
			if(tooCloseX(xcoordsused, xcoord) == false) {
				xcoordsused = append(xcoordsused, xcoord)
				break
			}		
		}
		
		ycoord := 0
		for ycoord == 0 || tooCloseY(ycoordsused, ycoord) {
			ycoord = random(2, 32)
			if(tooCloseY(ycoordsused, ycoord) == false) {
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

func tooCloseX(s []int, e int) bool {
    for _, a := range s {
        if a == e || a == (e-1) || a == (e+1) {
            return true
        }
    }
    return false
}

func tooCloseY(s []int, e int) bool {
    for _, a := range s {
        if a == e || a == (e+1) {
            return true
        }
    }
    return false
}

