package dsa

import (
    "math" 
)

func TwoCrystalBalls (breaks []bool) int {
    length := len(breaks)
    stepSize := int(math.Floor(math.Sqrt(float64(length))))
    
    index := -1 
    isFound := false 
    for i := stepSize;i < length; i += stepSize {
        if breaks[i] {
            isFound = true
            index = i
            break;
        }
    }

    if !isFound {
        return -1
    }

    for i := index; i > 0; i-- {
        if !breaks[i] {
            return i + 1
        }
    }

    return -1
 }
