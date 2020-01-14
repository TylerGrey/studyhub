package args


////////////////////////////////
// Mutation
////////////////////////////////

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

// AddHubIncorrectInfoInput 잘못된 정보 수정 요청 Args
type AddHubIncorrectInfoInput struct {
	Input AddHubIncorrectInfoArgs
}

// AddHubIncorrectInfoArgs ...
type AddHubIncorrectInfoArgs struct {
	HubID   string
	Message string
}

////////////////////////////////
// Query
////////////////////////////////

// HubsArgs ...
type HubsArgs struct {
	First   *int32
	Last    *int32
	After   *string
	Before  *string
	OrderBy *OrderBy
}