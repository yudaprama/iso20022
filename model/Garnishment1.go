package model

// Provides remittance information about a payment for garnishment-related purposes.
type Garnishment1 struct {

	// Specifies the type of garnishment.
	Type *GarnishmentType1 `xml:"Tp"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor, in this case, to the garnisher.
	Garnishee *PartyIdentification43 `xml:"Grnshee,omitempty"`

	// Party on the credit side of the transaction who administers the garnishment on behalf of the ultimate beneficiary.
	GarnishmentAdministrator *PartyIdentification43 `xml:"GrnshmtAdmstr,omitempty"`

	// Reference information that is specific to the agency receiving the garnishment.
	ReferenceNumber *Max140Text `xml:"RefNb,omitempty"`

	// Date of payment which garnishment was taken from.
	Date *ISODate `xml:"Dt,omitempty"`

	// Amount of money remitted for the referred document.
	RemittedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty"`

	// Indicates if the person to whom the garnishment applies (that is, the ultimate debtor) has family medical insurance coverage available.
	FamilyMedicalInsuranceIndicator *TrueFalseIndicator `xml:"FmlyMdclInsrncInd,omitempty"`

	// Indicates if the employment of the person to whom the garnishment applies (that is, the ultimate debtor) has been terminated.
	EmployeeTerminationIndicator *TrueFalseIndicator `xml:"MplyeeTermntnInd,omitempty"`
}

func (g *Garnishment1) AddType() *GarnishmentType1 {
	g.Type = new(GarnishmentType1)
	return g.Type
}

func (g *Garnishment1) AddGarnishee() *PartyIdentification43 {
	g.Garnishee = new(PartyIdentification43)
	return g.Garnishee
}

func (g *Garnishment1) AddGarnishmentAdministrator() *PartyIdentification43 {
	g.GarnishmentAdministrator = new(PartyIdentification43)
	return g.GarnishmentAdministrator
}

func (g *Garnishment1) SetReferenceNumber(value string) {
	g.ReferenceNumber = (*Max140Text)(&value)
}

func (g *Garnishment1) SetDate(value string) {
	g.Date = (*ISODate)(&value)
}

func (g *Garnishment1) SetRemittedAmount(value, currency string) {
	g.RemittedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (g *Garnishment1) SetFamilyMedicalInsuranceIndicator(value string) {
	g.FamilyMedicalInsuranceIndicator = (*TrueFalseIndicator)(&value)
}

func (g *Garnishment1) SetEmployeeTerminationIndicator(value string) {
	g.EmployeeTerminationIndicator = (*TrueFalseIndicator)(&value)
}
