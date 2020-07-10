// +build !nosteam

package bridgemap

import (
	bsteam "github.com/SecurityInsanity/matterbridge/bridge/steam"
)

func init() {
	FullMap["steam"] = bsteam.New
}
