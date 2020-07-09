package models

// UserAuthentication define una estructura basica de un usuario
type UserAuthentication struct {
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	Fingerprint bool   `json:"fingerprint,omitempty"`
}
