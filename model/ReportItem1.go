package model

// Account details of the report item.
type ReportItem1 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *SecuritiesAccount19 `xml:"AcctId"`

	// Level of the safekeeping account or sub-account of the report item.
	AccountLevel *HoldingAccountLevel1Code `xml:"AcctLvl"`

	// Financial instrument identification of the report item.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId,omitempty"`

	// Date of the report item.
	ItemDate *ISODate `xml:"ItmDt,omitempty"`
}

func (r *ReportItem1) AddAccountIdentification() *SecuritiesAccount19 {
	r.AccountIdentification = new(SecuritiesAccount19)
	return r.AccountIdentification
}

func (r *ReportItem1) SetAccountLevel(value string) {
	r.AccountLevel = (*HoldingAccountLevel1Code)(&value)
}

func (r *ReportItem1) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	r.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return r.FinancialInstrumentIdentification
}

func (r *ReportItem1) SetItemDate(value string) {
	r.ItemDate = (*ISODate)(&value)
}
