package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("invalid steps")
	}
	if duration.Seconds() <= 0 {
		return 0, errors.New("invalid duration")
	}
	if weight <= 0 {
		return 0, errors.New("invalid weight")
	}
	if height <= 0 {
		return 0, errors.New("invalid height")
	}
	calories := (weight * MeanSpeed(steps, height, duration) * duration.Minutes() * walkingCaloriesCoefficient) / minInH
	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("invalid steps")
	}
	if duration.Seconds() <= 0 {
		return 0, errors.New("invalid duration")
	}
	if weight <= 0 {
		return 0, errors.New("invalid weight")
	}
	if height <= 0 {
		return 0, errors.New("invalid height")
	}
	averageSpeed := MeanSpeed(steps, height, duration)
	calories := (weight * averageSpeed * duration.Minutes()) / minInH
	return calories, nil

}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps == 0 {
		return 0
	}
	if duration.Seconds() <= 0 {
		return 0
	}
	averageSpeed := Distance(steps, height) / duration.Hours()
	return averageSpeed
}

func Distance(steps int, height float64) float64 {
	dlinShaga := height * stepLengthCoefficient
	distance := (float64(steps) * dlinShaga) / mInKm
	return distance
}
