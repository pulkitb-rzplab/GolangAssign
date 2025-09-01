package main

import "fmt"

type pay interface {
	calculateSal() float32
}

type fte struct {
	days int
	rate float32
}

type ctr struct {
	days int
	rate float32
}

type frelan struct {
	hours int
	rate  float32
}

func (f fte) calculateSal() float32 {
	if f.days == 0 {
		return 0
	} else {
		return float32(f.days) * f.rate
	}

}

func (c ctr) calculateSal() float32 {
	if c.days == 0 {
		return 0
	} else {
		return float32(c.days) * c.rate
	}
}
func (fl frelan) calculateSal() float32 {
	if fl.hours == 0 {
		return 0
	} else {
		return float32(fl.hours) * fl.rate
	}
}

func calcsal(p pay) {
	fmt.Println("Calculating salary...")
	fmt.Println(p.calculateSal())
}

//func main() {

	ft := fte{days: 20, rate: 150}
	ctrtor := ctr{days: 15, rate: 80}
	freelance := frelan{hours: 35, rate: 60}
	calcsal(ft)
	calcsal(ctrtor)
	calcsal(freelance)
}
