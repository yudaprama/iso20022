package reda

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100104 struct {
	XMLName xml.Name        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.001.001.04 Document"`
	Message *PriceReportV04 `xml:"PricRpt"`
}

func (d *Document00100104) AddMessage() *PriceReportV04 {
	d.Message = new(PriceReportV04)
	return d.Message
}

// SCOPE
//
// A report provider, for example, a transfer agent, fund accountant or market data provider, sends the PriceReport message to the report recipient, for example, a fund management company, transfer agent, market data provider, regulator or any other interested party to provide the net asset value and price information for financial instruments on specific trade dates and, optionally, to quote price variation information.
//
// USAGE
//
// The PriceReport message is sent by the report provider to the report recipient to:
// - report prices for one or several different financial instruments for one or several different trade dates,
// - report statistical information about the valuation of a financial instrument,
// - inform another party that the quotation of a financial instrument is suspended,
// - report prices that are used for other purposes than the execution of investment funds orders.
//
type PriceReportV04 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Pagination of the message.
	MessagePagination *model.Pagination `xml:"MsgPgntn"`

	// Unique and unambiguous identifier for the price report, as assigned by the reporting party.
	PriceReportIdentification *model.Max35Text `xml:"PricRptId"`

	// Function of the price report, that is, whether the price report is a new price report or a replacement of some kind.
	Function *model.PriceReportFunction1Code `xml:"Fctn"`

	// Unique and unambiguous identifier for the cancellation of the previous price report, as assigned by the reporting party.
	CancellationIdentification *model.Max35Text `xml:"CxlId,omitempty"`

	// Information related to the price valuation of a financial instrument.
	PriceValuationDetails []*model.PriceValuation4 `xml:"PricValtnDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (p *PriceReportV04) AddMessageIdentification() *model.MessageIdentification1 {
	p.MessageIdentification = new(model.MessageIdentification1)
	return p.MessageIdentification
}

func (p *PriceReportV04) AddPoolReference() *model.AdditionalReference3 {
	p.PoolReference = new(model.AdditionalReference3)
	return p.PoolReference
}

func (p *PriceReportV04) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	p.PreviousReference = append(p.PreviousReference, newValue)
	return newValue
}

func (p *PriceReportV04) AddRelatedReference() *model.AdditionalReference3 {
	p.RelatedReference = new(model.AdditionalReference3)
	return p.RelatedReference
}

func (p *PriceReportV04) AddMessagePagination() *model.Pagination {
	p.MessagePagination = new(model.Pagination)
	return p.MessagePagination
}

func (p *PriceReportV04) SetPriceReportIdentification(value string) {
	p.PriceReportIdentification = (*model.Max35Text)(&value)
}

func (p *PriceReportV04) SetFunction(value string) {
	p.Function = (*model.PriceReportFunction1Code)(&value)
}

func (p *PriceReportV04) SetCancellationIdentification(value string) {
	p.CancellationIdentification = (*model.Max35Text)(&value)
}

func (p *PriceReportV04) AddPriceValuationDetails() *model.PriceValuation4 {
	newValue := new(model.PriceValuation4)
	p.PriceValuationDetails = append(p.PriceValuationDetails, newValue)
	return newValue
}

func (p *PriceReportV04) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	p.Extension = append(p.Extension, newValue)
	return newValue
}
