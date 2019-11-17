package model

// Specifies the voting entitlement.
type EligiblePosition2 struct {

	// Identification of the securities account.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Identifies party that legally owns the account.
	AccountOwner *PartyIdentification9Choice `xml:"AcctOwnr,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	HoldingBalance []*HoldingBalance3 `xml:"HldgBal,omitempty"`

	// Identifies owner of the voting rights.
	RightsHolder []*PartyIdentification9Choice `xml:"RghtsHldr,omitempty"`
}

func (e *EligiblePosition2) SetAccountIdentification(value string) {
	e.AccountIdentification = (*Max35Text)(&value)
}

func (e *EligiblePosition2) AddAccountOwner() *PartyIdentification9Choice {
	e.AccountOwner = new(PartyIdentification9Choice)
	return e.AccountOwner
}

func (e *EligiblePosition2) AddHoldingBalance() *HoldingBalance3 {
	newValue := new(HoldingBalance3)
	e.HoldingBalance = append(e.HoldingBalance, newValue)
	return newValue
}

func (e *EligiblePosition2) AddRightsHolder() *PartyIdentification9Choice {
	newValue := new(PartyIdentification9Choice)
	e.RightsHolder = append(e.RightsHolder, newValue)
	return newValue
}
