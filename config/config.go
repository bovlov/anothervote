package config

import (
	"strings"

	"github.com/bovlov/anothervote/logs/logs"
	"github.com/bovlov/anothervote/runtime/cache"
)

// 软件信息。
const (
	VERSION   string = "v1.2.0"                                      // 软件版本号
	AUTHOR    string = "another"                                     // 软件作者
	NAME      string = "anothervote"                                 // 软件名
	FULL_NAME string = NAME + "_" + VERSION + " （by " + AUTHOR + "）" // 软件全称
	TAG       string = "anothervote"                                 // 软件标识符
	ICON_PNG  string = `AAABAAEAEBAAAAEAIABoBAAAFgAAACgAAAAQAAAAIAAAAAEAIAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAD+/v8y/Pz8Mv39/TL9/fwy//7+Mv7+/zL9/v4y+/z8Mvv7/DL7/Pwy/f7+Mv7//jL+//4y/v7+Mv39/TL7+/sy+/38/8zLxv/Hycb/y87N/9HU0//g4uD/09LM/7mzpv+8uKz/vLiv/766s/+/u7X/1NPR//7+/v////////////n27v+0bwb/tm4B/7RsBP+1dxr/o2YM/7+MPf/RpGD/0KVi/8+jYP/MoVj/zJ9U/55rIv/v8e///v7+///////49uz/vHEB/71yAP+8cQL/172K/86GDP/X18///f78//3+/P/7/vz/+v36/826l/+sagr/p5V7//3+/P//////9vTq/7xxAf+9cgD/u3IA/7tyA//LgQn/4d/b//////////////////z++//19u//7ere/6p7Of/8/fv///////by6f+7cQH/vXIA/71xAP+8cgD/xn0J/+rp6P/////////////////+//7//P78/7aGPf+pbRf/0s/K//7+/v/18Of/vHEB/71yAP+9cgD/u3IA/795Cf/19PT///////////////////////z+/v/5/fn/4Mec/8Ovkf/+/v7/9PDl/71xAP+9cgD/vXIA/7xyAf+8eQ//+/v8///////////////////////9/v7/v6N0/6htFf+okGv//v7+//Lw4/+7cgP/vHIC/7xxA/+8cQP/tG8H/7ixpf/9/v3//v7+//7+/v/+/v7//v7+//z59P/08eb/qXkw//f5+P/6+/j/8ejW//Pr3P/17+P/9/Pn/+7n0/+gbSb/3uDd/97b1P++tav/vbWs/721rf++ta3/vrer/8CQSP/4/fr////////////////////////////+/v7/4Meb/5+BWf/hzqn/sXMW/9GlXv/RpF7/0aNe/9KjXv/p28D/+/77//////////////////////////////////r9+v+ygjX/1tXS/6Z1Lf/y8/L////////////////////////////////////////////////////////////9/v3/38Wb/7Wigv/NpGH/09DM//7+/v////////////////////////////////////////////////////////////Hm0/+2lmH/1axr/93Wz//+/v7////////////////////////////////////////////////////////////n1rj/qXUs/8GQRf/6+/n/////////////////////////////////////////////////////////////////+vn0/+XSsP/18OX//v79//////////////////////////////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==`
	DESC      string = "another vote 您的专业投票助手"
)

// 默认配置。
const (
	WORK_ROOT      string = TAG + "_pkg"                    // 运行时的目录名称
	CONFIG         string = WORK_ROOT + "/config.ini"       // 配置文件路径
	CACHE_DIR      string = WORK_ROOT + "/cache"            // 缓存文件目录
	LOG            string = WORK_ROOT + "/logs/pholcus.log" // 日志文件路径
	LOG_ASYNC      bool   = true                            // 是否异步输出日志
	PHANTOMJS_TEMP string = CACHE_DIR                       // Surfer-Phantom下载器：js文件临时目录
	HISTORY_TAG    string = "history"                       // 历史记录的标识符
	HISTORY_DIR    string = WORK_ROOT + "/" + HISTORY_TAG   // excel或csv输出方式下，历史记录目录
	SPIDER_EXT     string = ".pholcus.html"                 // 动态规则扩展名
)

