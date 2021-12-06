package main

import "github.com/XiaoSanGit/pixiv_daily_crawl/handler/pixiv"

func main() {
	// 默认客户端用环境变量 `PIXIV_PHPSESSID` 登录。
	// 并且 User-Agent 使用 `PIXIV_USER_AGENT` 或库内置的默认值。

	pixiv.DailyRankCrawler()
	return
	//// 使用 PHPSESSID Cookie 登录 (推荐)。
	//c := &client.Client{}
	//c.SetDefaultHeader("User-Agent", client.DefaultUserAgent)
	//c.SetPHPSESSID("first_visit_datetime_pc=2021-12-02+20%3A16%3A27; p_ab_id=5; p_ab_id_2=0; p_ab_d_id=1666672386; yuid_b=QIKIB1Q; _gcl_au=1.1.1718069623.1638443790; _ga=GA1.2.1660338144.1638444058; _gid=GA1.2.1093842282.1638444058; PHPSESSID=10269478_JkHuNEzBq7mBAEYEEspudTUgJr6ATMdU; device_token=c95812f06f24df7c92e86348118bac18; c_type=24; privacy_policy_notification=0; a_type=0; b_type=1; login_ever=yes; privacy_policy_agreement=3; __cf_bm=LQCutTT.eK0oX_.E_2oTMevDAjPfwE_QFh99jKFjVbU-1638504306-0-AVoJeYvlApl7zOqldyjlHDDUkKLZprzHaQ0ey2AvY1mwgXsVTH0tZ/1vgrUnJrb/Xp8UKLbUYK1gB04zcurfHZcekI6nW+Z3A23ARXUoW+TsmkpSoaE/w+JtjsaMXld+hB0/h4OmWuHFkdR0nkRXu4rdaBhEfR0vgw8Gk6wBoOUNG5hTgYrDXmgj4gmqYozrcQ==; MONITOR_WEB_ID=7253891b-21f2-4c9c-b53f-63ed03feea2c")

	// 启用免代理，环境变量 `PIXIV_BYPASS_SNI_BLOCKING` 不为空时自动为默认客户端启用免代理。
	// 当前实现需求一个 DNS over HTTPS 服务，默认使用 cloudflare，可通过 `PIXIV_DNS_QUERY_URL` 环境变量设置。
	// 必须在其他客户端选项前调用 `BypassSNIBlocking`，因为对于封锁的域名它会使用一个更改过的 Transport 进行请求，无视在它之前进行的的设置。

	//c.BypassSNIBlocking()

	//// 所有查询从 context 获取客户端设置, 如未设置将使用默认客户端。
	//var ctx = context.Background()
	//ctx = client.With(ctx, c)

	//// 搜索画作
	//result, err := artwork.Search(ctx, "パチュリー・ノーレッジ")
	//result.JSON // json return data.
	//result.Artworks() // []artwork.Artwork，只有部分数据，通过 `Fetch` `FetchPages` 方法获取完整数据。
	//artwork.Search(ctx, "パチュリー・ノーレッジ", artwork.SearchOptionPage(2)) // 获取第二页

	//// 画作详情
	//i := &artwork.Artwork{ID: "22238487"}
	//err := i.Fetch(ctx) // 获取画作详情(不含分页), 直接更新 struct 数据。
	//err := i.FetchPages(ctx) // 获取画作分页, 直接更新 struct 数据。

	//// 画作排行榜
	//rank := &artwork.Rank{
	//	Mode: "daily",
	//	Date: time.Date(2021,1,1,12,0,
	//		0,0,time.Local),
	//}
	//rank.Fetch(ctx)
	//print(rank.Items[0].Rank)
	//print(rank.Items[0].PreviousRank)

	//// 用户详情
	//i := &user.User{ID: "789096"}
	//err := i.Fetch(ctx) // 获取用户详情, 直接更新 struct 数据。

	//// Instantiate default collector
	//c := colly.NewCollector(
	//	colly.AllowedDomains("www.pixiv.net"),
	//	colly.MaxDepth(1),
	//)
	//
	//// 我们认为匹配该模式的是该网站的详情页
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
