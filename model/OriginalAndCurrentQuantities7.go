package model

// Signed face amount and amortised value.
type OriginalAndCurrentQuantities7 struct {

	// Sign of the quantity of security.
	ShortLongPosition *ShortLong1Code `xml:"ShrtLngPos"`

	// Quantity expressed as an amount representing the face amount, that is, the principal, of a debt instrument.
	FaceAmount *RestrictedFINImpliedCurrencyAndAmount `xml:"FaceAmt"`

	// Quantity expressed as an amount representing the current amortised face amount of a bond, for example, a periodic reduction/increase of a bond's principal amount.
	AmortisedValue *RestrictedFINImpliedCurrencyAndAmount `xml:"AmtsdVal"`
}

func (o *OriginalAndCurrentQuantities7) SetShortLongPosition(value string) {
	o.ShortLongPosition = (*ShortLong1Code)(&value)
}

func (o *OriginalAndCurrentQuantities7) SetFaceAmount(value, currency string) {
	o.FaceAmount = NewRestrictedFINImpliedCurrencyAndAmount(value, currency)
}

func (o *OriginalAndCurrentQuantities7) SetAmortisedValue(value, currency string) {
	o.AmortisedValue = NewRestrictedFINImpliedCurrencyAndAmount(value, currency)
}
