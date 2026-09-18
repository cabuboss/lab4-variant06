// Package waterbill содержит функции для расчёта оплаты за воду
// по показаниям счётчика: расход воды, стоимость расхода, пеня за
// просрочку платежа и формирование текстового отчёта для абонента.
package waterbill

import (
	"fmt"
	"strings"
)

// WaterUsage вычисляет расход воды по показаниям счётчика как
// разницу текущего показания curr и предыдущего показания prev
// (в кубических метрах).
//
// Возвращает ошибку, если prev или curr отрицательны, а также если
// curr меньше prev (счётчик физически не может "откручиваться назад").
func WaterUsage(prev, curr float64) (float64, error) {
	if prev < 0 {
		return 0, fmt.Errorf("показание счётчика не может быть отрицательным: prev=%.2f", prev)
	}
	if curr < 0 {
		return 0, fmt.Errorf("показание счётчика не может быть отрицательным: curr=%.2f", curr)
	}
	if curr < prev {
		return 0, fmt.Errorf("текущее показание не может быть меньше предыдущего: prev=%.2f, curr=%.2f", prev, curr)
	}

	return curr - prev, nil
}

// WaterCost вычисляет стоимость израсходованной воды по формуле
// cost = cubic * tariff, где cubic — расход в кубических метрах,
// tariff — цена за один кубический метр.
//
// Возвращает ошибку, если cubic отрицательный, а также если
// tariff меньше либо равен нулю.
func WaterCost(cubic, tariff float64) (float64, error) {
	if cubic < 0 {
		return 0, fmt.Errorf("расход воды не может быть отрицательным: cubic=%.2f", cubic)
	}
	if tariff <= 0 {
		return 0, fmt.Errorf("тариф должен быть положительным числом: tariff=%.2f", tariff)
	}

	return cubic * tariff, nil
}

// ApplyPenalty увеличивает стоимость *cost на penaltyPercent процентов
// как пеню за просроченный платёж. Изменение вносится напрямую в
// значение, на которое указывает cost.
//
// Возвращает ошибку, если cost равен nil, если *cost отрицательный,
// а также если penaltyPercent меньше 0 или больше 100.
func ApplyPenalty(cost *float64, penaltyPercent float64) error {
	if cost == nil {
		return fmt.Errorf("указатель на стоимость не может быть nil")
	}
	if *cost < 0 {
		return fmt.Errorf("стоимость не может быть отрицательной: cost=%.2f", *cost)
	}
	if penaltyPercent < 0 || penaltyPercent > 100 {
		return fmt.Errorf("процент пени должен быть в диапазоне от 0 до 100: penaltyPercent=%.2f", penaltyPercent)
	}

	*cost += *cost * penaltyPercent / 100

	return nil
}

// FormatWaterReport формирует текстовый многострочный отчёт для
// абонента owner с указанием расхода воды cubic (в кубических
// метрах) и итоговой стоимости cost (в тенге).
//
// Возвращает ошибку, если owner пустая строка (после удаления
// пробелов по краям), а также если cubic или cost отрицательны.
func FormatWaterReport(owner string, cubic, cost float64) (string, error) {
	if strings.TrimSpace(owner) == "" {
		return "", fmt.Errorf("имя абонента не может быть пустым")
	}
	if cubic < 0 {
		return "", fmt.Errorf("расход воды не может быть отрицательным: cubic=%.2f", cubic)
	}
	if cost < 0 {
		return "", fmt.Errorf("стоимость не может быть отрицательной: cost=%.2f", cost)
	}

	report := fmt.Sprintf(
		"=== Отчёт по оплате воды ===\n"+
			"%-20s %s\n"+
			"%-20s %10.3f м3\n"+
			"%-20s %10.2f ₸\n",
		"Абонент:", owner,
		"Расход воды:", cubic,
		"К оплате:", cost,
	)

	return report, nil
}
