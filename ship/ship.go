package ship

type Ship struct {
    civilization string
    speed int
    startx int
    starty int
    endx int
    endy int
    currentx int
    currenty int
    landed bool	
}

func New(civilization string, speed int, startx int, starty int, endx int, endy int, currentx int, currenty int, landed bool) Ship {
    s := Ship {civilization, speed, startx, starty, endx, endy, currentx, currenty, landed}
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

func (s Ship) Startx() int {
    return s.startx
}

func (s *Ship) SetStarty(starty int) {
    s.starty = starty
}

func (s Ship) Starty() int {
    return s.starty
}

func (s *Ship) SetEndx(endx int) {
    s.endx = endx
}

func (s Ship) Endx() int {
    return s.endx
}

func (s *Ship) SetEndy(endy int) {
    s.endy = endy
}

func (s Ship) Endy() int {
    return s.endy
}

func (s *Ship) SetCurrentx(currentx int) {
    s.currentx = currentx
}

func (s Ship) Currentx() int {
    return s.currentx
}

func (s *Ship) SetCurrenty(currenty int) {
    s.currenty = currenty
}

func (s Ship) Currenty() int {
    return s.currenty
}

func (s *Ship) SetLanded(landed bool) {
    s.landed = landed
}

func (s Ship) Landed() bool {
    return s.landed
}
