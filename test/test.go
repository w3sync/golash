package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/lrstanley/go-ytdlp"
)

const (
	bestQuality = "bestvideo[height>=480][height<=720]+bestaudio"
	URL         = "https://www.youtube.com/watch?v=dQw4w9WgXcQ&pp=ygUNcm9jayBhbmQgcm9sbA%3D%3D"
)

func Run() {
	fmt.Println("-------------------------- Main [START] --------------------------------------------")
	cmd := ytdlp.New().SkipDownload().DumpSingleJSON().Format(bestQuality).FormatSort("res,fps,codec:av1,codec:vp9,codec:avc1,br,asr,ext")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	res, err := cmd.Run(ctx, URL)
	if err != nil {
		panic(err)
	}

	var out bytes.Buffer

	err = json.Indent(&out, []byte(res.Stdout), "", "  ")
	if err != nil {
		panic(err)
	}

	if res.Stderr != "" {
		fmt.Println("yt-dlp stderr:", res.Stderr)
	}

	file, err := os.Create("./result.json")
	if err != nil {
		panic(err)
	}

	file.Write(out.Bytes())

	fmt.Println("-------------------------- Main [END] --------------------------------------------")
}
