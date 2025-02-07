package generate

import (
	"context"

	"github.com/containers/common/libimage"
)

func imageRunPath(ctx context.Context, img *libimage.Image) (string, error) {
	if img != nil {
		inspectData, err := img.Inspect(ctx, nil)
		if err != nil {
			return "", err
		}
		if inspectData.Os == "freebsd" {
			return "/var/run", nil
		}
		return "/run", nil
	} else {
		return "/var/run", nil
	}
}
