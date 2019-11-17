package model

// Information about third party rights.
type ThirdPartyRights1 struct {

	// Type of third party right.
	Type *Max35Text `xml:"Tp"`

	// Timestamp for the third party right.
	DateTime *ISODateTime `xml:"DtTm,omitempty"`

	// Party that holds the third party right.
	Holder *PartyIdentification70Choice `xml:"Hldr,omitempty"`

	// Identification of the holder with a Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`

	// Amount of the third party right.
	Amount *ActiveCurrencyAndAmount `xml:"Amt,omitempty"`

	// Description of the third party right.
	Description *Max350Text `xml:"Desc,omitempty"`
}

func (t *ThirdPartyRights1) SetType(value string) {
	t.Type = (*Max35Text)(&value)
}

func (t *ThirdPartyRights1) SetDateTime(value string) {
	t.DateTime = (*ISODateTime)(&value)
}

func (t *ThirdPartyRights1) AddHolder() *PartyIdentification70Choice {
	t.Holder = new(PartyIdentification70Choice)
	return t.Holder
}

func (t *ThirdPartyRights1) SetLegalEntityIdentifier(value string) {
	t.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}

func (t *ThirdPartyRights1) SetAmount(value, currency string) {
	t.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (t *ThirdPartyRights1) SetDescription(value string) {
	t.Description = (*Max350Text)(&value)
}
