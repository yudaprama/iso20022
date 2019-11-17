package model

// Provides information about the securities account.
type SecuritiesAccount6 struct {

	// Identification of the financial instrument.
	SecurityIdentification *SecurityIdentification7 `xml:"SctyId"`

	// Idenfitication of the account where financial instruments are maintained.
	SecuritiesAccountIdentification *Max35Text `xml:"SctiesAcctId"`

	// Identification of the party that owns the account.
	AccountOwnerIdentification *PartyIdentification2Choice `xml:"AcctOwnrId,omitempty"`

	// Identification of the place of safekeeping.
	SafekeepingPlace *PartyIdentification2Choice `xml:"SfkpgPlc"`

	// Provides information required for the registration of the security.
	RegistrationDetails *Max350Text `xml:"RegnDtls,omitempty"`
}

func (s *SecuritiesAccount6) AddSecurityIdentification() *SecurityIdentification7 {
	s.SecurityIdentification = new(SecurityIdentification7)
	return s.SecurityIdentification
}

func (s *SecuritiesAccount6) SetSecuritiesAccountIdentification(value string) {
	s.SecuritiesAccountIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount6) AddAccountOwnerIdentification() *PartyIdentification2Choice {
	s.AccountOwnerIdentification = new(PartyIdentification2Choice)
	return s.AccountOwnerIdentification
}

func (s *SecuritiesAccount6) AddSafekeepingPlace() *PartyIdentification2Choice {
	s.SafekeepingPlace = new(PartyIdentification2Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesAccount6) SetRegistrationDetails(value string) {
	s.RegistrationDetails = (*Max350Text)(&value)
}
