package otherrule

// 基础包
import (
	"github.com/bovlov/anothervote/app/downloader/request" //必需
	. "github.com/bovlov/anothervote/app/spider"           //必需
	//. "github.com/bovlov/anothervote/app/spider/common"    //选用
	"github.com/bovlov/anothervote/common/goquery" //DOM解析
	"github.com/bovlov/anothervote/logs"           //信息输出

	//"bufio"
	//"bytes"
	//"io/ioutil"
	// net包
	"net/http" //设置http.Header
	"net/url"

	// 编码包
	// "encoding/xml"
	// "encoding/json"

	// 字符串处理包
	// "regexp"
	"strconv"
	"strings"

	// 其他包
	//"fmt"
	// "math"
	//"time"
	// "io/ioutil"
)

func init() {
	Hejin.Register()
}

var Hejin = &Spider{
	Name:         "HEJIN",
	Description:  `HEJIN 自定义输入格式 url`,
	Pausetime:    2000,
	Keyin:        KEYIN,
	Limit:        LIMIT,
	EnableCookie: true,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			//ctx.Request is nil, dont use it here
			param := ctx.GetKeyin()
			if len(param) <= 12 {
				logs.Log.Warning("自定义输入的url参数不正确！ use default")
				//return
				param = `http://tzxts.lzyjdzsw.com/plugin.php?id=hejin_toupiao&model=detail&zid=20`
			}
			//logs.Log.Warning("param=%v", param)

			//http://tzxts.lzyjdzsw.com/plugin.php?id=hejin_toupiao&model=detail&zid=20
			urlParsed, _ := url.Parse(param)
			urlParams, _ := url.ParseQuery(urlParsed.RawQuery)
			logs.Log.Error("host=%v script=%v query=%v", urlParsed.Host, urlParsed.Path, urlParsed.RawPath)
			urlModel, modelExist := urlParams["model"]
			if !modelExist || len(urlModel) == 0 {
				logs.Log.Error("不是有效的url,model not exist,有效的url应该类似：http://tzxts.lzyjdzsw.com/plugin.php?id=hejin_toupiao&model=detail&zid=1")
				return
			}
			if urlParsed.Path != "/plugin.php" {
				logs.Log.Error("不是有效的url,plugin.php not exist,有效的url应该类似：http://tzxts.lzyjdzsw.com/plugin.php?id=hejin_toupiao&model=detail&zid=1 %v", urlParsed.Path)
				return
			}

			pluginId, pluginIdExist := urlParams["id"]
			if !pluginIdExist || len(pluginId) == 0 {
				logs.Log.Error("不是有效的url,pluginId not exist,有效的url应该类似：http://tzxts.lzyjdzsw.com/plugin.php?id=hejin_toupiao&model=detail&zid=1 ")
				return
				///} else if string(pluginId[0])[:5] != "hejin" {
				//logs.Log.Error("不是有效的url,pluginId != hejin*,有效的url应该类似：http://tzxts.lzyjdzsw.com/plugin.php?id=hejin_toupiao&model=detail&zid=1 %v", pluginId[0])
				//return
			}
			logs.Log.Error("pluginId=%v urlModel=%v", pluginId, urlModel)
			vid, vidExist := urlParams["vid"]
			if !vidExist || len(vid) == 0 {
				vid = make([]string, 1)
				vid[0] = "1"
			}

			logs.Log.Error("vid=%v", vid)
			zid, zidExist := urlParams["zid"]
			if !zidExist || len(zid) == 0 {
				logs.Log.Error("没有匹配到要投票的用户 请输入带zid的url %v", zid[0])
				return
			}

			//logs.Log.Error("zid=%v", zid)
			//ctx.SetTemp("pluginId", pluginId[0]) // cannot use ctx.SetTemp in root

			/* 思路：获取所有zid；根据zid导出openid；唯一openid；请求ticket带cookie */
			// urlTop300 := "http://tzxts.lzyjdzsw.com/plugin.php?id=hejin_toupiao&model=top300&vid=1#top300"
			urlPre := urlParsed.Scheme + "://" + urlParsed.Host + urlParsed.Path + "?id=" + pluginId[0]
			//ctx.SetTemp("urlPre", urlPre)
			urlTop300 := urlPre + "&model=top300&vid=" + vid[0]
			logs.Log.Warning("will top300: %v", urlTop300)

			ctx.AddQueue(&request.Request{
				Url:  urlTop300,
				Rule: "top300",
				Header: http.Header{
					"Cookie":     []string{},
					"Referer":    []string{param},
					"User-Agent": []string{"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36"},
				},
				DownloaderID: 0,
				Temp: map[string]interface{}{
					"zid":      zid[0],
					"model":    urlModel[0],
					"pluginId": pluginId[0],
					"urlPre":   urlPre,
					"vid":      vid[0],
				},
			})
		},

		Trunk: map[string]*Rule{
			"top300": {
				ParseFunc: func(ctx *Context) {
					logs.Log.Warning("start top300: url=%v", ctx.GetUrl())
					//query := ctx.GetDom()
					//logs.Log.Warning("got query when top300")
					if ctx.Response == nil {
						logs.Log.Error("no response!")
						return
					}
					//bodyContent := ctx.Response.Body
					//textContentBytes, bodyErr := ioutil.ReadAll(bodyContent)
					//if bodyErr != nil {
					//	logs.Log.Error("read body error: %v", bodyErr)
					//	return
					//}
					//textContent := string(textContentBytes)
					textContent := ctx.GetText()
					logs.Log.Warning("the textContent len=%v %v", len(textContent), textContent[:32])
					//bodyContent.Read(textContent)
					//logs.Log.Warning("the text len=%v %v", len(textContent), textContent[:32])
					tempContent := []byte(textContent)
					rankIdx := strings.Index(textContent, `<div class="rank300" id="top300">`)
					var rankContent string
					if rankIdx > 0 {
						logs.Log.Warning("find div rank300 at: %d", rankIdx)
						tempContent = tempContent[rankIdx:]
						textContent = string(tempContent)
						rankEndIdx := strings.Index(textContent, "</div>")
						if rankEndIdx > 0 {
							logs.Log.Warning("find /div at: %d", rankEndIdx)
							//tempContent = []byte(textContent)
							tempContent = tempContent[:rankEndIdx]
							rankContent = string(tempContent)
						}
					}
					logs.Log.Warning("the rankContent len=%v %v", len(rankContent), string(tempContent[:32]))
					if len(rankContent) == 0 {
						logs.Log.Error("no rank content in text:%v", string(tempContent))
						return
					}

					// 因为该网站的代码比较垃圾 编码混乱，gb2312和utf8混排，导致goquery无法解析，只能手动
					var uids []int
					tempContent = []byte(rankContent)
					for {
						spanIdx := strings.Index(rankContent, "</span><span>1")
						uid := 0
						if spanIdx > 0 {
							logs.Log.Warning("find a span: %v", string(tempContent[spanIdx:spanIdx+18]))
							uid, _ = strconv.Atoi(string(tempContent[spanIdx+14 : spanIdx+18]))
							uids = append(uids, uid)
							tempContent = tempContent[spanIdx+18:]
							rankContent = string(tempContent)
							continue
						}
						break
					}
					logs.Log.Warning("uids: len=%d %v", len(uids), uids)
					return

					//logs.Log.Warning("find rankContent:%v", rankContent)
					r := strings.NewReader(rankContent)
					rankQuery, _ := goquery.NewDocumentFromReader(r)

					// 找到文章链接 加入队列
					rankQuery.Add(rankContent).Find(".list li").Each(func(i int, s *goquery.Selection) {
						subUid := s.Find("span").Eq(1).Text()
						if len(subUid) > 0 {
							logs.Log.Warning("we find a uid:%v", subUid)
							uid, _ := strconv.Atoi(subUid)
							uids = append(uids, uid-10000)
							url := ctx.GetTemp("urlPre", "").(string) + "&model=dcexcel&zid=" + strconv.FormatInt(int64(uid), 10)
							logs.Log.Warning("will dcexcel: %v", url)

							ctx.AddQueue(&request.Request{
								Url:  url,
								Rule: "dcexcel",
								Header: http.Header{
									"Cookie":     []string{},
									"Referer":    []string{url},
									"User-Agent": []string{"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36"},
								},
								DownloaderID: 0,
							})
						}
					})

					ctx.SetTemp("uids", uids)
					logs.Log.Warning("has got uids:%v", uids)
				},
			},
			"dcexcel": {
				ItemFields: []string{
					"blogId",
					"文章标题",
					"文章内容",
				},
				ParseFunc: func(ctx *Context) {

					query := ctx.GetDom()
					var url = ctx.GetUrl()
					logs.Log.Warning("start dcexcel:%v", url)
					text := ctx.GetText()
					if len(text) > 0 {
						logs.Log.Warning("len=%v url=%v", len(text), url)
						return
					}
					return

					//var blogId int64 = 0
					var blogIdStr string
					if blogIdIdx := strings.Index(url, "blogid="); blogIdIdx > 0 {
						blogIdStr = url[blogIdIdx+9 : blogIdIdx+30]
						if commaIdx := strings.Index(blogIdStr, "&"); commaIdx > 0 {
							blogIdStr = blogIdStr[:commaIdx]
						}
						//fmt.Sscanf(string(blogIdStr), "%d", &blogId)
					}

					var detail = query.Find("#blogDetailDiv").Text()
					var title = query.Find(".blog_tit_detail").Eq(0).Text()
					logs.Log.Error("blogid=%v title=%v len(detail)=%v ", blogIdStr, title, len(detail))
					rowRet := map[int]interface{}{
						0: blogIdStr,
						1: title,
						2: detail,
					}

					// 结果输出
					ctx.Output(rowRet)
				},
			},
			"votea": {
				ItemFields: []string{
					"blogId",
					"文章标题",
					"文章内容",
				},
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					var url = ctx.GetUrl()

					//var blogId int64 = 0
					var blogIdStr string
					if blogIdIdx := strings.Index(url, "blogid="); blogIdIdx > 0 {
						blogIdStr = url[blogIdIdx+9 : blogIdIdx+30]
						if commaIdx := strings.Index(blogIdStr, "&"); commaIdx > 0 {
							blogIdStr = blogIdStr[:commaIdx]
						}
						//fmt.Sscanf(string(blogIdStr), "%d", &blogId)
					}

					var detail = query.Find("#blogDetailDiv").Text()
					var title = query.Find(".blog_tit_detail").Eq(0).Text()
					logs.Log.Error("blogid=%v title=%v len(detail)=%v ", blogIdStr, title, len(detail))
					rowRet := map[int]interface{}{
						0: blogIdStr,
						1: title,
						2: detail,
					}

					// 结果输出
					ctx.Output(rowRet)
				},
			},
		},
	},
}
