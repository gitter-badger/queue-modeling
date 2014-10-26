package statistics

import "testing"

///////////////////////////////////////////////////
//           Traffic Intensity Tests             //
///////////////////////////////////////////////////

func TestTrafficIntensityOK(t *testing.T) {
	stat := Stat{Lambda: 10, Mu: 10}

	result, _ := stat.TrafficIntensity()
	value := 1.0
	if result != value {
		t.Error("For {Lambda: ", stat.Lambda, " Mu: ", stat.Mu, "} ",
			"Expected ", value,
			"Got ", result)
	}

	stat.Lambda = 20

	result, _ = stat.TrafficIntensity()
	value = 2.0

	if result != value {
		t.Error("For {Lambda: ", stat.Lambda, " Mu: ", stat.Mu, "} ",
			"Expected ", value,
			"Got ", result)
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
	t.Error("Error")
}
