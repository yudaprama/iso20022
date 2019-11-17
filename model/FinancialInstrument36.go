package model

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument36 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account16 `xml:"TrfeeAcct,omitempty"`

	// Sub-accounts that are grouped in a master or omnibus account.
	SubAccountDetails *SubAccount1 `xml:"SubAcctDtls,omitempty"`
}

func (f *FinancialInstrument36) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument36) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument36) AddTransfereeAccount() *Account16 {
	f.TransfereeAccount = new(Account16)
	return f.TransfereeAccount
}

func (f *FinancialInstrument36) AddSubAccountDetails() *SubAccount1 {
	f.SubAccountDetails = new(SubAccount1)
	return f.SubAccountDetails
}
