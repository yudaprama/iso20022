package model

// Acceptor parameters dedicated to the acquirer protocol.
type AcquirerProtocolParameters2 struct {

	// Mode for the financial capture of the transaction by the acquirer.
	FinancialCapture *FinancialCapture1Code `xml:"FinCaptr"`

	// Configuration of the batch transfers.
	BatchTransfer *ExchangeConfiguration1 `xml:"BtchTrf,omitempty"`

	// Configuration parameters of completion exchanges.
	CompletionExchange *ExchangeConfiguration1 `xml:"CmpltnXchg,omitempty"`
}

func (a *AcquirerProtocolParameters2) SetFinancialCapture(value string) {
	a.FinancialCapture = (*FinancialCapture1Code)(&value)
}

func (a *AcquirerProtocolParameters2) AddBatchTransfer() *ExchangeConfiguration1 {
	a.BatchTransfer = new(ExchangeConfiguration1)
	return a.BatchTransfer
}

func (a *AcquirerProtocolParameters2) AddCompletionExchange() *ExchangeConfiguration1 {
	a.CompletionExchange = new(ExchangeConfiguration1)
	return a.CompletionExchange
}
