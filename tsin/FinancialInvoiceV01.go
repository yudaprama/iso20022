package tsin

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400101 struct {
	XMLName xml.Name             `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.004.001.01 Document"`
	Message *FinancialInvoiceV01 `xml:"FinInvc"`
}

func (d *Document00400101) AddMessage() *FinancialInvoiceV01 {
	d.Message = new(FinancialInvoiceV01)
	return d.Message
}

// Scope
// The FinancialInvoice message is used to support the provision of financial and related services where there is a requirement to exchange invoice information.
// Usage
// While the prime function of the FinancialInvoice message is as a request from the seller to the buyer for payment, the FinancialInvoice message can also serve to evidence an invoice in support of a financial service such as invoice factoring, letters of credit, and bank payment obligations, to enable Web based services such as electronic bill payment and presentment, and as the basis to transfer invoice information via third parties such as e-invoicing service providers.
// A consequence of the receipt of an invoice by the buyer is that it acts as a trigger for the use of related messages that are already defined in ISO 20022, notably where the information contained in the Financial Invoice enables payment for the goods or services received, and/or is provided in support of a request for invoice financing. While certain of these related messages, such as the CreditTransfer and PaymentInitiation messages, are shown in the sequence diagram they are out of scope. They are shown only to illustrate a given scenario and to place the invoice in the context of the financial banking processes that might be conducted between different financial institutions.
// The use of self-billing by the buyer to the seller, where the buyer acts as the invoice issuer or the process of handling an incorrect invoice, is not in scope.
type FinancialInvoiceV01 struct {

	// Collection of data that is exchanged between two or more parties in written, printed or electronic form. It contains general data relevant to the main body of the invoice such as date of issue, currency code and identification number.
	InvoiceHeader *model.InvoiceHeader1 `xml:"InvcHdr"`

	// Commercial information such as terms of commerce, parties, and documentation, related to the trading agreement under which this invoice is issued.
	TradeAgreement *model.TradeAgreement6 `xml:"TradAgrmt"`

	// Supply chain shipping arrangements for delivery of invoiced products and/or services.
	TradeDelivery *model.TradeDelivery1 `xml:"TradDlvry"`

	// Settlement information that enables the financial reconciliation and payment of this invoice.
	//
	TradeSettlement *model.TradeSettlement1 `xml:"TradSttlm"`

	// Unit of information in this invoice showning the related  provision of products and/or services and monetary summations reported as a discrete line item.
	//
	//
	//
	LineItem []*model.LineItem10 `xml:"LineItm,omitempty"`
}

func (f *FinancialInvoiceV01) AddInvoiceHeader() *model.InvoiceHeader1 {
	f.InvoiceHeader = new(model.InvoiceHeader1)
	return f.InvoiceHeader
}

func (f *FinancialInvoiceV01) AddTradeAgreement() *model.TradeAgreement6 {
	f.TradeAgreement = new(model.TradeAgreement6)
	return f.TradeAgreement
}

func (f *FinancialInvoiceV01) AddTradeDelivery() *model.TradeDelivery1 {
	f.TradeDelivery = new(model.TradeDelivery1)
	return f.TradeDelivery
}

func (f *FinancialInvoiceV01) AddTradeSettlement() *model.TradeSettlement1 {
	f.TradeSettlement = new(model.TradeSettlement1)
	return f.TradeSettlement
}

func (f *FinancialInvoiceV01) AddLineItem() *model.LineItem10 {
	newValue := new(model.LineItem10)
	f.LineItem = append(f.LineItem, newValue)
	return newValue
}
