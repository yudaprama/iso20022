package model

// Specifies the meta data associated with a net report.
type NetReportData1 struct {

	// Unique and unambiguous identifier for a message, as assigned by the Sender. This unique identifier can be used for cross-referencing purposes in subsequent messages.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the net report was generated.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Provides the cut off time that has been passed, resulting in the generation of the net report.
	NettingCutOffTime *ISOTime `xml:"NetgCutOffTm"`

	// Specifies the value date on which the net report was generated.
	ReportDate *ISODate `xml:"RptDt"`

	// Specifies the value date for the trades used in the generation of the net report.
	ValueDate *ISODate `xml:"ValDt"`

	// Specifies the type of net report, indicating how the obligations have been calculated.
	ReportType *Max4Text `xml:"RptTp"`

	// Describes the central system responsible for generating the net report.
	NetReportServicer *PartyIdentification73Choice `xml:"NetRptSvcr,omitempty"`

	// Describes the type of netting service supporting the net report.
	NetServiceType *Max35Text `xml:"NetSvcTp,omitempty"`

	// Page number of the message (within the net report) and continuation indicator to indicate that the report is to continue or that the message is the last page of the report.
	MessagePagination *Pagination `xml:"MsgPgntn,omitempty"`
}

func (n *NetReportData1) SetMessageIdentification(value string) {
	n.MessageIdentification = (*Max35Text)(&value)
}

func (n *NetReportData1) SetCreationDateTime(value string) {
	n.CreationDateTime = (*ISODateTime)(&value)
}

func (n *NetReportData1) SetNettingCutOffTime(value string) {
	n.NettingCutOffTime = (*ISOTime)(&value)
}

func (n *NetReportData1) SetReportDate(value string) {
	n.ReportDate = (*ISODate)(&value)
}

func (n *NetReportData1) SetValueDate(value string) {
	n.ValueDate = (*ISODate)(&value)
}

func (n *NetReportData1) SetReportType(value string) {
	n.ReportType = (*Max4Text)(&value)
}

func (n *NetReportData1) AddNetReportServicer() *PartyIdentification73Choice {
	n.NetReportServicer = new(PartyIdentification73Choice)
	return n.NetReportServicer
}

func (n *NetReportData1) SetNetServiceType(value string) {
	n.NetServiceType = (*Max35Text)(&value)
}

func (n *NetReportData1) AddMessagePagination() *Pagination {
	n.MessagePagination = new(Pagination)
	return n.MessagePagination
}
