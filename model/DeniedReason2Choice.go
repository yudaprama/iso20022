package model

// Choice of format for the denied reason.
type DeniedReason2Choice struct {

	// Specifies the reason why the request was denied.
	Code *DeniedReason4Code `xml:"Cd"`

	// Specifies the reason why the request was denied.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (d *DeniedReason2Choice) SetCode(value string) {
	d.Code = (*DeniedReason4Code)(&value)
}

func (d *DeniedReason2Choice) AddProprietary() *GenericIdentification20 {
	d.Proprietary = new(GenericIdentification20)
	return d.Proprietary
}
