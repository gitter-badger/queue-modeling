package statistics

import "fmt"

///////////////////////////////////////////////////
//                Structs                        //
///////////////////////////////////////////////////

//Stat is a struct that have the two input values
//used in the calc operations
type Stat struct {
	Lambda float64
	Mu     float64
}

///////////////////////////////////////////////////
//               Calculation functions           //
///////////////////////////////////////////////////

//TrafficIntensity calculate the minimum amount of servers
//to mantein the incoming of messages to que queue
//Queue M/M/1
func (stat Stat) TrafficIntensity() (result string) {
	rho := stat.Lambda / stat.Mu
	result = fmt.Sprintf("ρ = λ / μ = %5.2f", rho)
	return result
}
