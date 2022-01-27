package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type InputParams struct {
	nums   []int32
	target int32
}

type Result struct {
	pos   []int32
	valid bool
}

type TestCase struct {
	input    InputParams
	expected Result
}

func TestTwoSum(t *testing.T) {
	cases := []TestCase{
		{
			input: InputParams{
				nums:   []int32{3, 3},
				target: 6,
			},
			expected: Result{
				pos:   []int32{0, 1},
				valid: true,
			},
		},
		{
			input: InputParams{
				nums:   []int32{3, 2, 4},
				target: 6,
			},
			expected: Result{
				pos:   []int32{1, 2},
				valid: true,
			},
		},
		{
			input: InputParams{
				nums:   []int32{2, 7, 11, 15},
				target: 9,
			},
			expected: Result{
				pos:   []int32{0, 1},
				valid: true,
			},
		},
		{
			input: InputParams{
				nums:   []int32{1, 3, 7, 11, 15, 2},
				target: 9,
			},
			expected: Result{
				pos:   []int32{2, 5},
				valid: true,
			},
		},
		{
			input: InputParams{
				nums:   []int32{2, 7, 11, 15},
				target: 8,
			},
			expected: Result{
				pos:   nil,
				valid: false,
			},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			result, ok := twoSum(tc.input.nums, tc.input.target)

			if ok {
				require.Equal(t, result[0], tc.expected.pos[0])
				require.Equal(t, result[1], tc.expected.pos[1])
			}

			require.Equal(t, ok, tc.expected.valid)

		})
	}

}
