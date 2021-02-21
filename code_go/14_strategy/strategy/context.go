package strategy

type Context struct {
	strategy Strategy
}

func NewContext(s Strategy) *Context {
	return &Context{strategy: s}
}

func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.DoOperation(a, b)
}