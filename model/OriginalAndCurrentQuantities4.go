package model

// Original and current value of an asset-back instrument.
type OriginalAndCurrentQuantities4 struct {

	// Quantity expressed as an amount representing the face amount, ie, the principal, of a debt instrument.
	FaceAmount *RestrictedFINImpliedCurrencyAndAmount `xml:"FaceAmt"`

	// Quantity expressed as an amount representing the current amortised face amount of a bond, for example, a periodic reduction/increase of a bond's principal amount.
	AmortisedValue *RestrictedFINImpliedCurrencyAndAmount `xml:"AmtsdVal"`
}

func (o *OriginalAndCurrentQuantities4) SetFaceAmount(value, currency string) {
	o.FaceAmount = NewRestrictedFINImpliedCurrencyAndAmount(value, currency)
}

func (o *OriginalAndCurrentQuantities4) SetAmortisedValue(value, currency string) {
	o.AmortisedValue = NewRestrictedFINImpliedCurrencyAndAmount(value, currency)
}
