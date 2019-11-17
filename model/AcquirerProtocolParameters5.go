package model

// Acceptor parameters dedicated to the acquirer protocol.
type AcquirerProtocolParameters5 struct {

	// Mode for the financial capture of the transaction by the acquirer.
	FinancialCapture *FinancialCapture1Code `xml:"FinCaptr"`

	// Configuration of the batch transfers.
	BatchTransfer *ExchangeConfiguration4 `xml:"BtchTrf,omitempty"`

	// Configuration parameters of completion exchanges.
	CompletionExchange *ExchangeConfiguration5 `xml:"CmpltnXchg,omitempty"`

	// Configuration of the cancellation exchanges.
	CancellationExchange *CancellationProcess1Code `xml:"CxlXchg,omitempty"`
}

func (a *AcquirerProtocolParameters5) SetFinancialCapture(value string) {
	a.FinancialCapture = (*FinancialCapture1Code)(&value)
}

func (a *AcquirerProtocolParameters5) AddBatchTransfer() *ExchangeConfiguration4 {
	a.BatchTransfer = new(ExchangeConfiguration4)
	return a.BatchTransfer
}

func (a *AcquirerProtocolParameters5) AddCompletionExchange() *ExchangeConfiguration5 {
	a.CompletionExchange = new(ExchangeConfiguration5)
	return a.CompletionExchange
}

func (a *AcquirerProtocolParameters5) SetCancellationExchange(value string) {
	a.CancellationExchange = (*CancellationProcess1Code)(&value)
}
