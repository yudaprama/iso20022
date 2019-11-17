package model

// Tangible output or service produced by human or mechanical effort, or by a natural process for purposes of specifying a product.
type TradeProduct2 struct {

	// Identification of the product.
	Identification []*ProductIdentifier2Choice `xml:"Id,omitempty"`

	// Name of a product.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Information about the goods and/or services of a trade transaction.
	Description *Max140Text `xml:"Desc,omitempty"`

	// Country of origin of the product.
	CountryOfOrigin []*CountryCodeAndName1 `xml:"CtryOfOrgn,omitempty"`

	// Identifies the characteristic of a product.
	ProductCharacteristics []*ProductCharacteristics3 `xml:"PdctChrtcs,omitempty"`

	// Category of the product.
	ProductCategory []*ProductCategory1Choice `xml:"PdctCtgy,omitempty"`

	// Unique global serial identifier for this product instance.
	GlobalSerialIdentifier []*Max35Text `xml:"GblSrlIdr,omitempty"`
}

func (t *TradeProduct2) AddIdentification() *ProductIdentifier2Choice {
	newValue := new(ProductIdentifier2Choice)
	t.Identification = append(t.Identification, newValue)
	return newValue
}

func (t *TradeProduct2) SetName(value string) {
	t.Name = (*Max35Text)(&value)
}

func (t *TradeProduct2) SetDescription(value string) {
	t.Description = (*Max140Text)(&value)
}

func (t *TradeProduct2) AddCountryOfOrigin() *CountryCodeAndName1 {
	newValue := new(CountryCodeAndName1)
	t.CountryOfOrigin = append(t.CountryOfOrigin, newValue)
	return newValue
}

func (t *TradeProduct2) AddProductCharacteristics() *ProductCharacteristics3 {
	newValue := new(ProductCharacteristics3)
	t.ProductCharacteristics = append(t.ProductCharacteristics, newValue)
	return newValue
}

func (t *TradeProduct2) AddProductCategory() *ProductCategory1Choice {
	newValue := new(ProductCategory1Choice)
	t.ProductCategory = append(t.ProductCategory, newValue)
	return newValue
}

func (t *TradeProduct2) AddGlobalSerialIdentifier(value string) {
	t.GlobalSerialIdentifier = append(t.GlobalSerialIdentifier, (*Max35Text)(&value))
}
