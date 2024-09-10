package image

// Image defines Open Graph Image type
type Image struct {
	URL       string `json:"url,omitempty"`
	SecureURL string `json:"secure_url,omitempty"`
	Type      string `json:"type,omitempty"`
	Width     uint64 `json:"width,omitempty"`
	Height    uint64 `json:"height,omitempty"`
	Alt       string `json:"alt,omitempty"`
}

func NewImage() *Image {
	return &Image{}
}

func ensureHasImage(images []*Image) []*Image {
	if len(images) == 0 {
		images = append(images, NewImage())
	}
	return images
}

func AddURL(images []*Image, v string) []*Image {
	if len(images) == 0 || (images[len(images)-1].URL != "" && images[len(images)-1].URL != v) {
		images = append(images, NewImage())
	}
	images[len(images)-1].URL = v
	return images
}

func AddSecureURL(images []*Image, v string) []*Image {
	images = ensureHasImage(images)
	images[len(images)-1].SecureURL = v
	return images
}

func AddType(images []*Image, v string) []*Image {
	images = ensureHasImage(images)
	images[len(images)-1].Type = v
	return images
}

func AddWidth(images []*Image, v uint64) []*Image {
	images = ensureHasImage(images)
	images[len(images)-1].Width = v
	return images
}

func AddHeight(images []*Image, v uint64) []*Image {
	images = ensureHasImage(images)
	images[len(images)-1].Height = v
	return images
}

func AddAlt(images []*Image, v string) []*Image {
	images = ensureHasImage(images)
	images[len(images)-1].Alt = v
	return images
}
