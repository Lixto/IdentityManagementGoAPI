package models

//ResponseRegister of the server, it's only another approach we are not using it
type ResponseRegister struct {
	Ok         bool   `json:"ok"`
	Msg        string `json:"msg"`
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
	Address    string `json:"address"`
	Password   string `json:"password"`
	Contract   string `json:"contract"`
}
