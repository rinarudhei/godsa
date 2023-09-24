package dsa

import (
    "math" 
    "errors"
)

func TwoCrystalBalls (breaks []bool) (int, error) {
    length := len(breaks)

    err := validateInput(breaks)
    if err != nil {
        return -1, err
    }

    stepSize := countStepSize(length) 
    index := -1 
    isFound := false 

    for i := stepSize;i < length; i += stepSize {
        if breaks[i] {
            isFound = true
            index = i
            break
        }
    }

    if !isFound {
        return -1, nil
    }
    
    for i := index; i >= 0; i-- {
        if !breaks[i] {
            return i + 1, nil
        }
    }

    return -1, nil
 }

 
 func countStepSize(length int) int {
    floatLen := float64(length)
    sqrt := math.Sqrt(floatLen)
    rounded := math.Floor(sqrt)
    intStep := int(rounded) - 1
    
    if intStep < 1 {
        intStep = 1
    }

    return intStep
 }

 func validateInput(breaks []bool) error {
    length := len(breaks)
    if length <= 1 {
        return errors.New("length is less or equal than 1") 
    }

    if breaks[0] {
        return errors.New("breaking since first index") 
    }

    return nil
 }
