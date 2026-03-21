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

func init() {
	if FileExists(".env") {
		err := godotenv.Load() // Initialization env file
		if err != nil {
			log.Fatal("[FATAL Initialization package heartbot/degugbot]Not .env file.")
		}
	}
}

// SetDebugBot set status debug from env.
func SetDebugBot(bot *tgbotapi.BotAPI) {
	// Realization
	status := false // Default status false
	if statusFromEnv := os.Getenv("STATUS_DEBUG"); statusFromEnv != "" {
		parsed, err := strconv.ParseBool(statusFromEnv)
		if err != nil {
			log.Printf("[WARNING SetDebugBot]Invalid STATUS_DEBUG value: %q, error: %v.", statusFromEnv, err)
		} else {
			status = parsed
			bot.Debug = status
			log.Printf("[INFO SetDebugBot]Status debug set: %t.", status)
		}
	} else {
		log.Printf("[INFO SetDebugBot]Status debug set default: %t.", status)
	}
}
