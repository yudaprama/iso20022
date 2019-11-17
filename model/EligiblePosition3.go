package model

// Specifies the voting entitlement.
type EligiblePosition3 struct {

	// Identification of the securities account.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Identifies party that legally owns the account.
	AccountOwner *PartyIdentification9Choice `xml:"AcctOwnr,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	HoldingBalance []*HoldingBalance6 `xml:"HldgBal,omitempty"`

	// Identifies owner of the voting rights.
	RightsHolder []*PartyIdentification9Choice `xml:"RghtsHldr,omitempty"`
}

func (e *EligiblePosition3) SetAccountIdentification(value string) {
	e.AccountIdentification = (*Max35Text)(&value)
}

func (e *EligiblePosition3) AddAccountOwner() *PartyIdentification9Choice {
	e.AccountOwner = new(PartyIdentification9Choice)
	return e.AccountOwner
}

func (e *EligiblePosition3) AddHoldingBalance() *HoldingBalance6 {
	newValue := new(HoldingBalance6)
	e.HoldingBalance = append(e.HoldingBalance, newValue)
	return newValue
}

func (e *EligiblePosition3) AddRightsHolder() *PartyIdentification9Choice {
	newValue := new(PartyIdentification9Choice)
	e.RightsHolder = append(e.RightsHolder, newValue)
	return newValue
}
