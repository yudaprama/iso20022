package model

// Merchant performing the transaction.
type Organisation25 struct {

	// Identification of the merchant.
	Identification *GenericIdentification32 `xml:"Id,omitempty"`

	// Name of the merchant as appearing on the receipt.
	CommonName *Max70Text `xml:"CmonNm,omitempty"`

	// Location category of the place where the merchant actually performed the transaction.
	LocationCategory *LocationCategory1Code `xml:"LctnCtgy,omitempty"`

	// Location and contact information of the merchant performing the transaction.
	LocationAndContact *CommunicationAddress5 `xml:"LctnAndCtct,omitempty"`

	// Additional merchant data required by a card scheme.
	SchemeData *Max140Text `xml:"SchmeData,omitempty"`
}

func (o *Organisation25) AddIdentification() *GenericIdentification32 {
	o.Identification = new(GenericIdentification32)
	return o.Identification
}

func (o *Organisation25) SetCommonName(value string) {
	o.CommonName = (*Max70Text)(&value)
}

func (o *Organisation25) SetLocationCategory(value string) {
	o.LocationCategory = (*LocationCategory1Code)(&value)
}

func (o *Organisation25) AddLocationAndContact() *CommunicationAddress5 {
	o.LocationAndContact = new(CommunicationAddress5)
	return o.LocationAndContact
}

func (o *Organisation25) SetSchemeData(value string) {
	o.SchemeData = (*Max140Text)(&value)
}
