package model

// Image ...
type Image struct {
	URL    string
	Width  int32
	Height int32
}

// Address
type Address struct {
	Address string
	Lat     float64
	Lng     float64
}

// PageInfo
type PageInfo struct {
	StartCursor     *string
	EndCursor       *string
	HasNextPage     bool
	HasPreviousPage bool
}
