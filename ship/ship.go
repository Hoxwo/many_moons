package ship

type Ship struct {
    civilization string
    speed int
}

func New(civilization string, speed int) Ship {
    s := Ship {civilization, speed}
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

