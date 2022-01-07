package query

import "github.com/google/uuid"

type ListConstructionsFromCompanyQueryResponse struct {
	Constructions []struct{
		Id uuid.UUID
		Name string
	} `json:"constructions"`
}

type ListCompanyQueryResponse struct {
	Companies []struct{
		Id uuid.UUID
		Name string
	} `json:"companies"`
}
