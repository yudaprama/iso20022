package model

// Specifies the details relative to the submission of the certificate data set.
type RequiredSubmission5 struct {

	// Specifies with party(ies) is authorised to submit the data set as part of the transaction.
	Submitter []*BICIdentification1 `xml:"Submitr"`

	// Specifies the type of the certificate.
	CertificateType *TradeCertificateType2Code `xml:"CertTp"`
}

func (r *RequiredSubmission5) AddSubmitter() *BICIdentification1 {
	newValue := new(BICIdentification1)
	r.Submitter = append(r.Submitter, newValue)
	return newValue
}

func (r *RequiredSubmission5) SetCertificateType(value string) {
	r.CertificateType = (*TradeCertificateType2Code)(&value)
}
