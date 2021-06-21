package common

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetStatusCode(err error) int {
	if err != nil {
		logrus.Error(err)
	}
	switch err {
	case INTERNAL_SERVER_ERROR:
		return http.StatusInternalServerError
	case NOT_FOUND_ERROR:
		return http.StatusNotFound
	case CONFLIT_ERROR:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
