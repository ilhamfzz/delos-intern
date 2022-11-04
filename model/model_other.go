package model

type CountRespon struct {
	Endpoint   string `json:"endpoint"`
	Method     string `json:"method"`
	Count      int    `json:"count"`
	UniqueUser int    `json:"unique_user"`
}

type Farm_Binding struct {
	Name string `json:"name" binding:"required"`
}

type Pond_Binding struct {
	Name   string `json:"name" binding:"required"`
	FarmID uint   `json:"farm_id"`
}