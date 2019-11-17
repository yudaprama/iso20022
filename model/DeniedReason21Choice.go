package model

// Choice of format for the denied reason.
type DeniedReason21Choice struct {

	// Specifies the reason why the request was denied.
	Code *DeniedReason4Code `xml:"Cd"`

	// Specifies the reason why the request was denied.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (d *DeniedReason21Choice) SetCode(value string) {
	d.Code = (*DeniedReason4Code)(&value)
}

func (d *DeniedReason21Choice) AddProprietary() *GenericIdentification47 {
	d.Proprietary = new(GenericIdentification47)
	return d.Proprietary
}
