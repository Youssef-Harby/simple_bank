package events

type DMS struct {
	Data DMSContent `json:"data"`
}

type DMSContent struct {
	Details []DMSBody `json:"details"`
}

type DMSBody struct {
	CompanyId string        `json:"companyId"`
	URLs      []URL         `json:"urls"`
	Fillers   []interface{} `json:"fillers"`
}

type URL struct {
	URLID        int    `json:"urlId"`
	DocumentType string `json:"documentType"`
	DocumentName string `json:"documentName"`
	URL          string `json:"url"`
}
