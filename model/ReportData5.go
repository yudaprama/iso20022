package model

// Numerical representation of the nett increases and decreases in an account at a specific point in time. A cash balance is calculated from a sum of cash credits minus a sum of cash debits.
type ReportData5 struct {

	// Identification of the report assigned by the central system.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date by which the amount(s) requested must be settled.
	ValueDate *ISODate `xml:"ValDt"`

	// Date and time on which the report is generated. The offset with UTC may also be specified.
	DateAndTimeStamp *ISODateTime `xml:"DtAndTmStmp"`

	// Specifies the type of the Pay In Call.
	Type *CallIn1Code `xml:"Tp"`

	// Specifies the amount requested.
	PayInCallAmount []*PayInCallItem `xml:"PayInCallAmt,omitempty"`

	// To indicate the requested CLS Settlement Session that the related trade is part of.
	SettlementSessionIdentifier *Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// Specifies the requested amount in multiple currencies.
	AccountValue *Value `xml:"AcctVal,omitempty"`
}

func (r *ReportData5) SetMessageIdentification(value string) {
	r.MessageIdentification = (*Max35Text)(&value)
}

func (r *ReportData5) SetValueDate(value string) {
	r.ValueDate = (*ISODate)(&value)
}

func (r *ReportData5) SetDateAndTimeStamp(value string) {
	r.DateAndTimeStamp = (*ISODateTime)(&value)
}

func (r *ReportData5) SetType(value string) {
	r.Type = (*CallIn1Code)(&value)
}

func (r *ReportData5) AddPayInCallAmount() *PayInCallItem {
	newValue := new(PayInCallItem)
	r.PayInCallAmount = append(r.PayInCallAmount, newValue)
	return newValue
}

func (r *ReportData5) SetSettlementSessionIdentifier(value string) {
	r.SettlementSessionIdentifier = (*Exact4AlphaNumericText)(&value)
}

func (r *ReportData5) AddAccountValue() *Value {
	r.AccountValue = new(Value)
	return r.AccountValue
}
