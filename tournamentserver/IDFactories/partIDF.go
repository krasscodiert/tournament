package IDFactories

type PartIDF struct {
	currentID int
}

func (p* PartIDF) GetID() int {
	p.currentID += 1
	current := p.currentID
	return current
}