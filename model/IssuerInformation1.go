package model

// Unique and unambiguous way to identify an organisation.
type IssuerInformation1 struct {

	// Unique and unambiguous way to identify an organisation.
	Identification *PartyIdentification9Choice `xml:"Id"`

	// Address for the Universal Resource Locator (URL), eg, used over the www (HTTP) service.
	URLAddress *Max256Text `xml:"URLAdr,omitempty"`
}

func (i *IssuerInformation1) AddIdentification() *PartyIdentification9Choice {
	i.Identification = new(PartyIdentification9Choice)
	return i.Identification
}

func (i *IssuerInformation1) SetURLAddress(value string) {
	i.URLAddress = (*Max256Text)(&value)
}
