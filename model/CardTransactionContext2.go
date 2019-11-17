package model

// Context of the card transaction.
type CardTransactionContext2 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	// It correspond to the ISO 8583:1993 field number 22-6.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	// It correspond to the ISO 8583:1993 field number 22-5.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// Location category of the place where the transaction is actually performed.
	// It correspond partially to the ISO 8583:1993 field number 22-4.
	LocationCategory *LocationCategory2Code `xml:"LctnCtgy,omitempty"`

	// Human attendance at the terminal location during the transaction.
	// It correspond partially to the ISO 8583:1993 field number 22-4.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment2Code `xml:"TxEnvt,omitempty"`

	// Indicates the entity hosting the terminal performing the transaction.
	// It correspond partially to the ISO 8583:1993 field number 22-4.
	HostingCategory *TransactionEnvironment3Code `xml:"HstgCtgy,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	// It correspond to the ISO 8583:1993 field number 22-5.
	TransactionChannel *TransactionChannel3Code `xml:"TxChanl,omitempty"`

	// Entry mode of the card data.
	// It correspond to the ISO 8583 field number 25 for the version 87 (partially), field number 22-7 for the version 93, and field number 22-1 for the version 2003.
	CardDataEntryMode *CardDataReading2Code `xml:"CardDataNtryMd"`

	// Indicator of a card entry mode fall-back. It correspond to the ISO 8583:2003 field number 22-1.
	FallbackIndicator *CardFallback1Code `xml:"FllbckInd,omitempty"`

	// Payment options the card acceptor can support.
	SupportedOption []*SupportedPaymentOption1Code `xml:"SpprtdOptn,omitempty"`

	// Data used to assign specific condition such as liability shift or preferential interchange fees.
	SpecialConditions []*CardTransactionCondition1 `xml:"SpclConds,omitempty"`

	// Indicates to the issuer the level of risk of the transaction.
	RiskIndicator []*CardTransactionRiskIndicator1 `xml:"RskInd,omitempty"`
}

func (c *CardTransactionContext2) SetCardPresent(value string) {
	c.CardPresent = (*TrueFalseIndicator)(&value)
}

func (c *CardTransactionContext2) SetCardholderPresent(value string) {
	c.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (c *CardTransactionContext2) SetLocationCategory(value string) {
	c.LocationCategory = (*LocationCategory2Code)(&value)
}

func (c *CardTransactionContext2) SetAttendanceContext(value string) {
	c.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (c *CardTransactionContext2) SetTransactionEnvironment(value string) {
	c.TransactionEnvironment = (*TransactionEnvironment2Code)(&value)
}

func (c *CardTransactionContext2) SetHostingCategory(value string) {
	c.HostingCategory = (*TransactionEnvironment3Code)(&value)
}

func (c *CardTransactionContext2) SetTransactionChannel(value string) {
	c.TransactionChannel = (*TransactionChannel3Code)(&value)
}

func (c *CardTransactionContext2) SetCardDataEntryMode(value string) {
	c.CardDataEntryMode = (*CardDataReading2Code)(&value)
}

func (c *CardTransactionContext2) SetFallbackIndicator(value string) {
	c.FallbackIndicator = (*CardFallback1Code)(&value)
}

func (c *CardTransactionContext2) AddSupportedOption(value string) {
	c.SupportedOption = append(c.SupportedOption, (*SupportedPaymentOption1Code)(&value))
}

func (c *CardTransactionContext2) AddSpecialConditions() *CardTransactionCondition1 {
	newValue := new(CardTransactionCondition1)
	c.SpecialConditions = append(c.SpecialConditions, newValue)
	return newValue
}

func (c *CardTransactionContext2) AddRiskIndicator() *CardTransactionRiskIndicator1 {
	newValue := new(CardTransactionRiskIndicator1)
	c.RiskIndicator = append(c.RiskIndicator, newValue)
	return newValue
}
