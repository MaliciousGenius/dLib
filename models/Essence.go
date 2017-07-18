package models

type Essence struct {
	Id      string
	Name    string
	Content string
	Time    string
}

func NewEssence(id, name, content, time string) *Essence {
	return &Essence{id, name, content, time}
}
