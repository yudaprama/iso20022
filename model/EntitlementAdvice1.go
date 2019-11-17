package model

// Provides information about the CA entitlement.
type EntitlementAdvice1 struct {

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption1FormatChoice `xml:"OptnTp"`

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Date on which the holders of securities are/will be recorded for the income being paid or for entitlement to the rights or offer/privilege.
	RecordDate *DateFormat4Choice `xml:"RcrdDt,omitempty"`

	// Date on which securities/cash will be paid.
	PaymentDate *DateFormat4Choice `xml:"PmtDt,omitempty"`

	// Provides information about the entitlement and the entitled account.
	AccountAndDistributionDetails []*Entitlement1 `xml:"AcctAndDstrbtnDtls"`
}

func (e *EntitlementAdvice1) AddOptionType() *CorporateActionOption1FormatChoice {
	e.OptionType = new(CorporateActionOption1FormatChoice)
	return e.OptionType
}

func (e *EntitlementAdvice1) SetOptionNumber(value string) {
	e.OptionNumber = (*Exact3NumericText)(&value)
}

func (e *EntitlementAdvice1) AddRecordDate() *DateFormat4Choice {
	e.RecordDate = new(DateFormat4Choice)
	return e.RecordDate
}

func (e *EntitlementAdvice1) AddPaymentDate() *DateFormat4Choice {
	e.PaymentDate = new(DateFormat4Choice)
	return e.PaymentDate
}

func (e *EntitlementAdvice1) AddAccountAndDistributionDetails() *Entitlement1 {
	newValue := new(Entitlement1)
	e.AccountAndDistributionDetails = append(e.AccountAndDistributionDetails, newValue)
	return newValue
}
