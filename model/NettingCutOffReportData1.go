package model

// Specifies the meta data associated with a netting cut off report.
type NettingCutOffReportData1 struct {

	// Unique and unambiguous identifier for a message, as assigned by the Sender. This unique identifier can be used for cross-referencing purposes in subsequent messages.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the net report was generated.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Describes the type of net report, indicating how the obligations have been calculated.
	ReportType *Max4Text `xml:"RptTp"`

	// Date on which the proposed netting cut off will become active.
	ActivationDate *ISODate `xml:"ActvtnDt"`

	// Describes the participant receiving the net report.
	NetServiceParticipantIdentification *PartyIdentification73Choice `xml:"NetSvcPtcptId,omitempty"`

	// Describes the central system responsible for generating the net report.
	ReportServicer *PartyIdentification73Choice `xml:"RptSvcr,omitempty"`

	// Describes the type of netting service supporting the net report.
	NetServiceType *Max35Text `xml:"NetSvcTp,omitempty"`

	// Page number of the message (within the net cut off report) and continuation indicator to indicate that the report is to continue or that the message is the last page of the report.
	MessagePagination *Pagination `xml:"MsgPgntn,omitempty"`
}

func (n *NettingCutOffReportData1) SetMessageIdentification(value string) {
	n.MessageIdentification = (*Max35Text)(&value)
}

func (n *NettingCutOffReportData1) SetCreationDateTime(value string) {
	n.CreationDateTime = (*ISODateTime)(&value)
}

func (n *NettingCutOffReportData1) SetReportType(value string) {
	n.ReportType = (*Max4Text)(&value)
}

func (n *NettingCutOffReportData1) SetActivationDate(value string) {
	n.ActivationDate = (*ISODate)(&value)
}

func (n *NettingCutOffReportData1) AddNetServiceParticipantIdentification() *PartyIdentification73Choice {
	n.NetServiceParticipantIdentification = new(PartyIdentification73Choice)
	return n.NetServiceParticipantIdentification
}

func (n *NettingCutOffReportData1) AddReportServicer() *PartyIdentification73Choice {
	n.ReportServicer = new(PartyIdentification73Choice)
	return n.ReportServicer
}

func (n *NettingCutOffReportData1) SetNetServiceType(value string) {
	n.NetServiceType = (*Max35Text)(&value)
}

func (n *NettingCutOffReportData1) AddMessagePagination() *Pagination {
	n.MessagePagination = new(Pagination)
	return n.MessagePagination
}
