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
    p.civilization = civilization
}

func (s Ship) Civilization() string {
    return p.civilization
}

func (s *Ship) SetSpeed(speed int) {
    p.speed = speed
}

func (s Ship) Appearance() string {
    return p.speed
}

