package main

import "github.com/kingparra/myniceprogram/helpers"

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(10)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
}
