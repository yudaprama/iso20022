package model

// Identification of the account expressed with a data source scheme.
type AccountIdentification3 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id"`

	// Entity that assigns the information.
	Issuer *Max8Text `xml:"Issr"`

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Information *Exact4AlphaNumericText `xml:"Inf"`
}

func (a *AccountIdentification3) AddIdentification() *AccountIdentification1 {
	a.Identification = new(AccountIdentification1)
	return a.Identification
}

func (a *AccountIdentification3) SetIssuer(value string) {
	a.Issuer = (*Max8Text)(&value)
}

func (a *AccountIdentification3) SetInformation(value string) {
	a.Information = (*Exact4AlphaNumericText)(&value)
}
