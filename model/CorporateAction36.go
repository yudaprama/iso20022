package model

// Provides information about the corporate action event.
type CorporateAction36 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate50 `xml:"DtDtls,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage *CorporateActionEventStageFormat15Choice `xml:"EvtStag,omitempty"`

	// Indicates that the additional business process relates to a claim on the associated corporate action event.
	AdditionalBusinessProcessIndicator *AdditionalBusinessProcessFormat14Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat5Choice `xml:"LtryTp,omitempty"`
}

func (c *CorporateAction36) AddDateDetails() *CorporateActionDate50 {
	c.DateDetails = new(CorporateActionDate50)
	return c.DateDetails
}

func (c *CorporateAction36) AddEventStage() *CorporateActionEventStageFormat15Choice {
	c.EventStage = new(CorporateActionEventStageFormat15Choice)
	return c.EventStage
}

func (c *CorporateAction36) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat14Choice {
	c.AdditionalBusinessProcessIndicator = new(AdditionalBusinessProcessFormat14Choice)
	return c.AdditionalBusinessProcessIndicator
}

func (c *CorporateAction36) AddLotteryType() *LotteryTypeFormat5Choice {
	c.LotteryType = new(LotteryTypeFormat5Choice)
	return c.LotteryType
}
