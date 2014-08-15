package hbci

type Bank struct {
	Url     string
	Name    string
	Country [3]byte
	Blz     string
}

type Dialog struct {
	DialogId string
	NextMessageNumber int
	SystemId string
}

type BankResponse struct {
	StatusCode string
	StatusMsg  string
	Payload    string
}
