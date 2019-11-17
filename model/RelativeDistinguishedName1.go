package model

// Relative distinguished name defined by X.500 and X.509.
type RelativeDistinguishedName1 struct {

	// Type of attribute of a distinguished name (see X.500).
	AttributeType *AttributeType1Code `xml:"AttrTp"`

	// Value of the attribute of a distinguished name (see X.500).
	AttributeValue *Max140Text `xml:"AttrVal"`
}

func (r *RelativeDistinguishedName1) SetAttributeType(value string) {
	r.AttributeType = (*AttributeType1Code)(&value)
}

func (r *RelativeDistinguishedName1) SetAttributeValue(value string) {
	r.AttributeValue = (*Max140Text)(&value)
}
