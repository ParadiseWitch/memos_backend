package ex

type HttpEX struct {
	e         error
	DetailMsg string
	Code      *ExCode
}

func New(e error, code *ExCode, detailMsg string) *HttpEX {
	nativeErrorMsg := "nil"
	if e != nil {
		nativeErrorMsg = e.Error()
	}
	return &HttpEX{
		e:         e,
		Code:      code,
		DetailMsg: detailMsg + ", error: " + nativeErrorMsg,
	}
}
