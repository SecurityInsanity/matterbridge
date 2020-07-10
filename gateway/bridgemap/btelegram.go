// +build !notelegram

package bridgemap

import (
	btelegram "github.com/SecurityInsanity/matterbridge/bridge/telegram"
)

func init() {
	FullMap["telegram"] = btelegram.New
}
