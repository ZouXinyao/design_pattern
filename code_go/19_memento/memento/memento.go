package memento

type Memento struct {
	content string
}

func (m *Memento) GetSavedContent() string {
	return m.content
}