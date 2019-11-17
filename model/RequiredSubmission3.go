package model

// Specifies the details relative to the submission of the insurance data set.
type RequiredSubmission3 struct {

	// Specifies with party(ies) is authorised to submit the data set as part of the transaction.
	Submitter []*BICIdentification1 `xml:"Submitr"`

	// Specifies if the issuer must be matched as part of the validation of the data set.
	MatchIssuer *PartyIdentification27 `xml:"MtchIssr,omitempty"`

	// Specifies if the issue date must be matched as part of the validation of the data set.
	MatchIssueDate *YesNoIndicator `xml:"MtchIsseDt"`

	// Specifies if the transport information must be matched as part of the validation of the data set.
	MatchTransport *YesNoIndicator `xml:"MtchTrnsprt"`

	// Specifies if the insured amount must be matched as part of the validation of the data set.
	MatchAmount *YesNoIndicator `xml:"MtchAmt"`

	// Specifies which clauses are required in the insurance data set.
	ClausesRequired []*InsuranceClauses1Code `xml:"ClausesReqrd,omitempty"`

	// Specifies if the assured (insured) party must be matched as part of the validation of the data set.
	MatchAssuredParty *AssuredType1Code `xml:"MtchAssrdPty,omitempty"`
}

func (r *RequiredSubmission3) AddSubmitter() *BICIdentification1 {
	newValue := new(BICIdentification1)
	r.Submitter = append(r.Submitter, newValue)
	return newValue
}

func (r *RequiredSubmission3) AddMatchIssuer() *PartyIdentification27 {
	r.MatchIssuer = new(PartyIdentification27)
	return r.MatchIssuer
}

func (r *RequiredSubmission3) SetMatchIssueDate(value string) {
	r.MatchIssueDate = (*YesNoIndicator)(&value)
}

func (r *RequiredSubmission3) SetMatchTransport(value string) {
	r.MatchTransport = (*YesNoIndicator)(&value)
}

func (r *RequiredSubmission3) SetMatchAmount(value string) {
	r.MatchAmount = (*YesNoIndicator)(&value)
}

func (r *RequiredSubmission3) AddClausesRequired(value string) {
	r.ClausesRequired = append(r.ClausesRequired, (*InsuranceClauses1Code)(&value))
}

func (r *RequiredSubmission3) SetMatchAssuredParty(value string) {
	r.MatchAssuredParty = (*AssuredType1Code)(&value)
}
