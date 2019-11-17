package model

// Detailed account holdings information report for a corporate action event.
type CorporateActionEventAndBalance7 struct {

	// Provides general information related to a corporate action event.
	GeneralInformation *EventInformation5 `xml:"GnlInf"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *SecurityIdentification14 `xml:"UndrlygScty"`

	// Provides information about the balance related to a corporate action.
	Balance *CorporateActionBalanceDetails9 `xml:"Bal,omitempty"`

	// Provides additional information related to the event and the balance of the corporate action.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionEventAndBalance7) AddGeneralInformation() *EventInformation5 {
	c.GeneralInformation = new(EventInformation5)
	return c.GeneralInformation
}

func (c *CorporateActionEventAndBalance7) AddUnderlyingSecurity() *SecurityIdentification14 {
	c.UnderlyingSecurity = new(SecurityIdentification14)
	return c.UnderlyingSecurity
}

func (c *CorporateActionEventAndBalance7) AddBalance() *CorporateActionBalanceDetails9 {
	c.Balance = new(CorporateActionBalanceDetails9)
	return c.Balance
}

func (c *CorporateActionEventAndBalance7) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
