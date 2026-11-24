package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training содержит данные о тренировке.
type Training struct {
	Steps                 int           // количество шагов, проделанных за тренировку
	TrainingType          string        // тип тренировки (бег или ходьба)
	Duration              time.Duration // длительность тренировки
	personaldata.Personal               // встроенная структура Personal из пакета personaldata
}

// Parse разбирает строку вида "3456,Ходьба,3h00m" и заполняет поля структуры.
func (t *Training) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("неверный формат строки: ожидалось 3 элемента, получено %d", len(parts))
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("некорректное количество шагов %q: %w", parts[0], err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным: %d", steps)
	}
	t.Steps = steps

	t.TrainingType = parts[1]

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("некорректная продолжительность %q: %w", parts[2], err)
	}
	if duration <= 0 {
		return fmt.Errorf("продолжительность должна быть положительной: %v", duration)
	}
	t.Duration = duration

	return nil
}

// ActionInfo формирует строку с информацией о тренировке.
func (t Training) ActionInfo() (string, error) {
	dist := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var (
		calories float64
		err      error
	)
	switch strings.ToLower(t.TrainingType) {
	case "ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		dist,
		meanSpeed,
		calories,
	), nil
}