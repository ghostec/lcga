package ga

type Individual interface {
	CalculateFitness(inputs, outputs []interface{})
	Clone() Individual
	Execute(input interface{}) (output interface{})
	Fitness() float64
	Mate(Individual) Individual
}
