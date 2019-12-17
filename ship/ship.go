package ship

type Ship struct {
    civilization string
    speed int
    startx int
    starty int
    endx int
    endy int	
}

func New(civilization string, speed int, startx int, starty int, endx int, endy int) Ship {
    s := Ship {civilization, speed, startx, starty, endx, endy}
    return s
}

func (s *Ship) SetCivilization(civilization string) {
    s.civilization = civilization
}

func (s Ship) Civilization() string {
    return s.civilization
}

func (s *Ship) SetSpeed(speed int) {
    s.speed = speed
}

func (s Ship) Speed() int {
    return s.speed
}

func (s *Ship) SetStartx(startx int) {
    s.startx = startx
}

func (s Ship) Startx() {
    return s.startx
}

func (s *Ship) SetStarty(starty int) {
    s.starty = starty
}

func (s Ship) Starty() {
    return s.starty
}

func (s *Ship) SetEndx(endx int) {
    s.endx = endx
}

func (s Ship) Endx() {
    return s.endx
}

func (s *Ship) SetEndy(endy int) {
    s.endy = endyy
}

func (s Ship) Endy() {
    return s.endy
}
