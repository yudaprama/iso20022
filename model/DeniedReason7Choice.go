package model

// Choice of format for the denied reason.
type DeniedReason7Choice struct {

	// Specifies the reason why the request was denied.
	Code *DeniedReason6Code `xml:"Cd"`

	// Specifies the reason why the request was denied.
	Proprietary *GenericIdentification40 `xml:"Prtry"`
}

func (d *DeniedReason7Choice) SetCode(value string) {
	d.Code = (*DeniedReason6Code)(&value)
}

func (d *DeniedReason7Choice) AddProprietary() *GenericIdentification40 {
	d.Proprietary = new(GenericIdentification40)
	return d.Proprietary
}
