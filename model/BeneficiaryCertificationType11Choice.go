package model

// Choice between a standard code or proprietary code to specify the type of beneficiary certification required.
type BeneficiaryCertificationType11Choice struct {

	// Standard code to specify the type of certification required.
	Code *BeneficiaryCertificationType5Code `xml:"Cd"`

	// Proprietary identification of the type of certification required.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (b *BeneficiaryCertificationType11Choice) SetCode(value string) {
	b.Code = (*BeneficiaryCertificationType5Code)(&value)
}

func (b *BeneficiaryCertificationType11Choice) AddProprietary() *GenericIdentification47 {
	b.Proprietary = new(GenericIdentification47)
	return b.Proprietary
}
