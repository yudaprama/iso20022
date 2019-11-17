package model

// Provides information about the account identification and the account owner.
type SecuritiesAccount7 struct {

	// Identification of the party that owns the account.
	AccountOwnerIdentification *PartyIdentification2Choice `xml:"AcctOwnrId,omitempty"`

	// Idenfitication of the account where financial instruments are maintained.
	AccountIdentification *Max35Text `xml:"AcctId"`
}

func (s *SecuritiesAccount7) AddAccountOwnerIdentification() *PartyIdentification2Choice {
	s.AccountOwnerIdentification = new(PartyIdentification2Choice)
	return s.AccountOwnerIdentification
}

func (s *SecuritiesAccount7) SetAccountIdentification(value string) {
	s.AccountIdentification = (*Max35Text)(&value)
}
