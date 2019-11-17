package model

// Acceptor parameters dedicated to the acquirer protocol.
type AcquirerProtocolParameters8 struct {

	// Mode for the financial capture of the transaction by the acquirer.
	FinancialCapture *FinancialCapture1Code `xml:"FinCaptr"`

	// Configuration of the batch transfers.
	BatchTransfer *ExchangeConfiguration6 `xml:"BtchTrf,omitempty"`

	// Configuration parameters of completion exchanges.
	CompletionExchange *ExchangeConfiguration7 `xml:"CmpltnXchg,omitempty"`

	// Configuration of the cancellation exchanges.
	CancellationExchange *CancellationProcess1Code `xml:"CxlXchg,omitempty"`
}

func (a *AcquirerProtocolParameters8) SetFinancialCapture(value string) {
	a.FinancialCapture = (*FinancialCapture1Code)(&value)
}

func (a *AcquirerProtocolParameters8) AddBatchTransfer() *ExchangeConfiguration6 {
	a.BatchTransfer = new(ExchangeConfiguration6)
	return a.BatchTransfer
}

func (a *AcquirerProtocolParameters8) AddCompletionExchange() *ExchangeConfiguration7 {
	a.CompletionExchange = new(ExchangeConfiguration7)
	return a.CompletionExchange
}

func (a *AcquirerProtocolParameters8) SetCancellationExchange(value string) {
	a.CancellationExchange = (*CancellationProcess1Code)(&value)
}
