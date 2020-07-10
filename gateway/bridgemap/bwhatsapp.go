// +build !nowhatsapp

package bridgemap

import (
	bwhatsapp "github.com/SecurityInsanity/matterbridge/bridge/whatsapp"
)

func init() {
	FullMap["whatsapp"] = bwhatsapp.New
}
