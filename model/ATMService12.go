package model

// Deposit service provided by the ATM inside the session.
type ATMService12 struct {

	// Unique identification of the deposit service provided by the ATM inside the session.
	ServiceReference *Max35Text `xml:"SvcRef,omitempty"`

	// Codification of the type of service for the ATM.
	ATMServiceCode *Max35Text `xml:"ATMSvcCd,omitempty"`

	// Codification of the type of service for the host.
	HostServiceCode *Max35Text `xml:"HstSvcCd,omitempty"`

	// Describes the type of deposit service selected by the customer.
	ServiceType *ATMServiceType6Code `xml:"SvcTp"`

	// Identification of the variant of the service.
	ServiceVariantIdentification []*Max35Text `xml:"SvcVarntId,omitempty"`

	// True if deposit with cash back transaction.
	CashBack *TrueFalseIndicator `xml:"CshBck,omitempty"`

	// True if the deposit transaction is split in multiple accounts.
	MultiAccount *TrueFalseIndicator `xml:"MultiAcct,omitempty"`
}

func (a *ATMService12) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService12) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService12) SetHostServiceCode(value string) {
	a.HostServiceCode = (*Max35Text)(&value)
}

func (a *ATMService12) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType6Code)(&value)
}

func (a *ATMService12) AddServiceVariantIdentification(value string) {
	a.ServiceVariantIdentification = append(a.ServiceVariantIdentification, (*Max35Text)(&value))
}

func (a *ATMService12) SetCashBack(value string) {
	a.CashBack = (*TrueFalseIndicator)(&value)
}

func (a *ATMService12) SetMultiAccount(value string) {
	a.MultiAccount = (*TrueFalseIndicator)(&value)
}
