package model

// Specifies the details relative to the submission of the certificate data set.
type RequiredSubmission4 struct {

	// Specifies with party(ies) is authorised to submit the data set as part of the transaction.
	Submitter []*BICIdentification1 `xml:"Submitr"`

	// Specifies the type of the certificate.
	CertificateType *TradeCertificateType1Code `xml:"CertTp"`

	// Specifies if the issuer must be matched as part of the validation of the data set.
	MatchIssuer *PartyIdentification27 `xml:"MtchIssr,omitempty"`

	// Specifies if the issue date must be matched as part of the validation of the data set.
	MatchIssueDate *YesNoIndicator `xml:"MtchIsseDt"`

	// Specifies if the inspection date must be matched as part of the validation of the data set.
	MatchInspectionDate *YesNoIndicator `xml:"MtchInspctnDt"`

	// Specifies if the indication of an authorised inspector must be present as part of the validation of the data set.
	AuthorisedInspectorIndicator *YesNoIndicator `xml:"AuthrsdInspctrInd"`

	// Specifies if the consignee must be matched as part of the validation of the data set.
	MatchConsignee *YesNoIndicator `xml:"MtchConsgn"`

	// Specifies if the manufacturer must be matched as part of the validation of the data set.
	MatchManufacturer *PartyIdentification27 `xml:"MtchManfctr,omitempty"`

	// Specifies if the certificate data set is required in relation to specific line items, and which line items.
	LineItemIdentification []*Max70Text `xml:"LineItmId,omitempty"`
}

func (r *RequiredSubmission4) AddSubmitter() *BICIdentification1 {
	newValue := new(BICIdentification1)
	r.Submitter = append(r.Submitter, newValue)
	return newValue
}

func (r *RequiredSubmission4) SetCertificateType(value string) {
	r.CertificateType = (*TradeCertificateType1Code)(&value)
}

func (r *RequiredSubmission4) AddMatchIssuer() *PartyIdentification27 {
	r.MatchIssuer = new(PartyIdentification27)
	return r.MatchIssuer
}

func (r *RequiredSubmission4) SetMatchIssueDate(value string) {
	r.MatchIssueDate = (*YesNoIndicator)(&value)
}

func (r *RequiredSubmission4) SetMatchInspectionDate(value string) {
	r.MatchInspectionDate = (*YesNoIndicator)(&value)
}

func (r *RequiredSubmission4) SetAuthorisedInspectorIndicator(value string) {
	r.AuthorisedInspectorIndicator = (*YesNoIndicator)(&value)
}

func (r *RequiredSubmission4) SetMatchConsignee(value string) {
	r.MatchConsignee = (*YesNoIndicator)(&value)
}

func (r *RequiredSubmission4) AddMatchManufacturer() *PartyIdentification27 {
	r.MatchManufacturer = new(PartyIdentification27)
	return r.MatchManufacturer
}

func (r *RequiredSubmission4) AddLineItemIdentification(value string) {
	r.LineItemIdentification = append(r.LineItemIdentification, (*Max70Text)(&value))
}
