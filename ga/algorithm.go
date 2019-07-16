package ga

type IndividualFactory func() Individual

type AlgorithmConfig struct {
	IndividualFactory
	PopulationSize      int
	GenerationsPerEpoch int
	Epochs              int
	MutationProbability float64
}

type Algorithm struct {
	config     AlgorithmConfig
	population []Individual
}

func NewAlgorithm(config AlgorithmConfig) *Algorithm {
	return &Algorithm{config: config}
}

func (a *Algorithm) Execute(inputs, outputs []interface{}) {
	a.population = bootstrapPopulation(a.config.IndividualFactory, a.config.PopulationSize)
	for _, i := range a.population {
		i.CalculateFitness(inputs, outputs)
	}
	for epoch := 0; epoch < a.config.Epochs; epoch++ {
		for gen := 0; gen < a.config.GenerationsPerEpoch; gen++ {
			// Mate
			// double population
			// Mutate
			// Natural Selection
			// always keep the best and the worst, the other ones should be a random tournament
		}
	}
}

func bootstrapPopulation(factory IndividualFactory, size int) []Individual {
	population := make([]Individual, 0, size)
	for i := 0; i < size; i++ {
		population = append(population, factory())
	}
	return population
}
