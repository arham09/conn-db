package models

type Faktur struct {
	ID         int64  `form:"id" json:"id"`
	Code       string `form:"code" json:"code"`
	ExternalID string `form:"externalId" json:"externalId"`
	Name       string `form:"name" json:"name"`
	Status     string `form:"status" json:"status"`
}
