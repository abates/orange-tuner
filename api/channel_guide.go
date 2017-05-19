package api

// EPG data for a single channel
type ChannelGuide struct {
	Programs []Program `json:"programs,omitempty"`
}
