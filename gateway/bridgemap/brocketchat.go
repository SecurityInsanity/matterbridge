// +build !norocketchat

package bridgemap

import (
	brocketchat "github.com/SecurityInsanity/matterbridge/bridge/rocketchat"
)

func init() {
	FullMap["rocketchat"] = brocketchat.New
}
