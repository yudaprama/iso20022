package model

// Attribute of the certificate service to be put in the certificate extensions, or to be used for the request.
type RelativeDistinguishedName2 struct {

	// Type of attribute of a distinguished name (see X.500).
	AttributeType *AttributeType2Code `xml:"AttrTp"`

	// Value of the attribute of a distinguished name (see X.500).
	AttributeValue *Max140Text `xml:"AttrVal"`
}

func (r *RelativeDistinguishedName2) SetAttributeType(value string) {
	r.AttributeType = (*AttributeType2Code)(&value)
}

func (r *RelativeDistinguishedName2) SetAttributeValue(value string) {
	r.AttributeValue = (*Max140Text)(&value)
}
