package model

// Choice between a standard code or proprietary code to specify the type of beneficiary certification required.
type BeneficiaryCertificationType12Choice struct {

	// Beneficial owner certification type expressed in a coded form.
	Code *BeneficiaryCertificationType4Code `xml:"Cd"`

	// Beneficial owner certification type expressed in a proprietary form.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (b *BeneficiaryCertificationType12Choice) SetCode(value string) {
	b.Code = (*BeneficiaryCertificationType4Code)(&value)
}

func (b *BeneficiaryCertificationType12Choice) AddProprietary() *GenericIdentification47 {
	b.Proprietary = new(GenericIdentification47)
	return b.Proprietary
}
