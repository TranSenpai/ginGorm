package models

type ContractUpdateRequest struct {
	Filter   Filter   `json:"filter"`
	Contract Contract `json:"contract"`
}
