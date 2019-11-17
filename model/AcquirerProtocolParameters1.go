package model

// Acceptor parameters dedicated to the acquirer protocol.
type AcquirerProtocolParameters1 struct {

	// Identification of the acquirer using this protocol.
	AcquirerIdentification []*GenericIdentification32 `xml:"AcqrrId"`

	// Identification of the payment application, user of the acquirer protocol.
	ApplicationIdentification []*Max35Text `xml:"ApplId,omitempty"`

	// Acquirer host configuration.
	Host []*AcquirerHostConfiguration1 `xml:"Hst,omitempty"`

	// Acquirer protocol parameters of transactions performing an online authorisation.
	OnLineTransaction *AcquirerProtocolParameters2 `xml:"OnLineTx,omitempty"`

	// Acquirer protocol parameters of transactions performing an offline authorisation.
	OffLineTransaction *AcquirerProtocolParameters2 `xml:"OffLineTx,omitempty"`

	// Configuration parameters of reconciliation exchanges.
	ReconciliationExchange *ExchangeConfiguration1 `xml:"RcncltnXchg,omitempty"`

	// Indicates the reconciliation period is assigned by the acquirer instead of the acceptor.
	ReconciliationByAcquirer *TrueFalseIndicator `xml:"RcncltnByAcqrr,omitempty"`

	// Indicates the reconciliation total amounts are computed per currency.
	TotalsPerCurrency *TrueFalseIndicator `xml:"TtlsPerCcy,omitempty"`

	// Types of transaction to include in the batch.
	BatchTransferContent []*BatchTransactionType1Code `xml:"BtchTrfCntt,omitempty"`

	// Configuration of a message item.
	MessageItem []*MessageItemCondition1 `xml:"MsgItm,omitempty"`

	// Indicator to require protection of sensitive card data in messages.
	ProtectCardData *TrueFalseIndicator `xml:"PrtctCardData"`
}

func (a *AcquirerProtocolParameters1) AddAcquirerIdentification() *GenericIdentification32 {
	newValue := new(GenericIdentification32)
	a.AcquirerIdentification = append(a.AcquirerIdentification, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters1) AddApplicationIdentification(value string) {
	a.ApplicationIdentification = append(a.ApplicationIdentification, (*Max35Text)(&value))
}

func (a *AcquirerProtocolParameters1) AddHost() *AcquirerHostConfiguration1 {
	newValue := new(AcquirerHostConfiguration1)
	a.Host = append(a.Host, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters1) AddOnLineTransaction() *AcquirerProtocolParameters2 {
	a.OnLineTransaction = new(AcquirerProtocolParameters2)
	return a.OnLineTransaction
}

func (a *AcquirerProtocolParameters1) AddOffLineTransaction() *AcquirerProtocolParameters2 {
	a.OffLineTransaction = new(AcquirerProtocolParameters2)
	return a.OffLineTransaction
}

func (a *AcquirerProtocolParameters1) AddReconciliationExchange() *ExchangeConfiguration1 {
	a.ReconciliationExchange = new(ExchangeConfiguration1)
	return a.ReconciliationExchange
}

func (a *AcquirerProtocolParameters1) SetReconciliationByAcquirer(value string) {
	a.ReconciliationByAcquirer = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters1) SetTotalsPerCurrency(value string) {
	a.TotalsPerCurrency = (*TrueFalseIndicator)(&value)
}

func (a *AcquirerProtocolParameters1) AddBatchTransferContent(value string) {
	a.BatchTransferContent = append(a.BatchTransferContent, (*BatchTransactionType1Code)(&value))
}

func (a *AcquirerProtocolParameters1) AddMessageItem() *MessageItemCondition1 {
	newValue := new(MessageItemCondition1)
	a.MessageItem = append(a.MessageItem, newValue)
	return newValue
}

func (a *AcquirerProtocolParameters1) SetProtectCardData(value string) {
	a.ProtectCardData = (*TrueFalseIndicator)(&value)
}
