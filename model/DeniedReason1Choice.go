package model

// Choice of format for the denied reason.
type DeniedReason1Choice struct {

	// Specifies the reason why the request was denied.
	Code *DeniedReason5Code `xml:"Cd"`

	// Specifies the reason why the request was denied.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (d *DeniedReason1Choice) SetCode(value string) {
	d.Code = (*DeniedReason5Code)(&value)
}

func (d *DeniedReason1Choice) AddProprietary() *GenericIdentification20 {
	d.Proprietary = new(GenericIdentification20)
	return d.Proprietary
}
