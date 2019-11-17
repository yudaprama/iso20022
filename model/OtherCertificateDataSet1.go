package model

// Formal document used to record a fact and used as proof of the fact, in the context of a commercial trade transaction.
type OtherCertificateDataSet1 struct {

	// Identifies the certificate data set.
	DataSetIdentification *DocumentIdentification1 `xml:"DataSetId"`

	// Unique identifier of the document.
	CertificateIdentification *Max35Text `xml:"CertId"`

	// Specifies the type of the certificate.
	CertificateType *TradeCertificateType2Code `xml:"CertTp"`

	// Issue date of the document.
	IssueDate *ISODate `xml:"IsseDt"`

	// Issuer of the certificate, typically the inspection company or its agent.
	Issuer *PartyIdentification26 `xml:"Issr"`

	// Additional and important information that could not be captured by structured fields.
	CertificateInformation []*Max350Text `xml:"CertInf,omitempty"`
}

func (o *OtherCertificateDataSet1) AddDataSetIdentification() *DocumentIdentification1 {
	o.DataSetIdentification = new(DocumentIdentification1)
	return o.DataSetIdentification
}

func (o *OtherCertificateDataSet1) SetCertificateIdentification(value string) {
	o.CertificateIdentification = (*Max35Text)(&value)
}

func (o *OtherCertificateDataSet1) SetCertificateType(value string) {
	o.CertificateType = (*TradeCertificateType2Code)(&value)
}

func (o *OtherCertificateDataSet1) SetIssueDate(value string) {
	o.IssueDate = (*ISODate)(&value)
}

func (o *OtherCertificateDataSet1) AddIssuer() *PartyIdentification26 {
	o.Issuer = new(PartyIdentification26)
	return o.Issuer
}

func (o *OtherCertificateDataSet1) AddCertificateInformation(value string) {
	o.CertificateInformation = append(o.CertificateInformation, (*Max350Text)(&value))
}
