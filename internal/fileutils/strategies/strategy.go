package strategies

type CounterStrategy interface {
	Count(string) int
}
