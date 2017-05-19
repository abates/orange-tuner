package api

// Channel metadata
type Channel struct {
	AppLinkColor          string `json:"app_link_color,omitempty"`
	AppLinkIconUri        string `json:"app_link_icon_uri,omitempty"`
	AppLinkIntentUri      string `json:"app_link_intent_uri,omitempty"`
	AppLinkPosterArtUri   string `json:"app_link_poster_art_uri,omitempty"`
	AppLinkText           string `json:"app_link_text,omitempty"`
	Description           string `json:"description,omitempty"`
	DisplayName           string `json:"display_name,omitempty"`
	DisplayNumber         string `json:"display_number,omitempty"`
	InputId               string `json:"input_id,omitempty"`
	InternalProviderData  string `json:"internal_provider_data,omitempty"`
	InternalProviderFlag1 string `json:"internal_provider_flag1,omitempty"`
	InternalProviderFlag2 string `json:"internal_provider_flag2,omitempty"`
	InternalProviderFlag3 string `json:"internal_provider_flag3,omitempty"`
	InternalProviderFlag4 string `json:"internal_provider_flag4,omitempty"`
	NetworkAffiliation    string `json:"network_affiliation,omitempty"`
	OriginalNetworkId     string `json:"original_network_id,omitempty"`
	Searchable            string `json:"searchable,omitempty"`
	ServiceId             string `json:"service_id,omitempty"`
	ServiceType           string `json:"service_type,omitempty"`
	TransportStreamId     string `json:"transport_stream_id,omitempty"`
	Type_                 string `json:"type,omitempty"`
	VersionNumber         string `json:"version_number,omitempty"`
	VideoFormat           string `json:"video_format,omitempty"`
}
