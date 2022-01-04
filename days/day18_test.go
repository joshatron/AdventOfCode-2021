package days

import (
	"testing"
)

func TestParseSnailfishNumber(t *testing.T) {
	tests := []struct {
		input string
		want  SnailfishNumber
	}{
		{"[3,4]", SnailfishNumber{3, nil, 4, nil}},
		{"[[1,2],4]", SnailfishNumber{-100, &SnailfishNumber{1, nil, 2, nil}, 4, nil}},
		{"[2,[3,4]]", SnailfishNumber{2, nil, -100, &SnailfishNumber{3, nil, 4, nil}}},
		{"[[1,2],[3,4]]", SnailfishNumber{-100, &SnailfishNumber{1, nil, 2, nil}, -100, &SnailfishNumber{3, nil, 4, nil}}},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			ans := parseSnailfishNumber(test.input)
			if test.want.xLiteral != ans.xLiteral || test.want.yLiteral != ans.yLiteral {
				t.Error("from", test.input, ": got", ans, ", want", test.want)
			}
			if test.want.xLiteral == -100 {
				if test.want.xSnailfishNumber.xLiteral != ans.xSnailfishNumber.xLiteral ||
					test.want.xSnailfishNumber.yLiteral != ans.xSnailfishNumber.yLiteral {
					t.Error("from", test.input, ": got", ans.xSnailfishNumber, ", want", test.want.xSnailfishNumber)
				}
			}
			if test.want.yLiteral == -100 {
				if test.want.ySnailfishNumber.xLiteral != ans.ySnailfishNumber.xLiteral ||
					test.want.ySnailfishNumber.yLiteral != ans.ySnailfishNumber.yLiteral {
					t.Error("from", test.input, ": got", ans.ySnailfishNumber, ", want", test.want.ySnailfishNumber)
				}
			}
		})
	}
}
