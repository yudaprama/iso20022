package model

// Product characteristic applicable to this trade product.
type ProductCharacteristics3 struct {

	// Characteristics of the product.
	Characteristic *ProductCharacteristics1Choice `xml:"Chrtc,omitempty"`

	// Measurement value for this product characteristic.
	ValueMeasure *Quantity10 `xml:"ValMeasr,omitempty"`
}

func (p *ProductCharacteristics3) AddCharacteristic() *ProductCharacteristics1Choice {
	p.Characteristic = new(ProductCharacteristics1Choice)
	return p.Characteristic
}

func (p *ProductCharacteristics3) AddValueMeasure() *Quantity10 {
	p.ValueMeasure = new(Quantity10)
	return p.ValueMeasure
}
