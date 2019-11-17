package model

// Choice of identifiers for a clearing system member, as assigned by the clearing system. In some clearing systems, the accounts of the clearing system members are also assigned an identifier.
type ClearingSystemMemberIdentificationChoice struct {

	// (United States) Clearing House Interbank Payments System (CHIPS) Universal Identification (UID) - identifies entities that own accounts at CHIPS participating financial institutions, through which CHIPS payments are effected. The CHIPS UID is assigned by the New York Clearing House.
	CHIPSUniversalIdentification *CHIPSUniversalIdentifier `xml:"USCHU"`

	// New Zealand Bank/Branch Code - identifies New Zealand institutions on the New Zealand national clearing system. The code is assigned by the New Zealand Bankers' Association (NZBA).
	NewZealandNCCIdentification *NewZealandNCCIdentifier `xml:"NZNCC"`

	// Irish National Sorting Code - identifies Irish financial institutions on the Irish national clearing system. The code is assigned by the Irish Payments Services Organisation (IPSO).
	IrishNSCIdentification *IrishNSCIdentifier `xml:"IENSC"`

	// United Kingdom (UK) Sort Code - identifies British financial institutions on the British national clearing systems. The sort code is assigned by the Association for Payments and Clearing Services (APACS).
	UKDomesticSortCode *UKDomesticSortCodeIdentifier `xml:"GBSC"`

	// (United States) Clearing House Interbank Payment System (CHIPS) Participant Identifier (ID) - identifies financial institutions participating on CHIPS. The CHIPS Participant ID is assigned by the New York Clearing House.
	CHIPSParticipantIdentification *CHIPSParticipantIdentifier `xml:"USCH"`

	// Swiss Bank Code - identifies Swiss institutions on the Swiss national clearing system.
	SwissBCIdentification *SwissBCIdentifier `xml:"CHBC"`

	// Fedwire Routing Number - identifies financial institutions, in the US, on the FedWire system. The routing number is assigned by the American Bankers Association (ABA).
	FedwireRoutingNumberIdentification *FedwireRoutingNumberIdentifier `xml:"USFW"`

	// Portuguese National Clearing Code - identifies Portuguese financial institutions on the Portuguese national clearing system.
	PortugueseNCCIdentification *PortugueseNCCIdentifier `xml:"PTNCC"`

	// Russian Central Bank Identification Code - identifies Russian financial institutions on the Russian national clearing system.
	RussianCentralBankIdentificationCode *RussianCentralBankIdentificationCodeIdentifier `xml:"RUCB"`

	// Italian Domestic Identification Code - identifies Italian financial institutions on the Italian national clearing system. The code is assigned by the Associazione Bancaria Italiana (ABI).
	ItalianDomesticIdentificationCode *ItalianDomesticIdentifier `xml:"ITNCC"`

	// Austrian Bankleitzahl - identifies Austrian financial institutions on the Austrian national clearing system.
	AustrianBankleitzahlIdentification *AustrianBankleitzahlIdentifier `xml:"ATBLZ"`

	// Canadian Payments Association Routing Number - identifies Canadian financial institutions on the Canadian national clearing system.
	CanadianPaymentsAssociationRoutingNumberIdentification *CanadianPaymentsARNIdentifier `xml:"CACPA"`

	// Swiss Interbank Clearing (SIC) Code - identifies Swiss financial institutions domestically, on the Swiss national clearing system.
	SwissSICIdentification *SwissSICIdentifier `xml:"CHSIC"`

	// German Bankleitzahl - identifies German financial institutions on the German national clearing systems.
	GermanBankleitzahlIdentification *GermanBankleitzahlIdentifier `xml:"DEBLZ"`

	// Spanish Domestic Interbanking Code - identifies Spanish financial institutions on the Spanish national clearing system. The code is assigned by the Centro de Cooperacion Interbancaria (CCI).
	SpanishDomesticInterbankingIdentification *SpanishDomesticInterbankingIdentifier `xml:"ESNCC"`

	// South African National Clearing Code (NCC) - identifies South African financial institutions on the South African national clearing system. The code is assigned by the South African Bankers Services Company Ltd. (BankServ).
	SouthAfricanNCCIdentification *SouthAfricanNCCIdentifier `xml:"ZANCC"`

	// Hong Kong Bank Code - identifies Hong Kong financial institutions on the Hong Kong local clearing system.
	HongKongBankCode *HongKongBankIdentifier `xml:"HKNCC"`

	// Extensive branch network list of the Australian Bank State Branch (BSB) Code. The codes are used for identifying Australian financial institutions, as assigned by the Australian Payments Clearing Association (APCA).
	AustralianExtensiveBranchNetworkIdentification *ExtensiveBranchNetworkIdentifier `xml:"AUBSBx"`

	// Small network list of the Australian Bank State Branch (BSB) Code. The codes are used for identifying Australian financial institutions , as assigned by the Australian Payments Clearing Association (APCA).
	AustralianSmallNetworkIdentification *SmallNetworkIdentifier `xml:"AUBSBs"`
}

