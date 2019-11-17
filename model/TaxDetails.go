package model

// Set of characteristics defining the type of tax.
type TaxDetails struct {

	// Document issued by first agent on behalf of debtor to report withholding tax to taxing authority.
	CertificateIdentification *Max35Text `xml:"CertId,omitempty"`

	// Information on the type of tax.
	TaxType *TaxType `xml:"TaxTp,omitempty"`
}

func (t *TaxDetails) SetCertificateIdentification(value string) {
	t.CertificateIdentification = (*Max35Text)(&value)
}

func (t *TaxDetails) AddTaxType() *TaxType {
	t.TaxType = new(TaxType)
	return t.TaxType
}
