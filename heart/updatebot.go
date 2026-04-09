// Copyright (C) 2025-2026 my-app-s
// Licensed under the GNU AGPLv3

package heart

// CreateUpdatesChannel create channel updates.
func CreateUpdatesChannel(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	// Realization
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	return updates
}
