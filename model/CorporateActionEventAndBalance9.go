package model

// Detailed account holdings information report for a corporate action event.
type CorporateActionEventAndBalance9 struct {

	// Provides general information related to a corporate action event.
	GeneralInformation *EventInformation7 `xml:"GnlInf"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *SecurityIdentification19 `xml:"UndrlygScty"`

	// Provides information about the balance related to a corporate action.
	Balance *CorporateActionBalanceDetails30 `xml:"Bal,omitempty"`

	// Provides additional information related to the event and the balance of the corporate action.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionEventAndBalance9) AddGeneralInformation() *EventInformation7 {
	c.GeneralInformation = new(EventInformation7)
	return c.GeneralInformation
}

func (c *CorporateActionEventAndBalance9) AddUnderlyingSecurity() *SecurityIdentification19 {
	c.UnderlyingSecurity = new(SecurityIdentification19)
	return c.UnderlyingSecurity
}

func (c *CorporateActionEventAndBalance9) AddBalance() *CorporateActionBalanceDetails30 {
	c.Balance = new(CorporateActionBalanceDetails30)
	return c.Balance
}

func (c *CorporateActionEventAndBalance9) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
