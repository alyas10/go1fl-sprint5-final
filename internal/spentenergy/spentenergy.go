package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// WalkingSpentCalories возвращает количество калорий, потраченных при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validateInput(steps, weight, height, duration); err != nil {
		return 0, err
	}
	return weight * MeanSpeed(steps, height, duration) * duration.Minutes() / minInH * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories возвращает количество калорий, потраченных при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validateInput(steps, weight, height, duration); err != nil {
		return 0, err
	}
	return weight * MeanSpeed(steps, height, duration) * duration.Minutes() / minInH, nil
}

// MeanSpeed возвращает среднюю скорость.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps < 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

// Distance возвращает дистанцию в километрах по количеству шагов и росту в метрах.
func Distance(steps int, height float64) float64 {
	if steps <= 0 {
		return 0
	}
	return height * stepLengthCoefficient * float64(steps) / mInKm
}

// validateInput проверяет корректность входных параметров тренировки.
func validateInput(steps int, weight, height float64, duration time.Duration) error {
	if weight <= 0 {
		return fmt.Errorf("вес должен быть положительным числом: %.2f", weight)
	}
	if height <= 0 {
		return fmt.Errorf("рост должен быть положительным числом: %.2f", height)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным: %d", steps)
	}
	if duration <= 0 {
		return fmt.Errorf("продолжительность должна быть положительной: %v", duration)
	}
	return nil
}