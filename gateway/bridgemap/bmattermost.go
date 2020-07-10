// +build !nomattermost

package bridgemap

import (
	bmattermost "github.com/SecurityInsanity/matterbridge/bridge/mattermost"
)

func init() {
	FullMap["mattermost"] = bmattermost.New
}
