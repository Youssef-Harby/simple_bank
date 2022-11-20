package events

type Scoring struct {
	Data ScoringContent `json:"data"`
}

type ScoringContent struct {
	Scores []ScoringBody `json:"scores"`
}

type ScoringBody struct {
	CompanyId         string        `json:"companyId"`
	InceptionScore    float64       `json:"inceptionScore"`
	VCScore           float64       `json:"vCScore"`
	FundingScore      float64       `json:"fundingScore"`
	LikelihoodToRaise float64       `json:"likelihoodToRaise"`
	Fillers           []interface{} `json:"fillers"`
}
