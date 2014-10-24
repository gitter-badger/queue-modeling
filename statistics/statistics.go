package statistics

import "errors"

///////////////////////////////////////////////////
//                 Struct                        //
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
//needed to maintain the incoming of jobs to que queue
//Queue M/M/1
func (stat Stat) TrafficIntensity() (result float64, err error) {
	if stat.Mu == 0 {
		return 0.0, errors.New("Can't divide by zero")
	}

	rho := stat.Lambda / stat.Mu
	result = rho
	//result = fmt.Sprintf("ρ = λ / μ = %5.2f", rho)
	return result, nil
}
