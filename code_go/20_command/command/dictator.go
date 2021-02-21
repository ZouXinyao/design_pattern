package command

type Dictator struct {
	command Command
}

func (d *Dictator) Request() {
	d.command.execute()
}

func NewDictator(c Command) *Dictator {
	return &Dictator{
		command: c,
	}
}