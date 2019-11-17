package model

// Specifies the details relative to the submission of the certificate data set.
type RequiredSubmission6 struct {

	// Specifies with party(ies) is authorised to submit the data set as part of the transaction.
	Submitter []*BICIdentification1 `xml:"Submitr"`

	// Specifies the type of the certificate, in 4 letters, for example BENE for beneficiary certificate, SHIP for shipping line certifcate.
	CertificateType *Exact4AlphaNumericText `xml:"CertTp"`

	// Description of the certificate type required.
	CertificateTypeDescription *Max140Text `xml:"CertTpDesc"`
}

func (r *RequiredSubmission6) AddSubmitter() *BICIdentification1 {
	newValue := new(BICIdentification1)
	r.Submitter = append(r.Submitter, newValue)
	return newValue
}

func (r *RequiredSubmission6) SetCertificateType(value string) {
	r.CertificateType = (*Exact4AlphaNumericText)(&value)
}

func (r *RequiredSubmission6) SetCertificateTypeDescription(value string) {
	r.CertificateTypeDescription = (*Max140Text)(&value)
}
