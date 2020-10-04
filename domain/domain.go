package domain

import (
	"math"
)

type service struct {
}

var Storage = []int{0}

func NewFibonacciService() *service {
	return &service{}
}

type FibonacciService interface {
	FibonacciGenerator(request string) int
}

var _ FibonacciService = &service{}


func (s *service) FibonacciGenerator(request string) int {

	if request == "next" {

		latest := Storage[len(Storage)-1]

		if latest == 0 || (len(Storage) >= 2 && Storage[len(Storage)-2] == 0) {

			Storage = append(Storage, 1)

			return 1

		}

		next := Next(latest)

		Storage = append(Storage, next)

		return next

	} else if request == "current" {

		latest := Storage[len(Storage)-1]

		return latest

	} else if request == "previous" {

		if len(Storage) == 1 {

			return -1

		}

		latest := Storage[len(Storage)-1]

		if len(Storage) >= 2 && Storage[len(Storage)-2] == 0 {

			Storage = Storage[:1]

			return 0

		}

		var index int
		for i, v := range Storage {
			if latest == v {
				index = i
			}
		}

		Storage = Storage[:index]

		return Prev(latest)
	}

	return -1
}

func Next(n int) int {

	var num = float64(n) * (1 + math.Sqrt(float64(5))) / 2

	return int(math.Round(num))

}

func  Prev(n int) int {

	var num = float64(n) / ((1 + math.Sqrt(float64(5))) / 2)

	return int(math.Round(num))

}
