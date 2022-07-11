package strategy

// Strategy-Pattern Mode

type IStrategy interface {
	do(int, int) int
}

type add struct{}

// do implement the IStrategy interface and do addition calculation
func (*add) do(a, b int) int {
	return a + b
}

type reduce struct{}

// do implement the IStrategy interface and do reduction calculation
func (*reduce) do(a, b int) int {
	return a - b
}

type Operator struct {
	strategy IStrategy
}

// setStrategy set the field "strategy" in operator instance
func (operator *Operator) setStrategy(strategy IStrategy) {
	operator.strategy = strategy
}

// calculate do the specific calculation based on different implementation
// of strategy and return the result of that calculation
func (operator *Operator) calculate(a, b int) int {
	return operator.strategy.do(a, b)
}
