package model

// Payment context in which the transaction is performed.
type PaymentContext4 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// On-line or off-line context of the transaction.
	OnLineContext *TrueFalseIndicator `xml:"OnLineCntxt,omitempty"`

	// Human attendance at the POI location during the transaction.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment1Code `xml:"TxEnvt,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	TransactionChannel *TransactionChannel1Code `xml:"TxChanl,omitempty"`

	// Indicates whether a message can be sent or not on an attendant display (attendant display present or not).
	AttendantMessageCapable *TrueFalseIndicator `xml:"AttndntMsgCpbl,omitempty"`

	// Language used to display messages to the attendant.
	AttendantLanguage *ISO2ALanguageCode `xml:"AttndntLang,omitempty"`

	// Entry mode of the card data.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd,omitempty"`

	// Indicator of a card entry mode fallback.
	FallbackIndicator *TrueFalseIndicator `xml:"FllbckInd,omitempty"`
}

func (p *PaymentContext4) SetCardPresent(value string) {
	p.CardPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext4) SetCardholderPresent(value string) {
	p.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext4) SetOnLineContext(value string) {
	p.OnLineContext = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext4) SetAttendanceContext(value string) {
	p.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (p *PaymentContext4) SetTransactionEnvironment(value string) {
	p.TransactionEnvironment = (*TransactionEnvironment1Code)(&value)
}

func (p *PaymentContext4) SetTransactionChannel(value string) {
	p.TransactionChannel = (*TransactionChannel1Code)(&value)
}

func (p *PaymentContext4) SetAttendantMessageCapable(value string) {
	p.AttendantMessageCapable = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext4) SetAttendantLanguage(value string) {
	p.AttendantLanguage = (*ISO2ALanguageCode)(&value)
}

func (p *PaymentContext4) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (p *PaymentContext4) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*TrueFalseIndicator)(&value)
}