// 来自配置文件的配置项。
var (
	CRAWLS_CAP int = setting.DefaultInt("crawlcap", crawlcap) // 蜘蛛池最大容量
	// DATA_CHAN_CAP            int    = setting.DefaultInt("datachancap", datachancap)                               // 收集器容量
	PHANTOMJS                string = setting.String("phantomjs")                                          // Surfer-Phantom下载器：phantomjs程序路径
	PROXY                    string = setting.String("proxylib")                                           // 代理IP文件路径
	SPIDER_DIR               string = setting.String("spiderdir")                                          // 动态规则目录
	FILE_DIR                 string = setting.String("fileoutdir")                                         // 文件（图片、HTML等）结果的输出目录
	TEXT_DIR                 string = setting.String("textoutdir")                                         // excel或csv输出方式下，文本结果的输出目录
	DB_NAME                  string = setting.String("dbname")                                             // 数据库名称
	MGO_CONN_STR             string = setting.String("mgo::connstring")                                    // mongodb连接字符串
	MGO_CONN_CAP             int    = setting.DefaultInt("mgo::conncap", mgoconncap)                       // mongodb连接池容量
	MGO_CONN_GC_SECOND       int64  = setting.DefaultInt64("mgo::conngcsecond", mgoconngcsecond)           // mongodb连接池GC时间，单位秒
	MYSQL_CONN_STR           string = setting.String("mysql::connstring")                                  // mysql连接字符串
	MYSQL_CONN_CAP           int    = setting.DefaultInt("mysql::conncap", mysqlconncap)                   // mysql连接池容量
	MYSQL_MAX_ALLOWED_PACKET int    = setting.DefaultInt("mysql::maxallowedpacket", mysqlmaxallowedpacket) // mysql通信缓冲区的最大长度

	KAFKA_BORKERS string = setting.DefaultString("kafka::brokers", kafkabrokers) //kafka brokers

	LOG_CAP            int64 = setting.DefaultInt64("log::cap", logcap)          // 日志缓存的容量
	LOG_LEVEL          int   = logLevel(setting.String("log::level"))            // 全局日志打印级别（亦是日志文件输出级别）
	LOG_CONSOLE_LEVEL  int   = logLevel(setting.String("log::consolelevel"))     // 日志在控制台的显示级别
	LOG_FEEDBACK_LEVEL int   = logLevel(setting.String("log::feedbacklevel"))    // 客户端反馈至服务端的日志级别
	LOG_LINEINFO       bool  = setting.DefaultBool("log::lineinfo", loglineinfo) // 日志是否打印行信息                                  // 客户端反馈至服务端的日志级别
	LOG_SAVE           bool  = setting.DefaultBool("log::save", logsave)         // 是否保存所有日志到本地文件
)

func init() {
	// 主要运行时参数的初始化
	cache.Task = &cache.AppConf{
		Mode:           setting.DefaultInt("run::mode", mode),                 // 节点角色
		Port:           setting.DefaultInt("run::port", port),                 // 主节点端口
		Master:         setting.String("run::master"),                         // 服务器(主节点)地址，不含端口
		ThreadNum:      setting.DefaultInt("run::thread", thread),             // 全局最大并发量
		Pausetime:      setting.DefaultInt64("run::pause", pause),             // 暂停时长参考/ms(随机: Pausetime/2 ~ Pausetime*2)
		OutType:        setting.String("run::outtype"),                        // 输出方式
		DockerCap:      setting.DefaultInt("run::dockercap", dockercap),       // 分段转储容器容量
		Limit:          setting.DefaultInt64("run::limit", limit),             // 采集上限，0为不限，若在规则中设置初始值为LIMIT则为自定义限制，否则默认限制请求数
		ProxyMinute:    setting.DefaultInt64("run::proxyminute", proxyminute), // 代理IP更换的间隔分钟数
		SuccessInherit: setting.DefaultBool("run::success", success),          // 继承历史成功记录
		FailureInherit: setting.DefaultBool("run::failure", failure),          // 继承历史失败记录
	}
}

func logLevel(l string) int {
	switch strings.ToLower(l) {
	case "app":
		return logs.LevelApp
	case "emergency":
		return logs.LevelEmergency
	case "alert":
		return logs.LevelAlert
	case "critical":
		return logs.LevelCritical
	case "error":
		return logs.LevelError
	case "warning":
		return logs.LevelWarning
	case "notice":
		return logs.LevelNotice
	case "informational":
		return logs.LevelInformational
	case "info":
		return logs.LevelInformational
	case "debug":
		return logs.LevelDebug
	}
	return -10
}
