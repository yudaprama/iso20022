package model

// Acceptor parameters dedicated to the acquirer protocol.
type AcquirerProtocolParameters7 struct {

	// Identification of the acquirer using this protocol.
	AcquirerIdentification []*GenericIdentification53 `xml:"AcqrrId"`

	// Version of the acquirer protocol parameters.
	Version *Max256Text `xml:"Vrsn"`

	// Identification of the payment application, user of the acquirer protocol.
	ApplicationIdentification []*Max35Text `xml:"ApplId,omitempty"`

	// Acquirer host configuration.
	Host []*AcquirerHostConfiguration3 `xml:"Hst,omitempty"`

	// Acquirer protocol parameters of transactions performing an online authorisation.
	OnLineTransaction *AcquirerProtocolParameters8 `xml:"OnLineTx,omitempty"`

	// Acquirer protocol parameters of transactions performing an offline authorisation.
	OffLineTransaction *AcquirerProtocolParameters8 `xml:"OffLineTx,omitempty"`

	// Configuration parameters of reconciliation exchanges.
	ReconciliationExchange *ExchangeConfiguration6 `xml:"RcncltnXchg,omitempty"`

	// Indicates the reconciliation period is assigned by the acquirer instead of the acceptor.
	ReconciliationByAcquirer *TrueFalseIndicator `xml:"RcncltnByAcqrr,omitempty"`

	// Indicates the reconciliation total amounts are computed per currency.
	TotalsPerCurrency *TrueFalseIndicator `xml:"TtlsPerCcy,omitempty"`

	// Indicates that totals in reconciliation or batch must be split per group of points of interaction and card product profiles when these information are present in the transactions.
	SplitTotals *TrueFalseIndicator `xml:"SpltTtls,omitempty"`

	// True if the POI must send card data (protected or plain card data) in the AcceptorCompletionAdvice message following an authorisation exchange.
	CardDataVerification *TrueFalseIndicator `xml:"CardDataVrfctn,omitempty"`

	// Send a cancellation advice for offline transactions not yet captured.
	NotifyOffLineCancellation *TrueFalseIndicator `xml:"NtfyOffLineCxl,omitempty"`

	// Types of transaction to include in the batch.
	BatchTransferContent []*BatchTransactionType1Code `xml:"BtchTrfCntt,omitempty"`

	// BatchTransfer are exchanged per file transfer protocol rather than per message.
	FileTransferBatch *TrueFalseIndicator `xml:"FileTrfBtch,omitempty"`

	// BatchTransfer are authenticated by digital signature rather than a MAC (Message Authentication Code).
	BatchDigitalSignature *TrueFalseIndicator `xml:"BtchDgtlSgntr,omitempty"`

	// Configuration of a message item.
	MessageItem []*MessageItemCondition1 `xml:"MsgItm,omitempty"`

	// Indicator to require protection of sensitive card data in messages.
	ProtectCardData *TrueFalseIndicator `xml:"PrtctCardData"`
}

func (a *AcquirerProtocolParameters7) AddAcquirerIdentification() *GenericIdentification53 {
	newValue := new(GenericIdentification53)
	a.AcquirerIdentification = append(a.AcquirerIdentification, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters7) SetVersion(value string) {
	a.Version = (*Max256Text)(&value)
}

func (a *AcquirerProtocolParameters7) AddApplicationIdentification(value string) {
	a.ApplicationIdentification = append(a.ApplicationIdentification, (*Max35Text)(&value))
}

func (a *AcquirerProtocolParameters7) AddHost() *AcquirerHostConfiguration3 {
	newValue := new(AcquirerHostConfiguration3)
	a.Host = append(a.Host, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters7) AddOnLineTransaction() *AcquirerProtocolParameters8 {
	a.OnLineTransaction = new(AcquirerProtocolParameters8)
	return a.OnLineTransaction
}

func (a *AcquirerProtocolParameters7) AddOffLineTransaction() *AcquirerProtocolParameters8 {
	a.OffLineTransaction = new(AcquirerProtocolParameters8)
	return a.OffLineTransaction
}

func (a *AcquirerProtocolParameters7) AddReconciliationExchange() *ExchangeConfiguration6 {
	a.ReconciliationExchange = new(ExchangeConfiguration6)
	return a.ReconciliationExchange
}

func (a *AcquirerProtocolParameters7) SetReconciliationByAcquirer(value string) {
	a.ReconciliationByAcquirer = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters7) SetTotalsPerCurrency(value string) {
	a.TotalsPerCurrency = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters7) SetSplitTotals(value string) {
	a.SplitTotals = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters7) SetCardDataVerification(value string) {
	a.CardDataVerification = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters7) SetNotifyOffLineCancellation(value string) {
	a.NotifyOffLineCancellation = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters7) AddBatchTransferContent(value string) {
	a.BatchTransferContent = append(a.BatchTransferContent, (*BatchTransactionType1Code)(&value))
}

func (a *AcquirerProtocolParameters7) SetFileTransferBatch(value string) {
	a.FileTransferBatch = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters7) SetBatchDigitalSignature(value string) {
	a.BatchDigitalSignature = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters7) AddMessageItem() *MessageItemCondition1 {
	newValue := new(MessageItemCondition1)
	a.MessageItem = append(a.MessageItem, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters7) SetProtectCardData(value string) {
	a.ProtectCardData = (*TrueFalseIndicator)(&value)
}
