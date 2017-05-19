package api

// EPG data for a single program
type Program struct {
	AudioLanguage         string `json:"audio_language,omitempty"`
	BroadcastGenre        string `json:"broadcast_genre,omitempty"`
	CanonicalGenre        string `json:"canonical_genre,omitempty"`
	ChannelId             string `json:"channel_id,omitempty"`
	ContentRating         string `json:"content_rating,omitempty"`
	EndTimeUtcMillis      string `json:"end_time_utc_millis,omitempty"`
	EpisodeDisplayNumber  string `json:"episode_display_number,omitempty"`
	EpisodeTitle          string `json:"episode_title,omitempty"`
	InternalProviderData  string `json:"internal_provider_data,omitempty"`
	InternalProviderFlag1 string `json:"internal_provider_flag1,omitempty"`
	InternalProviderFlag2 string `json:"internal_provider_flag2,omitempty"`
	InternalProviderFlag3 string `json:"internal_provider_flag3,omitempty"`
	InternalProviderFlag4 string `json:"internal_provider_flag4,omitempty"`
	LongDescription       string `json:"long_description,omitempty"`
	Searchable            string `json:"searchable,omitempty"`
	SeasonDisplayNumber   string `json:"season_display_number,omitempty"`
	SeasonTitle           string `json:"season_title,omitempty"`
	ShortDescription      string `json:"short_description,omitempty"`
	StartTimeUtcMillis    string `json:"start_time_utc_millis,omitempty"`
	ThumbnailUri          string `json:"thumbnail_uri,omitempty"`
	Title                 string `json:"title,omitempty"`
	VersionNumber         string `json:"version_number,omitempty"`
	VideoHeight           string `json:"video_height,omitempty"`
	VideoWidth            string `json:"video_width,omitempty"`
}
