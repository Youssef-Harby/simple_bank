package events

type BSR struct {
	Data BSRContent `json:"data"`
}

type BSRContent struct {
	Statuses []BSRBody `json:"statuses"`
}

type BSRBody struct {
	CompanyID         string        `json:"companyId"`
	BankNames         string        `json:"bankNames"`
	CurrentCB         Money         `json:"currentCB"`
	Average3mdeposits Money         `json:"average3mdeposits"`
	Runway            float64       `json:"runway"`
	SameAsLegalName   *bool         `json:"sameAsLegalName"`
	BankAccountCount  *int          `json:"bankAccountCount"`
	Fillers           []interface{} `json:"fillers"`
}
