package model

// Merchant performing the transaction.
type Organisation5 struct {

	// Identification of the merchant.
	Identification *GenericIdentification32 `xml:"Id,omitempty"`

	// Name of the merchant as appearing on the receipt.
	CommonName *Max35Text `xml:"CmonNm,omitempty"`

	// Location category of the place where the merchant actually performed the transaction.
	LocationCategory *LocationCategory1Code `xml:"LctnCtgy,omitempty"`

	// Location of the merchant where the transaction took place, as appearing on the receipt.
	Address *Max70Text `xml:"Adr,omitempty"`

	// Country of the merchant where the transaction took place.
	CountryCode *ISO3ACountryCode `xml:"CtryCd,omitempty"`

	// Additional merchant data required by a card scheme.
	SchemeData *Max140Text `xml:"SchmeData,omitempty"`
}

func (o *Organisation5) AddIdentification() *GenericIdentification32 {
	o.Identification = new(GenericIdentification32)
	return o.Identification
}

func (o *Organisation5) SetCommonName(value string) {
	o.CommonName = (*Max35Text)(&value)
}

func (o *Organisation5) SetLocationCategory(value string) {
	o.LocationCategory = (*LocationCategory1Code)(&value)
}

func (o *Organisation5) SetAddress(value string) {
	o.Address = (*Max70Text)(&value)
}

func (o *Organisation5) SetCountryCode(value string) {
	o.CountryCode = (*ISO3ACountryCode)(&value)
}

func (o *Organisation5) SetSchemeData(value string) {
	o.SchemeData = (*Max140Text)(&value)
}
