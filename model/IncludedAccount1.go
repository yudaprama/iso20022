package model

// Provides information about the account that is impacted or not by the standing instruction.
type IncludedAccount1 struct {

	// Identification of the securities account.
	SecuritiesAccountIdentification *Max35Text `xml:"SctiesAcctId"`

	// Indicates whether the account is impacted or not by the standing instruction.
	//
	// Yes = The account is impacted by the standing instruction.
	// No = The account is not impacted by the standing instruction.
	IncludedIndicator *YesNoIndicator `xml:"InclInd"`
}

func (i *IncludedAccount1) SetSecuritiesAccountIdentification(value string) {
	i.SecuritiesAccountIdentification = (*Max35Text)(&value)
}

func (i *IncludedAccount1) SetIncludedIndicator(value string) {
	i.IncludedIndicator = (*YesNoIndicator)(&value)
}
