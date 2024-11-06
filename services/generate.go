package service

type IDGenerator struct {
	lastID int
}

func NewIDGenerator(initialID int) *IDGenerator {
	return &IDGenerator{
		lastID: initialID,
	}
}

func (g *IDGenerator) GenerateNextID() int {
	g.lastID++
	return g.lastID
}

func (g *IDGenerator) GetLastID() int {
	return g.lastID
}
