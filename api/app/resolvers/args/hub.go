package args

// CreateHubInput 사용자 생성 Args
type CreateHubInput struct {
	Input CreateHubArgs
}

// CreateHubArgs ...
type CreateHubArgs struct {
	Name       string
	Type       string
	CoverImage ImageInput
	Images     *[]ImageInput
	Tel        *string
	Address    AddressInput
	Hours      *[]HubHoursInput
}

// UpdateHubInput 사용자 생성 Args
type UpdateHubInput struct {
	Input UpdateHubArgs
}

// UpdateHubArgs ...
type UpdateHubArgs struct {
	ID         string
	Name       *string
	Type       *string
	CoverImage *ImageInput
	Images     *[]ImageInput
	Tel        *string
	Address    *AddressInput
	Hours      *[]HubHoursInput
}

// HubHoursInput ...
type HubHoursInput struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// HubsArgs ...
type HubsArgs struct {
	First   *int32
	Last    *int32
	After   *string
	Before  *string
	OrderBy *OrderBy
}
