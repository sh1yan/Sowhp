package scripts

import (
	log "Sowhp/concert/logger"
	"bufio"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// extractDomainAndIP 提取URL中的顶级域名
func extractDomainAndIP(url string) (domain string) {
	// 匹配域名
	domainPattern := `^(?:https?://)?([\w.-]+)`
	domainRegex := regexp.MustCompile(domainPattern)
	domainMatch := domainRegex.FindStringSubmatch(url)
	if len(domainMatch) > 1 {
		domain = domainMatch[1]
	}
	return domain
}

// FindTextUrl 获取本地text中url地址列表
func FindTextUrl(filepath string) []string {
	filePath := filepath // 替换为实际的txt文件路径

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		// fmt.Println("无法打开文件:", err)
		log.DebugError(err)
		return []string{}
	}
	defer file.Close()

	urls := []string{} // 用于存储URL地址的列表

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// 提取URL地址
		urllist := extractURL(line)
		for _, url := range urllist {
			if url != "" {
				urls = append(urls, url)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		// fmt.Println("读取文件错误:", err)
		log.DebugError(err)
		return []string{}
	}

	return urls
}

// extractURL 提取IP端口地址，并生成url地址
func extractURL(line string) []string {

	var urlTmpList []string
	var http string = "http://"
	var https string = "https://"

	// 判断当前输入的ip地址或者域名地址是否包含 http 或者 https
	if strings.Contains(line, "http://") || strings.Contains(line, "https://") {
		urllist := append(urlTmpList, line)
		return urllist
	} else if IsIPAddress(line) || IsDomainName(line) {
		line1 := http + line
		line2 := https + line
		urllist := append(urlTmpList, line1, line2)
		return urllist
	} else if IsIPAddressWithPort(line) || IsDomainNameWithPort(line) {
		line1 := http + line
		line2 := https + line
		urllist := append(urlTmpList, line1, line2)
		return urllist
	}

	return []string{}
}

// visitURL 访问URL地址，在Tsaks中按顺利执行
func visitURL(url string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
	}
}

// GetTimeStrin 获取当前时间字符串
func GetTimeStrin() string {
	currentTime := time.Now()
	timeString := currentTime.Format("20060102150405")
	return timeString
}

// GetUrlStatusCode 获取网站状态码
func GetUrlStatusCode(url string) string {

	// 创建一个自定义的 http.Client
	client := &http.Client{
		Timeout: 5 * time.Second, // 设置超时时间为 5 秒
	}

	resp, err := client.Head(url)
	if err != nil {
		log.DebugError(err)
		log.Debug(fmt.Sprintf("获取 %s 状态码超时，返回一个空的字符串", url))
		return " "
	}
	defer resp.Body.Close()
	statuscode := strconv.Itoa(resp.StatusCode)
	log.Debug(fmt.Sprintf("当前网站URL地址为：%s ,网站状态码为：%s", url, statuscode))
	return statuscode
}

// ChromeScreenshot 访问url地址并对地址进行截图
func ChromeScreenshot(URL string, resultname string) []string {

	// 创建一个上下文，并设置超时时间为 10 秒
	ct, cancel := context.WithTimeout(context.Background(), 13*time.Second)
	defer cancel()

	// 创建一个新的上下文
	ctx, cancel := chromedp.NewContext(ct)
	defer cancel()

	// 创建一个选项数组，配置浏览器行为
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("ignore-certificate-errors", true), // 忽略证书错误
	)
	// 创建一个自定义的执行器
	allocCtx, cancel := chromedp.NewExecAllocator(ctx, opts...)
	defer cancel()
	// 使用自定义执行器创建新的上下文
	ctx, cancel = chromedp.NewContext(allocCtx)
	defer cancel()

	// 定义一个变量来保存状态码和标题
	var infoarray []string

	// var statusCode int
	var pageTitle string

	// run
	var b2 []byte
	if err := chromedp.Run(ctx,
		// reset
		chromedp.Emulate(device.Reset),

		// 设置浏览器界面大小，并捕获/截取当前浏览器视口的屏幕截图
		chromedp.EmulateViewport(1920, 1080),
		//chromedp.Navigate(URL),
		visitURL(URL),
		chromedp.CaptureScreenshot(&b2),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.Evaluate(`document.title`, &pageTitle),
	); err != nil {
		log.DebugError(err)
		log.Debug(fmt.Sprintf("访问 %s 超时，返回一个空的 []string{} 字典集", URL))
		// 在这里处理无法访问URL的情况，例如记录错误、重试等
		cancel() // 强制关闭 chromedp 上下文
		return []string{}
	}

	urlname := extractDomainAndIP(URL)
	photoname := fmt.Sprintf("data/%s-%s.png", urlname, GetTimeStrin())
	log.Debug(fmt.Sprintf("生成截图照片成功，截图名称为：%s", photoname))
	result := fmt.Sprintf("./result/%s/%s", resultname, photoname)
	if err := os.WriteFile(result, b2, 0666); err != nil {
		// log.Fatal(err)
		log.DebugError(err)
	}
	log.Debug(fmt.Sprintf("本地存储截图照片成功，存储位置为：%s", result))

	statuscode := GetUrlStatusCode(URL)
	infoarray = append(infoarray, URL, pageTitle, statuscode, photoname)
	cancel() // 强制关闭 chromedp 上下文
	return infoarray
}
