package model

// Detailed account holdings information report for a corporate action event.
type CorporateActionEventAndBalance1 struct {

	// Provides general information related to a corporate action event.
	GeneralInformation *EventInformation1 `xml:"GnlInf"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *UnderlyingSecurity1 `xml:"UndrlygScty"`

	// Provides information about the balance related to a corporate action.
	Balance *CorporateActionBalanceDetails4 `xml:"Bal,omitempty"`

	// Provides additional information related to the event and the balance of the corporate action.
	Extension []*Extension2 `xml:"Xtnsn,omitempty"`
}

func (c *CorporateActionEventAndBalance1) AddGeneralInformation() *EventInformation1 {
	c.GeneralInformation = new(EventInformation1)
	return c.GeneralInformation
}

func (c *CorporateActionEventAndBalance1) AddUnderlyingSecurity() *UnderlyingSecurity1 {
	c.UnderlyingSecurity = new(UnderlyingSecurity1)
	return c.UnderlyingSecurity
}

func (c *CorporateActionEventAndBalance1) AddBalance() *CorporateActionBalanceDetails4 {
	c.Balance = new(CorporateActionBalanceDetails4)
	return c.Balance
}

func (c *CorporateActionEventAndBalance1) AddExtension() *Extension2 {
	newValue := new(Extension2)
	c.Extension = append(c.Extension, newValue)
	return newValue
}
