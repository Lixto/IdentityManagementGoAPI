package models

// User define una estructura basica de un usuario
type User struct {
	ID              int    `json:"ID,omitempty"` //omitempty : the field is omitted from the object if its value is empty
	Email           string `json:"Email,omitempty"`
	Password        string `json:"Password,omitempty"`
	Salt            string `json:"Salt,omitempty"`
	LongPassword    int    `json:"Long,omitempty"`
	MayusPassword   bool   `json:"Mayus,omitempty"`
	SpecialPassword bool   `json:"Special,omitempty"`
	NumbersPassword bool   `json:"Numbers,omitempty"`
	Address         string `json:"Address,omitempty"`
	Contract        string `json:"contract"`
	Fingerprint     bool   `json:"fingerprint,omitempty"`
}
