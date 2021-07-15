package util

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/constant"
)

var coordinateRefs = map[string]int{
	"N": 1,
	"S": -1,
	"E": 1,
	"W": -1,
}

func getFinalNumber(value string) (finalNumber float64, err error) {
	if strings.Contains(value, "/") {
		numbers := strings.Split(value, "/")
		number1, err := strconv.ParseInt(numbers[0], 0, 64)
		if err != nil {
			return 0, err
		}
		number2, err := strconv.ParseInt(numbers[1], 0, 64)
		if err != nil {
			return 0, err
		}
		return float64(number1 / number2), nil
	} else {
		result, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return 0, err
		}
		return result, nil
	}
}

func SeparateStringGPSCoordinate(coordinate string) (degrees float64, minutes float64, seconds float64, err error) {
	splited := strings.Split(coordinate, ",")
	if splited == nil || len(splited) != 3 {
		return 0, 0, 0, NewError(constant.GPSCoordinateError, coordinate)
	}
	for i, value := range splited {
		splited[i] = strings.ReplaceAll(splited[i], "\"", "")
		splited[i] = strings.ReplaceAll(splited[i], "[", "")
		splited[i] = strings.ReplaceAll(splited[i], "]", "")
		re := regexp.MustCompile(`([0-9]\/[0-9])`)
		reAux := regexp.MustCompile(`([0-9])`)
		if !re.MatchString(splited[i]) || !reAux.MatchString(splited[i]) {
			return 0, 0, 0, NewError(constant.GPSCoordinateValueError, value, coordinate)
		}
	}
	degrees, err = getFinalNumber(splited[0])
	if err != nil {
		return 0, 0, 0, err
	}
	minutes, err = getFinalNumber(splited[1])
	if err != nil {
		return 0, 0, 0, err
	}
	seconds, err = getFinalNumber(splited[2])
	if err != nil {
		return 0, 0, 0, err
	}
	return degrees, minutes, seconds, nil
}

func GetDecimalDegreeCoordinate(latitude string, longitude string, latitudeRef string, longitudeRef string) (dDLatitude float64, dDLongitude float64, err error) {
	latitudeDegree, latitudeMinutes, latitudeSeconds, err := SeparateStringGPSCoordinate(latitude)
	if err != nil {
		return 0, 0, err
	}
	longitudeDegree, longitudeMinutes, longitudeSeconds, err := SeparateStringGPSCoordinate(longitude)
	if err != nil {
		return 0, 0, err
	}
	dDLatitude = float64((latitudeDegree + (latitudeMinutes / 60) + (latitudeSeconds / 3600)) * float64(coordinateRefs[latitudeRef]))
	dDLongitude = float64((longitudeDegree + (longitudeMinutes / 60) + (longitudeSeconds / 3600)) * float64(coordinateRefs[longitudeRef]))
	return dDLatitude, dDLongitude, nil
}
