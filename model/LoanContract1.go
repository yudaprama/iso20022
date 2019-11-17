package model

// Contract by which an amount of money in exchange for future repayment of the principal amount along with interest or other finance charges.
type LoanContract1 struct {

	// Contract document referenced from this loan agreement.
	ContractDocumentIdentification *DocumentIdentification22 `xml:"CtrctDocId"`

	// Party that is specified as the buyer for this loan agreement.
	Buyer []*TradeParty2 `xml:"Buyr"`

	// Party that is specified as the seller for this loan agreement.
	Seller []*TradeParty2 `xml:"Sellr"`

	// Loan amount as defined in the contract.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *ISODate `xml:"MtrtyDt"`

	// Indicates whether the contract duration is extended or not.
	ProlongationFlag *TrueFalseIndicator `xml:"PrlngtnFlg"`

	// Start date of the loan contract.
	StartDate *ISODate `xml:"StartDt"`

	// Currency in which the loan is being settled.
	SettlementCurrency *ActiveCurrencyCode `xml:"SttlmCcy"`

	// When the amount is credited outside of the country, special conditions are applicable.
	SpecialConditions *SpecialCondition1 `xml:"SpclConds,omitempty"`

	// Loan duration in a coded form.
	DurationCode *Exact1NumericText `xml:"DrtnCd"`

	// Interest rate for the loan.
	InterestRate *InterestRate2Choice `xml:"IntrstRate"`

	// One part or division of the loan, used to define the repayment.
	Tranche []*LoanContractTranche1 `xml:"Trch,omitempty"`

	// Schedule of the payments defined for the loan contract.
	PaymentSchedule *PaymentSchedule1Choice `xml:"PmtSchdl,omitempty"`

	// Schedule of the interest payments defined for the loan contract.
	InterestSchedule *InterestPaymentSchedule1Choice `xml:"IntrstSchdl"`

	// Loan is an intra company loan.
	IntraCompanyLoan *TrueFalseIndicator `xml:"IntraCpnyLn"`

	// Details of the collateral for the loan.
	Collateral *ContractCollateral1 `xml:"Coll,omitempty"`

	// Loan offered by a group of lenders (called a syndicate) who work together to provide funds for a single borrower.
	SyndicatedLoan []*SyndicatedLoan1 `xml:"SndctdLn,omitempty"`

	// Documents provided as attachments to the loan contract.
	Attachment []*DocumentGeneralInformation3 `xml:"Attchmnt,omitempty"`
}

func (l *LoanContract1) AddContractDocumentIdentification() *DocumentIdentification22 {
	l.ContractDocumentIdentification = new(DocumentIdentification22)
	return l.ContractDocumentIdentification
}

func (l *LoanContract1) AddBuyer() *TradeParty2 {
	newValue := new(TradeParty2)
	l.Buyer = append(l.Buyer, newValue)
	return newValue
}

func (l *LoanContract1) AddSeller() *TradeParty2 {
	newValue := new(TradeParty2)
	l.Seller = append(l.Seller, newValue)
	return newValue
}

func (l *LoanContract1) SetAmount(value, currency string) {
	l.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (l *LoanContract1) SetMaturityDate(value string) {
	l.MaturityDate = (*ISODate)(&value)
}

func (l *LoanContract1) SetProlongationFlag(value string) {
	l.ProlongationFlag = (*TrueFalseIndicator)(&value)
}

func (l *LoanContract1) SetStartDate(value string) {
	l.StartDate = (*ISODate)(&value)
}

func (l *LoanContract1) SetSettlementCurrency(value string) {
	l.SettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (l *LoanContract1) AddSpecialConditions() *SpecialCondition1 {
	l.SpecialConditions = new(SpecialCondition1)
	return l.SpecialConditions
}

func (l *LoanContract1) SetDurationCode(value string) {
	l.DurationCode = (*Exact1NumericText)(&value)
}

func (l *LoanContract1) AddInterestRate() *InterestRate2Choice {
	l.InterestRate = new(InterestRate2Choice)
	return l.InterestRate
}

func (l *LoanContract1) AddTranche() *LoanContractTranche1 {
	newValue := new(LoanContractTranche1)
	l.Tranche = append(l.Tranche, newValue)
	return newValue
}

func (l *LoanContract1) AddPaymentSchedule() *PaymentSchedule1Choice {
	l.PaymentSchedule = new(PaymentSchedule1Choice)
	return l.PaymentSchedule
}

func (l *LoanContract1) AddInterestSchedule() *InterestPaymentSchedule1Choice {
	l.InterestSchedule = new(InterestPaymentSchedule1Choice)
	return l.InterestSchedule
}

func (l *LoanContract1) SetIntraCompanyLoan(value string) {
	l.IntraCompanyLoan = (*TrueFalseIndicator)(&value)
}

func (l *LoanContract1) AddCollateral() *ContractCollateral1 {
	l.Collateral = new(ContractCollateral1)
	return l.Collateral
}

func (l *LoanContract1) AddSyndicatedLoan() *SyndicatedLoan1 {
	newValue := new(SyndicatedLoan1)
	l.SyndicatedLoan = append(l.SyndicatedLoan, newValue)
	return newValue
}

func (l *LoanContract1) AddAttachment() *DocumentGeneralInformation3 {
	newValue := new(DocumentGeneralInformation3)
	l.Attachment = append(l.Attachment, newValue)
	return newValue
}
