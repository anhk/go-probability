package main

import (
	"fmt"
	"github.com/anhk/go-probability"
)

func test001(p *probability.Probability) {

	T, F := 0, 0
	for i := 0; i < 10000; i++ {
		if s := p.Should(); s {
			T++
		} else {
			F++
		}
	}

	fmt.Println("T:", T, "F", F)
}

func main() {
	p := probability.NewProbability(&probability.Option{
		Percent: -2,
	})
	test001(p)

	p.SetPercent(70)
	test001(p)
}
