package main

import (
	"flag"
	"fmt"
)

func main() {
	var lambda = flag.Float64("l", 0.0, "lambda")
	var mu = flag.Float64("m", 0.0, "Mu")
	flag.Parse()

	rho := *lambda / *mu

	fmt.Printf("ρ = λ / μ  => ρ = %.2f / %.2f => ρ = %.4f\n", *lambda, *mu, rho)
}
