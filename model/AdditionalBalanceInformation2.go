package model

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type AdditionalBalanceInformation2 struct {

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity1Choice `xml:"Qty"`

	// Reason a security is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	SubBalanceType *SecuritiesBalanceType2Code `xml:"SubBalTp"`

	// Reason a security is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	ExtendedSubBalanceType *Extended350Code `xml:"XtndedSubBalTp"`
}

func (a *AdditionalBalanceInformation2) AddQuantity() *SubBalanceQuantity1Choice {
	a.Quantity = new(SubBalanceQuantity1Choice)
	return a.Quantity
}

func (a *AdditionalBalanceInformation2) SetSubBalanceType(value string) {
	a.SubBalanceType = (*SecuritiesBalanceType2Code)(&value)
}

func (a *AdditionalBalanceInformation2) SetExtendedSubBalanceType(value string) {
	a.ExtendedSubBalanceType = (*Extended350Code)(&value)
}
