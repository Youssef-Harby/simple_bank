package events

import "time"

type Momentum struct {
	Data MomentumContent `json:"data"`
}

type MomentumContent struct {
	Companies []MomentumBody `json:"companies"`
}

type MomentumBody struct {
	CompanyId                   string        `json:"companyId"`
	InstagramFollowers          *int          `json:"instagramFollowers"`
	TwitterFollowers            *int          `json:"twitterFollowers"`
	AppStoreRatings             *int          `json:"appStoreRatings"`
	GoogleStoreRatings          *int          `json:"googleStoreRatings"`
	GoogleStoreInstalls         *int          `json:"googleStoreInstalls"`
	AppStoreDownloads           *int          `json:"appStoreDownloads"`
	Growth1month                *int          `json:"growth1month"`
	Growth6month                *int          `json:"growth6month"`
	Growth12month               *int          `json:"growth12month"`
	Growth24month               *int          `json:"growth24month"`
	LinkedinEmployees           *int          `json:"linkedinEmployees"`
	LinkedinJobs                *int          `json:"linkedinJobs"`
	LinkedinFollowers           *int          `json:"linkedinFollowers"`
	FacebookLikes               *int          `json:"facebookLikes"`
	FacebookFollowers           *int          `json:"facebookFollowers"`
	TwitterUrl                  *string       `json:"twitterUrl"`
	LinkedinProfileUrl          *string       `json:"linkedinProfileUrl"`
	InstagramUrl                *string       `json:"instagramUrl"`
	LinkedinUrl                 *string       `json:"linkedinUrl"`
	FacebookUrl                 *string       `json:"facebookUrl"`
	LinkedinApplicantProfileURL *string       `json:"linkedinApplicantProfileURL"`
	LinkedinFounderProfileURL   *string       `json:"linkedinFounderProfileURL"`
	CaptureDate                 time.Time     `json:"captureDate"`
	Notes                       []string      `json:"notes"`
	Fillers                     []interface{} `json:"fillers"`
}
