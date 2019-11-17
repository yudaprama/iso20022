package model

// Choice of format for the document type.
type UndertakingDocumentType1Choice struct {

	// Document type.
	//
	Code *ExternalUndertakingDocumentType1Code `xml:"Cd"`

	// Document type expressed as a proprietary code.
	Proprietary *GenericIdentification1 `xml:"Prtry"`
}

func (u *UndertakingDocumentType1Choice) SetCode(value string) {
	u.Code = (*ExternalUndertakingDocumentType1Code)(&value)
}

func (u *UndertakingDocumentType1Choice) AddProprietary() *GenericIdentification1 {
	u.Proprietary = new(GenericIdentification1)
	return u.Proprietary
}
