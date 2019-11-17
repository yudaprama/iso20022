package model

// Merchant performing the transaction.
type Organisation9 struct {

	// Identification of the merchant.
	Identification *GenericIdentification32 `xml:"Id"`

	// Name of the merchant as appearing on the receipt.
	CommonName *Max70Text `xml:"CmonNm,omitempty"`

	// Location category of the place where the merchant actually performed the transaction.
	LocationCategory *LocationCategory1Code `xml:"LctnCtgy,omitempty"`

	// Location of the merchant where the transaction took place, as appearing on the receipt.
	Address *Max140Text `xml:"Adr,omitempty"`

	// Country of the merchant where the transaction took place.
	CountryCode *ISO3NumericCountryCode `xml:"CtryCd,omitempty"`

	// Additional merchant data required by a card scheme.
	SchemeData *Max140Text `xml:"SchmeData,omitempty"`
}

func (o *Organisation9) AddIdentification() *GenericIdentification32 {
	o.Identification = new(GenericIdentification32)
	return o.Identification
}

func (o *Organisation9) SetCommonName(value string) {
	o.CommonName = (*Max70Text)(&value)
}

func (o *Organisation9) SetLocationCategory(value string) {
	o.LocationCategory = (*LocationCategory1Code)(&value)
}

func (o *Organisation9) SetAddress(value string) {
	o.Address = (*Max140Text)(&value)
}

func (o *Organisation9) SetCountryCode(value string) {
	o.CountryCode = (*ISO3NumericCountryCode)(&value)
}

func (o *Organisation9) SetSchemeData(value string) {
	o.SchemeData = (*Max140Text)(&value)
}
