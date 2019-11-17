package model

// Choice of format for the document type.
type UndertakingDocumentType2Choice struct {

	// Document type.
	//
	Code *ExternalUndertakingDocumentType2Code `xml:"Cd"`

	// Document type expressed as a proprietary code.
	Proprietary *GenericIdentification1 `xml:"Prtry"`
}

func (u *UndertakingDocumentType2Choice) SetCode(value string) {
	u.Code = (*ExternalUndertakingDocumentType2Code)(&value)
}

func (u *UndertakingDocumentType2Choice) AddProprietary() *GenericIdentification1 {
	u.Proprietary = new(GenericIdentification1)
	return u.Proprietary
}
