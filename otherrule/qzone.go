package otherrule

// 基础包
import (
	"github.com/henrylee2cn/pholcus/app/downloader/request" //必需
	. "github.com/henrylee2cn/pholcus/app/spider"           //必需
	//. "github.com/henrylee2cn/pholcus/app/spider/common"    //选用
	"github.com/henrylee2cn/pholcus/common/goquery" //DOM解析
	"github.com/henrylee2cn/pholcus/logs"           //信息输出

	// net包
	"net/http" //设置http.Header
	// "net/url"

	// 编码包
	// "encoding/xml"
	// "encoding/json"

	// 字符串处理包
	// "regexp"
	"strconv"
	"strings"

	// 其他包
	"fmt"
	// "math"
	"time"
	// "io/ioutil"
)

func init() {
	QzoneArticles.Register()
}

var QzoneArticles = &Spider{
	Name:         "QZONE",
	Description:  `QZONE [自定义输入格式 "ID"::"Cookie"][最多支持250页，内设定时1~2s]`,
	Pausetime:    2000,
	Keyin:        KEYIN,
	Limit:        LIMIT,
	EnableCookie: true,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			param := strings.Split(ctx.GetKeyin(), "::")
			if len(param) != 2 {
				logs.Log.Error("自定义输入的参数不正确！")
				return
			}
			id := strings.Trim(param[0], " ")
			cookie := strings.Trim(param[1], " ")

			var i = 0
			for _, blogid := range BlogIds {
				i++
				if i > 2 {
					break
				}
				//https://h5.qzone.qq.com/proxy/domain/b.qzone.qq.com/cgi-bin/blognew/blog_output_data?uin=545845496&blogid=1344636005&styledm=qzonestyle.gtimg.cn&imgdm=qzs.qq.com&bdm=b.qzone.qq.com&mode=2&numperpage=15&timestamp=1502288560&dprefix=&blogseed=0.6215156952384859&inCharset=gb2312&outCharset=gb2312&ref=qzone&entertime=1502288565153&cdn_use_https=1
				urlHost := "https://h5.qzone.qq.com"
				urlPath := "/proxy/domain/b.qzone.qq.com/cgi-bin/blognew/blog_output_data?uin=" + id + "&blogid=" + strconv.Itoa(blogid) + "&styledm=qzonestyle.gtimg.cn&imgdm=qzs.qq.com&bdm=b.qzone.qq.com&mode=2&numperpage=15&timestamp=" + strconv.FormatInt(time.Now().Unix(), 10) + "&dprefix=&blogseed=0.6215156952384859&inCharset=gb2312&outCharset=gb2312&ref=qzone&entertime=1502288565153&cdn_use_https=1"
				ctx.AddQueue(&request.Request{
					Url:  urlHost + urlPath,
					Rule: "文章详情",
					Header: http.Header{
						//":authority":                []string{"h5.qzone.qq.com"},
						//":method":                   []string{"GET"},
						//":path":                     []string{urlPath},
						//":scheme":                   []string{"https"},
						//"upgrade-insecure-requests": []string{"1"},
						"Cookie":     []string{cookie},
						"Referer":    []string{"https://qzs.qq.com/qzone/newblog/blogcanvas.html"},
						"User-Agent": []string{"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36"},
					},
					DownloaderID: 0,
				})
			}
		},

		Trunk: map[string]*Rule{
			"文章列表": {
				ItemFields: []string{
					"文章名",
					"blogid",
					"url",
				},
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					// 找到文章链接 加入队列
					结果 := map[int]interface{}{
						0: ctx.GetTemp("好友名", ""),
						1: ctx.GetTemp("好友ID", ""),
						2: ctx.GetTemp("认证", ""),
					}
					//var i int = 0
					query.Find(".article").Each(func(i int, s *goquery.Selection) {
						logs.Log.Error("this is eq %d", i)
						if i >= 3 {
							return
						}

						//fmt.Println("222")
						artLink := s.Find(".c_tx2 a")
						title, _ := artLink.Attr("title")
						name := artLink.Find("span").Text()
						fmt.Println(name)
						url, _ := artLink.Attr("href")
						blogid, _ := artLink.Attr("blogid")

						logs.Log.Error("i=%d title=%s name=%v url=%v blogid=%v\n", i, title, name, url, blogid)
						结果[i] = title

						/*
							x := &request.Request{
								Url:          url,
								Rule:         "文章详情",
								DownloaderID: 0,
								Temp: map[string]interface{}{
									"好友名":  name,
									"好友ID": uid,
									"认证":   认证,
									"关注":   关注,
									"粉丝":   粉丝,
									"微博":   微博,
								},
							}
							ctx.AddQueue(x)
						*/
					})
					// 结果输出
					ctx.Output(结果)
				},
			},
			"文章详情": {
				ItemFields: []string{
					"文章标题",
					"文章内容",
					"url",
				},
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					var url = ctx.GetUrl()
					var detailHtml, e1 = query.Find("#blogDetailDiv").Html()
					//var detailHtmlEach = query.Find("#blogDetailDiv").Each(func(n int){})
					var detail = query.Find("#blogDetailDiv").Text()
					var title = query.Find(".blog_tit_detail").Eq(0).Text()
					var bodyHtml, e2 = query.Html()
					logs.Log.Error("title=%v detail=%v detailHtml=%v e1=%v body=%v e2=%v", title, detail, detailHtml, e1, bodyHtml, e2)
					结果 := map[int]interface{}{
						0: title,
						1: detail,
						2: url,
					}

					// 结果输出
					ctx.Output(结果)
				},
			},
		},
	},
}
