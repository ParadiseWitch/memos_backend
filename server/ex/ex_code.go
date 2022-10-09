package ex

import (
	"memos/server/dto"
	"net/http"
)

type ExCode struct {
	LogLevel dto.Level
	HttpCode int
	LogMsg   string
	TipMsg   string
}

func EC_INVALID_PARAM() ExCode {
	return ExCode{
		LogLevel: dto.LOGLEVEL_WARN,
		HttpCode: http.StatusBadRequest,
		LogMsg:   "invalid param",
		TipMsg:   "无效的参数",
	}
}

func EC_DATA_NOT_EXIST(s string) ExCode {
	return ExCode{
		LogLevel: dto.LOGLEVEL_WARN,
		HttpCode: http.StatusOK,
		LogMsg:   "data not exist",
		TipMsg:   s + "不存在",
	}
}

func EC_DB_OP_ERROR() ExCode {
	return ExCode{
		LogLevel: dto.LOGLEVEL_ERROR,
		HttpCode: http.StatusInternalServerError,
		LogMsg:   "database operation error",
		TipMsg:   "服务器错误",
	}
}

func EC_PSD_MISTAKE() ExCode {
	return ExCode{
		LogLevel: dto.LOGLEVEL_WARN,
		HttpCode: http.StatusOK,
		LogMsg:   "password mistake",
		TipMsg:   "密码错误",
	}
}

func EC_USER_HAS_REGISTERED() ExCode {
	return ExCode{
		LogLevel: dto.LOGLEVEL_WARN,
		HttpCode: http.StatusOK,
		LogMsg:   "the user has registered",
		TipMsg:   "账号已被注册",
	}
}
