package model

// Choice of formats to  express the type of beneficiary's certification.
type BeneficiaryCertificationType1FormatChoice struct {

	// Standard code to  specify the type of beneficiary's certification.
	Code *BeneficiaryCertificationType1Code `xml:"Cd"`

	// Proprietary code to  express the type of beneficiary's certification.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (b *BeneficiaryCertificationType1FormatChoice) SetCode(value string) {
	b.Code = (*BeneficiaryCertificationType1Code)(&value)
}

func (b *BeneficiaryCertificationType1FormatChoice) AddProprietary() *GenericIdentification13 {
	b.Proprietary = new(GenericIdentification13)
	return b.Proprietary
}
