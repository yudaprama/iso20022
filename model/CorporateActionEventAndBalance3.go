package model

// Detailed account holdings information report for a corporate action event.
type CorporateActionEventAndBalance3 struct {

	// Provides general information related to a corporate action event.
	GeneralInformation *EventInformation1 `xml:"GnlInf"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *UnderlyingSecurity3 `xml:"UndrlygScty"`

	// Provides information about the balance related to a corporate action.
	Balance *CorporateActionBalanceDetails4 `xml:"Bal,omitempty"`

	// Provides additional information related to the event and the balance of the corporate action.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionEventAndBalance3) AddGeneralInformation() *EventInformation1 {
	c.GeneralInformation = new(EventInformation1)
	return c.GeneralInformation
}

func (c *CorporateActionEventAndBalance3) AddUnderlyingSecurity() *UnderlyingSecurity3 {
	c.UnderlyingSecurity = new(UnderlyingSecurity3)
	return c.UnderlyingSecurity
}

func (c *CorporateActionEventAndBalance3) AddBalance() *CorporateActionBalanceDetails4 {
	c.Balance = new(CorporateActionBalanceDetails4)
	return c.Balance
}

func (c *CorporateActionEventAndBalance3) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
