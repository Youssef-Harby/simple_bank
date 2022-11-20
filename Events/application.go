package events

import "time"

type Application struct {
	Data ApplicationContent `json:"data"`
}

type ApplicationContent struct {
	Applications []ApplicationBody `json:"applications"`
}

type ApplicationBody struct {
	CompanyID               string        `json:"companyId"`
	TaxID                   string        `json:"taxId"`
	ApplicationDate         time.Time     `json:"applicationDate"`
	CompanyName             string        `json:"companyName"`
	CompanyLegalName        string        `json:"companyLegalName"`
	LanguageId              string        `json:"languageId"`
	City                    string        `json:"city"`
	State                   string        `json:"state"`
	Website                 string        `json:"website"`
	OperationalHeadquarters string        `json:"operationalHeadquarters"`
	CountryOfIncorporation  string        `json:"countryOfIncorporation"`
	Industry                string        `json:"industry"`
	CompanySize             string        `json:"companySize"`
	FundingType             string        `json:"fundingType"`
	Investors               []string      `json:"investors"`
	StreetAddress           string        `json:"streetAddress"`
	ZipCode                 string        `json:"zipCode"`
	KYB                     string        `json:"KYB"`
	UTM                     string        `json:"UTM"`
	PreApproval             string        `json:"preApproval"`
	FoundingDate            time.Time     `json:"foundingDate"`
	Applicants              []Applicant   `json:"applicants"`
	RequestedCreditLine     Money         `json:"requestedCreditLine"`
	StatedCashBalance       Money         `json:"statedCashBalance"`
	StatedAverageRevenue    Money         `json:"statedAverageRevenue"`
	StatedAverageExpense    Money         `json:"statedAverageExpense"`
	StatedBurnRate          Money         `json:"statedBurnRate"`
	TotalFunding            Money         `json:"TotalFunding"`
	AccountExecutive        *string       `json:"accountExecutive"`
	Fillers                 []interface{} `json:"fillers"`
}

type Applicant struct {
	ApplicantDetailId int    `json:"applicantDetailId"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	JobTitle          string `json:"jobTitle"`
	Mobile            string `json:"mobile"`
	Email             string `json:"email"`
	ControlPerson     bool   `json:"controlPerson"`
}

type Money struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
