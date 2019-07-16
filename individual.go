package ga

type Individual interface {
	CalculateFitness(inputs []interface{}, outputs []interface{})
	Execute(input interface{}) (output interface{})
	Fitness() float64
}
