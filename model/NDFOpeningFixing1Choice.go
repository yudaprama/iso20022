package model

// Choice of providing the opening conditions or fixing conditions for an NDF instruction.
type NDFOpeningFixing1Choice struct {

	// Used to provide the opening information associated with an NDF trade.
	OpeningConditions *OpeningConditions1 `xml:"OpngConds"`

	// Reference of the opening confirmation provided on an NDF fixing instruction to link back to the original NDF opening confirmation.
	OpeningConfirmationReference *Max35Text `xml:"OpngConfRef"`
}

func (n *NDFOpeningFixing1Choice) AddOpeningConditions() *OpeningConditions1 {
	n.OpeningConditions = new(OpeningConditions1)
	return n.OpeningConditions
}

func (n *NDFOpeningFixing1Choice) SetOpeningConfirmationReference(value string) {
	n.OpeningConfirmationReference = (*Max35Text)(&value)
}
