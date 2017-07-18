package saleissue

type SaleIssue struct {
	Issue     string `json:"issue"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	OpenTime  int64  `json:"open_time"`
}
