package model

// Choice of format for the denied reason.
type DeniedReason23Choice struct {

	// Specifies the reason why the request was denied.
	Code *DeniedReason7Code `xml:"Cd"`

	// Specifies the reason why the request was denied.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (d *DeniedReason23Choice) SetCode(value string) {
	d.Code = (*DeniedReason7Code)(&value)
}

func (d *DeniedReason23Choice) AddProprietary() *GenericIdentification47 {
	d.Proprietary = new(GenericIdentification47)
	return d.Proprietary
}
