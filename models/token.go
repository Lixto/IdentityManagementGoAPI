package models

//TokenAuthentication struct to marshal
type TokenAuthentication struct {
	Token                 string `json:"token" form:"token"`
	ExpirationTimeForUser string `json:"userExpitationTime" form:"userExpitationTime"`
	ExpirationTime        int64  `json:"expitationTime" form:"expitationTime"`
}
