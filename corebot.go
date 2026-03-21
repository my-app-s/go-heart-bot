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

	"github.com/joho/godotenv" // For work getenv
)

func init() {
	if FileExists(".env") {
		err := godotenv.Load() // Initialization env file
		if err != nil {
			log.Fatal("[FATAL Initialization package heartbot/corebot]Not .env file.")
		}
	}
}

// GlobalHeartBot get bot by token from scope env.
func GlobalHeartBot() *tgbotapi.BotAPI {
	// Realization
	token := os.Getenv("TOKEN_API")
	if token == "" {
		log.Fatal("[FATAL GlobalHeartBot]TOKEN is not found.")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panicf("[PANIC GlobalHeartBot]TOKEN error: %v.", err)
	}

	token = "*****" // Token clear

	return bot
}

// LocalHeartBot get bot by token from local file env.
func LocalHeartBot() *tgbotapi.BotAPI {
	// Realization
	token := os.Getenv("TOKEN_API")
	if token == "" {
		log.Fatal("[FATAL LocalHeartBot]TOKEN is not found.")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panicf("[PANIC LocalHeartBot]TOKEN error: %v.", err)
	}

	token = "*****" // Token clear

	return bot
}
