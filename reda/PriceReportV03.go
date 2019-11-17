package reda

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100103 struct {
	XMLName xml.Name        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.001.001.03 Document"`
	Message *PriceReportV03 `xml:"PricRptV03"`
}

func (d *Document00100103) AddMessage() *PriceReportV03 {
	d.Message = new(PriceReportV03)
	return d.Message
}

// Scope
// A report provider, eg, a transfer agent, fund accountant or market data provider, sends the PriceReport message to the report recipient, eg, a fund management company, transfer agent, market data provider, regulator or other interested party to provide the net asset value and price information for financial instruments on specific trade dates and, optionally, to quote price variation information.
// Usage
// The PriceReport message is used to:
// - report prices for one or several different financial instruments for one or several different trade dates,
// - report statistical information about the valuation of a financial instrument,
// - inform another party that the quotation of a financial instrument is suspended,
// - report prices that are used for purposes other than the execution of investment funds orders.
type PriceReportV03 struct {

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

	// Information related to the price valuation of a financial instrument.
	PriceValuationDetails []*model.PriceValuation3 `xml:"PricValtnDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (p *PriceReportV03) AddMessageIdentification() *model.MessageIdentification1 {
	p.MessageIdentification = new(model.MessageIdentification1)
	return p.MessageIdentification
}

func (p *PriceReportV03) AddPoolReference() *model.AdditionalReference3 {
	p.PoolReference = new(model.AdditionalReference3)
	return p.PoolReference
}

func (p *PriceReportV03) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	p.PreviousReference = append(p.PreviousReference, newValue)
	return newValue
}

func (p *PriceReportV03) AddRelatedReference() *model.AdditionalReference3 {
	p.RelatedReference = new(model.AdditionalReference3)
	return p.RelatedReference
}

func (p *PriceReportV03) AddMessagePagination() *model.Pagination {
	p.MessagePagination = new(model.Pagination)
	return p.MessagePagination
}

func (p *PriceReportV03) AddPriceValuationDetails() *model.PriceValuation3 {
	newValue := new(model.PriceValuation3)
	p.PriceValuationDetails = append(p.PriceValuationDetails, newValue)
	return newValue
}

func (p *PriceReportV03) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	p.Extension = append(p.Extension, newValue)
	return newValue
}
