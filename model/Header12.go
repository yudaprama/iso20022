package model

// Set of characteristics related to the transfer of transactions.
type Header12 struct {

	// Indicates if the file transfer is a download or an upload.
	DownloadTransfer *TrueFalseIndicator `xml:"DwnldTrf"`

	// Version of file format.
	FormatVersion *Max6Text `xml:"FrmtVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	// Date and time at which the file or message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification53 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the exchange.
	RecipientParty *GenericIdentification53 `xml:"RcptPty,omitempty"`
}

func (h *Header12) SetDownloadTransfer(value string) {
	h.DownloadTransfer = (*TrueFalseIndicator)(&value)
}

func (h *Header12) SetFormatVersion(value string) {
	h.FormatVersion = (*Max6Text)(&value)
}

func (h *Header12) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header12) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header12) AddInitiatingParty() *GenericIdentification53 {
	h.InitiatingParty = new(GenericIdentification53)
	return h.InitiatingParty
}

func (h *Header12) AddRecipientParty() *GenericIdentification53 {
	h.RecipientParty = new(GenericIdentification53)
	return h.RecipientParty
}
