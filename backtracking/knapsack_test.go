package backtracking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryKnapSack(t *testing.T) {
	type args struct {
		objects  []*Object
		knapSack *Knapsack
	}
	tests := []struct {
		name          string
		args          args
		expectedValue int
	}{
		{
			name: "TestBinaryKnapSack capacity 10",
			args: args{
				objects: []*Object{
					{value: 5, weight: 3},
					{value: 10, weight: 5},
					{value: 15, weight: 7},
					{value: 20, weight: 9},
					{value: 25, weight: 11},
				},
				knapSack: &Knapsack{totalWeightCapacity: 10},
			},
			expectedValue: 20,
		},
		{
			name: "TestBinaryKnapSack capacity 4",
			args: args{
				objects: []*Object{
					{value: 8, weight: 1},
					{value: 4, weight: 2},
					{value: 0, weight: 3},
					{value: 5, weight: 2},
					{value: 3, weight: 2},
				},
				knapSack: &Knapsack{totalWeightCapacity: 4},
			},
			expectedValue: 13,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			BestSolution = 0
			BinaryKnapSack(test.args.objects, 0, test.args.knapSack)
			assert.Equal(t, test.expectedValue, BestSolution)
		})
	}
}
