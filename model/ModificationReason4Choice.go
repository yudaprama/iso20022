package model

// Choice of format for the Modification reason.
type ModificationReason4Choice struct {

	// Specifies the reason why the related instruction is modified, or the related modification request is executed.
	Code *ModifiedStatusReason1Code `xml:"Cd"`

	// Specifies the reason why the related instruction is modified, or the related modification request is executed.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (m *ModificationReason4Choice) SetCode(value string) {
	m.Code = (*ModifiedStatusReason1Code)(&value)
}

func (m *ModificationReason4Choice) AddProprietary() *GenericIdentification30 {
	m.Proprietary = new(GenericIdentification30)
	return m.Proprietary
}
