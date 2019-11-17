package model

// Set of characteristics related to the transfer of transactions.
type Header27 struct {

	// Indicates if the file transfer is a download or an upload.
	DownloadTransfer *TrueFalseIndicator `xml:"DwnldTrf"`

	// Version of file format.
	FormatVersion *Max6Text `xml:"FrmtVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Number `xml:"XchgId"`

	// Date and time at which the file or message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification71 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the exchange.
	RecipientParty *GenericIdentification92 `xml:"RcptPty,omitempty"`
}

func (h *Header27) SetDownloadTransfer(value string) {
	h.DownloadTransfer = (*TrueFalseIndicator)(&value)
}

func (h *Header27) SetFormatVersion(value string) {
	h.FormatVersion = (*Max6Text)(&value)
}

func (h *Header27) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Number)(&value)
}

func (h *Header27) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header27) AddInitiatingParty() *GenericIdentification71 {
	h.InitiatingParty = new(GenericIdentification71)
	return h.InitiatingParty
}

func (h *Header27) AddRecipientParty() *GenericIdentification92 {
	h.RecipientParty = new(GenericIdentification92)
	return h.RecipientParty
}
