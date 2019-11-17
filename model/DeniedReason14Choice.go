package model

// Choice of format for the denied reason.
type DeniedReason14Choice struct {

	// Specifies the reason why the request was denied.
	Code *DeniedReason7Code `xml:"Cd"`

	// Specifies the reason why the request was denied.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (d *DeniedReason14Choice) SetCode(value string) {
	d.Code = (*DeniedReason7Code)(&value)
}

func (d *DeniedReason14Choice) AddProprietary() *GenericIdentification30 {
	d.Proprietary = new(GenericIdentification30)
	return d.Proprietary
}
