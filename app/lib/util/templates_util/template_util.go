package templates_util

import (
	"fmt"
	"html/template"
	"math"
)

func GetTemplateFunctions() template.FuncMap {
	return template.FuncMap{
		"add": func(i int, j int) int {
			return i + j // 足し算
		},
		"subtraction": func(i int, j int) int {
			return i - j // 引き算
		},
		"division": func(i int, j int) float64 {
			return float64(i / j) // 割り算
		},
		"floor": func(i float64) float64 {
			return math.Floor(i) // 切り捨て
		},
		"ceil": func(i float64) float64 {
			return math.Ceil(i) // 切り上げ
		},
		"round": func(i float64) float64 {
			return math.Round(i) // 四捨五入
		},
		"generatePaginatorLink": generatePaginatorLink,
	}
}

/**
 * パジネーターのリンクを生成する
 */
func generatePaginatorLink(path string, rowCount int, offset int, limit int) template.HTML {

	htmlString := "<nav aria-label=\"Page navigation example\">"

	htmlString = htmlString + "<ul class=\"pagination justify-content-center\">"

	if offset == 0 {
		htmlString = htmlString + "<li class=\"page-item disabled\">\n<a class=\"page-link\" href=\"#\" tabindex=\"-1\">Previous</a>\n</li>"
	} else {
		htmlString = htmlString + "<li class=\"page-item\">\n<a class=\"page-link\" href=\"#\" tabindex=\"-1\">Previous</a>\n</li>"
	}

	pageNumber := 1
	htmlString = htmlString + fmt.Sprintf(""+
		"<li class=\"page-item\">"+
		"<a class=\"page-link\" href=\"%s?offset=%d&amp;limit=%d\">%d</a>"+
		"</li>",
		path, offset, limit, pageNumber)

	htmlString = htmlString + "<li class=\"page-item\">\n<a class=\"page-link\" href=\"/admin/articles/list?offset=5&limit=5\">Next</a>\n</li>"

	htmlString = htmlString + "</ul>"
	htmlString = htmlString + "</nav>"

	return template.HTML(htmlString)
}
