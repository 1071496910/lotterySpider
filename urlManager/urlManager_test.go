package urlManager

import (
	"github.com/willf/bloom"
	"testing"
)

func Test_bloomUFilter(t *testing.T) {
	type args struct {
		url  string
		getK string
	}
	tests := []struct {
		name   string
		args   args
		expect string
	}{
		{
			"1",
			args{
				"1",
				"2",
			},
			"",
		},
		{
			"2",
			args{
				"2",
				"3",
			},
			"",
		}, {
			"3",
			args{
				"3",
				"4",
			},
			"",
		}, {
			"4",
			args{
				"4",
				"5",
			},
			"",
		}, {
			"5",
			args{
				"5",
				"6",
			},
			"",
		}, {
			"7",
			args{
				"7",
				"8",
			},
			"",
		},
	}
	b := bloom.New(10, 2)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b.Add([]byte(tt.args.url))

			if b.Test([]byte(tt.args.getK)) {
				t.Errorf("test error, k: %v", tt.args.getK)
			}
		})
	}
}
