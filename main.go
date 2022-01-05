package main

import (
	"context"
	"fmt"
	"github.com/NateScarlet/pixiv/pkg/client"
	"github.com/XiaoSanGit/pixiv_daily_crawl/handler/pixiv"
	"github.com/XiaoSanGit/pixiv_daily_crawl/setting"
	"time"
)

func main() {
	// 默认客户端用环境变量 `PIXIV_PHPSESSID` 登录。
	// 并且 User-Agent 使用 `PIXIV_USER_AGENT` 或库内置的默认值。
	//artIds := make(chan string)
	// 使用 PHPSESSID Cookie 登录 (推荐)。
	var year, month, day int
	fmt.Println("输入年（num）：")
	fmt.Scanln(&year)
	fmt.Println("输入月（num）：")
	fmt.Scanln(&month)
	fmt.Println("输入日（num）：")
	fmt.Scanln(&day)

	c := &client.Client{}
	c.SetDefaultHeader("User-Agent", client.DefaultUserAgent)
	c.SetPHPSESSID(setting.GetCookie())
	c.Timeout = time.Hour * 5
	// 启用免代理，环境变量 `PIXIV_BYPASS_SNI_BLOCKING` 不为空时自动为默认客户端启用免代理。
	// 当前实现需求一个 DNS over HTTPS 服务，默认使用 cloudflare，可通过 `PIXIV_DNS_QUERY_URL` 环境变量设置。
	// 必须在其他客户端选项前调用 `BypassSNIBlocking`，因为对于封锁的域名它会使用一个更改过的 Transport 进行请求，无视在它之前进行的的设置。

	c.BypassSNIBlocking()

	// 所有查询从 context 获取客户端设置, 如未设置将使用默认客户端。
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Hour*1)
	defer cancelFunc()
	ctx = client.With(ctx, c)
	print("start crawler...")
	pixiv.TopRankCrawlerYearly(ctx, year, month, day)

	//// Instantiate default collector
	//c := colly.NewCollector(
	//	colly.AllowedDomains("www.pixiv.net"),
	//	colly.MaxDepth(1),
	//)
	//页
	//detailRegex, _ := regexp.Compile(`/go/go\?p=\d+$`)
	//// 匹配下面模式的是该网站的列表页
	//listRegex, _ := regexp.Compile(`/t/\d+#\w+`)
	//
	//// 所有a标签，上设置回调函数
	//c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	//	link := e.Attr("href")
	//
	//	// 已访问过的详情页或列表页，跳过
	//	if visited[link] && (detailRegex.Match([]byte(link)) || listRegex.Match([]byte(link))) {
	//		return
	//	}
	//
	//	// 既不是列表页，也不是详情页
	//	// 那么不是我们关心的内容，要跳过
	//	if !detailRegex.Match([]byte(link)) && !listRegex.Match([]byte(link)) {
	//		println("not match", link)
	//		return
	//	}
	//
	//	// 因为大多数网站有反爬虫策略
	//	// 所以爬虫逻辑中应该有 sleep 逻辑以避免被封杀
	//	time.Sleep(time.Second)
	//	println("match", link)
	//
	//	visited[link] = true
	//
	//	time.Sleep(time.Millisecond * 2)
	//	c.Visit(e.Request.AbsoluteURL(link))
	//})
	//
	//err := c.Visit("https://www.abcdefg.com/go/go")
	//if err != nil {fmt.Println(err)}
}
