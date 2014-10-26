package statistics

import (
	"fmt"
	"testing"
)

///////////////////////////////////////////////////
//                   Utility                     //
///////////////////////////////////////////////////
func errorOutput(lambda, mu, correctValue, result float64) (out string) {
	out = fmt.Sprintf("For {Lambda: %.4f Mu: %.4f} Expected %.4f Got %.4f", lambda, mu, correctValue, result)
	return out
}

///////////////////////////////////////////////////
//           Traffic Intensity Tests             //
///////////////////////////////////////////////////

func TestTrafficIntensityOK(t *testing.T) {
	stat := Stat{Lambda: 10, Mu: 10}

	result, _ := stat.TrafficIntensity()
	correctValue := 1.0
	if result != correctValue {
		t.Error(errorOutput(stat.Lambda, stat.Mu, correctValue, result))
	}

	stat.Lambda = 20

	result, _ = stat.TrafficIntensity()
	correctValue = 2.0

	if result != correctValue {
		t.Error(errorOutput(stat.Lambda, stat.Mu, correctValue, result))
	}
}

func TestTrafficIntensityError(t *testing.T) {
	stat := Stat{Lambda: 10, Mu: 0}

	_, err := stat.TrafficIntensity()

	if err == nil {
		t.Error("For {Lambda: ", stat.Lambda, " Mu: ", stat.Mu, "} ",
			"Expected error not nil",
			"Got error nil")
	}
}

///////////////////////////////////////////////////
//            Zero Jobs In System Tests          //
///////////////////////////////////////////////////

func TestZeroJobsInSystemOK(t *testing.T) {
	stat := Stat{Lambda: 20, Mu: 10}

	result, _ := stat.ZeroJobsInSystem()
	correctValue := -1.0
	if result != correctValue {
		t.Error(errorOutput(stat.Lambda, stat.Mu, correctValue, result))
	}
}
