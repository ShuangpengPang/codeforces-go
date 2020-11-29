// Code generated by copypasta/template/nowcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`2`,`[1,2,3,4,5]`,
			`14`,
		},
		{
			`3`,`[1,2,3,4,5]`,
			`13`,
		},
	}
	targetCaseNum := -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, tree6, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://ac.nowcoder.com/acm/contest/9247/b