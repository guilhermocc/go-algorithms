package search

import (
	"go-algorithms/graphs"
	"testing"
)

func Test_breadthSearch(t *testing.T) {
	type args struct {
		currentNode  *graphs.Node
		targetNodeID int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "breadthSearch with first node as target",
			args: args{
				targetNodeID: 0,
				currentNode: graphs.NewNode(0, []*graphs.Node{
					graphs.NewNode(1, []*graphs.Node{
						graphs.NewNode(2, []*graphs.Node{}),
						graphs.NewNode(3, []*graphs.Node{}),
					}),
					graphs.NewNode(4, []*graphs.Node{
						graphs.NewNode(5, nil),
						graphs.NewNode(6, nil),
					}),
				}),
			},
			want: 0,
		},
		{
			name: "breadthSearch with a non existent target",
			args: args{
				targetNodeID: 10,
				currentNode: graphs.NewNode(0, []*graphs.Node{
					graphs.NewNode(1, []*graphs.Node{
						graphs.NewNode(2, []*graphs.Node{}),
						graphs.NewNode(3, []*graphs.Node{}),
					}),
					graphs.NewNode(4, []*graphs.Node{
						graphs.NewNode(5, nil),
						graphs.NewNode(6, nil),
					}),
				}),
			},
			want: -1,
		},
		{
			name: "breadthSearch with nested node",
			args: args{
				targetNodeID: 6,
				currentNode: graphs.NewNode(0, []*graphs.Node{
					graphs.NewNode(1, []*graphs.Node{
						graphs.NewNode(2, []*graphs.Node{}),
						graphs.NewNode(3, []*graphs.Node{}),
					}),
					graphs.NewNode(4, []*graphs.Node{
						graphs.NewNode(5, nil),
						graphs.NewNode(6, nil),
					}),
				}),
			},
			want: 2,
		},
		{
			name: "breadthSearch with first neighbour node",
			args: args{
				targetNodeID: 4,
				currentNode: graphs.NewNode(0, []*graphs.Node{
					graphs.NewNode(1, []*graphs.Node{
						graphs.NewNode(2, []*graphs.Node{}),
						graphs.NewNode(3, []*graphs.Node{}),
					}),
					graphs.NewNode(4, []*graphs.Node{
						graphs.NewNode(5, nil),
						graphs.NewNode(6, nil),
					}),
				}),
			},
			want: 1,
		},
		{
			name: "breadthSearch with nested node",
			args: args{
				targetNodeID: 6,
				currentNode: graphs.NewNode(0, []*graphs.Node{
					graphs.NewNode(1, []*graphs.Node{
						graphs.NewNode(2, []*graphs.Node{}),
						graphs.NewNode(3, []*graphs.Node{}),
					}),
					graphs.NewNode(4, []*graphs.Node{
						graphs.NewNode(5, nil),
						graphs.NewNode(6, nil),
					}),
				}),
			},
			want: 2,
		},
		{
			name: "depthSearchRec with optimal solution",
			args: args{
				targetNodeID: 2,
				currentNode: graphs.NewNode(5, []*graphs.Node{
					graphs.NewNode(3, []*graphs.Node{}),
					graphs.NewNode(4, []*graphs.Node{
						graphs.NewNode(1, []*graphs.Node{graphs.NewNode(2, nil)}),
						graphs.NewNode(2, nil),
					}),
				}),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := breadthSearch(tt.args.currentNode, tt.args.targetNodeID); got != tt.want {
				t.Errorf("breadthSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
