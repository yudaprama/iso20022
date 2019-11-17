package model

// Product characteristic applicable to this trade product.
type ProductCharacteristics2 struct {

	// Characteristics of the product.
	Characteristic *ProductCharacteristics1Choice `xml:"Chrtc,omitempty"`

	// Measurement value for this product characteristic.
	ValueMeasure *Quantity3 `xml:"ValMeasr,omitempty"`
}

func (p *ProductCharacteristics2) AddCharacteristic() *ProductCharacteristics1Choice {
	p.Characteristic = new(ProductCharacteristics1Choice)
	return p.Characteristic
}

func (p *ProductCharacteristics2) AddValueMeasure() *Quantity3 {
	p.ValueMeasure = new(Quantity3)
	return p.ValueMeasure
}
