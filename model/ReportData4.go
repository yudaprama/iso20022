package model

// Numerical representation of the nett increases and decreases in an account at a specific point in time. A cash balance is calculated from a sum of cash credits minus a sum of cash debits.
type ReportData4 struct {

	// Identification of the report as assigned by the sender.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Value date for which the pay-in schedule is generated.
	ValueDate *ISODate `xml:"ValDt"`

	// Date and time on which the report is generated. The offset with UTC may also be specified.
	DateAndTimeStamp *ISODateTime `xml:"DtAndTmStmp"`

	// Type of pay-in schedule.
	Type *Entry2Code `xml:"Tp"`

	// Defines the  schedule timing that is, whether it is an initial or a revised schedule.
	ScheduleType *Exact4AlphaNumericText `xml:"SchdlTp"`

	// To indicate the requested CLS Settlement Session that the related trade is part of.
	SettlementSessionIdentifier *Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`
}

func (r *ReportData4) SetMessageIdentification(value string) {
	r.MessageIdentification = (*Max35Text)(&value)
}

func (r *ReportData4) SetValueDate(value string) {
	r.ValueDate = (*ISODate)(&value)
}

func (r *ReportData4) SetDateAndTimeStamp(value string) {
	r.DateAndTimeStamp = (*ISODateTime)(&value)
}

func (r *ReportData4) SetType(value string) {
	r.Type = (*Entry2Code)(&value)
}

func (r *ReportData4) SetScheduleType(value string) {
	r.ScheduleType = (*Exact4AlphaNumericText)(&value)
}

func (r *ReportData4) SetSettlementSessionIdentifier(value string) {
	r.SettlementSessionIdentifier = (*Exact4AlphaNumericText)(&value)
}
