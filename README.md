# 🫀 Go Heart Bot

Библиотека для упрощённого создания Telegram-ботов на языке Go.
Проект ориентирован на **простоту, надежность и удобство расширения**.

> [!NOTE]
> 
> ![Go Version](https://img.shields.io/badge/go-1.18%2B-blue.svg)
> ![License](https://img.shields.io/badge/license-GNU%20AGPLv3-red.svg)
> ![status: dev](https://img.shields.io/badge/status-dev-orange)

---

## 📦 Что это

`go-heart-bot` — это каркас для разработки Telegram-ботов:

- Упрощает настройку и запуск бота
- Содержит готовые модули для работы с Telegram API (`heart/accessbot.go`, `corebot.go`, `degugbot.go`, `infobot.go`, `updatebot.go`)
- Легко расширяемый и тестируемый

---

## ⚙️ Пример `.env`

```env
TOKEN_API=your_bot_token
ACCESS_CODE=12345
ALLOWED_USER_ID=987654321
STATUS_DEBUG=true
STATUS_CHECK_ACCESS=false
TZ=UTC
```

---

## 💻 Пример использования

```go
package main

import (
    "github.com/my-app-s/go-heart-bot/heart"
)

func main() {
    bot := heart.NewCoreBot()
    
    bot.Start()
}
```

> Здесь `NewCoreBot` — функция для инициализации бота. Другие модули (`AccessBot`, `DebugBot`, `InfoBot`, `UpdateBot`) можно подключать по необходимости.

---

## 🧪 Тестирование

```bash
go test ./...
```

- Запускает все тесты внутри пакета `heart`
- Использует стандартный `testing` пакет Go

---

## ⚠️ Disclaimer / Отказ от ответственности

### English Version
This project is an **independent development** provided on an **"AS IS"** basis.

* **Liability:** In no event shall the author be liable for any errors, bugs, or data loss arising from the use of this software.
* **Status:** This is an experimental tool. Always verify the generated HTML output before deployment.

> [!CAUTION]
> Any use (operation) of this code is at your own risk.

---

### Русская версия
Данный проект является **независимой разработкой** и предоставляется «как есть».

* **Ответственность:** Автор не несет ответственности за любые ошибки, баги или потерю данных, возникшие в результате использования данного кода.
* **Статус:** Это экспериментальный инструмент. Проверяйте сгенерированный HTML перед деплоем.

> [!CAUTION]
> Любое использование (эксплуатация) данного кода осуществляется на ваш страх и риск.

---

## 📜 Лицензия

Проект лицензирован под **GNU Affero General Public License v3.0 (AGPLv3)**.
См. [LICENSE](./LICENSE) для подробностей.
