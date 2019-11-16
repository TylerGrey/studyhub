package args

// ImageInput ...
type ImageInput struct {
	File   string
	Width  *int32
	Height *int32
}

// AddressInput ...
type AddressInput struct {
	Address string
	Lat     float64
	Lng     float64
}

// OrderBy ...
type OrderBy struct {
	Field     string
	Direction string
}
