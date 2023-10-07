package entity

type SampleSetRequest struct {
	AppID  string      `json:"app_id" binding:"required"`
	UserID string      `json:"user_id"`
	Key    string      `json:"key" binding:"required"`
	Value  interface{} `json:"value" binding:"required"`
}

type SampleGetRequest struct {
	AppID  string `json:"app_id" binding:"required"`
	UserID string `json:"user_id"`
	Key    string `json:"key" binding:"required"`
}

type SomeData struct {
	Value interface{} `json:"value"`
}
