package model

// Account to or from which a securities entry is made.
type SecuritiesAccount4 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm"`

	// Additional information about a financial instrument to help identify the instrument.
	FinancialInstrumentSupplementaryIdentification *Max35Text `xml:"FinInstrmSplmtryId,omitempty"`

	// Identification of a security, as assigned under a formal or proprietary identification scheme.
	FinancialInstrumentIdentification *SecurityIdentification3Choice `xml:"FinInstrmId,omitempty"`

	// Name of the financial instrument in free format text.
	FinancialInstrumentName *Max350Text `xml:"FinInstrmNm,omitempty"`

	// Specifies the current state of an account, eg, enabled or deleted.
	Status *AccountStatus1Code `xml:"Sts"`
}

func (s *SecuritiesAccount4) AddIdentification() *AccountIdentification1 {
	s.Identification = new(AccountIdentification1)
	return s.Identification
}

func (s *SecuritiesAccount4) SetName(value string) {
	s.Name = (*Max35Text)(&value)
}

func (s *SecuritiesAccount4) SetFinancialInstrumentSupplementaryIdentification(value string) {
	s.FinancialInstrumentSupplementaryIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount4) AddFinancialInstrumentIdentification() *SecurityIdentification3Choice {
	s.FinancialInstrumentIdentification = new(SecurityIdentification3Choice)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesAccount4) SetFinancialInstrumentName(value string) {
	s.FinancialInstrumentName = (*Max350Text)(&value)
}

func (s *SecuritiesAccount4) SetStatus(value string) {
	s.Status = (*AccountStatus1Code)(&value)
}
