// +build !nokeybase

package bridgemap

import (
	bkeybase "github.com/SecurityInsanity/matterbridge/bridge/keybase"
)

func init() {
	FullMap["keybase"] = bkeybase.New
}
