package models

type CreateReq struct {
	Words []string `json:"words" validate:"required"`
}
