package planet

type Planet struct {
    name string
    appearance string
    resources int
    occupied string
    planettype string
    xcoord int
    ycoord int
    homeplanet bool
    timetocolonize int
}

func New(name string, appearance string, resources int, occupied string, planettype string, xcoord int, ycoord int, homeplanet bool, timetocolonize int) Planet {
    p := Planet {name, appearance, resources, occupied, planettype, xcoord, ycoord, homeplanet, timetocolonize}
    return p
}

func (p *Planet) SetName(name string) {
    p.name = name
}

func (p Planet) Name() string {
    return p.name
}

func (p *Planet) SetAppearance(appearance string) {
    p.appearance = appearance
}

func (p Planet) Appearance() string {
    return p.appearance
}

func (p *Planet) SetResources(resources int) {
    p.resources = resources
}

func (p Planet) Resources() int {
    return p.resources
}

func (p *Planet) SetOccupied(occupied string) {
    p.occupied = occupied
}

func (p Planet) Occupied() string {
    return p.occupied
}

func (p *Planet) SetPlanettype(planettype string) {
    p.planettype = planettype
}

func (p Planet) Planettype() string {
    return p.planettype
}

func (p *Planet) SetXcoord(xcoord int) {
    p.xcoord = xcoord
}

func (p Planet) Xcoord() int {
    return p.xcoord
}

func (p *Planet) SetYcoord(ycoord int) {
    p.ycoord = ycoord
}

func (p Planet) Ycoord() int {
    return p.ycoord
}

func (p *Planet) SetHomeplanet(homeplanet bool) {
    p.homeplanet = homeplanet
}

func (p Planet) Homeplanet() bool {
    return p.homeplanet
}

func (p *Planet) SetTimetocolonize(timetocolonize int) {
    p.timetocolonize = timetocolonize
}

func (p Planet) Timetocolonize() int {
    return p.timetocolonize
}

