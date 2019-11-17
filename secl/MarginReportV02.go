package secl

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00500102 struct {
	XMLName xml.Name         `xml:"urn:iso:std:iso:20022:tech:xsd:secl.005.001.02 Document"`
	Message *MarginReportV02 `xml:"MrgnRpt"`
}

func (d *Document00500102) AddMessage() *MarginReportV02 {
	d.Message = new(MarginReportV02)
	return d.Message
}

// Scope
// The MarginReport message is sent by the central counterparty (CCP) to a clearing member to report on:
// - the exposure resulting from the trade positions
// - the value of the collateral held by the CCP (market value of this collateral) and
// - the resulting difference representing the risk encountered by the CCP.
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// There are four possibilities to report the above information. Indeed, the margin report may be structured as follows:
// - per clearing member: the report will only show the information for the clearing member, or
// - per clearing member and per financial instrument: the report will show the information for the clearing member, structured by security identification, or
// - per clearing member and per non clearing member: the report will show the information for the clearing member (that is for global clearing member only) structured by non clearing member(s), or
// - per clearing member and per non clearing member and per security identification: the report will show the information for the clearing member (global clearing member only) structured by non clearing member(s) and by security identification.
type MarginReportV02 struct {

	// Provides parameters of the margin report such as the creation date and time, the report currency or the calculation date and time.
	ReportParameters *model.ReportParameters3 `xml:"RptParams"`

	// Page number of the message (within a report) and continuation indicator to indicate that the report is to continue or that the message is the last page of the report.
	Pagination *model.Pagination `xml:"Pgntn"`

	// Provides the identification of the account owner, that is the clearing member (individual clearing member or general clearing member).
	ClearingMember *model.PartyIdentification35Choice `xml:"ClrMmb"`

	// Provides details on the valuation of the collateral on deposit.
	ReportSummary *model.MarginCalculation1 `xml:"RptSummry,omitempty"`

	// Provides the margin report details.
	ReportDetails []*model.MarginReport2 `xml:"RptDtls"`

	// Additional information that can't be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MarginReportV02) AddReportParameters() *model.ReportParameters3 {
	m.ReportParameters = new(model.ReportParameters3)
	return m.ReportParameters
}

func (m *MarginReportV02) AddPagination() *model.Pagination {
	m.Pagination = new(model.Pagination)
	return m.Pagination
}

func (m *MarginReportV02) AddClearingMember() *model.PartyIdentification35Choice {
	m.ClearingMember = new(model.PartyIdentification35Choice)
	return m.ClearingMember
}

func (m *MarginReportV02) AddReportSummary() *model.MarginCalculation1 {
	m.ReportSummary = new(model.MarginCalculation1)
	return m.ReportSummary
}

func (m *MarginReportV02) AddReportDetails() *model.MarginReport2 {
	newValue := new(model.MarginReport2)
	m.ReportDetails = append(m.ReportDetails, newValue)
	return newValue
}

func (m *MarginReportV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
