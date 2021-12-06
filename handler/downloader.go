package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/XiaoSanGit/pixiv_daily_crawl/model"
	"github.com/XiaoSanGit/pixiv_daily_crawl/setting"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

//FetchPixivImage 从一组映射获得真实的实体
func FetchPixivImage(ctx context.Context, id string, links []string) (*model.PixivOneIdImage, error) {
	pixivOneIdImage := &model.PixivOneIdImage{
		Id: id,
	}
	var err error
	for _, link := range links {
		thisImage := &model.PixivImage{}
		thisImage.Data, err = getPixivOriginalImage(link)
		thisImage.Filename = filepath.Base(link)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		pixivOneIdImage.Images = append(pixivOneIdImage.Images, thisImage)
	}
	return pixivOneIdImage, nil
}

//getPixivOriginalImage 根据origin image url下载某张图，以byte形式出现
func getPixivOriginalImage(link string) (b []byte, err error) {
	client := &http.Client{
		Timeout: time.Minute * 5,
	}

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Set("Referer", "https://www.pixiv.net")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

//GetPixivImgUrlsByPixivId 给ids去下载，返回下载链接
func GetPixivImgUrlsByPixivId(ids []string) (map[string][]string, error) {
	mapId2Url := map[string][]string{}

	for _, id := range ids {
		link := fmt.Sprintf(setting.DetailPixivItemUrl, id)
		imgUrls, err := getPixivItem(link)
		mapId2Url[id] = imgUrls
		if err != nil {
			return nil, err
		}
	}
	return mapId2Url, nil
}

func getPixivItem(url string) (originals []string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	page := &model.PixivResponse{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(page)
	if err != nil {
		log.Println(err)
		return
	}

	if page.Error {
		err = errors.New(translateError(page.ErrorMessage))
		return
	}

	for _, l := range page.Pages {
		originals = append(originals, l.Urls.Original)
	}

	return
}

func translateError(japanese string) string {
	var errorTranslations = map[string]string{
		"該当作品は削除されたか、存在しない作品IDです。": "The work ID has been deleted or does not exist.",
	}

	english, ok := errorTranslations[japanese]
	if !ok {
		return japanese
	}
	return english
}
