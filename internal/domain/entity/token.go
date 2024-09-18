package entity

type Token struct {
	Title    string
	Password string
	Token    string
}

type TokenDecrypted struct {
	Title string
	Token string
}
