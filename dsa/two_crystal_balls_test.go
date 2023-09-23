package dsa

import "testing"


func TestTwoCrystalBalls(t *testing.T) {
    breaks := []bool{false, false, false, false, true, true, true, true, true, true}
    
    result := TwoCrystalBalls(breaks)
    expect := 4
    
    if result != expect {
        t.Errorf("expected result is %d, but the actual result is %d", expect, result)
    }
}
