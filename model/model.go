package model

type PageAccessTokenModel struct {
	PageAccessToken string `json:"access_token"`
	ID              string `json:"id"`
}

type LeadFromModel struct {
	Data   []LeadFromDataModel `json:"data"`
	Paging LeadFormPagingModel `json:"paging"`
}

type LeadModel struct {
	Data   []LeadDataModel     `json:"data"`
	Paging LeadFormPagingModel `json:"paging"`
}

type LeadFromDataModel struct {
	ID     string `json:"id"`
	Locale string `json:"locale"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
type LeadFormPagingModel struct {
	Cursors CursorsModel `json:"cursors"`
}

type CursorsModel struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

type LeadDataModel struct {
	CreatedTime string          `json:"created_time"`
	ID          string          `json:"id"`
	FieldData   []LeadFieldData `json:"field_data"`
}
type LeadFieldData struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}
