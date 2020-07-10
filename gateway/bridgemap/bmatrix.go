// +build !nomatrix

package bridgemap

import (
	bmatrix "github.com/SecurityInsanity/matterbridge/bridge/matrix"
)

func init() {
	FullMap["matrix"] = bmatrix.New
}
