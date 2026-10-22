<div align="center">
  <img src="assets/cart-flow.svg" width="100%" alt="DemoBlaze cart flow — Go и Playwright" />
</div>

<h1 align="center">DemoBlaze · Cart Flow</h1>

<p align="center">
  Учебный end-to-end сценарий: проверить, что после добавления товара в корзину там ровно один нужный товар.
</p>

<p align="center">
  <a href="https://img.shields.io/badge/Go-1.25-00ADD8?style=for-the-badge&logo=go&logoColor=white">
    <img src="https://img.shields.io/badge/Go-1.25-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go 1.25" />
  </a>
  <img src="https://img.shields.io/badge/Playwright-0.6201.1-2EAD33?style=for-the-badge&logo=playwright&logoColor=white" alt="Playwright Go 0.6201.1" />
  <a href="https://github.com/olegval26/vk-demoblaze-playwright/actions/workflows/playwright.yml">
    <img src="https://github.com/olegval26/vk-demoblaze-playwright/actions/workflows/playwright.yml/badge.svg" alt="GitHub Actions workflow status" />
  </a>
</p>

## Путь теста

1. Открыть [DemoBlaze](https://www.demoblaze.com/).
2. Перейти к **Samsung galaxy s6**.
3. Нажать **Add to cart** и принять диалог **Product added**.
4. Открыть корзину и проверить, что в ней одна строка с **Samsung galaxy s6**.

## Что здесь практикую

- Поиск элементов по их ролям и точным названиям.
- Обработку браузерного диалога после добавления товара.
- Отдельный контекст браузера для чистой корзины.
- Проверки результата через ожидание нужного состояния интерфейса.

## Запустить локально

Нужен Go 1.25 или новее и доступ в интернет для установки Playwright и Chromium.

```bash
go mod download
go run github.com/mxschmitt/playwright-go/cmd/playwright@v0.6201.1 install chromium
go test -v ./...
```

Тест запускает Chromium в headless-режиме. GitHub Actions выполняет тот же тест при push и pull request.

## Структура

```text
.
├── .github/workflows/playwright.yml
├── cart_test.go
├── go.mod
└── README.md
```

<p align="center">
  <sub>Один сценарий целиком: открыть → добавить → подтвердить → проверить.</sub>
</p>
