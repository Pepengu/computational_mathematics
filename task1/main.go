package main

import (
	"fmt"
	"math"
)

/*Calculating the sin(x) value using the series expansion until expected accuracy reached*/
func sin(x, accuracy float64) float64 {
	term, sum := x, x

	for i := 1; math.Abs(term) >= accuracy; i++ {
		term *= (-1) * x * x / float64((2*i)*(2*i+1))

		sum += term
	}

	return sum
}

/*Calculating the ch(x) value using the series expansion until expected accuracy reached*/
func ch(x, accuracy float64) float64 {
	term, sum := 1., 1.

	for i := 1; math.Abs(term) >= accuracy; i++ {
		term *= x * x / float64((2*i)*(2*i-1))

		sum += term
	}

	return sum
}

/*Calculating the sqrt(x) value using the Heron's method until expected accuracy reached*/
func sqrt(x, accuracy float64) float64 {
	z := math.Max(x, 1)

	for math.Abs(z*z-x) >= accuracy {
		z = (z + x/z) / 2
	}

	return z
}

var (
	eps1 = 1e-6 / (3 * 2.43)
	eps2 = 1e-6 / (3 * 3.6)
	eps3 = 1e-6 / (3 * 1.62)
)

func calculate(start, end, step float64) {
	for x := start; x < end; x += step {
		my := ch(sqrt(x*x+0.3, eps2)/(1+x), eps1) * sin((1+x)/(0.6*x), eps3)
		std := math.Cosh(math.Sqrt(x*x+0.3)/(1+x)) * math.Sin((1+x)/(0.6*x))
		err := math.Abs(my - std)
		fmt.Printf("step = %.2g\t| my_calc = %.10g \t| stand_calc = %.10g\t| error = %.10g\n", x, my, std, err)
	}
}

func main() {
	calculate(0.2, 0.3, 0.01)
}
