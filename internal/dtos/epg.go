package dtos

type EpgResult struct {
	Id             string `json:"id"`
	EpgId          string `json:"epg_id"`
	Title          string `json:"title"`
	Lang           string `json:"lang"`
	Start          string `json:"start"`
	End            string `json:"end"`
	Description    string `json:"description"`
	ChannelId      string `json:"channel_id"`
	StartTimestamp string `json:"start_timestamp"`
	StopTimestamp  string `json:"stop_timestamp"`
}

type EpgListingResult struct {
	EpgListings []EpgResult `json:"epg_listings"`
}
