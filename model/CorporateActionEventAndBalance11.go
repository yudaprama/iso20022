package model

// Detailed account holdings information report for a corporate action event.
type CorporateActionEventAndBalance11 struct {

	// Provides general information related to a corporate action event.
	GeneralInformation *EventInformation9 `xml:"GnlInf"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *SecurityIdentification19 `xml:"UndrlygScty"`

	// Provides information about the balance related to a corporate action.
	Balance *CorporateActionBalanceDetails30 `xml:"Bal,omitempty"`

	// Provides additional information related to the event and the balance of the corporate action.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionEventAndBalance11) AddGeneralInformation() *EventInformation9 {
	c.GeneralInformation = new(EventInformation9)
	return c.GeneralInformation
}

func (c *CorporateActionEventAndBalance11) AddUnderlyingSecurity() *SecurityIdentification19 {
	c.UnderlyingSecurity = new(SecurityIdentification19)
	return c.UnderlyingSecurity
}

func (c *CorporateActionEventAndBalance11) AddBalance() *CorporateActionBalanceDetails30 {
	c.Balance = new(CorporateActionBalanceDetails30)
	return c.Balance
}

func (c *CorporateActionEventAndBalance11) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
