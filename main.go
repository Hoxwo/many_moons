package main

import "fmt"
import "log"
import "math/rand"
import "time"
//import "strings"
//import planet "many_moons/planet"
//import ship "many_moons/ship"
import civilization "many_moons/civilization"
import ui "github.com/gizak/termui/v3"
import "github.com/gizak/termui/v3/widgets"

func main() {
	// set the current year to 0
	// currentyear := 0
	// for randos
	rand.Seed(time.Now().UTC().UnixNano())
	// array of civilizations 
	civilizations := make([]*civilization.Civilization, 0)

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
	c0 := civilization.New("Balanced",   "cyan",    2, 2, 2, 2, 2, 2, 1, 1, 30, 30, 1, 2, -1)
	c1 := civilization.New("Warlike",    "red",     3, 2, 2, 2, 1, 1, 1, 1, 30, 30, 0, 0, 0)
	c2 := civilization.New("Defensive",  "magenta", 2, 3, 2, 2, 1, 1, 1, 1, 30, 30, 0, 0, 0)
	c3 := civilization.New("Explorer",   "green",   1, 1, 3, 2, 2, 2, 1, 1, 30, 30, 0, 0, 0)
	//c4 := civilization.New("Autocracy",  "blue",    2, 1, 2, 3, 1, 2, 1, 1, 30, 30, 0, 0, 0)
	//c5 := civilization.New("Technology", "yellow",  1, 2, 2, 1, 3, 2, 1, 1, 30, 30, 0, 0, 0)				

	//add them to master list
	civilizations = append(civilizations, &c0, &c1, &c2, &c3) //&c4, &c5)	


	//select first planet
	//selected = 0

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	
	//planetmap := ui.NewCanvas()
	//planetmap.SetRect(0, 0, 78, 34)
	//planetmap.SetLine(image.Pt(0, 0), image.Pt(10, 20), ui.ColorWhite)

	//datapane := ui.NewCanvas()
	//datapane.SetRect(78, 0, 126, 34)
	//datapane.SetLine(image.Pt(0, 0), image.Pt(10, 20), ui.ColorWhite)

	//ui.Render(planetmap, datapane)

	planetmap := widgets.NewParagraph()
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

	yourcivstatspane := widgets.NewParagraph()
	yourcivstatspane.Title = fmt.Sprintf("%s", civilizations[0].Name())
	yourcivstatspane.Text = ""
	yourcivstatspane.SetRect(78, 14, 116, 18)
	yourcivstatspane.BorderStyle.Fg = ui.ColorBlue

	selectplanetinfopane := widgets.NewParagraph()
	selectplanetinfopane.Title = "Selected Planet"
	selectplanetinfopane.Text = "" //"Simple colored text\nwith label. It [can be](fg:red) multilined with \\n or [break automatically](fg:red,fg:bold)"
	selectplanetinfopane.SetRect(78, 18, 116, 34)
	selectplanetinfopane.BorderStyle.Fg = ui.ColorBlue

	ui.Render(planetmap, civstatspane, sliderspane, yourcivstatspane, selectplanetinfopane)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
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
