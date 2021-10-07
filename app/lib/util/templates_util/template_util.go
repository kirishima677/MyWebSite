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

	previousOffset := offset - limit
	if previousOffset < 0 {
		previousOffset = 0
	}

	nextOffset := offset + limit

	htmlString := "<nav aria-label=\"Page navigation example\">"
	htmlString = htmlString + "<ul class=\"pagination justify-content-center\">"

	if offset == 0 {
		htmlString = htmlString + "<li class=\"page-item disabled\">\n<a class=\"page-link\" href=\"#\" tabindex=\"-1\">Previous</a>\n</li>"
	} else {
		htmlString = htmlString + fmt.Sprintf("<li class=\"page-item\">\n<a class=\"page-link\" href=\"%s?offset=%d&limit=%d\" tabindex=\"-1\">Previous</a>\n</li>",
			path, previousOffset, limit)
	}

	pageNumber := int(math.Ceil(float64(rowCount / limit)))
	nowPageNumber := int(math.Ceil(float64(offset/limit))) + 1
	pageOffset := 0
	for i := 0; i <= pageNumber; i++ {
		disabled := ""
		if (i + 1) == nowPageNumber {
			disabled = "disabled"
		}

		htmlString = htmlString + fmt.Sprintf(""+
			"<li class=\"page-item %s\">"+
			"<a class=\"page-link\" href=\"%s?offset=%d&amp;limit=%d\">%d</a>"+
			"</li>",
			disabled, path, pageOffset, limit, (i+1))
		pageOffset = pageOffset + limit
	}

	if rowCount <= limit+offset {
		htmlString = htmlString + "<li class=\"page-item disabled\">\n<a class=\"page-link\" href=\"#\">Next</a>\n</li>"
	} else {
		htmlString = htmlString + fmt.Sprintf("<li class=\"page-item\">\n<a class=\"page-link\" href=\"/admin/articles/list?offset=%d&limit=%d\">Next</a>\n</li>", nextOffset, limit)
	}

	htmlString = htmlString + "</ul>"
	htmlString = htmlString + "</nav>"

	return template.HTML(htmlString)
}
