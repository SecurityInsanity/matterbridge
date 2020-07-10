// +build !nogitter

package bridgemap

import (
	bgitter "github.com/SecurityInsanity/matterbridge/bridge/gitter"
)

func init() {
	FullMap["gitter"] = bgitter.New
}
