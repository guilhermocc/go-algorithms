package backtracking

type Object struct {
	value  int
	weight int
}

type Knapsack struct {
	totalWeightCapacity int
	currentValue        int
	currentWeight       int
}

func (b *Knapsack) putObject(object *Object) {
	b.currentWeight += object.weight
	b.currentValue += object.value
}

func (b *Knapsack) removeObject(object *Object) {
	b.currentWeight -= object.weight
	b.currentValue -= object.value
}

var BestSolution = 0

func BinaryKnapSack(objects []*Object, currentObjectIndex int, knapSack *Knapsack) {
	if currentObjectIndex+1 == len(objects) {
		if knapSack.currentValue > BestSolution {
			BestSolution = knapSack.currentValue
		}
		return
	}
	// Evaluation of the first possibility
	BinaryKnapSack(objects, currentObjectIndex+1, knapSack)
	// Evaluation of the second possibility
	if knapSack.currentWeight+objects[currentObjectIndex].weight <= knapSack.totalWeightCapacity {
		knapSack.putObject(objects[currentObjectIndex])
		BinaryKnapSack(objects, currentObjectIndex+1, knapSack)
		knapSack.removeObject(objects[currentObjectIndex])
	}
	return
}
