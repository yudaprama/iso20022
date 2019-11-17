package model

// Choice of format for the denied reason.
type DeniedReason16Choice struct {

	// Specifies the reason why the request was denied.
	Code *DeniedReason4Code `xml:"Cd"`

	// Specifies the reason why the request was denied.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (d *DeniedReason16Choice) SetCode(value string) {
	d.Code = (*DeniedReason4Code)(&value)
}

func (d *DeniedReason16Choice) AddProprietary() *GenericIdentification30 {
	d.Proprietary = new(GenericIdentification30)
	return d.Proprietary
}
