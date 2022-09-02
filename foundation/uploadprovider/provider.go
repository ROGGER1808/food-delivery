package uploadprovider

import (
	"context"
	"gitlab.com/genson1808/food-delivery/foundation/fimage"
)

type UploadProvider interface {
	SaveFileUpload(ctx context.Context, data []byte, dst string) (*fimage.Image, error)
}
