package model

// Indicates the details of a guarantee.
type GuaranteeDetails1 struct {

	// Party issuing the guarantee.
	Issuer *QualifiedPartyIdentification1 `xml:"Issr,omitempty"`

	// Rank of the guarantee provider. A value of 1 means highest rank. Providers may have the same position.
	Position *positiveInteger `xml:"Pos,omitempty"`

	// Textual description of guarantee details.
	Description *Max2000Text `xml:"Desc,omitempty"`

	// Amount by time periods, maximum value applies at any given date.
	GuaranteedAmount []*AmountAndPeriod1 `xml:"GrntedAmt,omitempty"`

	// Amount not covered by the guarantee. Maximum value applies at any given date.
	Excess []*AmountAndPeriod1 `xml:"Xcss,omitempty"`

	// Covered percentage, the maximum value applies at any given date.
	CoveredPercentage []*PercentageAndPeriod1 `xml:"CvrdPctg,omitempty"`

	// Associated free form document.
	AssociatedDocument []*QualifiedDocumentInformation1 `xml:"AssoctdDoc,omitempty"`

	// Additional information related to the demand.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (g *GuaranteeDetails1) AddIssuer() *QualifiedPartyIdentification1 {
	g.Issuer = new(QualifiedPartyIdentification1)
	return g.Issuer
}

func (g *GuaranteeDetails1) SetPosition(value string) {
	g.Position = (*positiveInteger)(&value)
}

func (g *GuaranteeDetails1) SetDescription(value string) {
	g.Description = (*Max2000Text)(&value)
}

func (g *GuaranteeDetails1) AddGuaranteedAmount() *AmountAndPeriod1 {
	newValue := new(AmountAndPeriod1)
	g.GuaranteedAmount = append(g.GuaranteedAmount, newValue)
	return newValue
}

func (g *GuaranteeDetails1) AddExcess() *AmountAndPeriod1 {
	newValue := new(AmountAndPeriod1)
	g.Excess = append(g.Excess, newValue)
	return newValue
}

func (g *GuaranteeDetails1) AddCoveredPercentage() *PercentageAndPeriod1 {
	newValue := new(PercentageAndPeriod1)
	g.CoveredPercentage = append(g.CoveredPercentage, newValue)
	return newValue
}

func (g *GuaranteeDetails1) AddAssociatedDocument() *QualifiedDocumentInformation1 {
	newValue := new(QualifiedDocumentInformation1)
	g.AssociatedDocument = append(g.AssociatedDocument, newValue)
	return newValue
}

func (g *GuaranteeDetails1) AddAdditionalInformation(value string) {
	g.AdditionalInformation = append(g.AdditionalInformation, (*Max2000Text)(&value))
}
