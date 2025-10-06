package dtos

type VodStreamResult struct {
	Num                int    `json:"num"`
	Name               string `json:"name"`
	StreamType         string `json:"stream_type"`
	StreamId           int    `json:"stream_id"`
	StreamIcon         string `json:"stream_icon"`
	Rating             Rating `json:"rating"`
	Rating5Based       Rating `json:"rating_5based"`
	Added              string `json:"added"`
	IsAdult            string `json:"is_adult"`
	CategoryId         string `json:"category_id"`
	ContainerExtension string `json:"container_extension"`
	CustomSid          string `json:"custom_sid"`
	DirectSource       string `json:"direct_source"`
}
