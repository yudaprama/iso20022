package model

// Specifies the owner type and mandate type.
type OwnerType1 struct {

	// Type of ownership.
	Type *AccountOwnerType1Code `xml:"Tp"`

	// Type of mandate.
	MandateType *AccountPermissionType1Code `xml:"MndtTp,omitempty"`

	// Additional information about owner type or mandate type in proprietary format.
	Proprietary *GenericIdentification1 `xml:"Prtry,omitempty"`
}

func (o *OwnerType1) SetType(value string) {
	o.Type = (*AccountOwnerType1Code)(&value)
}

func (o *OwnerType1) SetMandateType(value string) {
	o.MandateType = (*AccountPermissionType1Code)(&value)
}

func (o *OwnerType1) AddProprietary() *GenericIdentification1 {
	o.Proprietary = new(GenericIdentification1)
	return o.Proprietary
}
