package model

// Choice of format for the denied reason.
type DeniedReason15Choice struct {

	// Specifies the reason why the request was denied.
	Code *DeniedReason6Code `xml:"Cd"`

	// Specifies the reason why the request was denied.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (d *DeniedReason15Choice) SetCode(value string) {
	d.Code = (*DeniedReason6Code)(&value)
}

func (d *DeniedReason15Choice) AddProprietary() *GenericIdentification30 {
	d.Proprietary = new(GenericIdentification30)
	return d.Proprietary
}
