package idfactories

// PartIDF - PartIDF  struct
type PartIDF struct {
	currentID int
}

//GetID - returns ID
func (p *PartIDF) GetID() int {
	p.currentID++
	current := p.currentID
	return current
}
