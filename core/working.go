package core

import (
	dire "Sowhp/concert"
	log "Sowhp/concert/logger"
	"Sowhp/scripts"
	"flag"
	"fmt"
	"os"
)

// 当前版本信息
var version = "1.0.2"

// logo
var slogan = `

  _____               _           
 / ____|             | |          
| (___   _____      _| |__  _ __  
 \___ \ / _ \ \ /\ / / '_ \| '_ \ 
 ____) | (_) \ V  V /| | | | |_) |
|_____/ \___/ \_/\_/ |_| |_| .__/ 
                           | |    
                           |_|    

			Sowhp version: ` + version + `

`

var (
	txtfilepath  string
	resultmap    map[string]map[string][]string
	arraymap     map[string][]string
	count        = 0 // 用于统计当前截图的顺序
	count_result = 0 // 用于统计计数当前拍摄成功的条数
)

func Flag() {
	flag.StringVar(&txtfilepath, "f", "", "URL文件路径地址，请参照格式输入, -f D://url.txt")
	flag.IntVar(&log.LogLevel, "logl", 3, "设置日志输出等级，默认为3级，-logl 3")
	print(slogan)
	flag.Parse()
}

func WorkIng() {
	Flag()
	if isCharacterEmpty(txtfilepath) { // 判断输入的URL路径是否为空，若为空，则停止运行。
		log.Debug(fmt.Sprintf("当前输入路径为：%s", txtfilepath))
		run(txtfilepath)
	} else {
		flag.Usage()
		os.Exit(0)
	}

}

func isCharacterEmpty(char string) bool {
	return len(char) != 0
}

func run(path string) {
	log.Common("正在运行 Sowhp 网站首页截图工具")
	urls := scripts.FindTextUrl(path) // 获取TXT里url列表
	log.Debug(fmt.Sprint("以获取本地TXT文本里URL列表：", urls))
	resultname := fmt.Sprintf("result_%s", scripts.GetTimeStrin())
	total := len(urls) // 获取本地列表中的数据长度
	log.Common(fmt.Sprintf("以获取本地URL列表信息数据，共涉及%v条，正在截图拍照中...", total))
	arraymap = make(map[string][]string)             // 初始化arraymap空间地址
	resultmap = make(map[string]map[string][]string) // 初始化resultmap空间地址
	// 格式参考样例 arraymap["Website URL Address"] = []string{"TItle Name", "status", "网站截图路径"}
	dire.MkdirResport()                                            // 创建默认结果目录
	dire.Dir_mk(fmt.Sprintf("./result/%s", resultname))            // 创建本次扫描结果目录文件夹
	dire.Dir_mk(fmt.Sprintf("./result/%s/%s", resultname, "data")) // 创建本次扫描图片存放目录
	for _, url := range urls {
		urlresultlist := scripts.ChromeScreenshot(url, resultname)
		count = count + 1 // 用于展示当前截图拍照的顺序位置
		// 判断是否拍照成功，并根据拍照结果进行对应的结果填充
		if len(urlresultlist) == 0 {
			log.Common(fmt.Sprintf("[%v/%v] 访问 %s 地址超时，无法进行首页截图拍照！", count, total, url))
			arraymap[url] = []string{"x_x!", "连接失败", "data/"}
		} else {
			count_result = count_result + 1
			log.Common(fmt.Sprintf("[%v/%v] 已完成对 %s 地址的首页截图拍照！", count, total, url))
			arraymap[urlresultlist[0]] = []string{urlresultlist[1], urlresultlist[2], urlresultlist[3]}
		}
	}
	resultmap[resultname] = arraymap
	log.Common(fmt.Sprintf("已完成所有URL地址的截图拍照，网站首页拍摄成功数量总数为 %v 条", count_result))
	log.Common("正在生成最终结果报告文件...")
	scripts.CreateHtml(resultmap)
}
