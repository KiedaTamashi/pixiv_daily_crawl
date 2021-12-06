package pixiv

import (
	"bytes"
	"context"
	"fmt"
	"github.com/XiaoSanGit/pixiv_daily_crawl/handler"
	"github.com/XiaoSanGit/pixiv_daily_crawl/model"
	"image"
	"image/jpeg"
	"log"
	"os"
	"sync"
	"time"
)

var ctx, cancel = context.WithTimeout(context.Background(), time.Minute*10)

type result struct {
	img *model.PixivOneIdImage
	err error
}

func DailyRankCrawler() {
	ids := []string{"94544184"}
	IdUrlMap, err := handler.GetPixivImgUrlsByPixivId(ids)
	if err != nil {
		return
	}
	wg := sync.WaitGroup{}
	results := make(chan result)
	wg.Add(len(IdUrlMap))
	go func() { wg.Wait(); close(results) }()

	fetch := func(id string, links []string) {
		defer wg.Done()
		img, errImg := handler.FetchPixivImage(ctx, id, links)
		if errImg != nil {
			err := fmt.Errorf("%v: %w", links, errImg)
			results <- result{err: err}
			return
		}
		results <- result{img: img}
	}

	for id, urls := range IdUrlMap {
		go fetch(id, urls)
	}

	for r := range results {
		switch {
		case r.err != nil:
			log.Printf("fetching image: %v", r.err)
		default:
			//ArtworId := r.img.Id
			basePath := "./data/"
			for _, item := range r.img.Images {
				serveFrames(item.Data, basePath+item.Filename)
			}
		}
	}
}

func serveFrames(imgByte []byte, name string) {

	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		log.Fatalln(err)
	}
	//os.MkdirAll(name)

	out, _ := os.Create(name)
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 9

	err = jpeg.Encode(out, img, &opts)
	//jpeg.Encode(out, img, nil)
	if err != nil {
		log.Println(err)
	}

}
