package pixiv

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/NateScarlet/pixiv/pkg/artwork"
	"github.com/XiaoSanGit/pixiv_daily_crawl/handler"
	"github.com/XiaoSanGit/pixiv_daily_crawl/handler/common"
	"github.com/XiaoSanGit/pixiv_daily_crawl/model"
	"github.com/XiaoSanGit/pixiv_daily_crawl/setting"
	_ "image/gif"
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

func DownloadAllResultByIds(ids []string) {
	//ids := []string{"94544184"}
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
				common.ServeFrames(item.Data, basePath+item.Filename)
			}
		}
	}
}

//DownloadOneImageFromOneId 从某个id下载一张图片
func DownloadOneImageFromOneId(id string, lastId string, f *os.File) {
	//ids := []string{"94544184"}
	writeStream := csv.NewWriter(f)
	if lastId != "" && id == lastId {
		writeStream.Write([]string{id, ""})
		writeStream.Flush()
		return
	}

	IdUrlMap, err := handler.GetPixivImgUrlsByPixivId([]string{id})
	if err != nil {
		return
	}
	for _, urls := range IdUrlMap { //就一个
		img, errImg := handler.FetchPixivImage(ctx, id, urls[0:1]) //就取一张而已
		if errImg != nil {
			err = fmt.Errorf("%v: %w", urls[0:1], errImg)
			log.Printf("fetching image error: %v", err)
			//return
		} else {
			for _, item := range img.Images {
				fileName := setting.DataFolder + item.Filename
				common.ServeFrames(item.Data, fileName)
				writeStream.Write([]string{id, fileName})
			}
		}
	}
	writeStream.Flush()
	return
}

func TopRankCrawlerYearly(ctx context.Context, month int, day int) {
	var lastArtId string   //日榜可能蝉联
	var lastArtName string //日榜可能蝉联
	var err error
	var f *os.File
	if _, err = os.Stat(setting.CsvName); errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does not exist
		f, err = os.Create(setting.CsvName)
		_, _ = f.WriteString("\xEF\xBB\xBF")
	} else {
		f, err = os.OpenFile(setting.CsvName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	}
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for ; month < 13; month++ {
		for ; day < common.DayMap[month]+1; day++ {
			// 画作排行榜
			rank := &artwork.Rank{
				Mode:    "daily",
				Content: "illust",
				Date: time.Date(2021, time.Month(month), day, 18, 0,
					0, 0, time.Local),
			}
			_ = rank.Fetch(ctx)
			thisID := rank.Items[0].Artwork.ID
			fmt.Printf("%d-%d,%v \n", month, day, thisID)
			dateString := fmt.Sprintf("%d-%d", month, day)
			thisName, err := DownloadOneFromTodayTop(thisID, lastArtId, lastArtName, f, dateString)
			if err != nil {
				print(err)
				return
			}
			lastArtId = thisID
			lastArtName = thisName
		}
	}
}

//DownloadOneFromTodayTop 从某个id下载一张图片
func DownloadOneFromTodayTop(id string, lastId string, lastName string, f *os.File, dateString string) (thisName string, err error) {
	//ids := []string{"94544184"}
	writeStream := csv.NewWriter(f)
	if lastId != "" && id == lastId {
		print(lastId, lastName)
		newName := lastName[0:len(lastName)-4] + "_dup" + lastName[len(lastName)-4:]
		writeStream.Write([]string{dateString, id, newName})
		writeStream.Flush()
		_, err = common.CopyFile(newName, lastName)
		if err != nil {
			return "", err
		}
		return lastName, nil
	}

	var fileName string
	IdUrlMap, err := handler.GetPixivImgUrlsByPixivId([]string{id})
	if err != nil {
		return "", err
	}
	for _, urls := range IdUrlMap { //就一个
		img, errImg := handler.FetchPixivImage(ctx, id, urls[0:1]) //就取一张而已
		if errImg != nil {
			err = fmt.Errorf("%v: %w", urls[0:1], errImg)
			log.Printf("fetching image error: %v", err)
			//return
		} else {
			for _, item := range img.Images {
				fileName = setting.DataFolder + item.Filename
				common.ServeFrames(item.Data, fileName)
				writeStream.Write([]string{dateString, id, fileName})
			}
		}
	}
	writeStream.Flush()
	return fileName, nil
}
