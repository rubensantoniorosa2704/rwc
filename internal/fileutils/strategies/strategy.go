package strategies

type CounterStrategy interface {
	Count(<-chan string) (int, error)
}
