package metadata

import (
	"context"
	"encoding/json"
	"time"

	"github.com/lrstanley/go-ytdlp"
)

const (
	bestQuality = "bv*[height>=480]+ba/b"
	sortFormat  = "res,fps,codec:av1,codec:vp9,codec:avc1,br,asr,ext"
)

type Format struct {
	FormatID       string      `json:"format_id"`
	FormatNote     string      `json:"format_note"`
	Ext            string      `json:"ext"`
	Protocol       string      `json:"protocol"`
	ACodec         string      `json:"acodec"`
	VCodec         string      `json:"vcodec"`
	URL            string      `json:"url"`
	Width          int         `json:"width"`
	Height         int         `json:"height"`
	FPS            float64     `json:"fps"`
	Rows           int         `json:"rows"`
	Columns        int         `json:"columns"`
	Fragments      []Fragment  `json:"fragments"`
	AudioExt       string      `json:"audio_ext"`
	VideoExt       string      `json:"video_ext"`
	VBR            float64     `json:"vbr"`
	ABR            float64     `json:"abr"`
	TBR            *float64    `json:"tbr"`
	Resolution     string      `json:"resolution"`
	AspectRatio    float64     `json:"aspect_ratio"`
	FilesizeApprox *int64      `json:"filesize_approx"`
	HTTPHeaders    HTTPHeaders `json:"http_headers"`
	Format         string      `json:"format"`
}

type Fragment struct {
	URL      string  `json:"url"`
	Duration float64 `json:"duration"`
}

type HTTPHeaders struct {
	UserAgent      string `json:"User-Agent"`
	Accept         string `json:"Accept"`
	AcceptLanguage string `json:"Accept-Language"`
	SecFetchMode   string `json:"Sec-Fetch-Mode"`
}

type Info struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Formats []Format `json:"formats"`
}

func GetMetadata(url string) (*Info, error) {
	cmd := ytdlp.New().SkipDownload().DumpSingleJSON().Format(bestQuality).FormatSort(sortFormat)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	res, err := cmd.Run(ctx, url)
	if err != nil {
		panic(err)
	}

	var metadata Info

	err = json.Unmarshal([]byte(res.Stdout), &metadata)
	if err != nil {
		panic(err)
	}

	filterData := make([]Format, 0)

	for _, v := range metadata.Formats {
		if len(v.FormatID) >= 2 && v.FormatID[:2] == "sb" {
			continue
		}
		filterData = append(filterData, v)
	}

	metadata.Formats = filterData

	return &metadata, nil
}
