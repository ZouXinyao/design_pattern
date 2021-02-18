package prototype

type mainframe struct {
	cpu string
	gpu string
}

func (pc *mainframe) SetCpu(cpuBrand string) {
	pc.cpu = cpuBrand
}

func (pc *mainframe) GetCpu() string {
	return pc.cpu
}

func (pc *mainframe) SetGpu(gpuBrand string) {
	pc.gpu = gpuBrand
}

func (pc *mainframe) GetGpu() string {
	return pc.gpu
}

func (pc *mainframe) clone() *mainframe {
	pcClone := mainframe{}
	pcClone.cpu = pc.cpu
	pcClone.gpu = pc.gpu
	return &pcClone
}

func NewMainFrame() *mainframe {
	return &mainframe{}
}
