package models

type Train struct {
	ID       int64  `json:"id"`
	Number   string `json:"number"`
	Type     string `json:"type"`
	Capacity string `json:"capacity"`
	Status   string `json:"status"`
}
