//////////////////////////////////////////////////
//             author: Raphael Jardim Lopes     //
//             all rights reserved              //
//             date: 24/10/2014                 //
//////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

///////////////////////////////////////////////////
//                Structs                        //
///////////////////////////////////////////////////
type statistic struct {
	lambda float64
	mu     float64
}

///////////////////////////////////////////////////
//                Main function                  //
///////////////////////////////////////////////////
func main() {
	app := cli.NewApp()
	app.Name = "Queue Modeling Tool"
	app.Usage = "Gives a list with the formulas and answers with the given λ and μ"

	//flags
	app.Flags = []cli.Flag{
		cli.Float64Flag{
			Name:  "lambda",
			Value: 0.0,
			Usage: "lambda measured",
		},
		cli.Float64Flag{
			Name:  "mu",
			Value: 1.0,
			Usage: "mu measured",
		},
	}

	//actions
	app.Action = func(c *cli.Context) {

		l := c.Float64("lambda")
		m := c.Float64("mu")

		stats := statistic{lambda: l, mu: m}

		fmt.Printf("%s", stats.trafficIntensity())

	}
	app.Run(os.Args)
}

///////////////////////////////////////////////////
//               Calculation functions           //
///////////////////////////////////////////////////

func (stats statistic) trafficIntensity() (result string) {
	rho := stats.lambda / stats.mu
	result = fmt.Sprintf("ρ = λ / μ = %5.2f", rho)
	return result
}

