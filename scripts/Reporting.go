package scripts

import (
	log "Sowhp/concert/logger"
	"fmt"
	"github.com/klarkxy/gohtml"
	"os"
)

var resultname string

func CreateHtml(m map[string]map[string][]string) {

	for result_name := range m {
		resultname = result_name // 获取html结果标题头
	}

	bootstrap := gohtml.NewHtml()
	bootstrap.Html()
	bootstrap.Head().Title().Text(resultname)
	bootstrap.Meta().Charset("utf-8")
	bootstrap1 := bootstrap.Body().Tag("div").Align("center").Id("content").Tag("table").Align("center").Border("1").Tag("tbody")
	bootstrap1.Tr().Align("center").Td().Colspan("7").Text("url探测结果")
	tr := bootstrap1.Tr().Align("center").Bgcolor("#0080FF").Attr("style", "color:white")
	tr.Td().Text("URL详情")
	tr.Td().Text("截图")

	filetxt, err := os.OpenFile(fmt.Sprintf("./result/%s/%s.txt", resultname, resultname), os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.FileMode(0666))
	if err != nil {
		// log.Fatal(err)
		log.DebugError(err)

	}
	_, err = filetxt.WriteString(fmt.Sprintf("%s  %s  %s  %s \r\n", "Website URL Address", "TItle Name", "status", "Website screenshot path"))
	if err != nil {
		// log.Fatal(err)
		log.DebugError(err)
	}
	defer filetxt.Close()

	for urlippath := range m[resultname] {
		titlename := m[resultname][urlippath][0]
		statevalue := m[resultname][urlippath][1]
		photopath := m[resultname][urlippath][2]
		tr1 := bootstrap1.Tr().Align("center")
		td := tr1.Td()
		td.A().Href(urlippath).Target("_blank").Text(urlippath)
		td.Text(fmt.Sprintf("%s  %s", statevalue, titlename))
		tr1.Td().Img().Src(photopath).Attr("style", "width:800px;hight:200px")
		_, err = filetxt.WriteString(fmt.Sprintf("%s  %s  %s  %s", urlippath, titlename, statevalue, photopath))
		if err != nil {
			// log.Fatal(err)
			log.DebugError(err)
		}
		filetxt.WriteString("\r\n")
	}

	resultHtml := fmt.Sprintf("./result/%s/%s.html", resultname, resultname)
	// 生成最终结果HTML展示情况
	filehtml, err := os.OpenFile(resultHtml, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.FileMode(0666))
	if err != nil {
		// log.Fatal(err)
		log.DebugError(err)
	}
	// 关闭文件
	defer filehtml.Close()

	_, err = filehtml.WriteString(bootstrap.String())
	if err != nil {
		// log.Fatal(err)
		log.DebugError(err)
	}
	filehtml.WriteString("\r\n")
	log.Common(fmt.Sprintf("以生成最终结果报告，报告存放路径为：%s", resultHtml))
}
