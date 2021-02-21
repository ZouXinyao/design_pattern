package mediator

type Mediator interface {
	processingReq(Request) bool
	notifyAboutRelease()
}

type ConcreteMediator struct {
	isFree   bool
	reqQueue []Request
}

func NewMediator() *ConcreteMediator {
	return &ConcreteMediator{
		isFree: true,
		reqQueue: make([]Request, 0),
	}
}

func (c *ConcreteMediator) processingReq(req Request) bool {
	if c.isFree {
		c.isFree = false
		return true
	}

	if req.isHighPriority() {
		c.reqQueue = append([]Request{req}, c.reqQueue...)
	} else {
		c.reqQueue = append(c.reqQueue, req)
	}
	return false
}

func (c *ConcreteMediator) notifyAboutRelease() {
	c.isFree = true
	if len(c.reqQueue) > 0 {
		req := c.reqQueue[0]
		c.reqQueue = c.reqQueue[1:]
		req.Process()
	}
}
