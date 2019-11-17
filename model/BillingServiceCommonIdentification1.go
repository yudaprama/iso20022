package model

// Common identification of a service to be billed.
type BillingServiceCommonIdentification1 struct {

	// Defines the issuer of the common code, such as "AFP".
	Issuer *Max6Text `xml:"Issr"`

	// Standard reference code used to uniquely identify this service across financial institutions. This is not the financial institution’s internal bank service identification.
	Identification *Max8Text `xml:"Id"`
}

func (b *BillingServiceCommonIdentification1) SetIssuer(value string) {
	b.Issuer = (*Max6Text)(&value)
}

func (b *BillingServiceCommonIdentification1) SetIdentification(value string) {
	b.Identification = (*Max8Text)(&value)
}
