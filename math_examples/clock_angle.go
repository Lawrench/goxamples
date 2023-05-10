package math_examples

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// ClockProblem is a programming exercise.
// The goal is to get the smallest angle between
// hour and minute hands for any given time, taking
// into account the hour hand will move during
// the hour (e.g., at 4:45, the hour hand will be
// 3/4 of the way to the 5, rather than directly
// on the 4.)

func GetAngle(t string) (string, error) {
	h, m, err := parseTime(t)
	if err != nil {
		return "", err
	}

	angle := calcAngle(h, m)
	return fmt.Sprintf("Least angle for %s is %v.", t, angle), nil
}

func parseTime(t string) (float64, float64, error) {
	parts := strings.Split(t, ":")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("unable to parse time")
	}

	h, err := strconv.ParseFloat(parts[0], 64)
	if err != nil || h < 0 || h > 24 {
		return 0, 0, fmt.Errorf("unable to parse hour")
	}

	m, err := strconv.ParseFloat(parts[1], 64)
	if err != nil || m < 0 || m > 59 {
		return 0, 0, fmt.Errorf("unable to parse minutes")
	}

	if h == 24 && m > 0 {
		return 0, 0, fmt.Errorf("invalid time")
	}

	if h > 12 {
		h -= 12
	}

	return h, m, nil
}

func calcAngle(h float64, m float64) float64 {
	const degreesPerHour = 30
	const degreesPerMin = 6
	hAngle := h * degreesPerHour
	hAngle += (m / 60) * degreesPerHour
	mAngle := m * degreesPerMin
	total := math.Abs(hAngle - mAngle)
	if total > 180 {
		total = 360 - total
	}
	return total
}
