package models

//Response of the server, it's only another approach we are not using it
type Response struct {
	Ok  bool   // true -> correcto, false -> error
	Msg string // mensaje adicional
}
