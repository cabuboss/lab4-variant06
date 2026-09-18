# Лабораторная работа №4 — Вариант 6: `waterbill` (вода)

Собственные пакеты, документация и подключение пакетов из GitHub.

## Структура проекта

```
lab4-variant06/
  go.mod
  cmd/app/main.go              — демонстрационная программа
  pkg/waterbill/waterbill.go   — собственный пакет
  README.md
```

## Пакет `waterbill`

Расчёт оплаты за воду по показаниям счётчика.

| Функция | Сигнатура | Назначение |
|---------|-----------|------------|
| `WaterUsage` | `(prev, curr float64) (float64, error)` | Расход воды: `curr - prev` (м³) |
| `WaterCost` | `(cubic, tariff float64) (float64, error)` | Стоимость: `cubic * tariff` (₸) |
| `ApplyPenalty` | `(cost *float64, penaltyPercent float64) error` | Начисление пени через **указатель** |
| `FormatWaterReport` | `(owner string, cubic, cost float64) (string, error)` | Отчёт через `fmt.Sprintf` |

Все функции экспортируемые (PascalCase), у каждой есть doc-комментарий.
Doc-комментарий есть и перед объявлением `package`.

### Обработка ошибок

Все ошибки возвращаются через `fmt.Errorf(...)`:

- `WaterUsage` — отрицательные показания, `curr < prev` (счётчик не крутится назад)
- `WaterCost` — отрицательный расход, тариф `<= 0`
- `ApplyPenalty` — `nil`-указатель, отрицательная стоимость, процент вне диапазона `0..100`
- `FormatWaterReport` — пустое имя абонента, отрицательные расход или стоимость

## Внешние пакеты из GitHub

| Пакет | Зачем | Где используется |
|-------|-------|------------------|
| `github.com/fatih/color` | Цветной вывод в терминал | Заголовки секций, вывод ошибок красным |
| `github.com/google/uuid` | Генерация ID операции | `uuid.New().String()` в начале программы |

Установка:

```bash
go get github.com/fatih/color
go get github.com/google/uuid
go mod tidy
```

## Запуск

```bash
go run ./cmd/app
```

Пример вывода:

```
=== Лабораторная работа №4 — waterbill (вариант 6) ===
ID операции: 8de929ae-0017-4e76-ac6c-6e5cbfca1acc

--- Успешный расчёт оплаты за воду ---
Предыдущее показание:        145.250 м3
Текущее показание:           172.800 м3
Израсходовано:                27.550 м3
Тариф за м3:                  128.50 ₸
Стоимость до пени:           3540.18 ₸
Пеня за просрочку:              8.50 %
Стоимость после пени:        3841.09 ₸

=== Отчёт по оплате воды ===
Абонент:             Иванов И.И.
Расход воды:             27.550 м3
К оплате:               3841.09 ₸

--- Демонстрация обработки ошибок ---
сценарий 1 (curr < prev): текущее показание не может быть меньше предыдущего: prev=200.00, curr=150.00
сценарий 2 (тариф = 0): тариф должен быть положительным числом: tariff=0.00
сценарий 3 (пеня = 150%): процент пени должен быть в диапазоне от 0 до 100: penaltyPercent=150.00
```

Форматирование `fmt.Printf`: ширина поля `%-25s` / `%10.2f` для выравнивания,
точность `%.3f` для кубометров и `%.2f` для денег.

## Документация

Консоль:

```bash
go doc ./pkg/waterbill              # пакет и список функций
go doc -all ./pkg/waterbill         # полная документация
go doc ./pkg/waterbill WaterUsage   # одна функция
```

> Примечание: `go doc ./...` выдаёт ошибку `cannot find package "."` —
> команда `go doc` не принимает шаблон `...` и работает с одним пакетом.
> Рабочий вариант — `go doc ./pkg/waterbill`.

Веб-документация:

```bash
go install golang.org/x/tools/cmd/godoc@latest
godoc -http=:6060
```

Далее открыть <http://localhost:6060/pkg/github.com/cabuboss/lab4-variant06/pkg/waterbill/>

## Проверки

```bash
go build ./...   # сборка
go vet ./...     # статический анализ
gofmt -l .       # форматирование (пустой вывод = всё отформатировано)
```

## Заливка в GitHub

```bash
git init
git add .
git commit -m "Лабораторная работа №4, вариант 6: waterbill"
git branch -M main
git remote add origin https://github.com/<логин>/lab4-variant06.git
git push -u origin main
```
