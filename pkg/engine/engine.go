package engine

type Engine struct {
	Cache Cache
}

func NewEngine() *Engine {
	return &Engine{
		Cache: NewCache(),
	}
}

func (e *Engine) StartGC() {

}
