package model

// Provides information about the securities account.
type SecuritiesAccount12 struct {

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Identification of the party that owns the account.
	AccountOwnerIdentification *PartyIdentification2Choice `xml:"AcctOwnrId,omitempty"`

	// Idenfitication of the account where financial instruments are maintained.
	AccountIdentification *Max35Text `xml:"AcctId"`

	// Type of balance.
	BalanceType *SecuritiesBalanceType6FormatChoice `xml:"BalTp,omitempty"`

	// Specifies the form of the financial instrument.
	SecurityHoldingForm *FormOfSecurity1Code `xml:"SctyHldgForm,omitempty"`
}

func (s *SecuritiesAccount12) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesAccount12) AddAccountOwnerIdentification() *PartyIdentification2Choice {
	s.AccountOwnerIdentification = new(PartyIdentification2Choice)
	return s.AccountOwnerIdentification
}

func (s *SecuritiesAccount12) SetAccountIdentification(value string) {
	s.AccountIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount12) AddBalanceType() *SecuritiesBalanceType6FormatChoice {
	s.BalanceType = new(SecuritiesBalanceType6FormatChoice)
	return s.BalanceType
}

func (s *SecuritiesAccount12) SetSecurityHoldingForm(value string) {
	s.SecurityHoldingForm = (*FormOfSecurity1Code)(&value)
}
