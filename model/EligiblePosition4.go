package model

// Specifies the voting entitlement.
type EligiblePosition4 struct {

	// Identification of the securities account.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Identifies party that legally owns the account.
	AccountOwner *PartyIdentification39 `xml:"AcctOwnr,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	HoldingBalance []*HoldingBalance6 `xml:"HldgBal,omitempty"`

	// Identifies owner of the voting rights.
	RightsHolder []*PartyIdentification39 `xml:"RghtsHldr,omitempty"`
}

func (e *EligiblePosition4) SetAccountIdentification(value string) {
	e.AccountIdentification = (*Max35Text)(&value)
}

func (e *EligiblePosition4) AddAccountOwner() *PartyIdentification39 {
	e.AccountOwner = new(PartyIdentification39)
	return e.AccountOwner
}

func (e *EligiblePosition4) AddHoldingBalance() *HoldingBalance6 {
	newValue := new(HoldingBalance6)
	e.HoldingBalance = append(e.HoldingBalance, newValue)
	return newValue
}

func (e *EligiblePosition4) AddRightsHolder() *PartyIdentification39 {
	newValue := new(PartyIdentification39)
	e.RightsHolder = append(e.RightsHolder, newValue)
	return newValue
}
