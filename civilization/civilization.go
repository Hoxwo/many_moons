package civilization

type Civilization struct {
    name string
    color string
    attack int
    defense int
    navigation int
    government int
    technology int
    resources int
    shipsavailable int 
    maxshipsavailable int // (base value 1 + (1 for tech level 4, 2 for tech level 7, 3 for tech level 10) start phase with these ships available 
    shiptimer int // base value - (multiplier * technology level + resource level)
    colonizationtime int //base value - (multiplier * government level + resource level)
}

func New(civilization string, speed int, attack int, defense int, navigation int, government int, technology int, resources int, shipsavailable int, maxshipsavailable int, shiptimer int, colonizationtime int) Civilization {
    c := Civilization {civilization, speed, attack, defense, navigation, government, technology, resources, shipsavilable, maxshipsavailable, shiptimer, colonizationtime}
    return c
}

func (c *Civilization) SetName(name string) {
    c.name = civilization
}

func (c Civilization) Name() string {
    return c.name
}

func (c *Civilization) SetColor(name string) {
    c.color = color
}

func (c Civilization) Color() string {
    return c.color
}

func (c *Civilization) SetAttack(attack int) {
    c.attack = attack
}

func (c Civilization) Attack() int {
    return c.attack
}

func (c *Civilization) SetDefense(defense int) {
    c.defense = defense
}

func (c Civilization) Defense() int {
    return c.defense
}

func (c *Civilization) SetNavigation(navigation int) {
    c.navigation = navigation
}

func (c Civilization) Navigation() int {
    return c.navigation
}

func (c *Civilization) SetGovernment(government int) {
    c.government = government
}

func (c Civilization) Government() int {
    return c.government
}

func (c *Civilization) SetTechnology(technology int) {
    c.technology = technology
}

func (c Civilization) Technology() int {
    return c.technology
}

func (c *Civilization) SetResources(resources int) {
    c.resources = resources
}

func (c Civilization) Resources() int {
    return c.resources
}

func (c *Civilization) SetShipsavailable(shipsavailable int) {
    c.shipsavailable = shipsavailable
}

func (c Civilization) Shipsavailable() int {
    return c.shipsavailable
}

func (c *Civilization) SetMaxshipsavailable(maxshipsavailable int) {
    c.maxshipsavailable = maxshipsavailable
}

func (c Civilization) Maxshipsavailable() int {
    return c.maxshipsavailable
}

func (c *Civilization) SetShiptimer(shiptimer int) {
    c.shiptimer = shiptimer
}

func (c Civilization) Shiptimer() int {
    return c.shiptimer
}

func (c *Civilization) SetColonizationtime(colonizationtime int) {
    c.colonizationtime = colonizationtime
}

func (c Civilization) Colonizationtime() int {
    return c.colonizationtime
}





