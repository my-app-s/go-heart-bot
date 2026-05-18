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

## 📜 Disclaimer

**English**: Materials are provided ***as is*** under the LICENSE file. No warranties, no rights granted unless explicitly stated. Authors are not liable for damages. No partnership or obligations created.  

**Русский**: Материалы предоставляются ***как есть*** и регулируются LICENSE. Гарантий нет, права не передаются без явного указания. Автор(ы) не несут ответственности. Партнёрство или обязательства не создаются.  

📌 See full disclaimer in [DISCLAIMER.md](https://github.com/my-app-s/my-app-s/blob/main/DISCLAIMER.md)

---

## 📜 Лицензия

Проект лицензирован под **GNU Affero General Public License v3.0 (AGPLv3)**.
См. [LICENSE](./LICENSE) для подробностей.
