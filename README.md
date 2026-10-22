<div align="center">
  <img src="assets/cart-flow.svg" width="100%" alt="Сценарий корзины DemoBlaze — end-to-end тест на Go и Playwright" />
  <p><strong>🇷🇺 Русский</strong> &nbsp;·&nbsp; <a href="README.en.md">🇬🇧 English</a></p>
</div>

<h1 align="center">DemoBlaze · Cart Flow</h1>

<p align="center">
  Небольшой end-to-end тест на Go и Playwright: добавить один товар в корзину и проверить результат.
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.25-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go 1.25" />
  <img src="https://img.shields.io/badge/Playwright-0.6201.1-2EAD33?style=for-the-badge&logo=playwright&logoColor=white" alt="Playwright Go 0.6201.1" />
  <a href="https://github.com/olegval26/vk-demoblaze-playwright/actions/workflows/playwright.yml">
    <img src="https://github.com/olegval26/vk-demoblaze-playwright/actions/workflows/playwright.yml/badge.svg" alt="Статус GitHub Actions" />
  </a>
</p>

## Сценарий

`Каталог DemoBlaze → Samsung galaxy s6 → Add to cart → подтвердить диалог → Cart`

Тест открывает [DemoBlaze](https://www.demoblaze.com/), добавляет **Samsung galaxy s6**, принимает диалог **Product added**, затем проверяет корзину.

## Что проверяется

- На странице товара отображается нужный продукт.
- В браузерном диалоге приходит ожидаемое подтверждение, и тест его принимает.
- В корзине ровно одна строка — **Samsung galaxy s6**.

Для каждого запуска создаётся отдельный контекст браузера. Проверки ждут нужного состояния интерфейса, фиксированные паузы не используются.

## Запуск локально

Нужен Go 1.25 или новее и доступ в интернет для установки Playwright и Chromium.

```bash
go mod download
go run github.com/mxschmitt/playwright-go/cmd/playwright@v0.6201.1 install chromium
go test -v ./...
```

Тест запускает Chromium в headless-режиме.

## GitHub Actions

GitHub Actions запускает тот же тест при каждом push и pull request. Запуск также можно выполнить вручную во вкладке **Actions**.

## Структура проекта

```text
.
├── .github/workflows/playwright.yml
├── assets/cart-flow.svg
├── cart_test.go
├── go.mod
├── go.sum
├── README.md
└── README.en.md
```

<p align="center">
  <sub>Один понятный маршрут: открыть → добавить → подтвердить → проверить.</sub>
</p>
