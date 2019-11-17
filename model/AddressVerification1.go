package model

// Numeric characters of the cardholder's address for verification.
type AddressVerification1 struct {

	// Numeric characters from the cardholder's address excluding the postal code (that is street number).
	AddressDigits *Max5NumericText `xml:"AdrDgts,omitempty"`

	// Numeric characters from the cardholder's postal code.
	PostalCodeDigits *Max5NumericText `xml:"PstlCdDgts,omitempty"`
}

func (a *AddressVerification1) SetAddressDigits(value string) {
	a.AddressDigits = (*Max5NumericText)(&value)
}

func (a *AddressVerification1) SetPostalCodeDigits(value string) {
	a.PostalCodeDigits = (*Max5NumericText)(&value)
}
