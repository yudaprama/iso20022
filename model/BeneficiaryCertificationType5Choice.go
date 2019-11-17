package model

// Choice between a standard code or proprietary code to specify the type of beneficiary certification required.
type BeneficiaryCertificationType5Choice struct {

	// Standard code to specify the type of certification required.
	Code *BeneficiaryCertificationType4Code `xml:"Cd"`

	// Proprietary identification of the type of certification required.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (b *BeneficiaryCertificationType5Choice) SetCode(value string) {
	b.Code = (*BeneficiaryCertificationType4Code)(&value)
}

func (b *BeneficiaryCertificationType5Choice) AddProprietary() *GenericIdentification20 {
	b.Proprietary = new(GenericIdentification20)
	return b.Proprietary
}
