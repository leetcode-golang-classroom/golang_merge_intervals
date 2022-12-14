package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	for idx := 0; idx < b.N; idx++ {
		merge(intervals)
	}
}
func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "intervals = [[1,3],[2,6],[8,10],[15,18]]",
			args: args{intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name: "intervals = [[1,4],[4,5]]",
			args: args{intervals: [][]int{{1, 4}, {4, 5}}},
			want: [][]int{{1, 5}},
		},
		{
			name: "intervals = [[1,4],[0,4]]",
			args: args{intervals: [][]int{{1, 4}, {0, 4}}},
			want: [][]int{{0, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
