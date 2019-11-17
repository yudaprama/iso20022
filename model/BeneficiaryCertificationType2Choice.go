package model

// Choice between a standard code or proprietary code to specify the type of beneficiary certification required.
type BeneficiaryCertificationType2Choice struct {

	// Standard code to specify the type of certification required.
	Code *BeneficiaryCertificationType3Code `xml:"Cd"`

	// Proprietary identification of the type of certification required.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (b *BeneficiaryCertificationType2Choice) SetCode(value string) {
	b.Code = (*BeneficiaryCertificationType3Code)(&value)
}

func (b *BeneficiaryCertificationType2Choice) AddProprietary() *GenericIdentification20 {
	b.Proprietary = new(GenericIdentification20)
	return b.Proprietary
}
