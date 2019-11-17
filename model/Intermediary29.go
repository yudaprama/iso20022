package model

// Party that provides services to investors relating to financial products.
type Intermediary29 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification100 `xml:"Id"`

	// Function performed by the intermediary.
	Role *Role5Choice `xml:"Role"`

	// Counterparties eligibility as defined by article 24 of the EU MiFID Directive applicable to transactions executed by investment firms for eligible counterparties.
	OrderOriginatorEligibility *OrderOriginatorEligibility1Code `xml:"OrdrOrgtrElgblty,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *Intermediary29) AddIdentification() *PartyIdentification100 {
	i.Identification = new(PartyIdentification100)
	return i.Identification
}

func (i *Intermediary29) AddRole() *Role5Choice {
	i.Role = new(Role5Choice)
	return i.Role
}

func (i *Intermediary29) SetOrderOriginatorEligibility(value string) {
	i.OrderOriginatorEligibility = (*OrderOriginatorEligibility1Code)(&value)
}

func (i *Intermediary29) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
