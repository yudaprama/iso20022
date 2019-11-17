package model

// Choice of format for the Modification reason.
type ModificationReason2Choice struct {

	// Specifies the reason why the related instruction is modified, or the related modification request is executed.
	Code *ModifiedStatusReason1Code `xml:"Cd"`

	// Specifies the reason why the related instruction is modified, or the related modification request is executed.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (m *ModificationReason2Choice) SetCode(value string) {
	m.Code = (*ModifiedStatusReason1Code)(&value)
}

func (m *ModificationReason2Choice) AddProprietary() *GenericIdentification20 {
	m.Proprietary = new(GenericIdentification20)
	return m.Proprietary
}
