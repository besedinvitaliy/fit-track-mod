package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring, ",")

	if len(parts) != 3 {
		return fmt.Errorf("некорректное количество данных тренировки")
	}

	steps, err := strconv.Atoi(parts[0])
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным")
	}
	t.Steps = steps

	t.TrainingType = parts[1]

	duration, err := time.ParseDuration(parts[2])
	if duration <= 0 {
		return fmt.Errorf("продолжительность должна быть положительной")
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(
			t.Steps, t.Weight, t.Height, t.Duration,
		)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(
			t.Steps, t.Weight, t.Height, t.Duration,
		)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	return fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		meanSpeed,
		calories,
	), nil

}
