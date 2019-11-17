package model

// Globalised card transaction entry details.
type CardAggregated1 struct {

	// Service in addition to the main service.
	AdditionalService *CardPaymentServiceType2Code `xml:"AddtlSvc,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	TransactionCategory *ExternalCardTransactionCategory1Code `xml:"TxCtgy,omitempty"`

	// Unique identification of the sales reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	SaleReconciliationIdentification *Max35Text `xml:"SaleRcncltnId,omitempty"`

	// Range of sequence numbers on which the globalisation applies.
	SequenceNumberRange *CardSequenceNumberRange1 `xml:"SeqNbRg,omitempty"`

	// Date range on which the globalisation applies.
	TransactionDateRange *DateOrDateTimePeriodChoice `xml:"TxDtRg,omitempty"`
}

func (c *CardAggregated1) SetAdditionalService(value string) {
	c.AdditionalService = (*CardPaymentServiceType2Code)(&value)
}

func (c *CardAggregated1) SetTransactionCategory(value string) {
	c.TransactionCategory = (*ExternalCardTransactionCategory1Code)(&value)
}

func (c *CardAggregated1) SetSaleReconciliationIdentification(value string) {
	c.SaleReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardAggregated1) AddSequenceNumberRange() *CardSequenceNumberRange1 {
	c.SequenceNumberRange = new(CardSequenceNumberRange1)
	return c.SequenceNumberRange
}

func (c *CardAggregated1) AddTransactionDateRange() *DateOrDateTimePeriodChoice {
	c.TransactionDateRange = new(DateOrDateTimePeriodChoice)
	return c.TransactionDateRange
}
