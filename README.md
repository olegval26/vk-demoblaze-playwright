# VK QA test: DemoBlaze cart

Автотест на Go и Playwright для сценария из тестового задания: добавить **Samsung galaxy s6** в корзину и проверить, что в ней ровно один товар — он же.

## Что делает тест

1. Открывает [DemoBlaze](https://www.demoblaze.com/).
2. Открывает карточку **Samsung galaxy s6**.
3. Нажимает **Add to cart** и принимает системное окно сайта.
4. Переходит в **Cart**.
5. Проверяет, что в корзине ровно одна строка и в ней указан **Samsung galaxy s6**.

## Требования

- Go 1.25 или новее
- Доступ в интернет для установки Playwright и Chromium

## Запуск

```bash
go mod download
go run github.com/mxschmitt/playwright-go/cmd/playwright@v0.6201.1 install chromium
go test -v ./...
```

Тест запускает Chromium в headless-режиме. Для Go Playwright драйвер и браузер устанавливаются командой CLI выше.

## Структура проекта

```text
.
├── .github/workflows/playwright.yml  # запуск теста в GitHub Actions
├── cart_test.go                      # сквозной тест корзины
├── go.mod                            # версия Go и зависимость Playwright
└── README.md                         # описание и запуск
```

Каждый запуск создаёт отдельный браузерный контекст, чтобы корзина начиналась с чистого состояния. Проверки Playwright ждут нужное состояние интерфейса; фиксированные паузы не используются, потому что DemoBlaze загружает товары асинхронно.

## CI

GitHub Actions запускает тест при каждом push, pull request и вручную через вкладку **Actions**. Для запуска в CI устанавливается Chromium и системные зависимости, затем выполняется `go test -v ./...`.
