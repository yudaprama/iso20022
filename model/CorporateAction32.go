package model

// Provides information about the corporate action event.
type CorporateAction32 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate45 `xml:"DtDtls,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage *CorporateActionEventStageFormat14Choice `xml:"EvtStag,omitempty"`

	// Indicates whether the message is related to a claim on the associated corporate action event.
	AdditionalBusinessProcessIndicator *AdditionalBusinessProcessFormat10Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat4Choice `xml:"LtryTp,omitempty"`
}

func (c *CorporateAction32) AddDateDetails() *CorporateActionDate45 {
	c.DateDetails = new(CorporateActionDate45)
	return c.DateDetails
}

func (c *CorporateAction32) AddEventStage() *CorporateActionEventStageFormat14Choice {
	c.EventStage = new(CorporateActionEventStageFormat14Choice)
	return c.EventStage
}

func (c *CorporateAction32) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat10Choice {
	c.AdditionalBusinessProcessIndicator = new(AdditionalBusinessProcessFormat10Choice)
	return c.AdditionalBusinessProcessIndicator
}

func (c *CorporateAction32) AddLotteryType() *LotteryTypeFormat4Choice {
	c.LotteryType = new(LotteryTypeFormat4Choice)
	return c.LotteryType
}
