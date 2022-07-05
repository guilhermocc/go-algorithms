package search

import (
	"go-algorithms/graphs"
	"testing"
)

func Test_depthSearchRec(t *testing.T) {
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
			name: "depthSearchRec with first node as target",
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
			name: "depthSearchRec with a non existent target",
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
			name: "depthSearchRec with nested node",
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
			name: "depthSearchRec with no-optimal solution",
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
			want: 3,
		},
		{
			name: "depthSearchRec with first neighbour node",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := depthSearchRec(tt.args.currentNode, tt.args.targetNodeID); got != tt.want {
				t.Errorf("depthSearchRec() = %v, want %v", got, tt.want)
			}
		})
	}
}
