package dsa

import "testing"

type testcase struct {
    breaks []bool
    expect int
}

var testcases = []testcase{
    {[]bool{false, false, false, false, true, true, true, true, true}, 4},
    //already breaking from the beginning
    {[]bool{true, true, true, true, true, true, true, true, true, true}, -1},
    
    {[]bool{false, false, false, false, false, true}, 5},
    //invalid input
    {[]bool{}, -1},
    
    {[]bool{false, false}, -1},
    //invalid input
    {[]bool{false}, -1},
    //invalid input
    {[]bool{true}, -1},
}

func TestTwoCrystalBalls(t *testing.T) {
    for _, test := range testcases {
        if result, _ := TwoCrystalBalls(test.breaks); result != test.expect {
            t.Errorf("expected result is %d, but the actual result is %d", test.expect, result)
        }
    } 
}
