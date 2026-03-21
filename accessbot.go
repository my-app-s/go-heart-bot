/*
 * Copyright (C) 2026 my-app-s
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */
 package heartbot

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv" // For work getenv
)

// FileExists check exists file .env.
func FileExists(filename string) bool {
	// Realization
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func init() {
	if FileExists(".env") {
		err := godotenv.Load() // Initialization env file
		if err != nil {
			log.Fatal("[FATAL Initialization package heartbot/accessbot]Not .env file.")
		}
	}
}

// AllowedUser set status debug from env.
func AllowedUser() int64 {
	// Realization
	allowedUserID := os.Getenv("ALLOWED_USER_ID")
	parseID, err := strconv.Atoi(allowedUserID)
	if err != nil {
		log.Panicf("[WARNING AllowedUser]Invalid ALLOWED_USER_ID value error: %v.", err)
	}
	return int64(parseID)
}

// SetDebugBot set status debug from env.
func AccessCode() int64 {
	// Realization
	code := os.Getenv("ACCESS_CODE")
	parseCode, err := strconv.Atoi(code)
	if err != nil {
		log.Panicf("[PANIC AccessCode]%v.", err)
	}
	return int64(parseCode)
}

// GetStatusCheck get status check access from env.
func GetStatusCheck() bool {
	// Realization
	statusCheckAccess := os.Getenv("STATUS_CHECK_ACCESS")
	parse, err := strconv.ParseBool(statusCheckAccess)
	if err != nil {
		log.Panicf("[PANIC AccessCode]%v.", err)
	}
	return parse
}

// CheckUserAccess request and check access code.
func CheckUserAccess(bot *tgbotapi.BotAPI, text string, chatID, ownerID, accessCode int64) bool {
	// Realization
	if ownerID == chatID {
		log.Println("[INFO checkUserAccess]🔑 Good access, user admin.']")
		return true
	} else {
		if text != strconv.FormatInt(accessCode, 10) {
			bot.Send(tgbotapi.NewMessage(chatID,
				"Sorry, this bot is restricted. "+
					"🔒 Only the owner or users with an access code can continue."+
					"(Извините, бот доступен только владельцу или по коду доступа.)"))
			bot.Send(tgbotapi.NewMessage(chatID,
				"👀 Please enter your access code.(Введите код доступа):"))
			return false
		} else {
			bot.Send(tgbotapi.NewMessage(chatID, "Good enter code."))
			return true
		}
	}
	return false
}
