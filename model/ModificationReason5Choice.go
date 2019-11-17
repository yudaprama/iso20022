package model

// Choice of format for the Modification reason.
type ModificationReason5Choice struct {

	// Specifies the reason why the related instruction is modified, or the related modification request is executed.
	Code *ModifiedStatusReason1Code `xml:"Cd"`

	// Specifies the reason why the related instruction is modified, or the related modification request is executed.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (m *ModificationReason5Choice) SetCode(value string) {
	m.Code = (*ModifiedStatusReason1Code)(&value)
}

func (m *ModificationReason5Choice) AddProprietary() *GenericIdentification47 {
	m.Proprietary = new(GenericIdentification47)
	return m.Proprietary
}
