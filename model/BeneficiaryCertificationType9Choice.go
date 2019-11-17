package model

// Choice between a standard code or proprietary code to specify the type of beneficiary certification required.
type BeneficiaryCertificationType9Choice struct {

	// Beneficial owner certification type expressed in a coded form.
	Code *BeneficiaryCertificationType4Code `xml:"Cd"`

	// Beneficial owner certification type expressed in a proprietary form.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (b *BeneficiaryCertificationType9Choice) SetCode(value string) {
	b.Code = (*BeneficiaryCertificationType4Code)(&value)
}

func (b *BeneficiaryCertificationType9Choice) AddProprietary() *GenericIdentification30 {
	b.Proprietary = new(GenericIdentification30)
	return b.Proprietary
}
