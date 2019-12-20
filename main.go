package main

import "fmt"
import "log"
import "math/rand"
import math "math"
import "time"
import "image"
import "sort"
import planet "many_moons/planet"
import ship "many_moons/ship"
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
	centerofmapx := 0
 	centerofmapy := 0

	//array of ships
	ships := make([]*ship.Ship, 0)	

	//array of messages
	messageHistory := make([]string,0)
	//save message

	//All colors
	//termui.ColorCyan
	//termui.ColorGreen
	//termui.ColorYellow
	//termui.ColorWhite	
	//termui.ColorBlue
	//termui.ColorMagenta
	//termui.ColorRed
	
	//generate all planets, add them to master list
	planets = GenerateSpace(planets, centerofmapx, centerofmapy)

	// set up civilizations
	// {name, color, atk, def, nav, gov, tec, base res, res, shipsavail, maxshipsavail, shiptimer, maxshiptimer, colonizationtime, adsli, ademsli, eisli}
	//c0 := civilization.New("Balanced",   "cyan",    2, 2, 2, 2, 2, 2, 0, 1, 0, 30, 10, 1, 2, -1)
	//c1 := civilization.New("Warlike",    "red",     3, 2, 2, 2, 1, 1, 1, 1, 0, 30, 10, 0, 0, 0)
	//c2 := civilization.New("Defensive",  "magenta", 2, 3, 2, 2, 1, 1, 1, 1, 0, 30, 10, 0, 0, 0)
	//c3 := civilization.New("Explorer",   "green",   1, 1, 3, 2, 2, 2, 1, 1, 0, 30, 10, 0, 0, 0)
	//c4 := civilization.New("Autocracy",  "blue",    2, 1, 2, 3, 1, 2, 1, 1, 30, 30, 0, 0, 0)
	//c5 := civilization.New("Technology", "yellow",  1, 2, 2, 1, 3, 2, 1, 1, 30, 30, 0, 0, 0)
	//super powered version of the nations for testing				
	c0 := civilization.New("Balanced",   "cyan",    5, 5, 5, 5, 5, 5, 5, 1, 3, 0, 20, 10, 1, 2, -1)
	c1 := civilization.New("Warlike",    "red",     7, 6, 5, 5, 3, 4, 4, 1, 1, 0, 30, 10, 0, 0, 0)
	c2 := civilization.New("Defensive",  "magenta", 6, 7, 5, 5, 4, 4, 4, 1, 1, 0, 30, 10, 0, 0, 0)
	c3 := civilization.New("Explorer",   "green",   3, 4, 7, 6, 5, 5, 5, 1, 1, 0, 30, 10, 0, 0, 0)


	//add them to master list
	civilizations = append(civilizations, &c0, &c1, &c2, &c3) //&c4, &c5)	

	//set a home planet for each civ
	civnumber :=0
	for _, p := range planets {
		if(p.Planettype() == "Small Terrestrial" && civnumber < 4) {
			p.SetHomeplanet(true)
			p.SetOccupied(civilizations[civnumber].Name())
			civnumber++
		}
	}

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
				if(len(planets[i].Occupied()) > 0) {
					if(findCivilizationColorByName(planets[i].Occupied(), civilizations) == "cyan") {
						planetmap.SetPoint(pt1, ui.ColorCyan)
						planetmap.SetPoint(pt2, ui.ColorCyan)
						planetmap.SetPoint(pt3, ui.ColorCyan)
						planetmap.SetPoint(pt4, ui.ColorCyan)	
					} else if(findCivilizationColorByName(planets[i].Occupied(), civilizations) == "red") {
						planetmap.SetPoint(pt1, ui.ColorRed)
						planetmap.SetPoint(pt2, ui.ColorRed)
						planetmap.SetPoint(pt3, ui.ColorRed)
						planetmap.SetPoint(pt4, ui.ColorRed)	
					} else if(findCivilizationColorByName(planets[i].Occupied(), civilizations) == "magenta") {
						planetmap.SetPoint(pt1, ui.ColorMagenta)
						planetmap.SetPoint(pt2, ui.ColorMagenta)
						planetmap.SetPoint(pt3, ui.ColorMagenta)
						planetmap.SetPoint(pt4, ui.ColorMagenta)	
					} else if(findCivilizationColorByName(planets[i].Occupied(), civilizations) == "green") {
						planetmap.SetPoint(pt1, ui.ColorGreen)
						planetmap.SetPoint(pt2, ui.ColorGreen)
						planetmap.SetPoint(pt3, ui.ColorGreen)
						planetmap.SetPoint(pt4, ui.ColorGreen)	
					} 
				} else {
					planetmap.SetPoint(pt1, ui.ColorWhite)
					planetmap.SetPoint(pt2, ui.ColorWhite)
					planetmap.SetPoint(pt3, ui.ColorWhite)
					planetmap.SetPoint(pt4, ui.ColorWhite)	
				}
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
	playercivstatspane.SetRect(78, 19, 116, 24)
	playercivstatspane.BorderStyle.Fg = ui.ColorCyan

	//messages for player
	messages := widgets.NewList()	
	if(len(messageHistory) == 0) {		
		messages.Rows = make([]string,0)
	} else if(len(messageHistory) < 3) {
		messages.Rows = messageHistory[:len(messageHistory)]
	} else {
		messages.Rows = messageHistory[len(messageHistory)-3:len(messageHistory)]
	}
	messages.TextStyle = ui.NewStyle(ui.ColorWhite)
	messages.WrapText = false
	messages.BorderStyle.Fg = ui.ColorCyan
	messages.Title = "Messages"
	messages.SetRect(78, 19, 116, 24)

	selectplanetinfopane := widgets.NewParagraph()
	selectplanetinfopane.Title = "Planet Info"
	selectplanetinfopane.Text = SelectedPlanetText(planets, selectedplanet, ships, civilizations[0].Name(), civilizations[0].Navigation(), 
							civilizations[0].Technology())
	selectplanetinfopane.SetRect(78, 24, 116, 34)
	selectplanetinfopane.BorderStyle.Fg = ui.ColorBlue

	ui.Render(planetmapframe, planetmap, civstatspane, sliderspane, playercivstatspane, messages, selectplanetinfopane)

	tickerCount := 0
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Millisecond).C
	for {
		tickerCount++
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case ",":
				if(selectedplanet > 0) {
					selectedplanet--
				}
			case ".":
				if(selectedplanet < 15) {
					selectedplanet++
				}
			case "s":
				if(civilizations[0].Shipsavailable() > 0) {
					distance := int(math.Floor(
							distance(findPlanetXcoordByName(findCivilizationBasePlanetByName(civilizations[0].Name(), planets), planets),
							findPlanetYcoordByName(findCivilizationBasePlanetByName(civilizations[0].Name(), planets), planets),
							findPlanetXcoordByName(planets[selectedplanet].Name(), planets),
							findPlanetYcoordByName(planets[selectedplanet].Name(), planets))))
					navdistance := int(math.Floor(float64(distance/5)))
					tecneeded := 0
					if(planets[selectedplanet].Planettype() == "Ice Giant") {
						tecneeded = 4
					} else if(planets[selectedplanet].Planettype() == "Gas Giant") {
						tecneeded = 7
					}

					if(civilizations[0].Navigation() >= navdistance && civilizations[0].Technology() >= tecneeded) {
						civilizations[0].SetShipsavailable(civilizations[0].Shipsavailable() - 1)
						planetfromname := findCivilizationBasePlanetByName(civilizations[0].Name(), planets)
						newship := sendShip(civilizations[0].Name(), planetfromname, planets[selectedplanet].Name(), planets)
						ships = append(ships, &newship)
						messageHistory = append(messageHistory, fmt.Sprintf("Ship sent to %s", planets[selectedplanet].Name()))
					} else if(civilizations[0].Navigation() < navdistance && civilizations[0].Technology() >= tecneeded) {
						messageHistory = append(messageHistory, fmt.Sprintf("Higher NAV required"))
					} else if(civilizations[0].Navigation() >= navdistance && civilizations[0].Technology() < tecneeded) {
						messageHistory = append(messageHistory, fmt.Sprintf("Higher TEC required"))
					} else {
						messageHistory = append(messageHistory, fmt.Sprintf("Higher NAV and TEC required"))
					}
				} else {
					messageHistory = append(messageHistory, fmt.Sprintf("No ships available"))
				}				
			}
		case <-ticker:
			if(tickerCount % 100 == 1) {
				if(tickerCount % 400 == 1) {
					currentyear = currentyear + 1
					EveryCivilizationCountdown(civilizations, planets)
					message := advanceAllShips(ships, planets, civilizations)
					if(len(message) > 0) {
						messageHistory = append(messageHistory, message)
					}
					advanceColonizing(planets)
				}
			
			//render map and stat panes
			planetmapframe := widgets.NewParagraph()
			planetmapframe.Title = CurrentYearText(currentyear)
			planetmapframe.SetRect(0, 0, 78, 34)
			planetmapframe.BorderStyle.Fg = ui.ColorWhite

			//canvas-style planet map
			planetmap := ui.NewCanvas()
			planetmap.SetRect(2, 1, 75, 33)
			for i := 0; i < len(planets); i++ {
				//this is the ring that shows selected planet
				if(selectedplanet == i) {
					spt1 := image.Pt(((planets[i].Xcoord()*2)-3), ((planets[i].Ycoord()*4)-3))
					spt2 := image.Pt(((planets[i].Xcoord()*2)+3), ((planets[i].Ycoord()*4)-3))
					spt3 := image.Pt(((planets[i].Xcoord()*2)-3), ((planets[i].Ycoord()*4)+3))
					spt4 := image.Pt(((planets[i].Xcoord()*2)+3), ((planets[i].Ycoord()*4)+3))
					planetmap.SetPoint(spt1, ui.ColorYellow)
					planetmap.SetPoint(spt2, ui.ColorYellow)
					planetmap.SetPoint(spt3, ui.ColorYellow)
					planetmap.SetPoint(spt4, ui.ColorYellow)
				}
				
				//these points constitute the planet
				pt1 := image.Pt(((planets[i].Xcoord()*2)+0), ((planets[i].Ycoord()*4)+0))
				pt2 := image.Pt(((planets[i].Xcoord()*2)+1), ((planets[i].Ycoord()*4)+0))
				pt3 := image.Pt(((planets[i].Xcoord()*2)+0), ((planets[i].Ycoord()*4)+1))
				pt4 := image.Pt(((planets[i].Xcoord()*2)+1), ((planets[i].Ycoord()*4)+1))
				if(len(planets[i].Occupied()) > 0) {
				if(planets[i].Colonizing() == true) {
					planetmap.SetPoint(pt1, ui.ColorBlue)
					planetmap.SetPoint(pt2, ui.ColorBlue)
					planetmap.SetPoint(pt3, ui.ColorBlue)
					planetmap.SetPoint(pt4, ui.ColorBlue)	
				} else if(planets[i].Colonizing() == false && findCivilizationColorByName(planets[i].Occupied(), civilizations) == "cyan") {
					planetmap.SetPoint(pt1, ui.ColorCyan)
					planetmap.SetPoint(pt2, ui.ColorCyan)
					planetmap.SetPoint(pt3, ui.ColorCyan)
					planetmap.SetPoint(pt4, ui.ColorCyan)
					
					//satellites
					satellitenumber := 0
					if(planets[i].Homeplanet() == true) {
						satellitenumber = 1
					}
					
					deflevel := findCivilizationDEFByName(planets[i].Occupied(), civilizations)
					if(deflevel <= 1 || deflevel == 2) {
						satellitenumber += 1		
					} else if(deflevel == 3 || deflevel == 4) {
						satellitenumber += 2		
					} else if(deflevel == 5 || deflevel == 6) {
						satellitenumber += 3		
					} else if(deflevel == 7 || deflevel == 8) {
						satellitenumber += 4		
					} else {
						satellitenumber += 5
					}
				
					xcoordsused := make([]int, 0)
					ycoordsused := make([]int, 0) 
					for ii := 0; ii < satellitenumber; ii++ {
						xcoord := 0
						for xcoord == 0 || contains(xcoordsused, xcoord) {
							xcoord = random(planets[i].Xcoord() - 2,planets[i].Xcoord() + 3)
							if(contains(xcoordsused, xcoord) == false) {
								xcoordsused = append(xcoordsused, xcoord)
								break
							}		
						}
							
						ycoord := 0
						for ycoord == 0 || contains(ycoordsused, ycoord) {
							ycoord = random(planets[i].Ycoord() - 2, planets[i].Ycoord() + 4)
							if(contains(ycoordsused, ycoord) == false) {
								ycoordsused = append(ycoordsused, ycoord)
								break
							}		
						}
						satpt := image.Pt((xcoord*2), (ycoord*4))
						planetmap.SetPoint(satpt, ui.ColorCyan)
					}	
				} else if(planets[i].Colonizing() == false && findCivilizationColorByName(planets[i].Occupied(), civilizations) == "red") {
					planetmap.SetPoint(pt1, ui.ColorRed)
					planetmap.SetPoint(pt2, ui.ColorRed)
					planetmap.SetPoint(pt3, ui.ColorRed)
					planetmap.SetPoint(pt4, ui.ColorRed)

					//satellites
					satellitenumber := 0
					if(planets[i].Homeplanet() == true) {
						satellitenumber = 1
					}
					
					deflevel := findCivilizationDEFByName(planets[i].Occupied(), civilizations)
					if(deflevel <= 1 || deflevel == 2) {
						satellitenumber += 1		
					} else if(deflevel == 3 || deflevel == 4) {
						satellitenumber += 2		
					} else if(deflevel == 5 || deflevel == 6) {
						satellitenumber += 3		
					} else if(deflevel == 7 || deflevel == 8) {
						satellitenumber += 4		
					} else {
						satellitenumber += 5
					}
				
					xcoordsused := make([]int, 0)
					ycoordsused := make([]int, 0) 
					for ii := 0; ii < satellitenumber; ii++ {
						xcoord := 0
						for xcoord == 0 || contains(xcoordsused, xcoord) {
							xcoord = random(planets[i].Xcoord() - 2,planets[i].Xcoord() + 3)
							if(contains(xcoordsused, xcoord) == false) {
								xcoordsused = append(xcoordsused, xcoord)
								break
							}		
						}
							
						ycoord := 0
						for ycoord == 0 || contains(ycoordsused, ycoord) {
							ycoord = random(planets[i].Ycoord() - 2, planets[i].Ycoord() + 4)
							if(contains(ycoordsused, ycoord) == false) {
								ycoordsused = append(ycoordsused, ycoord)
								break
							}		
						}
						satpt := image.Pt((xcoord*2), (ycoord*4))
						planetmap.SetPoint(satpt, ui.ColorRed)
					}
	
				} else if(planets[i].Colonizing() == false && findCivilizationColorByName(planets[i].Occupied(), civilizations) == "magenta") {
					planetmap.SetPoint(pt1, ui.ColorMagenta)
					planetmap.SetPoint(pt2, ui.ColorMagenta)
					planetmap.SetPoint(pt3, ui.ColorMagenta)
					planetmap.SetPoint(pt4, ui.ColorMagenta)
					
					//satellites
					satellitenumber := 0
					if(planets[i].Homeplanet() == true) {
						satellitenumber = 1
					}
					
					deflevel := findCivilizationDEFByName(planets[i].Occupied(), civilizations)
					if(deflevel <= 1 || deflevel == 2) {
						satellitenumber += 1		
					} else if(deflevel == 3 || deflevel == 4) {
						satellitenumber += 2		
					} else if(deflevel == 5 || deflevel == 6) {
						satellitenumber += 3		
					} else if(deflevel == 7 || deflevel == 8) {
						satellitenumber += 4		
					} else {
						satellitenumber += 5
					}
				
					xcoordsused := make([]int, 0)
					ycoordsused := make([]int, 0) 
					for ii := 0; ii < satellitenumber; ii++ {
						xcoord := 0
						for xcoord == 0 || contains(xcoordsused, xcoord) {
							xcoord = random(planets[i].Xcoord() - 2,planets[i].Xcoord() + 3)
							if(contains(xcoordsused, xcoord) == false) {
								xcoordsused = append(xcoordsused, xcoord)
								break
							}		
						}
							
						ycoord := 0
						for ycoord == 0 || contains(ycoordsused, ycoord) {
							ycoord = random(planets[i].Ycoord() - 2, planets[i].Ycoord() + 4)
							if(contains(ycoordsused, ycoord) == false) {
								ycoordsused = append(ycoordsused, ycoord)
								break
							}		
						}
						satpt := image.Pt((xcoord*2), (ycoord*4))
						planetmap.SetPoint(satpt, ui.ColorMagenta)
					}
	
				} else if(planets[i].Colonizing() == false && findCivilizationColorByName(planets[i].Occupied(), civilizations) == "green") {
					planetmap.SetPoint(pt1, ui.ColorGreen)
					planetmap.SetPoint(pt2, ui.ColorGreen)
					planetmap.SetPoint(pt3, ui.ColorGreen)
					planetmap.SetPoint(pt4, ui.ColorGreen)	
					
					//satellites
					satellitenumber := 0
					if(planets[i].Homeplanet() == true) {
						satellitenumber = 1
					}
					
					deflevel := findCivilizationDEFByName(planets[i].Occupied(), civilizations)
					if(deflevel <= 1 || deflevel == 2) {
						satellitenumber += 1		
					} else if(deflevel == 3 || deflevel == 4) {
						satellitenumber += 2		
					} else if(deflevel == 5 || deflevel == 6) {
						satellitenumber += 3		
					} else if(deflevel == 7 || deflevel == 8) {
						satellitenumber += 4		
					} else {
						satellitenumber += 5
					}
				
					xcoordsused := make([]int, 0)
					ycoordsused := make([]int, 0) 
					for ii := 0; ii < satellitenumber; ii++ {
						xcoord := 0
						for xcoord == 0 || contains(xcoordsused, xcoord) {
							xcoord = random(planets[i].Xcoord() - 2,planets[i].Xcoord() + 3)
							if(contains(xcoordsused, xcoord) == false) {
								xcoordsused = append(xcoordsused, xcoord)
								break
							}		
						}
							
						ycoord := 0
						for ycoord == 0 || contains(ycoordsused, ycoord) {
							ycoord = random(planets[i].Ycoord() - 2, planets[i].Ycoord() + 4)
							if(contains(ycoordsused, ycoord) == false) {
								ycoordsused = append(ycoordsused, ycoord)
								break
							}		
						}
						satpt := image.Pt((xcoord*2), (ycoord*4))
						planetmap.SetPoint(satpt, ui.ColorGreen)
					}
				}
				} else {
					planetmap.SetPoint(pt1, ui.ColorWhite)
					planetmap.SetPoint(pt2, ui.ColorWhite)
					planetmap.SetPoint(pt3, ui.ColorWhite)
					planetmap.SetPoint(pt4, ui.ColorWhite)	
				}	
			}
				//ships
				for i := 0; i < len(ships); i++ {
					if(ships[i].Landed() == false) {
						shippt1 := image.Pt(((ships[i].Currentx() * 2)), ((ships[i].Currenty()*4)))
						shippt2 := image.Pt(((ships[i].Currentx() * 2) + 2), ((ships[i].Currenty()*4) + 2))	
						if(findCivilizationColorByName(ships[i].Civilization(), civilizations) == "cyan") {
							planetmap.SetLine(shippt1, shippt2, ui.ColorCyan)	
						} else if(findCivilizationColorByName(ships[i].Civilization(), civilizations) == "red") {
							planetmap.SetLine(shippt1, shippt2, ui.ColorRed)	
						} else if(findCivilizationColorByName(ships[i].Civilization(), civilizations) == "magenta") {
							planetmap.SetLine(shippt1, shippt2, ui.ColorMagenta)	
						} else if(findCivilizationColorByName(ships[i].Civilization(), civilizations) == "green") {
							planetmap.SetLine(shippt1, shippt2, ui.ColorGreen)	
						}		
					}
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

			//messages for player
			messages := widgets.NewList()	
			if(len(messageHistory) == 0) {		
				messages.Rows = make([]string,0)
			} else if(len(messageHistory) < 3) {
				messages.Rows = messageHistory[:len(messageHistory)]
			} else {
				messages.Rows = messageHistory[len(messageHistory)-3:len(messageHistory)]
			}
			messages.TextStyle = ui.NewStyle(ui.ColorWhite)
			messages.WrapText = false
			messages.BorderStyle.Fg = ui.ColorCyan
			messages.Title = "Messages"
			messages.SetRect(78, 19, 116, 24)

			selectplanetinfopane := widgets.NewParagraph()
			selectplanetinfopane.Title = "Planet Info"
			selectplanetinfopane.Text = SelectedPlanetText(planets, selectedplanet, ships, civilizations[0].Name(), civilizations[0].Navigation(),
									civilizations[0].Technology())
			selectplanetinfopane.SetRect(78, 24, 116, 34)
			selectplanetinfopane.BorderStyle.Fg = ui.ColorBlue
			//	

			ui.Render(planetmapframe, planetmap, civstatspane, sliderspane, playercivstatspane, messages, selectplanetinfopane)
			}
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

func EveryCivilizationCountdown(civilizations []*civilization.Civilization, planets []*planet.Planet) {
	for _, c := range civilizations {
		if(c.Shiptimer() < c.Maxshiptimer()) { 
			c.SetShiptimer(c.Shiptimer() + 1)
		} else {
			c.SetShiptimer(0)
			if(c.Shipsavailable() < c.Maxshipsavailable()) {
				c.SetShipsavailable(c.Shipsavailable() + 1)			
			}
			
			resourcesfromplanets := 0			
			for _, p := range planets {
				if(p.Occupied() == c.Name()) {
					resourcesfromplanets = resourcesfromplanets + p.Resources()
				}	
			}
			totalresources := c.Baseresources() + resourcesfromplanets
			if(totalresources > 9) {
				totalresources = 9			
			} 
			c.SetResources(totalresources)
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

func SelectedPlanetText(planets []*planet.Planet, selectedplanet int, ships []*ship.Ship, playercivilizationname string, playercivilizationnav int, playercivilizationtec int) string {
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
	distance := int(math.Floor(distance(findPlanetXcoordByName(findCivilizationBasePlanetByName(playercivilizationname, planets), planets),
				findPlanetYcoordByName(findCivilizationBasePlanetByName(playercivilizationname, planets), planets),
				xcoord,
				ycoord)))
	navdistance := int(math.Floor(float64(distance/5)))
	if(navdistance > 9) {
		navdistance = 9
	}
	if(planettype == "Ice Giant") {
		if(playercivilizationnav >= navdistance && playercivilizationtec >= 4) {
			selectedPlanetText = selectedPlanetText + fmt.Sprintf("\n[NAV required: %d](fg:blue) [TEC required: 4](fg:blue)", navdistance )
		} else if(playercivilizationnav < navdistance && playercivilizationtec >= 4) {
			selectedPlanetText = selectedPlanetText + fmt.Sprintf("\n[NAV required: %d](fg:red) [TEC required: 4](fg:blue)", navdistance )
		} else if(playercivilizationnav >= navdistance && playercivilizationtec < 4) {
			selectedPlanetText = selectedPlanetText + fmt.Sprintf("\n[NAV required: %d](fg:blue) [TEC required: 4](fg:red)", navdistance )
		} else if(playercivilizationnav < navdistance && playercivilizationtec < 4) {
			selectedPlanetText = selectedPlanetText + fmt.Sprintf("\n[NAV required: %d](fg:red) [TEC required: 4](fg:red)", navdistance )
		} 
	} else if(planettype == "Gas Giant") {
		if(playercivilizationnav >= navdistance && playercivilizationtec >= 7) {
			selectedPlanetText = selectedPlanetText + fmt.Sprintf("\n[NAV required: %d](fg:blue) [TEC required: 7](fg:blue)", navdistance )
		} else if(playercivilizationnav < navdistance && playercivilizationtec >= 7) {
			selectedPlanetText = selectedPlanetText + fmt.Sprintf("\n[NAV required: %d](fg:red) [TEC required: 7](fg:blue)", navdistance )
		} else if(playercivilizationnav >= navdistance && playercivilizationtec < 7) {
			selectedPlanetText = selectedPlanetText + fmt.Sprintf("\n[NAV required: %d](fg:blue) [TEC required: 7](fg:red)", navdistance )
		} else if(playercivilizationnav < navdistance && playercivilizationtec < 7) {
			selectedPlanetText = selectedPlanetText + fmt.Sprintf("\n[NAV required: %d](fg:red) [TEC required: 7](fg:red)", navdistance )
		} 
	} else {
		if(playercivilizationnav >= navdistance) {
			selectedPlanetText = selectedPlanetText + fmt.Sprintf("\n[NAV required: %d](fg:blue) [TEC required: 0](fg:blue)", navdistance )
		} else if(playercivilizationnav < navdistance) {
			selectedPlanetText = selectedPlanetText + fmt.Sprintf("\n[NAV required: %d](fg:red) [TEC required: 0](fg:blue)", navdistance )
		}
	}

	selectedPlanetText = selectedPlanetText + fmt.Sprintf("\n[<, >] scan [s] send ship")
	return selectedPlanetText
}

func GenerateSpace(planets []*planet.Planet, centerofmapx int, centerofmapy int) []*planet.Planet {
	numberplanets := 15 //random(16, 20)
	xcoordsused := make([]int, 0)
	ycoordsused := make([]int, 0) 
	for i := 0;  i <= numberplanets; i++ {
		// 1 - small rocky planet, 2 - large rocky planet, 3 - ice giant, 4 - gas giant
		planettype := random(1,5)
		// planets 2, 6, 9, 12  are small terrestrial		
		if(i == 2 || i == 6 || i == 9 || i == 12) {
			planettype = 1
		}
		
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
			p0 := planet.New(fmt.Sprintf("%s%d","ST ", i), "o",   1, "", "Small Terrestrial", xcoord, ycoord, false, 0, false)
			planets = append(planets, &p0)
		} else if(planettype == 2) {
			p0 := planet.New(fmt.Sprintf("%s%d","LT ", i), "O",   2, "", "Large Terrestrial", xcoord, ycoord, false, 0, false)
			planets = append(planets, &p0)		
		} else if(planettype == 3) {
			p0 := planet.New(fmt.Sprintf("%s%d","IG ", i), "(o)",   2, "", "Ice Giant", xcoord, ycoord, false, 0, false)
			planets = append(planets, &p0)
		} else {
			p0 := planet.New(fmt.Sprintf("%s%d","GG ", i), "(O)",   3, "", "Gas Giant", xcoord, ycoord, false, 0, false)
			planets = append(planets, &p0)
		}
	}
	
	//sort it left to right
	sort.Slice(planets, func(i, j int) bool {
  		return planets[i].Xcoord() < planets[j].Xcoord()
	})

	//get center of map for use in making base planets
	totalx := 0
	for _, x := range xcoordsused {
		totalx += x
	}
	centerofmapx = int(math.Floor(float64(totalx/16)))

	totaly := 0
        for _, y := range ycoordsused {
                totaly += y
        }
	centerofmapy = int(math.Floor(float64(totaly/16)))	

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

func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func findCivilizationColorByName(searchname string, civilizations []*civilization.Civilization) string {
	for _, c := range civilizations {
		if(c.Name() == searchname) {
			return c.Color()
		}
	}
	
	return "ERROR"
}

func findCivilizationDEFByName(searchname string, civilizations []*civilization.Civilization) int {
	for _, c := range civilizations {
		if(c.Name() == searchname) {
			return c.Defense()
		}
	}
	
	return -7
}

func findCivilizationATKByName(searchname string, civilizations []*civilization.Civilization) int {
	for _, c := range civilizations {
		if(c.Name() == searchname) {
			return c.Attack()
		}
	}
	
	return -7
} 

func findCivilizationRESByName(searchname string, civilizations []*civilization.Civilization) int {
	for _, c := range civilizations {
		if(c.Name() == searchname) {
			return c.Resources()
		}
	}

	return -7
}

func findCivilizationTimeToColonizeByName(searchname string, civilizations []*civilization.Civilization) int {
	for _, c := range civilizations {
		if(c.Name() == searchname) {
			return c.Colonizationtime()
		}
	}
	
	return -7
}

func findCivilizationBasePlanetByName(searchname string, planets []*planet.Planet) string {
	for _, p := range planets {
		if(p.Occupied() == searchname && p.Homeplanet() == true) {
			return p.Name()
		}
	}
	
	return "ERROR"
}


func findPlanetXcoordByName(searchname string, planets []*planet.Planet) int {
	for _, p := range planets {
		if(p.Name() == searchname) {
			return p.Xcoord()
		}
	}
	
	return -7
}

func findPlanetYcoordByName(searchname string,  planets []*planet.Planet) int {
	for _, p := range planets {
		if(p.Name() == searchname) {
			return p.Ycoord()
		}
	}
	
	return -7
}

func findPlanetOccupiedByName(searchname string,  planets []*planet.Planet) string {
	for _, p := range planets {
		if(p.Name() == searchname) {
			return p.Occupied()
		}
	}
	
	return "ERROR"
}

func findPlanetByCoords(xcoord int, ycoord int, planets []*planet.Planet) string {
	for _, p := range planets {
		if(p.Xcoord() == xcoord && p.Ycoord() == ycoord) {
			return p.Name()
		}
	}
	
	return "ERROR"
}

func sendShip(civilizationfromname string, planetfromname string, planettoname string, planets []*planet.Planet) ship.Ship {
	return ship.New(civilizationfromname, 10, 
		findPlanetXcoordByName(planetfromname,  planets), 
		findPlanetYcoordByName(planetfromname,  planets),
		findPlanetXcoordByName(planettoname,  planets), 
		findPlanetYcoordByName(planettoname, planets), 
		findPlanetXcoordByName(planetfromname,  planets), 
		findPlanetYcoordByName(planetfromname,  planets),
		false)		
}

func advanceAllShips(ships []*ship.Ship, planets []*planet.Planet, civilizations []*civilization.Civilization) string {
	for _, s := range ships { 
		//if advancing, -1 from x or y each time
		if(willShipAdvance(s.Speed()) == true) {
			rollforxory := random(1,3)
			if(rollforxory == 1 && s.Currentx() < s.Endx() || (s.Currenty() == s.Endy() && s.Currentx() < s.Endx())) {
				s.SetCurrentx(s.Currentx() + 1)
			} else if(rollforxory == 1 && s.Currentx() > s.Endx() || (s.Currenty() == s.Endy() && s.Currentx() > s.Endx())) {
				s.SetCurrentx(s.Currentx() - 1)
			} else if(rollforxory == 2 && s.Currenty() < s.Endy() || (s.Currentx() == s.Endx() && s.Currenty() < s.Endy())) {
				s.SetCurrenty(s.Currenty() + 1)
			} else if(rollforxory == 2 && s.Currenty() > s.Endy() || (s.Currentx() == s.Endx() && s.Currenty() > s.Endy())) {
				s.SetCurrenty(s.Currenty() - 1)
			}
		}

		if(s.Currentx() == s.Endx() && s.Currenty() == s.Endy() && s.Landed() != true) {
			return landShip(s, ships, planets, civilizations)
		}
	}
	return ""
}

func willShipAdvance(speed int) bool {
	//get number from 1 to 10
	rolltomove := random(1,11)

	//ship speed determines chance to move. speed of 10 means move every time, speed of 1 means 1/10 chance to move
	if(rolltomove <= speed) {
		return true
	}

	return false
}

func landShip(s *ship.Ship, ships []*ship.Ship, planets []*planet.Planet, civilizations []*civilization.Civilization) string {
	planetname := findPlanetByCoords(s.Endx(), s.Endy(), planets)
	if(s.Landed() == false) {
		colonize(s.Civilization(), planetname, planets, civilizations)
		attackmessage := attack(planetname, s.Civilization(), planets, civilizations)
		s.SetLanded(true) 
		return attackmessage
	}
	return ""
}

func colonize(civilizationname string, planetname string, planets []*planet.Planet, civilizations []*civilization.Civilization) { 
	for i, p := range planets {
		if(p.Name() == planetname && p.Occupied() == "") {
			updatedplanet := planet.New(p.Name(), p.Appearance(),   p.Resources(), civilizationname, p.Planettype(), p.Xcoord(), p.Ycoord(), 							    false, findCivilizationTimeToColonizeByName(civilizationname, civilizations), true)
			p = &updatedplanet
			planets[i] = p
		}
	}
}	

func advanceColonizing(planets []*planet.Planet) { 
	for _, p := range planets {
		if(p.Occupied() != "" && p.Colonizing() == true) {
			if(p.Timetocolonize() > 1) {
				p.SetTimetocolonize(p.Timetocolonize() - 1)
			} else {
			 	p.SetColonizing(false)
				p.SetTimetocolonize(0)
			}
		}
	}
}

func changePlanetOwner(oldowner string, newowner string, planets []*planet.Planet) {
	for _, p := range planets {
		if(p.Occupied() == oldowner) {
			p.SetOccupied(newowner)
		}
	}	
}

func attack(defenderplanetname string, attackername string, planets []*planet.Planet, civilizations []*civilization.Civilization) string {
	defendername := findPlanetOccupiedByName(defenderplanetname, planets)
	defenderdef := findCivilizationDEFByName(defendername, civilizations)
	attackeratk := findCivilizationATKByName(attackername, civilizations)
	difference := attackeratk - defenderdef

	//attacker gets ATK rolls of 10 sided die and defender gets DEF+1 rolls
	//attacker gets 2 additional rolls of die with guaranteed 10 if 3 levels above DEF
	//attacker has their resource value added
	attacksum := 0
	for i := 0; i < attackeratk; i++ {
		attacksum = attacksum + random(1, 11)
	}	
	attacksum = attacksum + findCivilizationRESByName(attackername, civilizations)

	defendsum := 0
	for i := 0; i < defenderdef; i++ {
		defendsum = defendsum + random(1, 11)
	}
	defendsum = defendsum + 10
		
	if(difference >= 3) {
		attacksum = attacksum + 10
	}
	
	if(attacksum > defendsum) {
		changePlanetOwner(defendername, attackername, planets)
		return fmt.Sprintf("Conquered - %s: %d - %s: %d", attackername, attacksum, defendername, defendsum)
	} else {
		return fmt.Sprintf("%s: %d - %s: %d", attackername, attacksum, defendername, defendsum)
	} 
}	

func distance(fromx int, fromy int, tox int, toy int) float64 {
	return math.Sqrt(math.Pow(float64(fromx - tox), 2) + math.Pow(float64(fromy - toy), 2))
}
