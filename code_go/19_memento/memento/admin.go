package memento

type Admin struct {
	mementoArray []*Memento
}

func (a *Admin) Add(m *Memento) {
	a.mementoArray = append(a.mementoArray, m)
}

func (a *Admin) GetMemento(index int) *Memento {
	return a.mementoArray[index]
}

func NewAdmin() *Admin {
	return &Admin{
		mementoArray: make([]*Memento, 0),
	}
}