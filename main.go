package main

//import "image"
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
	//for randos
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
	c0 := civilization.New("Balanced",   "cyan",    2, 2, 2, 2, 2, 2, 1, 1, 30, 30, 0, 0, 0)
	c1 := civilization.New("Warlike",    "red",     3, 2, 2, 2, 1, 1, 1, 1, 30, 30, 0, 0, 0)
	c2 := civilization.New("Defensive",  "magenta", 2, 3, 2, 2, 1, 1, 1, 1, 30, 30, 0, 0, 0)
	c3 := civilization.New("Explorer",   "green",   1, 1, 3, 2, 2, 2, 1, 1, 30, 30, 0, 0, 0)
	c4 := civilization.New("Autocracy",  "blue",    2, 1, 2, 3, 1, 2, 1, 1, 30, 30, 0, 0, 0)
	c5 := civilization.New("Technology", "yellow",  1, 2, 2, 1, 3, 2, 1, 1, 30, 30, 0, 0, 0)				

	//add them to master list
	civilizations = append(civilizations, &c0, &c1, &c2, &c3, &c4, &c5)	


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

	p0 := widgets.NewParagraph()
	p0.Text = "Borderless Text"
	p0.SetRect(0, 0, 20, 5)
	p0.Border = false

	p1 := widgets.NewParagraph()
	p1.Title = "标签"
	p1.Text = "你好，世界。"
	p1.SetRect(20, 0, 35, 5)

	p2 := widgets.NewParagraph()
	p2.Title = "Multiline"
	p2.Text = "Simple colored text\nwith label. It [can be](fg:red) multilined with \\n or [break automatically](fg:red,fg:bold)"
	p2.SetRect(0, 5, 35, 10)
	p2.BorderStyle.Fg = ui.ColorYellow

	p3 := widgets.NewParagraph()
	p3.Title = "Auto Trim"
	p3.Text = "Long text with label and it is auto trimmed."
	p3.SetRect(0, 10, 40, 15)

	p4 := widgets.NewParagraph()
	p4.Title = "Text Box with Wrapping"
	p4.Text = "Press q to QUIT THE DEMO. [There](fg:blue,mod:bold) are other things [that](fg:red) are going to fit in here I think. What do you think? Now is the time for all good [men to](bg:blue) come to the aid of their country. [This is going to be one really really really long line](fg:green) that is going to go together and stuffs and things. Let's see how this thing renders out.\n    Here is a new paragraph and stuffs and things. There should be a tab indent at the beginning of the paragraph. Let's see if that worked as well."
	p4.SetRect(40, 0, 70, 20)
	p4.BorderStyle.Fg = ui.ColorBlue

	ui.Render(p0, p1, p2, p3, p4)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
