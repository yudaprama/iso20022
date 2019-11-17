package model

// Choice of format for the denied reason.
type DeniedReason24Choice struct {

	// Specifies the reason why the request was denied.
	Code *DeniedReason6Code `xml:"Cd"`

	// Specifies the reason why the request was denied.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (d *DeniedReason24Choice) SetCode(value string) {
	d.Code = (*DeniedReason6Code)(&value)
}

func (d *DeniedReason24Choice) AddProprietary() *GenericIdentification47 {
	d.Proprietary = new(GenericIdentification47)
	return d.Proprietary
}
