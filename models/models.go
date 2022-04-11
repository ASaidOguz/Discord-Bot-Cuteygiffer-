package models

type GifSearch struct {
	Data struct {
		Type             string `json:"type"`
		ID               string `json:"id"`
		URL              string `json:"url"`
		Slug             string `json:"slug"`
		BitlyGifURL      string `json:"bitly_gif_url"`
		BitlyURL         string `json:"bitly_url"`
		EmbedURL         string `json:"embed_url"`
		Username         string `json:"username"`
		Source           string `json:"source"`
		Title            string `json:"title"`
		Rating           string `json:"rating"`
		ContentURL       string `json:"content_url"`
		SourceTld        string `json:"source_tld"`
		SourcePostURL    string `json:"source_post_url"`
		IsSticker        int    `json:"is_sticker"`
		ImportDatetime   string `json:"import_datetime"`
		TrendingDatetime string `json:"trending_datetime"`
		Images           struct {
			FixedWidthStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_width_still"`
			PreviewGif struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"preview_gif"`
			FixedHeightDownsampled struct {
				Height   string `json:"height"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_height_downsampled"`
			Preview struct {
				Height  string `json:"height"`
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
				Width   string `json:"width"`
			} `json:"preview"`
			FixedHeightSmall struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_height_small"`
			Downsized struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized"`
			FixedWidthDownsampled struct {
				Height   string `json:"height"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_width_downsampled"`
			FixedWidth struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_width"`
			DownsizedStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized_still"`
			DownsizedMedium struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized_medium"`
			OriginalMp4 struct {
				Height  string `json:"height"`
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
				Width   string `json:"width"`
			} `json:"original_mp4"`
			DownsizedLarge struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized_large"`
			PreviewWebp struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"preview_webp"`
			Original struct {
				Frames   string `json:"frames"`
				Hash     string `json:"hash"`
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"original"`
			OriginalStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"original_still"`
			FixedHeightSmallStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_height_small_still"`
			FixedWidthSmall struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_width_small"`
			Looping struct {
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
			} `json:"looping"`
			DownsizedSmall struct {
				Height  string `json:"height"`
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
				Width   string `json:"width"`
			} `json:"downsized_small"`
			FixedWidthSmallStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_width_small_still"`
			FixedHeightStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_height_still"`
			FixedHeight struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_height"`
			Four80WStill struct {
				URL    string `json:"url"`
				Width  string `json:"width"`
				Height string `json:"height"`
			} `json:"480w_still"`
		} `json:"images"`
		User struct {
			AvatarURL    string `json:"avatar_url"`
			BannerImage  string `json:"banner_image"`
			BannerURL    string `json:"banner_url"`
			ProfileURL   string `json:"profile_url"`
			Username     string `json:"username"`
			DisplayName  string `json:"display_name"`
			Description  string `json:"description"`
			IsVerified   bool   `json:"is_verified"`
			WebsiteURL   string `json:"website_url"`
			InstagramURL string `json:"instagram_url"`
		} `json:"user"`
	} `json:"data"`
	Meta struct {
		Msg        string `json:"msg"`
		Status     int    `json:"status"`
		ResponseID string `json:"response_id"`
	} `json:"meta"`
}
