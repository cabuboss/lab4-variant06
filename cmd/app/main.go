// Программа демонстрирует работу пакета waterbill: расчёт расхода
// воды, стоимости, пени за просрочку и формирование отчёта, а также
// показывает обработку ошибок на нескольких некорректных сценариях.
package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/google/uuid"

	"lab4-variant06/pkg/waterbill"
)

func main() {
	// 1. Заголовок программы.
	color.New(color.FgCyan, color.Bold).Println("=== Лабораторная работа №4 — waterbill (вариант 6) ===")

	// 2. Идентификатор операции (демонстрация пакета uuid).
	operationID := uuid.New().String()
	color.New(color.FgBlue, color.Faint).Printf("ID операции: %s\n\n", operationID)

	// 3. Успешный сценарий расчёта.
	color.New(color.FgGreen, color.Bold).Println("--- Успешный расчёт оплаты за воду ---")

	prev := 145.250
	curr := 172.800
	tariff := 128.50

	cubic, err := waterbill.WaterUsage(prev, curr)
	if err != nil {
		color.Red("ошибка расчёта расхода воды: %v", err)
		return
	}
	fmt.Printf("%-25s %10.3f м3\n", "Предыдущее показание:", prev)
	fmt.Printf("%-25s %10.3f м3\n", "Текущее показание:", curr)
	fmt.Printf("%-25s %10.3f м3\n", "Израсходовано:", cubic)

	cost, err := waterbill.WaterCost(cubic, tariff)
	if err != nil {
		color.Red("ошибка расчёта стоимости: %v", err)
		return
	}
	fmt.Printf("%-25s %10.2f ₸\n", "Тариф за м3:", tariff)
	fmt.Printf("%-25s %10.2f ₸\n", "Стоимость до пени:", cost)

	penaltyPercent := 8.5
	if err := waterbill.ApplyPenalty(&cost, penaltyPercent); err != nil {
		color.Red("ошибка начисления пени: %v", err)
		return
	}
	fmt.Printf("%-25s %10.2f %%\n", "Пеня за просрочку:", penaltyPercent)
	fmt.Printf("%-25s %10.2f ₸\n", "Стоимость после пени:", cost)

	report, err := waterbill.FormatWaterReport("Иванов И.И.", cubic, cost)
	if err != nil {
		color.Red("ошибка формирования отчёта: %v", err)
		return
	}
	fmt.Println()
	fmt.Println(report)

	// 4. Демонстрация обработки ошибок на некорректных сценариях.
	color.New(color.FgYellow, color.Bold).Println("--- Демонстрация обработки ошибок ---")

	if _, err := waterbill.WaterUsage(200.0, 150.0); err != nil {
		color.Red("сценарий 1 (curr < prev): %v", err)
	}

	if _, err := waterbill.WaterCost(10.0, 0); err != nil {
		color.Red("сценарий 2 (тариф = 0): %v", err)
	}

	badCost := 1000.0
	if err := waterbill.ApplyPenalty(&badCost, 150.0); err != nil {
		color.Red("сценарий 3 (пеня = 150%%): %v", err)
	}
}
