package http

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type Handler interface {
	WrapFunc(func(Handler)) http.HandlerFunc
	Response(status int, response interface{})
	Ok(response interface{})
	BadRequest(response interface{})
	Error()
	ParseJson(obj interface{}) (error interface{})
	GetParam(param string) string
	GetRequestId() string
	GetLogger() *logrus.Entry
	GetDB() *gorm.DB
}

// handler helper
type handler struct {
	logger    *logrus.Entry
	w         http.ResponseWriter
	r         *http.Request
	requestID string
	db        *gorm.DB
}

func NewWebHandler(db *gorm.DB) Handler {
	return &handler{db: db}
}

func (h *handler) WrapFunc(f func(h Handler)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(&handler{
			logger:    r.Context().Value(LoggerID).(*logrus.Entry),
			w:         w,
			r:         r,
			requestID: r.Context().Value(RequestID).(string),
			db:        h.db,
		})
	}
}

// Answer
func (h *handler) Response(status int, response interface{}) {
	jsonStr, _ := json.MarshalIndent(struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data"`
	}{status, response}, "", "\t")
	h.w.WriteHeader(status)
	h.w.Write(jsonStr) // nolint:errcheck
}

func (h *handler) Ok(response interface{}) {
	h.Response(http.StatusOK, response)
}

func (h *handler) Error() {
	h.Response(http.StatusInternalServerError, "INTERNAL_ERROR")
}

func (h *handler) BadRequest(response interface{}) {
	h.Response(http.StatusBadRequest, response)
}

// Unmarshal JSON request
func (h *handler) ParseJson(obj interface{}) interface{} {
	input, err := ioutil.ReadAll(h.r.Body)
	if err != nil {
		return "NOT_JSON"
	}

	// Unmarshal
	err = json.Unmarshal(input, &obj)
	if err != nil {
		return "NOT_JSON"
	}

	return nil
}

func (h *handler) GetParam(param string) string {
	return h.r.URL.Query().Get(param)
}

// RequestId
func (h *handler) GetRequestId() string {
	return h.requestID
}

// Logger
func (h *handler) GetLogger() *logrus.Entry {
	return h.logger
}

// DB
func (h *handler) GetDB() *gorm.DB {
	return h.db
}
