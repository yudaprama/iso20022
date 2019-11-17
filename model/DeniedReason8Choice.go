package model

// Choice of format for the denied reason.
type DeniedReason8Choice struct {

	// Specifies the reason why the request was denied.
	Code *DeniedReason7Code `xml:"Cd"`

	// Specifies the reason why the request was denied.
	Proprietary *GenericIdentification40 `xml:"Prtry"`
}

func (d *DeniedReason8Choice) SetCode(value string) {
	d.Code = (*DeniedReason7Code)(&value)
}

func (d *DeniedReason8Choice) AddProprietary() *GenericIdentification40 {
	d.Proprietary = new(GenericIdentification40)
	return d.Proprietary
}
