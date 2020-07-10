// +build !nozulip

package bridgemap

import (
	bzulip "github.com/SecurityInsanity/matterbridge/bridge/zulip"
)

func init() {
	FullMap["zulip"] = bzulip.New
}
