package model

// Formal document used to record a fact and used as proof of the fact that goods have been insured under an insurance policy.
type InsuranceDataSet1 struct {

	// Identifies the insurancedata set
	DataSetIdentification *DocumentIdentification1 `xml:"DataSetId"`

	// Issuer of the certificate, typically the insurance company or its agent.
	Issuer *PartyIdentification26 `xml:"Issr"`

	// Issue date of the document.
	IssueDate *ISODate `xml:"IsseDt"`

	// Date upon which cover under an insurance policy becomes effective.
	EffectiveDate *ISODate `xml:"FctvDt,omitempty"`

	// Place where the insurance certificate was issued.
	PlaceOfIssue *PostalAddress5 `xml:"PlcOfIsse,omitempty"`

	// Unique identifier of the document.
	InsuranceDocumentIdentification *Max35Text `xml:"InsrncDocId"`

	// Transport information relative to the goods that are insured under the insurance policy.
	Transport *SingleTransport3 `xml:"Trnsprt,omitempty"`

	// Value of the goods as insured under the insurance policy.
	InsuredAmount *CurrencyAndAmount `xml:"InsrdAmt"`

	// Information about the goods and/or services of a trade transaction.
	InsuredGoodsDescription *Max70Text `xml:"InsrdGoodsDesc,omitempty"`

	// Description of the conditions and exclusion clauses under which insurance is granted.
	InsuranceConditions []*Max350Text `xml:"InsrncConds,omitempty"`

	// Standard insurance clauses defined by the Institute of London Underwriters (or the American Institute of marine Underwriters).
	InsuranceClauses []*InsuranceClauses1Code `xml:"InsrncClauses,omitempty"`

	// Party that is covered under the assurance policy.
	Assured *PartyIdentification29Choice `xml:"Assrd"`

	// Place where claims under the insurance policy will be paid.
	ClaimsPayableAt *PostalAddress5 `xml:"ClmsPyblAt"`

	// Currency in which claims, if valid, will be paid.
	ClaimsPayableIn *CurrencyCode `xml:"ClmsPyblIn,omitempty"`
}

func (i *InsuranceDataSet1) AddDataSetIdentification() *DocumentIdentification1 {
	i.DataSetIdentification = new(DocumentIdentification1)
	return i.DataSetIdentification
}

func (i *InsuranceDataSet1) AddIssuer() *PartyIdentification26 {
	i.Issuer = new(PartyIdentification26)
	return i.Issuer
}

func (i *InsuranceDataSet1) SetIssueDate(value string) {
	i.IssueDate = (*ISODate)(&value)
}

func (i *InsuranceDataSet1) SetEffectiveDate(value string) {
	i.EffectiveDate = (*ISODate)(&value)
}

func (i *InsuranceDataSet1) AddPlaceOfIssue() *PostalAddress5 {
	i.PlaceOfIssue = new(PostalAddress5)
	return i.PlaceOfIssue
}

func (i *InsuranceDataSet1) SetInsuranceDocumentIdentification(value string) {
	i.InsuranceDocumentIdentification = (*Max35Text)(&value)
}

func (i *InsuranceDataSet1) AddTransport() *SingleTransport3 {
	i.Transport = new(SingleTransport3)
	return i.Transport
}

func (i *InsuranceDataSet1) SetInsuredAmount(value, currency string) {
	i.InsuredAmount = NewCurrencyAndAmount(value, currency)
}

func (i *InsuranceDataSet1) SetInsuredGoodsDescription(value string) {
	i.InsuredGoodsDescription = (*Max70Text)(&value)
}

func (i *InsuranceDataSet1) AddInsuranceConditions(value string) {
	i.InsuranceConditions = append(i.InsuranceConditions, (*Max350Text)(&value))
}

func (i *InsuranceDataSet1) AddInsuranceClauses(value string) {
	i.InsuranceClauses = append(i.InsuranceClauses, (*InsuranceClauses1Code)(&value))
}

func (i *InsuranceDataSet1) AddAssured() *PartyIdentification29Choice {
	i.Assured = new(PartyIdentification29Choice)
	return i.Assured
}

func (i *InsuranceDataSet1) AddClaimsPayableAt() *PostalAddress5 {
	i.ClaimsPayableAt = new(PostalAddress5)
	return i.ClaimsPayableAt
}

func (i *InsuranceDataSet1) SetClaimsPayableIn(value string) {
	i.ClaimsPayableIn = (*CurrencyCode)(&value)
}
