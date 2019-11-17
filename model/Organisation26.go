package model

// Merchant performing the transaction.
type Organisation26 struct {

	// Name of the merchant.
	CommonName *Max70Text `xml:"CmonNm"`

	// Location of the merchant.
	Address *Max140Text `xml:"Adr,omitempty"`

	// Country of the merchant.
	CountryCode *ISO3NumericCountryCode `xml:"CtryCd"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Identifier of the sponsored merchant assigned by the payment facilitator of their acquirer.
	RegisteredIdentifier *Max35Text `xml:"RegdIdr"`
}

func (o *Organisation26) SetCommonName(value string) {
	o.CommonName = (*Max70Text)(&value)
}

func (o *Organisation26) SetAddress(value string) {
	o.Address = (*Max140Text)(&value)
}

func (o *Organisation26) SetCountryCode(value string) {
	o.CountryCode = (*ISO3NumericCountryCode)(&value)
}

func (o *Organisation26) SetMerchantCategoryCode(value string) {
	o.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (o *Organisation26) SetRegisteredIdentifier(value string) {
	o.RegisteredIdentifier = (*Max35Text)(&value)
}
