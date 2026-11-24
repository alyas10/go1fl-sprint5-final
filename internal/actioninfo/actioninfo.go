package actioninfo

import (
	"log"
)

// DataParser описывает общий интерфейс для тренировок и прогулок.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info обрабатывает набор строк dataset и выводит информацию о каждой активности.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("Ошибка парсинга: %v", err)
			continue
		}

		result, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Ошибка формирования информации: %v", err)
			continue
		}

		log.Println(result)
	}
}