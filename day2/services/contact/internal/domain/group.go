package domain

type Group struct {
	ID   int
	name string
}

func (g *Group) Name() string {
	return g.name
}

func (g *Group) SetName(name string) {
	if len(name) > 250 {
		name = name[:250]
	}

	g.name = name
}

func (g *Group) GetName() string {
	return g.name
}
