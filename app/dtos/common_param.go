package dtos

type CommonParam struct {
	Where  string `json:"where"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}

type UriUuid struct {
	Id string `uri:"id" binding:"required,uuid4_rfc4122"`
}

type UriName struct {
	Name string `uri:"name" binding:"required"`
}
