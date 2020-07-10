// +build !nomsteams

package bridgemap

import (
	bmsteams "github.com/SecurityInsanity/matterbridge/bridge/msteams"
)

func init() {
	FullMap["msteams"] = bmsteams.New
}
