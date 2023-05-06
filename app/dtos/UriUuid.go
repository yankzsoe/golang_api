package dtos

type UriUuid struct {
	Id string `uri:"id" binding:"required,uuid"`
}
