package planet

type Planet struct {
    name string
    appearance string
    resources int
    occupied string
    planettype string
}

func New(name string, appearance string, resources int, occupied string, planettype string) Planet {
    p := Planet {name, appearance, resources, occupied, planettype}
    return p
}

func (p *Planet) SetName(name string) {
    p.name = name
}

func (p Planet) Name() string {
    return p.name
}

func (p *Planet) SetAppearance(name string) {
    p.appearance = appearance
}

func (p Planet) Appearance() string {
    return p.appearance
}

func (p *Planet) SetResources(name string) {
    p.resources = resources
}

func (p Planet) Resources() string {
    return p.resources
}

func (p *Planet) SetOccupied(name string) {
    p.occupied = occupied
}

func (p Planet) Occupied() string {
    return p.occupied
}

func (p *Planet) SetPlanettype(name string) {
    p.planettype = planettype
}

func (p Planet) Planettype() string {
    return p.planettype
}
