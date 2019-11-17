package model

// Specifies the voting entitlement.
type EligiblePosition5 struct {

	// Identification of the securities account.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Identifies the party that legally owns the account.
	AccountOwner *PartyIdentification40Choice `xml:"AcctOwnr,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, for example, sub-balance per status.
	HoldingBalance []*HoldingBalance7 `xml:"HldgBal,omitempty"`

	// Identifies the owner of the voting rights.
	RightsHolder []*PartyIdentification40Choice `xml:"RghtsHldr,omitempty"`
}

func (e *EligiblePosition5) SetAccountIdentification(value string) {
	e.AccountIdentification = (*Max35Text)(&value)
}

func (e *EligiblePosition5) AddAccountOwner() *PartyIdentification40Choice {
	e.AccountOwner = new(PartyIdentification40Choice)
	return e.AccountOwner
}

func (e *EligiblePosition5) AddHoldingBalance() *HoldingBalance7 {
	newValue := new(HoldingBalance7)
	e.HoldingBalance = append(e.HoldingBalance, newValue)
	return newValue
}

func (e *EligiblePosition5) AddRightsHolder() *PartyIdentification40Choice {
	newValue := new(PartyIdentification40Choice)
	e.RightsHolder = append(e.RightsHolder, newValue)
	return newValue
}
