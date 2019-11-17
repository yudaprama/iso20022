package model

// Specifies the voting entitlement.
type EligiblePosition6 struct {

	// Identification of the securities account.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Identifies the party that legally owns the account.
	AccountOwner *PartyIdentification71 `xml:"AcctOwnr,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, for example, sub-balance per status.
	HoldingBalance []*HoldingBalance7 `xml:"HldgBal,omitempty"`

	// Identifies the owner of the voting rights.
	RightsHolder []*PartyIdentification71 `xml:"RghtsHldr,omitempty"`
}

func (e *EligiblePosition6) SetAccountIdentification(value string) {
	e.AccountIdentification = (*Max35Text)(&value)
}

func (e *EligiblePosition6) AddAccountOwner() *PartyIdentification71 {
	e.AccountOwner = new(PartyIdentification71)
	return e.AccountOwner
}

func (e *EligiblePosition6) AddHoldingBalance() *HoldingBalance7 {
	newValue := new(HoldingBalance7)
	e.HoldingBalance = append(e.HoldingBalance, newValue)
	return newValue
}

func (e *EligiblePosition6) AddRightsHolder() *PartyIdentification71 {
	newValue := new(PartyIdentification71)
	e.RightsHolder = append(e.RightsHolder, newValue)
	return newValue
}
