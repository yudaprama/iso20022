package model

// Choice between a standard code or proprietary code to specify the type of beneficiary certification required.
type BeneficiaryCertificationType10Choice struct {

	// Standard code to specify the type of certification required.
	Code *BeneficiaryCertificationType5Code `xml:"Cd"`

	// Proprietary identification of the type of certification required.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (b *BeneficiaryCertificationType10Choice) SetCode(value string) {
	b.Code = (*BeneficiaryCertificationType5Code)(&value)
}

func (b *BeneficiaryCertificationType10Choice) AddProprietary() *GenericIdentification30 {
	b.Proprietary = new(GenericIdentification30)
	return b.Proprietary
}
