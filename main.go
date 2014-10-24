package main

import (
	"flag"
	"fmt"
)

type statistic struct {
	lambda float64
	mu     float64
}

var l, m *float64

func init() {
	l = flag.Float64("l", 0.0, "lambda")
	m = flag.Float64("m", 0.0, "Mu")
	
}

func main() {
        flag.Parse()
	stats := statistic{*l, *m}

	fmt.Printf("%s", stats.trafficIntensity())
}

func (stats statistic) trafficIntensity() (result string) {
	rho := stats.lambda / stats.mu
	fmt.Sprintf(result, "ρ = λ / μ = %5.2f", rho)
	return
}
