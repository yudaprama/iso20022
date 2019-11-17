package model

// Deposit service provided by the ATM inside the session.
type ATMService13 struct {

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

	// True if this is not the final deposit.
	PartialDeposit *TrueFalseIndicator `xml:"PrtlDpst,omitempty"`
}

func (a *ATMService13) SetServiceReference(value string) {
	a.ServiceReference = (*Max35Text)(&value)
}

func (a *ATMService13) SetATMServiceCode(value string) {
	a.ATMServiceCode = (*Max35Text)(&value)
}

func (a *ATMService13) SetHostServiceCode(value string) {
	a.HostServiceCode = (*Max35Text)(&value)
}

func (a *ATMService13) SetServiceType(value string) {
	a.ServiceType = (*ATMServiceType6Code)(&value)
}

func (a *ATMService13) AddServiceVariantIdentification(value string) {
	a.ServiceVariantIdentification = append(a.ServiceVariantIdentification, (*Max35Text)(&value))
}

func (a *ATMService13) SetCashBack(value string) {
	a.CashBack = (*TrueFalseIndicator)(&value)
}

func (a *ATMService13) SetMultiAccount(value string) {
	a.MultiAccount = (*TrueFalseIndicator)(&value)
}

func (a *ATMService13) SetPartialDeposit(value string) {
	a.PartialDeposit = (*TrueFalseIndicator)(&value)
}