func (c *ClearingSystemMemberIdentificationChoice) SetCHIPSUniversalIdentification(value string) {
	c.CHIPSUniversalIdentification = (*CHIPSUniversalIdentifier)(&value)
}

func (c *ClearingSystemMemberIdentificationChoice) SetNewZealandNCCIdentification(value string) {
	c.NewZealandNCCIdentification = (*NewZealandNCCIdentifier)(&value)
}

func (c *ClearingSystemMemberIdentificationChoice) SetIrishNSCIdentification(value string) {
	c.IrishNSCIdentification = (*IrishNSCIdentifier)(&value)
}

func (c *ClearingSystemMemberIdentificationChoice) SetUKDomesticSortCode(value string) {
	c.UKDomesticSortCode = (*UKDomesticSortCodeIdentifier)(&value)
}

func (c *ClearingSystemMemberIdentificationChoice) SetCHIPSParticipantIdentification(value string) {
	c.CHIPSParticipantIdentification = (*CHIPSParticipantIdentifier)(&value)
}

func (c *ClearingSystemMemberIdentificationChoice) SetSwissBCIdentification(value string) {
	c.SwissBCIdentification = (*SwissBCIdentifier)(&value)
}

func (c *ClearingSystemMemberIdentificationChoice) SetFedwireRoutingNumberIdentification(value string) {
	c.FedwireRoutingNumberIdentification = (*FedwireRoutingNumberIdentifier)(&value)
}

func (c *ClearingSystemMemberIdentificationChoice) SetPortugueseNCCIdentification(value string) {
	c.PortugueseNCCIdentification = (*PortugueseNCCIdentifier)(&value)
}

func (c *ClearingSystemMemberIdentificationChoice) SetRussianCentralBankIdentificationCode(value string) {
	c.RussianCentralBankIdentificationCode = (*RussianCentralBankIdentificationCodeIdentifier)(&value)
}

func (c *ClearingSystemMemberIdentificationChoice) SetItalianDomesticIdentificationCode(value string) {
	c.ItalianDomesticIdentificationCode = (*ItalianDomesticIdentifier)(&value)
}

func (c *ClearingSystemMemberIdentificationChoice) SetAustrianBankleitzahlIdentification(value string) {
	c.AustrianBankleitzahlIdentification = (*AustrianBankleitzahlIdentifier)(&value)
}

func (c *ClearingSystemMemberIdentificationChoice) SetCanadianPaymentsAssociationRoutingNumberIdentification(value string) {
	c.CanadianPaymentsAssociationRoutingNumberIdentification = (*CanadianPaymentsARNIdentifier)(&value)
}

func (c *ClearingSystemMemberIdentificationChoice) SetSwissSICIdentification(value string) {
	c.SwissSICIdentification = (*SwissSICIdentifier)(&value)
}

func (c *ClearingSystemMemberIdentificationChoice) SetGermanBankleitzahlIdentification(value string) {
	c.GermanBankleitzahlIdentification = (*GermanBankleitzahlIdentifier)(&value)
}

func (c *ClearingSystemMemberIdentificationChoice) SetSpanishDomesticInterbankingIdentification(value string) {
	c.SpanishDomesticInterbankingIdentification = (*SpanishDomesticInterbankingIdentifier)(&value)
}

func (c *ClearingSystemMemberIdentificationChoice) SetSouthAfricanNCCIdentification(value string) {
	c.SouthAfricanNCCIdentification = (*SouthAfricanNCCIdentifier)(&value)
}

func (c *ClearingSystemMemberIdentificationChoice) SetHongKongBankCode(value string) {
	c.HongKongBankCode = (*HongKongBankIdentifier)(&value)
}

func (c *ClearingSystemMemberIdentificationChoice) SetAustralianExtensiveBranchNetworkIdentification(value string) {
	c.AustralianExtensiveBranchNetworkIdentification = (*ExtensiveBranchNetworkIdentifier)(&value)
}

func (c *ClearingSystemMemberIdentificationChoice) SetAustralianSmallNetworkIdentification(value string) {
	c.AustralianSmallNetworkIdentification = (*SmallNetworkIdentifier)(&value)
}
