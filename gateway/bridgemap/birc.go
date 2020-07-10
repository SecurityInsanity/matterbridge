// +build !noirc

package bridgemap

import (
	birc "github.com/SecurityInsanity/matterbridge/bridge/irc"
)

func init() {
	FullMap["irc"] = birc.New
}
