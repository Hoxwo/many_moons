package main

import "fmt"
import "time"
import "math/rand"
import "strings"
import planet "many_moons/planet"
import ship "many_moons/ship"
import civilization "many_moons/civilization"
import termui "github.com/gizak/termui"

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
	c0 := coin.New("Balanced",   "cyan",    2, 2, 2, 2, 2, 2, 1, 1, 30, 30, 0, 0, 0)
	c1 := coin.New("Warlike",    "red",     3, 2, 2, 2, 1, 1, 1, 1, 30, 30, 0, 0, 0)
	c2 := coin.New("Defensive",  "magenta", 2, 3, 2, 2, 1, 1, 1, 1, 30, 30, 0, 0, 0)
	c3 := coin.New("Explorer",   "green",   1, 1, 3, 2, 2, 2, 1, 1, 30, 30, 0, 0, 0)
	c4 := coin.New("Autocracy",  "blue",    2, 1, 2, 3, 1, 2, 1, 1, 30, 30, 0, 0, 0)
	c5 := coin.New("Technology", "yellow",  1, 2, 2, 1, 3, 2, 1, 1, 30, 30, 0, 0, 0)				

	//add them to master list
	civilizations = append(civilizations, &c0, &c1, &c2, &c3, &c4, &c5)	


	//select first planet
	selected = 0

	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	//termui.Render(spls1, singlespl0, par0)

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})

	//if in buying state, player gains 1 of selected coin
	//	player loses savings equal to value of 1 coin
	//if in selling state, player loses 1 of selected coin
	//	player gains savings equal to value of 1 coin
	termui.Handle("/sys/kbd/1", func(termui.Event) {
		lastPlayerQuantityInput = 1		
		message := ""
		if(coins[selected].LaunchDay() > dayCounter) {
			message = "Coin not available"
		} else {
			if(state == 2) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 1) 
			} else if(state == 3) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 2)
			}
		}
		//save message
		if(strings.Compare(message,"") != 0) {
				messageHistory = append(messageHistory, message)
		}
	})
	
	termui.Handle("/sys/kbd/2", func(termui.Event) {
		lastPlayerQuantityInput = 2
		message := ""
		if(coins[selected].LaunchDay() > dayCounter) {
			message = "Coin not available"
		} else {
			if(state == 2) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 1) 
			} else if(state == 3) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 2)
			}
		}
		//save message
		if(strings.Compare(message,"") != 0) {
				messageHistory = append(messageHistory, message)
		}
	})
	
	termui.Handle("/sys/kbd/3", func(termui.Event) {
		lastPlayerQuantityInput = 3
		message := ""
		if(coins[selected].LaunchDay() > dayCounter) {
			message = "Coin not available"
		} else {
			if(state == 2) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 1) 
			} else if(state == 3) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 2)
			}
		}
		//save message
		if(strings.Compare(message,"") != 0) {
				messageHistory = append(messageHistory, message)
		}
	})
	
	termui.Handle("/sys/kbd/4", func(termui.Event) {
		lastPlayerQuantityInput = 4
		message := ""
		if(coins[selected].LaunchDay() > dayCounter) {
			message = "Coin not available"
		} else {
			if(state == 2) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 1) 
			} else if(state == 3) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 2)
			}
		}
		//save message
		if(strings.Compare(message,"") != 0) {
				messageHistory = append(messageHistory, message)
		}
	})
	
	termui.Handle("/sys/kbd/5", func(termui.Event) {
		lastPlayerQuantityInput = 5
		message := ""
		if(coins[selected].LaunchDay() > dayCounter) {
			message = "Coin not available"
		} else {
			if(state == 2) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 1) 
			} else if(state == 3) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 2)
			}
		}
		//save message
		if(strings.Compare(message,"") != 0) {
				messageHistory = append(messageHistory, message)
		}
	})
	
	termui.Handle("/sys/kbd/6", func(termui.Event) {
		lastPlayerQuantityInput = 6
		message := ""
		if(coins[selected].LaunchDay() > dayCounter) {
			message = "Coin not available"
		} else {
			if(state == 2) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 1) 
			} else if(state == 3) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 2)
			}
		}
		//save message
		if(strings.Compare(message,"") != 0) {
				messageHistory = append(messageHistory, message)
		}
	})
	
	termui.Handle("/sys/kbd/7", func(termui.Event) {
		lastPlayerQuantityInput = 7
		message := ""
		if(coins[selected].LaunchDay() > dayCounter) {
			message = "Coin not available"
		} else {
			if(state == 2) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 1) 
			} else if(state == 3) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 2)
			}
		}
		//save message
		if(strings.Compare(message,"") != 0) {
				messageHistory = append(messageHistory, message)
		}
	})
	
	termui.Handle("/sys/kbd/8", func(termui.Event) {
		lastPlayerQuantityInput = 8
		message := ""
		if(coins[selected].LaunchDay() > dayCounter) {
			message = "Coin not available"
		} else {
			if(state == 2) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 1) 
			} else if(state == 3) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 2)
			}
		}
		//save message
		if(strings.Compare(message,"") != 0) {
				messageHistory = append(messageHistory, message)
		}
	})

	termui.Handle("/sys/kbd/9", func(termui.Event) {
		lastPlayerQuantityInput = 9
		message := ""
		if(coins[selected].LaunchDay() > dayCounter) {
			message = "Coin not available"
		} else {
			if(state == 2) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 1) 
			} else if(state == 3) {
				message = t.ModifyCoinAndSavingsBalance(coins[selected].Name(), lastPlayerQuantityInput, 
									(float64(lastPlayerQuantityInput)*coins[selected].Price()), 2)
			}
		}
		//save message
		if(strings.Compare(message,"") != 0) {
				messageHistory = append(messageHistory, message)
		}
	})		

	termui.Handle("/sys/kbd/b", func(termui.Event) {
		//player buy
		state = 2
	})
	
	termui.Handle("/sys/kbd/s", func(termui.Event) {
		//player sell
		state = 3
	})

	termui.Handle("/sys/kbd/x", func(termui.Event) {
		//player back to normal
		state = 1
	})

	termui.Handle("/sys/kbd/k", func(termui.Event) {
		if(selected > 0) {		
			selected--
		}
	})

	termui.Handle("/sys/kbd/m", func(termui.Event) {
		if(selected < 13) {		
			selected++
		}
	})

	termui.Handle("/sys/kbd/,", func(termui.Event) {
		if(selected > 0) {		
			selected--
		}
	})

	termui.Handle("/sys/kbd/.", func(termui.Event) {
		if(selected < 13) {		
			selected++
		}
	})
	
	//"/sys/kbd/g"
	termui.Handle("/timer/1s", func(termui.Event) {
		currentTime = currentTime.Add(time.Hour * 24 * 1)
		dayCounter++

		sixSidedDie := random(1,7)
		m := int(currentTime.Month())		
		if( m == 8 || m == 9 || m == 10 || m == 2 || m == 3 || m == 4) {
			// make trend sticky
			if( marketTrend == 3 || marketTrend == 4) {
				sixSidedDie = sixSidedDie - 1
				if (strings.Compare(news, "") != 0) {
					sixSidedDie = sixSidedDie + 1
				}
			}	
	
			if (sixSidedDie <= 2) {
				marketTrend = 4
			} else if (sixSidedDie == 3 || sixSidedDie == 4) {
				marketTrend = 3
			} else if (sixSidedDie == 5) {
				marketTrend = 2
			} else {
				marketTrend = 1
			}	
		} else { 
			// make trend sticky
			if( marketTrend == 1 || marketTrend == 2) {
				sixSidedDie = sixSidedDie - 1
				if (strings.Compare(news, "") != 0) {
					sixSidedDie = sixSidedDie - 1
				}		
			}	

			if (sixSidedDie <= 2) {
					marketTrend = 1
			} else if (sixSidedDie == 3 || sixSidedDie == 4) {
				marketTrend = 2
			} else if (sixSidedDie == 5) {
				marketTrend = 3
			} else {
				marketTrend = 4
			}	
		} 
		
		totalMarketCap = AdvanceOneDay(coins, exchanges, coinPrices, exchangeValues, coinPriceHistory, exchangeValueHistory, 							    coinMarketShares, dayCounter, marketTrend, news, &newsHistory, &t)
		
		//selected coin info	
		selectedInfo := termui.NewPar("")
		if(state == 1) {		
			selectedInfo = termui.NewPar(SelectedCoinTextState1(coins, selected))
		} else if(state == 2) {
			selectedInfo = termui.NewPar(SelectedCoinTextState2(coins, selected))
		} else {
			selectedInfo = termui.NewPar(SelectedCoinTextState3(coins, selected))
		}
		selectedInfo.Height = 8
		selectedInfo.Width = 21
		selectedInfo.X = 30
		selectedInfo.Y = 16
		selectedInfo.BorderLabel = ""
		selectedInfo.BorderFg = termui.ColorCyan
		selectedInfo.TextFgColor = termui.ColorGreen

		//trader info
		traderInfo := termui.NewPar("")		
		traderInfo = termui.NewPar(TraderInfoText(t, coins))
		traderInfo.Height = 18
		traderInfo.Width = 27
		traderInfo.X = 51
		traderInfo.Y = 10
		traderInfo.BorderLabel = fmt.Sprintf("Balances -- %s", currentTime.Format("01-02-2006"))
		traderInfo.BorderFg = termui.ColorCyan
		traderInfo.TextFgColor = termui.ColorWhite			
	
		//savings history
		savingsHistory := termui.NewLineChart()
		savingsHistory.BorderLabel = "Savings"
		savingsHistory.Mode = "dot"
		if(dayCounter < 27) {
			savingsHistory.Data = t.SavingsBalanceHistory()[:dayCounter]
		} else {
			savingsHistory.Data = t.SavingsBalanceHistory()[dayCounter-27:dayCounter]
		}
		savingsHistory.Width = 27	
		savingsHistory.Height = 6
		savingsHistory.X = 51
		savingsHistory.Y = 28
		savingsHistory.DotStyle = '.'
		savingsHistory.AxesColor = termui.ColorCyan
		savingsHistory.LineColor = termui.ColorBlue | termui.AttrBold

		//Coin Worth $
		coinWorth := termui.NewLineChart()
		coinWorth.BorderLabel = "Crypto Net Worth $"
		if(dayCounter < 21) {
			coinWorth.Data = GetHistoricTotalMarketCapAsFloatArray(exchangeValueHistory)[:dayCounter]
		} else {
			coinWorth.Data = GetTraderDollarValueForAllCoins(t, coinPriceHistory, dayCounter)[dayCounter-21:dayCounter]			
		}
		coinWorth.Width = 21
		coinWorth.Height = 10
		coinWorth.X = 30
		coinWorth.Y = 24
		coinWorth.AxesColor = termui.ColorWhite
		coinWorth.LineColor = termui.ColorYellow

		// Short term dollar amounts, or estimate of day until launch
		shorttermhisttitle0 := ShortTermCoinTitle(coins[0], dayCounter, 0, selected)
		shorttermhist0 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist0.Data = FloatToInts(GetHistoricPriceDataForCoin("Bitcoin", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist0.Data = FloatToInts(GetHistoricPriceDataForCoin("Bitcoin", coinPriceHistory)[dayCounter-30:dayCounter])
		}			
		shorttermhist0.Title = shorttermhisttitle0
		shorttermhist0.LineColor = termui.ColorCyan

		shorttermhisttitle1 := ShortTermCoinTitle(coins[1], dayCounter, 1, selected)
		shorttermhist1 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist1.Data = FloatToInts(GetHistoricPriceDataForCoin("LightCoin", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist1.Data = FloatToInts(GetHistoricPriceDataForCoin("LightCoin", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist1.Title = shorttermhisttitle1
		shorttermhist1.LineColor = termui.ColorGreen

		shorttermhisttitle2 := ShortTermCoinTitle(coins[2], dayCounter, 2, selected)
		shorttermhist2 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist2.Data = FloatToInts(GetHistoricPriceDataForCoin("Nethereum", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist2.Data = FloatToInts(GetHistoricPriceDataForCoin("Nethereum", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist2.Title = shorttermhisttitle2
		shorttermhist2.LineColor = termui.ColorMagenta

		shorttermhisttitle3 := ShortTermCoinTitle(coins[3], dayCounter, 3, selected)
		shorttermhist3 := termui.NewSparkline()	
		if(dayCounter < 30) {
			shorttermhist3.Data = FloatToInts(GetHistoricPriceDataForCoin("Nethereum Vintage", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist3.Data = FloatToInts(GetHistoricPriceDataForCoin("Nethereum Vintage", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist3.Title = shorttermhisttitle3
		shorttermhist3.LineColor = termui.ColorCyan

		shorttermhisttitle4 := ShortTermCoinTitle(coins[4], dayCounter, 4, selected)
		shorttermhist4 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist4.Data = FloatToInts(GetHistoricPriceDataForCoin("Riddle", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist4.Data = FloatToInts(GetHistoricPriceDataForCoin("Riddle", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist4.Title = shorttermhisttitle4
		shorttermhist4.LineColor = termui.ColorGreen
	
		shorttermhisttitle5 := ShortTermCoinTitle(coins[5], dayCounter, 5, selected)
		shorttermhist5 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist5.Data = FloatToInts(GetHistoricPriceDataForCoin("Ledge", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist5.Data = FloatToInts(GetHistoricPriceDataForCoin("Ledge", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist5.Title = shorttermhisttitle5
		shorttermhist5.LineColor = termui.ColorMagenta
	
		shorttermhisttitle6 := ShortTermCoinTitle(coins[6], dayCounter, 6, selected)
		shorttermhist6 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist6.Data = FloatToInts(GetHistoricPriceDataForCoin("Bancem", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist6.Data = FloatToInts(GetHistoricPriceDataForCoin("Bancem", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist6.Title = shorttermhisttitle6
		shorttermhist6.LineColor = termui.ColorCyan
	
		shorttermhisttitle7 := ShortTermCoinTitle(coins[7], dayCounter, 7, selected)
		shorttermhist7 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist7.Data = FloatToInts(GetHistoricPriceDataForCoin("ZEO", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist7.Data = FloatToInts(GetHistoricPriceDataForCoin("ZEO", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist7.Title = shorttermhisttitle7
		shorttermhist7.LineColor = termui.ColorGreen

		shorttermhisttitle8 := ShortTermCoinTitle(coins[8], dayCounter, 8, selected)
		shorttermhist8 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist8.Data = FloatToInts(GetHistoricPriceDataForCoin("YCash", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist8.Data = FloatToInts(GetHistoricPriceDataForCoin("YCash", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist8.Title = shorttermhisttitle8
		shorttermhist8.LineColor = termui.ColorMagenta

		shorttermhisttitle9 := ShortTermCoinTitle(coins[9], dayCounter, 9, selected)
		shorttermhist9 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist9.Data = FloatToInts(GetHistoricPriceDataForCoin("Interstellar", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist9.Data = FloatToInts(GetHistoricPriceDataForCoin("Interstellar", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist9.Title = shorttermhisttitle9
		shorttermhist9.LineColor = termui.ColorCyan

		shorttermhisttitle10 := ShortTermCoinTitle(coins[10], dayCounter, 10, selected)
		shorttermhist10 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist10.Data = FloatToInts(GetHistoricPriceDataForCoin("Bitbeets", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist10.Data = FloatToInts(GetHistoricPriceDataForCoin("Bitbeets", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist10.Title = shorttermhisttitle10
		shorttermhist10.LineColor = termui.ColorGreen

		shorttermhisttitle11 := ShortTermCoinTitle(coins[11], dayCounter, 11, selected)
		shorttermhist11 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist11.Data = FloatToInts(GetHistoricPriceDataForCoin("TRAM", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist11.Data = FloatToInts(GetHistoricPriceDataForCoin("TRAM", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist11.Title = shorttermhisttitle11
		shorttermhist11.LineColor = termui.ColorMagenta
		
		shorttermhisttitle12 := ShortTermCoinTitle(coins[12], dayCounter, 12, selected)
		shorttermhist12 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist12.Data = FloatToInts(GetHistoricPriceDataForCoin("DigiLink", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist12.Data = FloatToInts(GetHistoricPriceDataForCoin("DigiLink", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist12.Title = shorttermhisttitle12
		shorttermhist12.LineColor = termui.ColorCyan

		shorttermhisttitle13 := ShortTermCoinTitle(coins[13], dayCounter, 13, selected)
		shorttermhist13 := termui.NewSparkline()
		if(dayCounter < 30) {
			shorttermhist13.Data = FloatToInts(GetHistoricPriceDataForCoin("XTRAbits", coinPriceHistory)[:dayCounter])
		} else {
			shorttermhist13.Data = FloatToInts(GetHistoricPriceDataForCoin("XTRAbits", coinPriceHistory)[dayCounter-30:dayCounter])
		}	
		shorttermhist13.Title = shorttermhisttitle13
		shorttermhist13.LineColor = termui.ColorGreen

		// put them together
		shorttermhistograms := termui.NewSparklines(shorttermhist0, shorttermhist1, shorttermhist2, 
					shorttermhist3, shorttermhist4, shorttermhist5,
					shorttermhist6, shorttermhist7, shorttermhist8,
					shorttermhist9, shorttermhist10, shorttermhist11,
					shorttermhist12, shorttermhist13)
		shorttermhistograms.Height = 30
		shorttermhistograms.Width = 30
		shorttermhistograms.Y = 4
		shorttermhistograms.X = 0
		shorttermhistograms.BorderLabel = "Coin - supply - price"
		
		//List of exchanges - presented as gauges of total cap
		exchangeGauge0 := termui.NewGauge()
		exchangeGauge0.Percent = ExchangeInfoPercent(exchanges[0])
		exchangeGauge0.Width = 28
		exchangeGauge0.Height = 3
		exchangeGauge0.Y = 10
		exchangeGauge0.X = 78
		exchangeGauge0.BorderLabel = ExchangeInfoString(exchanges[0], dayCounter)
		exchangeGauge0.Label = ExchangeInfoLabel(exchanges[0])
		exchangeGauge0.BarColor = termui.ColorMagenta
		exchangeGauge0.BorderFg = termui.ColorWhite
		exchangeGauge0.LabelAlign = termui.AlignRight

		exchangeGauge1 := termui.NewGauge()
		exchangeGauge1.Percent = ExchangeInfoPercent(exchanges[1])
		exchangeGauge1.Width = 28
		exchangeGauge1.Height = 3
		exchangeGauge1.Y = 13
		exchangeGauge1.X = 78
		exchangeGauge1.BorderLabel = ExchangeInfoString(exchanges[1], dayCounter)
		exchangeGauge1.Label = ExchangeInfoLabel(exchanges[1])
		exchangeGauge1.BarColor = termui.ColorMagenta
		exchangeGauge1.BorderFg = termui.ColorWhite
		exchangeGauge1.LabelAlign = termui.AlignRight
		
		exchangeGauge2 := termui.NewGauge()
		exchangeGauge2.Percent = ExchangeInfoPercent(exchanges[2])
		exchangeGauge2.Width = 28
		exchangeGauge2.Height = 3
		exchangeGauge2.Y = 16
		exchangeGauge2.X = 78
		exchangeGauge2.BorderLabel = ExchangeInfoString(exchanges[2], dayCounter)
		exchangeGauge2.Label = ExchangeInfoLabel(exchanges[2])
		exchangeGauge2.BarColor = termui.ColorMagenta
		exchangeGauge2.BorderFg = termui.ColorWhite
		exchangeGauge2.LabelAlign = termui.AlignRight

		exchangeGauge3 := termui.NewGauge()
		exchangeGauge3.Percent = ExchangeInfoPercent(exchanges[3])
		exchangeGauge3.Width = 28
		exchangeGauge3.Height = 3
		exchangeGauge3.Y = 19
		exchangeGauge3.X = 78
		exchangeGauge3.BorderLabel = ExchangeInfoString(exchanges[3], dayCounter)
		exchangeGauge3.Label = ExchangeInfoLabel(exchanges[3])
		exchangeGauge3.BarColor = termui.ColorMagenta
		exchangeGauge3.BorderFg = termui.ColorWhite
		exchangeGauge3.LabelAlign = termui.AlignRight

		exchangeGauge4 := termui.NewGauge()
		exchangeGauge4.Percent = ExchangeInfoPercent(exchanges[4])
		exchangeGauge4.Width = 28
		exchangeGauge4.Height = 3
		exchangeGauge4.Y = 22
		exchangeGauge4.X = 78
		exchangeGauge4.BorderLabel = ExchangeInfoString(exchanges[4], dayCounter)
		exchangeGauge4.Label = ExchangeInfoLabel(exchanges[4])
		exchangeGauge4.BarColor = termui.ColorMagenta
		exchangeGauge4.BorderFg = termui.ColorWhite
		exchangeGauge4.LabelAlign = termui.AlignRight

		exchangeGauge5 := termui.NewGauge()
		exchangeGauge5.Percent = ExchangeInfoPercent(exchanges[5])
		exchangeGauge5.Width = 28
		exchangeGauge5.Height = 3
		exchangeGauge5.Y = 25
		exchangeGauge5.X = 78
		exchangeGauge5.BorderLabel = ExchangeInfoString(exchanges[5], dayCounter)
		exchangeGauge5.Label = ExchangeInfoLabel(exchanges[5])
		exchangeGauge5.BarColor = termui.ColorMagenta
		exchangeGauge5.BorderFg = termui.ColorWhite
		exchangeGauge5.LabelAlign = termui.AlignRight

		exchangeGauge6 := termui.NewGauge()
		exchangeGauge6.Percent = ExchangeInfoPercent(exchanges[6])
		exchangeGauge6.Width = 28
		exchangeGauge6.Height = 3
		exchangeGauge6.Y = 28
		exchangeGauge6.X = 78
		exchangeGauge6.BorderLabel = ExchangeInfoString(exchanges[6], dayCounter)
		exchangeGauge6.Label = ExchangeInfoLabel(exchanges[6])
		exchangeGauge6.BarColor = termui.ColorMagenta
		exchangeGauge6.BorderFg = termui.ColorWhite
		exchangeGauge6.LabelAlign = termui.AlignRight

		exchangeGauge7 := termui.NewGauge()
		exchangeGauge7.Percent = ExchangeInfoPercent(exchanges[7])
		exchangeGauge7.Width = 28
		exchangeGauge7.Height = 3
		exchangeGauge7.Y = 31
		exchangeGauge7.X = 78
		exchangeGauge7.BorderLabel = ExchangeInfoString(exchanges[7], dayCounter)
		exchangeGauge7.Label = ExchangeInfoLabel(exchanges[7])
		exchangeGauge7.BarColor = termui.ColorMagenta
		exchangeGauge7.BorderFg = termui.ColorWhite
		exchangeGauge7.LabelAlign = termui.AlignRight

		marketCap := termui.NewLineChart()
		marketCap.BorderLabel = MarketCapInfoString(totalMarketCap)
		marketCap.Mode = "dot"
		marketCapWindow := make([]float64, 30)
		if(dayCounter < 28) {
			marketCapWindow = GetHistoricTotalMarketCapAsFloatArray(exchangeValueHistory)[:dayCounter]
		} else {
			marketCapWindow = GetHistoricTotalMarketCapAsFloatArray(exchangeValueHistory)[dayCounter-28:dayCounter]			
		}
		marketCap.Data = marketCapWindow	
		marketCap.Width = 28
		marketCap.Height = 10
		marketCap.X = 78
		marketCap.Y = 0
		marketCap.DotStyle = '.'
		marketCap.AxesColor = termui.ColorWhite
		marketCap.LineColor = termui.ColorGreen | termui.AttrBold

		marketShares := termui.NewBarChart()
		data := []int{MarketShareForCoin(coinMarketShares,coins[0]), MarketShareForCoin(coinMarketShares,coins[1]), 
				MarketShareForCoin(coinMarketShares,coins[2]),MarketShareForCoin(coinMarketShares,coins[3]),
				 MarketShareForCoin(coinMarketShares,coins[4]),MarketShareForCoin(coinMarketShares,coins[5]),
					MarketShareForCoin(coinMarketShares,coins[6]), MarketShareForCoin(coinMarketShares,coins[7]), 						MarketShareForCoin(coinMarketShares,coins[8]),MarketShareForCoin(coinMarketShares,coins[9]),
					 MarketShareForCoin(coinMarketShares,coins[10]), MarketShareForCoin(coinMarketShares,coins[11]),
					MarketShareForCoin(coinMarketShares,coins[12]), MarketShareForCoin(coinMarketShares,coins[13])}
		labels := []string{coins[0].Symbol(), coins[1].Symbol(), coins[2].Symbol(), coins[3].Symbol(),
					coins[4].Symbol(), coins[5].Symbol(), coins[6].Symbol(), coins[7].Symbol(),
					coins[8].Symbol(), coins[9].Symbol(), coins[10].Symbol(), coins[11].Symbol(),
					coins[12].Symbol(), coins[13].Symbol()}
		marketShares.BorderLabel = "Market Share"
		marketShares.Data = data
		marketShares.Width = 64
		marketShares.Height = 4
		marketShares.X=0
		marketShares.Y=0
		marketShares.DataLabels = labels
		marketShares.TextColor = termui.ColorWhite
		marketShares.BarColor = termui.ColorBlue
		marketShares.NumColor = termui.ColorWhite
		
		//news
		recentNews := termui.NewList()	
		if(len(newsHistory) == 0) {		
			recentNews.Items = make([]string,0)
		} else if(len(newsHistory) < 4) {
			recentNews.Items = newsHistory[:len(newsHistory)]
		} else {
			recentNews.Items = newsHistory[len(newsHistory)-4:len(newsHistory)]
		}
		recentNews.ItemFgColor = termui.ColorWhite
		recentNews.BorderLabel = fmt.Sprintf("Latest News")
		recentNews.Height = 6
		recentNews.Width = 48
		recentNews.Y = 4
		recentNews.X = 30

		//messages for player
		messages := termui.NewList()	
		if(len(messageHistory) == 0) {		
			recentNews.Items = make([]string,0)
		} else if(len(messageHistory) < 4) {
			messages.Items = messageHistory[:len(messageHistory)]
		} else {
			messages.Items = messageHistory[len(messageHistory)-4:len(messageHistory)]
		}
		messages.ItemFgColor = termui.ColorCyan
		messages.BorderLabel = fmt.Sprintf("Messages")
		messages.Height = 6
		messages.Width = 21
		messages.Y = 10
		messages.X = 30

		//market sentiment
		sentiment := termui.NewPar(SentimentString(marketTrend))
		sentiment.Height = 4
		sentiment.Width = 14
		sentiment.X = 64
		sentiment.Y = 0
		sentiment.BorderLabel = "Sentiment"
		sentiment.BorderFg = termui.ColorCyan
		if(marketTrend == 1 || marketTrend == 2) {
			sentiment.TextFgColor = termui.ColorGreen
		} else {
			sentiment.TextFgColor = termui.ColorRed
		}


		debug := termui.NewPar(fmt.Sprintf(" day %d", dayCounter))
		debug.Height = 1
		debug.Width = 20
		debug.X = 34
		debug.Y = 30
		debug.Border = false		
	
	termui.Render( shorttermhistograms, debug, exchangeGauge0, exchangeGauge1, exchangeGauge2, exchangeGauge3,
				exchangeGauge4, exchangeGauge5, exchangeGauge6, exchangeGauge7, marketCap, marketShares, 
				recentNews, sentiment, selectedInfo, traderInfo, messages, savingsHistory, coinWorth)
	})

	termui.Loop()

}
