//////////////////////////////////////////////////
//             author: Raphael Jardim Lopes     //
//             all rights reserved              //
//             date: 24/10/2014                 //
//////////////////////////////////////////////////

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/raphaeljlps/queue-modeling/statistics"
)

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

		stats := statistics.Stat{Lambda: l, Mu: m}
		result, err := stats.TrafficIntensity()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		fmt.Printf("ρ = λ / μ = %.4f", result)
		
		result, err := stats.ZeroJobsInSystem()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("p0 = 1 - p = %.4f", result)
		}
	}
	app.Run(os.Args)
}
