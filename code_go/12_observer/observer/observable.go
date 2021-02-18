package observer

type Subject struct {
	observers []Observer // 观察者切片
	context   string     // 信息
}

// 创建新的被观察者
func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

// 被观察者类中添加观察者
func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

// 发布更改信息
func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}

func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}
