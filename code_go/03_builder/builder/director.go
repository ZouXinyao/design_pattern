package builder

type director struct {
	builder pcBuilder
}

func NewDirector(builder pcBuilder) *director {
	return &director{builder: builder}
}

func (dir *director) Construct(cpu string, gpu string, mainBoard string, power string) {
	dir.builder.setCpu(cpu)
	dir.builder.setGpu(gpu)
	dir.builder.setMainBoard(mainBoard)
	dir.builder.setPower(power)
}