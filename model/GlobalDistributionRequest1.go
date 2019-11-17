package model

// Provides information about the global distribution.
type GlobalDistributionRequest1 struct {

	// Indicates wether is message is an advice or pre-advice.
	PreadviceIndicator *YesNoIndicator `xml:"PradvcInd"`

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption1FormatChoice `xml:"OptnTp"`

	// Date on which the holders of securities are/will be recorded for the income being paid or for entitlement to the rights or offer/privilege.
	RecordDate *DateFormat4Choice `xml:"RcrdDt"`

	// Date on which securities/cash will be paid.
	PaymentDate *DateFormat4Choice `xml:"PmtDt"`

	// Provides information about the securities movement.
	SecuritiesMovement []*SecurityMovement1 `xml:"SctiesMvmnt,omitempty"`

	// Provides information about the cash movement.
	CashMovement []*CashMovement1 `xml:"CshMvmnt,omitempty"`
}

func (g *GlobalDistributionRequest1) SetPreadviceIndicator(value string) {
	g.PreadviceIndicator = (*YesNoIndicator)(&value)
}

func (g *GlobalDistributionRequest1) SetOptionNumber(value string) {
	g.OptionNumber = (*Exact3NumericText)(&value)
}

func (g *GlobalDistributionRequest1) AddOptionType() *CorporateActionOption1FormatChoice {
	g.OptionType = new(CorporateActionOption1FormatChoice)
	return g.OptionType
}

func (g *GlobalDistributionRequest1) AddRecordDate() *DateFormat4Choice {
	g.RecordDate = new(DateFormat4Choice)
	return g.RecordDate
}

func (g *GlobalDistributionRequest1) AddPaymentDate() *DateFormat4Choice {
	g.PaymentDate = new(DateFormat4Choice)
	return g.PaymentDate
}

func (g *GlobalDistributionRequest1) AddSecuritiesMovement() *SecurityMovement1 {
	newValue := new(SecurityMovement1)
	g.SecuritiesMovement = append(g.SecuritiesMovement, newValue)
	return newValue
}

func (g *GlobalDistributionRequest1) AddCashMovement() *CashMovement1 {
	newValue := new(CashMovement1)
	g.CashMovement = append(g.CashMovement, newValue)
	return newValue
}
