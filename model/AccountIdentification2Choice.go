package model

// Choice between the identification of a cash or securities account.
type AccountIdentification2Choice struct {

	// Identification of a cash account.
	CashAccountIdentification *Max35Text `xml:"CshAcctId"`

	// Identification of a securities account.
	SecuritiesAccountIdentification *Max35Text `xml:"SctiesAcctId"`
}

func (a *AccountIdentification2Choice) SetCashAccountIdentification(value string) {
	a.CashAccountIdentification = (*Max35Text)(&value)
}

func (a *AccountIdentification2Choice) SetSecuritiesAccountIdentification(value string) {
	a.SecuritiesAccountIdentification = (*Max35Text)(&value)
}
