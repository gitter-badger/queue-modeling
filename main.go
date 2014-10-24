//////////////////////////////////////////////////
//             author: Raphael Jardim Lopes     //
//             all rights reserved              //
//             date: 24/10/2014                 //
//////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

	"github.com/raphaeljlps/queue-modeling/statistics"
	"github.com/codegangsta/cli"
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

		fmt.Printf("%s", stats.TrafficIntensity())

	}
	app.Run(os.Args)
}
