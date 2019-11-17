package model

// Identification of a party. The party can be identified by providing the party's name and optionally, the BIC, account number, address, clearing system identification or LEI can also be provided.
type PartyIdentification59 struct {

	// Identification of the party expressed as the party's name.
	PartyName *Max34Text `xml:"PtyNm,omitempty"`

	// Identification of the party expressed as a BIC and an optional alternative identifier.
	AnyBIC *PartyIdentification44 `xml:"AnyBIC,omitempty"`

	// Identification of the party's account number.
	AccountNumber *Max34Text `xml:"AcctNb,omitempty"`

	// Identification of the party's address.
	Address *Max105Text `xml:"Adr,omitempty"`

	// Choice of a clearing system identifier.
	ClearingSystemIdentification *ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty"`

	// Identification of the Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`
}

func (p *PartyIdentification59) SetPartyName(value string) {
	p.PartyName = (*Max34Text)(&value)
}

func (p *PartyIdentification59) AddAnyBIC() *PartyIdentification44 {
	p.AnyBIC = new(PartyIdentification44)
	return p.AnyBIC
}

func (p *PartyIdentification59) SetAccountNumber(value string) {
	p.AccountNumber = (*Max34Text)(&value)
}

func (p *PartyIdentification59) SetAddress(value string) {
	p.Address = (*Max105Text)(&value)
}

func (p *PartyIdentification59) AddClearingSystemIdentification() *ClearingSystemIdentification2Choice {
	p.ClearingSystemIdentification = new(ClearingSystemIdentification2Choice)
	return p.ClearingSystemIdentification
}

func (p *PartyIdentification59) SetLegalEntityIdentifier(value string) {
	p.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}
