// Cálculo simultâneo de pi.
// Veja https://goo.gl/la6Kli.
//
// Isso demonstra a capacidade de Go de lidar com
// grande número de processos concorrentes.
// É uma forma irracional de calcular pi.
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(pi(5000))
}

// pi lança n goroutines para calcular um
// aproximação de pi.
func pi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k <= n; k++ {
		go term(ch, float64(k))
	}
	f := 0.0
	for k := 0; k <= n; k++ {
		f += <-ch
	}
	return f
}

func term(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}
