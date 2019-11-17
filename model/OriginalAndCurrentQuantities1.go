package model

// Original and current value of an asset-back instrument.
type OriginalAndCurrentQuantities1 struct {

	// Quantity expressed as an amount representing the face amount, ie, the principal, of a debt instrument.
	FaceAmount *ImpliedCurrencyAndAmount `xml:"FaceAmt"`

	// Quantity expressed as an amount representing the current amortised face amount of a bond, for example, a periodic reduction/increase of a bond's principal amount.
	AmortisedValue *ImpliedCurrencyAndAmount `xml:"AmtsdVal"`
}

func (o *OriginalAndCurrentQuantities1) SetFaceAmount(value, currency string) {
	o.FaceAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (o *OriginalAndCurrentQuantities1) SetAmortisedValue(value, currency string) {
	o.AmortisedValue = NewImpliedCurrencyAndAmount(value, currency)
}
