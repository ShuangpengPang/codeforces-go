// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		{
			`[[-2,0],[2,0],[0,2],[0,-2]]`, `2`, 
			`4`, 
		},
		{
			`[[-3,0],[3,0],[2,6],[5,4],[0,9],[7,8]]`, `5`, 
			`5`, 
		},
		{
			`[[-2,0],[2,0],[0,2],[0,-2]]`, `1`, 
			`1`, 
		},
		{
			`[[1,2],[3,5],[1,-1],[2,3],[4,1],[1,3]]`, `2`, 
			`4`, 
		},
		// TODO 测试参数的下界和上界
		{
			`[[6596,-1720],[37,-1237],[2243,1289],[6499,1860],[-9932,-307],[-740,7499],[7139,232],[4748,-623],[7448,-9396],[-6738,-5464],[-7338,5717],[-7773,-9791],[906,8441],[-7250,-1932],[-6035,-4705],[-7319,-4499],[-4363,-806],[6572,2055],[312,-6791],[-5288,9215],[3897,1546],[-5402,7559],[1806,-892],[7306,-1581],[-1644,8544],[-9907,-9400],[9825,-5920],[5984,-8576],[4398,7423],[2264,-7300],[-1143,-5084],[-8524,-9409],[8321,-1752],[168,9691],[3161,-6205],[8206,-9002],[-2982,-241],[7067,-3743],[-3032,-9940],[-6558,1527],[-4276,1034],[-714,7003],[-3461,5580],[1165,112],[-3802,-3650],[-9756,-4006],[-7890,-1356],[-6671,9159],[-1395,-4835],[9172,-5083],[5600,-8131],[-1235,6549],[7631,-2481],[8910,3914],[-2836,-1792],[-6256,9600],[4034,4502],[-1543,5934],[6195,2294],[-148,7633],[-8855,7000],[1310,-1715],[1291,-2927],[6712,2151],[-6932,4832],[4780,7065],[-648,-6109],[2639,506],[7916,-9521],[-2591,3575],[8015,7559],[589,-4130],[-6927,5167],[5306,1944],[-1796,2463],[-583,1534],[278,-4239],[-2055,332],[-3678,2510],[-6248,6121],[5984,9979],[4531,-4380],[5147,8960],[-8223,-8222],[-193,-5365],[9564,4708],[-9944,8781],[7892,1721],[688,-8216],[6988,4908],[1435,4774],[6136,4612],[4193,-1752],[4416,3460],[-7738,-4820]]`, `4196`,
			`23`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, numPoints, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-189/problems/maximum-number-of-darts-inside-of-a-circular-dartboard/