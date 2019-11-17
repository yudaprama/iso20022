package model

// Acceptor parameters dedicated to the acquirer protocol.
type AcquirerProtocolParameters4 struct {

	// Mode for the financial capture of the transaction by the acquirer.
	FinancialCapture *FinancialCapture1Code `xml:"FinCaptr"`

	// Configuration of the batch transfers.
	BatchTransfer *ExchangeConfiguration2 `xml:"BtchTrf,omitempty"`

	// Configuration parameters of completion exchanges.
	CompletionExchange *ExchangeConfiguration3 `xml:"CmpltnXchg,omitempty"`

	// Configuration of the cancellation exchanges.
	CancellationExchange *CancellationProcess1Code `xml:"CxlXchg,omitempty"`
}

func (a *AcquirerProtocolParameters4) SetFinancialCapture(value string) {
	a.FinancialCapture = (*FinancialCapture1Code)(&value)
}

func (a *AcquirerProtocolParameters4) AddBatchTransfer() *ExchangeConfiguration2 {
	a.BatchTransfer = new(ExchangeConfiguration2)
	return a.BatchTransfer
}

func (a *AcquirerProtocolParameters4) AddCompletionExchange() *ExchangeConfiguration3 {
	a.CompletionExchange = new(ExchangeConfiguration3)
	return a.CompletionExchange
}

func (a *AcquirerProtocolParameters4) SetCancellationExchange(value string) {
	a.CancellationExchange = (*CancellationProcess1Code)(&value)
}
