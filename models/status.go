package models

type GStates struct {
	RunningQueries    uint64 `xml:"RunningQueries" json:"RunningQueries"`
	TotalQueries      uint64 `xml:"TotalQueries" json:"TotalQueries"`
	FailedQueries     uint64 `xml:"FailedQueries" json:"FailedQueries"`
	SuccessfulQueries uint64 `xml:"SuccessfulQueries" json:"SuccessfulQueries"`
}
