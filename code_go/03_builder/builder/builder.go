package builder

type pcBuilder interface {
	setCpu(str string)
	setGpu(str string)
	setMainBoard(str string)
	setPower(str string)
	GetResult() string
}

type DellBuilder struct {
	cpu string
	gpu string
	mainBoard string
	power string
}

func (builder *DellBuilder) setCpu(cpuBrand string) {
	builder.cpu = cpuBrand
}

func (builder *DellBuilder) setGpu(gpuBrand string) {
	builder.gpu = gpuBrand
}

func (builder *DellBuilder) setMainBoard(boardBrand string) {
	builder.mainBoard = boardBrand
}

func (builder *DellBuilder) setPower(powerBrand string) {
	builder.power = powerBrand
}

func (builder *DellBuilder) GetResult() string {
	return builder.cpu + " " + builder.gpu + " " + builder.mainBoard + " " + builder.power
}

type LenovoBuilder struct {
	cpu string
	gpu string
	mainBoard string
	power string
}

func (builder *LenovoBuilder) setCpu(cpuBrand string) {
	builder.cpu = cpuBrand
}

func (builder *LenovoBuilder) setGpu(gpuBrand string) {
	builder.gpu = gpuBrand
}

func (builder *LenovoBuilder) setMainBoard(boardBrand string) {
	builder.mainBoard = boardBrand
}

func (builder *LenovoBuilder) setPower(powerBrand string) {
	builder.power = powerBrand
}

func (builder *LenovoBuilder) GetResult() string {
	return builder.cpu + " " + builder.gpu + " " + builder.mainBoard + " " + builder.power
}