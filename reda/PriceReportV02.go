package reda

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100102 struct {
	XMLName xml.Name        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.001.001.02 Document"`
	Message *PriceReportV02 `xml:"reda.001.001.02"`
}

func (d *Document00100102) AddMessage() *PriceReportV02 {
	d.Message = new(PriceReportV02)
	return d.Message
}

// Scope
// The PriceReport message is sent by a report provider, eg, a fund accountant, transfer agent, market data provider, or any other interested party, to a report user, eg, a fund management company, a transfer agent, market data provider, regulator or any other interested party.
// This message is used to provide net asset value and price information for financial instruments on given trade dates and, optionally, to quote price variation information.
// Usage
// The PriceReport message can be used to:
// - report prices for one or several different financial instruments for one or several different trade dates,
// - report statistical information about the valuation of a financial instrument,
// - inform another party that the quotation of a financial instrument is suspended,
// - report prices that are used for other purposes than the execution of investment funds orders.
type PriceReportV02 struct {

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Information related to the price valuation of a financial instrument.
	PriceValuationDetails []*model.PriceValuation2 `xml:"PricValtnDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (p *PriceReportV02) AddPoolReference() *model.AdditionalReference3 {
	p.PoolReference = new(model.AdditionalReference3)
	return p.PoolReference
}

func (p *PriceReportV02) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	p.PreviousReference = append(p.PreviousReference, newValue)
	return newValue
}

func (p *PriceReportV02) AddRelatedReference() *model.AdditionalReference3 {
	p.RelatedReference = new(model.AdditionalReference3)
	return p.RelatedReference
}

func (p *PriceReportV02) AddPriceValuationDetails() *model.PriceValuation2 {
	newValue := new(model.PriceValuation2)
	p.PriceValuationDetails = append(p.PriceValuationDetails, newValue)
	return newValue
}

func (p *PriceReportV02) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	p.Extension = append(p.Extension, newValue)
	return newValue
}
