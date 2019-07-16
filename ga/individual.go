package ga

type Individual interface {
	CalculateFitness(inputs, outputs []interface{})
	Execute(input interface{}) (output interface{})
	Fitness() float64
}
