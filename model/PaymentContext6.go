package model

// Payment context in which the transaction is performed.
type PaymentContext6 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// On-line or off-line context of the transaction.
	OnLineContext *TrueFalseIndicator `xml:"OnLineCntxt,omitempty"`

	// Human attendance at the POI (Point Of Interaction) location during the transaction.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment1Code `xml:"TxEnvt,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	TransactionChannel *TransactionChannel1Code `xml:"TxChanl,omitempty"`

	// Entry mode of the card data.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd"`

	// Indicator of a card entry mode fallback.
	FallbackIndicator *TrueFalseIndicator `xml:"FllbckInd,omitempty"`

	// Payment options the card acceptor can support.
	SupportedOption []*SupportedPaymentOption1Code `xml:"SpprtdOptn,omitempty"`
}

func (p *PaymentContext6) SetCardPresent(value string) {
	p.CardPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext6) SetCardholderPresent(value string) {
	p.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext6) SetOnLineContext(value string) {
	p.OnLineContext = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext6) SetAttendanceContext(value string) {
	p.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (p *PaymentContext6) SetTransactionEnvironment(value string) {
	p.TransactionEnvironment = (*TransactionEnvironment1Code)(&value)
}

func (p *PaymentContext6) SetTransactionChannel(value string) {
	p.TransactionChannel = (*TransactionChannel1Code)(&value)
}

func (p *PaymentContext6) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (p *PaymentContext6) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext6) AddSupportedOption(value string) {
	p.SupportedOption = append(p.SupportedOption, (*SupportedPaymentOption1Code)(&value))
}
