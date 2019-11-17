package model

// Detailed account holdings information report for a corporate action event.
type CorporateActionEventAndBalance12 struct {

	// Provides general information related to a corporate action event.
	GeneralInformation *EventInformation10 `xml:"GnlInf"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *SecurityIdentification20 `xml:"UndrlygScty"`

	// Provides information about the balance related to a corporate action.
	Balance *CorporateActionBalanceDetails33 `xml:"Bal,omitempty"`

	// Provides additional information related to the event and the balance of the corporate action.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionEventAndBalance12) AddGeneralInformation() *EventInformation10 {
	c.GeneralInformation = new(EventInformation10)
	return c.GeneralInformation
}

func (c *CorporateActionEventAndBalance12) AddUnderlyingSecurity() *SecurityIdentification20 {
	c.UnderlyingSecurity = new(SecurityIdentification20)
	return c.UnderlyingSecurity
}

func (c *CorporateActionEventAndBalance12) AddBalance() *CorporateActionBalanceDetails33 {
	c.Balance = new(CorporateActionBalanceDetails33)
	return c.Balance
}

func (c *CorporateActionEventAndBalance12) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
