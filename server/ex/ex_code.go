package ex

import (
	"memos/server/dto"
	"net/http"
)

type ExCode struct {
	LogLevel dto.Level
	HttpCode int
	Msg      string
	Handler  func()
}

var EC_INVALID_PARAM = &ExCode{
	LogLevel: dto.LOGLEVEL_WARN,
	HttpCode: http.StatusBadRequest,
	Msg:      "invalid param",
}

var EC_DATA_NOT_EXIST = &ExCode{
	LogLevel: dto.LOGLEVEL_WARN,
	HttpCode: http.StatusOK,
	Msg:      "data not exist",
}

var EC_DB_OP_ERROR = &ExCode{
	LogLevel: dto.LOGLEVEL_ERROR,
	HttpCode: http.StatusInternalServerError,
	Msg:      "database operation error",
}

var EC_PSD_MISTAKE = &ExCode{
	LogLevel: dto.LOGLEVEL_WARN,
	HttpCode: http.StatusOK,
	Msg:      "password mistake",
}

var EC_USER_HAS_REGISTERED = &ExCode{
	LogLevel: dto.LOGLEVEL_WARN,
	HttpCode: http.StatusOK,
	Msg:      "the user has registered",
}
