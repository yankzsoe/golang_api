package dtos

type CommonParam struct {
	Where  string `json:"where"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}
