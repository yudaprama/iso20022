package model

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument25 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Specifies whether the financial instrument is transferred as an asset or as cash.
	TransferType *TransferType1Code `xml:"TrfTp"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account6 `xml:"TrfeeAcct,omitempty"`
}

func (f *FinancialInstrument25) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument25) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument25) SetTransferType(value string) {
	f.TransferType = (*TransferType1Code)(&value)
}

func (f *FinancialInstrument25) AddTransfereeAccount() *Account6 {
	f.TransfereeAccount = new(Account6)
	return f.TransfereeAccount
}
