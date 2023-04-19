package object

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) FindOuter(name string) (*Environment, bool) {
	_, ok := e.store[name]
	if ok {
		return e, true
	} else if !ok && e.outer != nil {
		return e.outer.FindOuter(name)
	}
	return nil, false
}

func (e *Environment) Set(name string, val Object) (Object, bool) {
	_, ok := e.store[name]
	if ok {
		return nil, false
	}
	e.store[name] = val
	return val, true
}
