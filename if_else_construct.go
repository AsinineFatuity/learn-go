package main

import "math/rand"

//initialization statement with condition evaluation

func InitializaAndEvaluate() bool {
	maxVal := 10
	if val := generateRandomNumber(); val > maxVal {
		return true
	}
	return false

}

func generateRandomNumber() int {
	randomGenerator := rand.New(rand.NewSource(42))
	return randomGenerator.Intn(100)
}
