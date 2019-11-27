package resolver

import (
	"github.com/TylerGrey/studyhub/internal/resolver/model"
)

// Image 공통 이미지 모델
type Image struct {
	// TODO: json map 파싱방법 고민해볼것...
	Data model.Image
}

// URL ...
func (i Image) URL() string {
	return i.Data.URL
}

// Width ...
func (i Image) Width() int32 {
	return i.Data.Width
}

// Height ...
func (i Image) Height() int32 {
	return i.Data.Height
}
