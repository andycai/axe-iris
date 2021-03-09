package e

var codeText = map[int]string{
	CodeOk:                 "OK",
	CodeParamErr:           "参数错误",
	CodeNotFoundData:       "没查询到数据",
	CodeOpErr:              "操作错误",
}

func CodeText(code int) string {
	return codeText[code]
}
