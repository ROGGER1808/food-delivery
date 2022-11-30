package uploadprovider

import (
	"bytes"
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"gitlab.com/genson1808/food-delivery/component/fimage"
)

type cloudDinaryProvider struct {
	cloudName    string
	apiKey       string
	secretKey    string
	uploadFolder string
}

func NewCloudinaryProvider(cloudName, apiKey, secretKey, uploadFolder string) *cloudDinaryProvider {
	return &cloudDinaryProvider{cloudName: cloudName, apiKey: apiKey, secretKey: secretKey, uploadFolder: uploadFolder}
}

func (provider *cloudDinaryProvider) SaveFileUpload(ctx context.Context, data []byte, dst string) (*fimage.Image, error) {
	fileBytes := bytes.NewReader(data)
	//fileType := http.DetectContentType(data)

	cld, err := cloudinary.NewFromParams(provider.cloudName, provider.apiKey, provider.secretKey)
	if err != nil {
		return nil, err
	}

	resp, err := cld.Upload.Upload(ctx, fileBytes, uploader.UploadParams{Folder: provider.uploadFolder})
	if err != nil {
		return nil, err
	}

	img := &fimage.Image{Url: resp.SecureURL, CloudName: "Cloudinary"}
	return img, nil
}
