package statistics

import (
	"errors"
	"math"
)

/////////////////////////////////////////////////////////////////////////////
//                                Struct  MM1                              //
/////////////////////////////////////////////////////////////////////////////

//Stat is a struct that have the two input values
//used in the calc operations.
type Stat struct {
	Lambda float64
	Mu     float64
}

//////////////////////////////////////////////////////////////////////////////
//                    Performance Measurement Formulas                      //
//////////////////////////////////////////////////////////////////////////////

//TrafficIntensity calculate the minimum amount of servers
//needed to maintain the incoming of jobs to que queue.
//Queue type M/M/1
func (stat Stat) TrafficIntensity() (result float64, err error) {
	if stat.Mu == 0 {
		return 0, errors.New("Can't divide by zero")
	}

	rho := stat.Lambda / stat.Mu
	result = rho
	return result, nil
}

//ZeroJobsInSystem calculate the
//TODO: add doc to this function
func (stat Stat) ZeroJobsInSystem() (result float64, err error) {
	p, er := stat.TrafficIntensity()
	if er != nil {
		return 0, er
	}

	p0 := 1 - p
	result = p0

	return result, nil
}

//MeanNumberJobsInSystem is a function that calculate the mean number of jobs
//in the system.
func (stat Stat) MeanNumberJobsInSystem() (result float64, err error) {
	p, er := stat.TrafficIntensity()
	if er != nil {
		return 0, er
	}

	pminus, e := stat.ZeroJobsInSystem()
	if e != nil {
		return 0, e
	}

	result = p / pminus

	return result, nil
}

//TODO ADD COMMENT
func (stat Stat) VarianceNumberJobsInSystem() (result float64, err error) {
	p, er := stat.TrafficIntensity()
	if er != nil {
		return 0, er
	}

	pminus, e := stat.ZeroJobsInSystem()
	if e != nil {
		return 0, e
	}

	if pminus == 0 {
		return 0, errors.New("Can't divide by zero")
	}

	result = p / math.Pow(pminus, 2)

	return result, nil
}

//TODO: ADD COMMENT
func (stat Stat) MeanNumberJobsInQueue() (result float64, err error) {
	p, er := stat.TrafficIntensity()
	if er != nil {
		return 0, er
	}

	pminus, e := stat.ZeroJobsInSystem()
	if e != nil {
		return 0, e
	}

	if pminus == 0 {
		return 0, errors.New("Can't divide by zero")
	}

	result = (math.Pow(p, 2)) / pminus

	return result, nil
}
