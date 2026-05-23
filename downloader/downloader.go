package downloader

import (
	"context"
	"fmt"
	"os"
	"path"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/lrstanley/go-ytdlp"
	"github.com/w3sync/golash.git/metadata"
	"github.com/w3sync/golash.git/utils"
)

func Download(url string) error {
	info, err := metadata.GetMetadata(url)
	if err != nil {
		panic(err)
	}

	fmt.Println(info.Title)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	for k, v := range info.Formats {
		if v.ACodec != "none" && v.VCodec != "none" {
			fmt.Fprintf(w,
				"idx: %d | FORMAT_ID: %s | EXT: %s | ACODEC: %s | VCODEC: %s | WIDTH: %s | HEIGHT: %s | FPS: %s | RESOLUTIO: %s | FORMAT: %s \n",
				k,
				v.FormatID,
				v.Ext,
				v.ACodec,
				v.VCodec,
				v.Width,
				v.Height,
				v.FPS,
				v.Resolution,
				v.Format,
			)
		}
	}

	var selectedID string
	fmt.Scanf("%s", &selectedID)
	idx, _ := strconv.Atoi(selectedID)

	loc := path.Join(utils.GetConfigMetadata().DownloadPath, utils.GetConfigMetadata().FilenameTemplate)

	dl := ytdlp.
		New().
		Format(string(info.Formats[idx].FormatID)).
		Output(loc).
		Progress().
		ProgressFunc(time.Second, func(p ytdlp.ProgressUpdate) {
			fmt.Printf(
				"\r%.2f%% | %d / %d bytes | ETA: %s",
				p.Percent(),
				p.DownloadedBytes,
				p.TotalBytes,
				p.ETA(),
			)

			if p.Status == ytdlp.ProgressStatusFinished {
				fmt.Println("\nDone")
			}
		})

	_, err = dl.Run(context.Background(), url)
	if err != nil {
		panic(err)
	}

	return nil
}
