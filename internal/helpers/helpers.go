package helpers

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

const (
	defaultLimit = 50
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type QueryParams struct {
	Offset int
	Limit  int
}

type HttpResp struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func GetQueryParams(offset, limit string) (*QueryParams, error) {
	var (
		off, lim int
		err      error
	)
	if offset != "" {
		off, err = strconv.Atoi(offset)
		if err != nil {
			return nil, err
		}
	}
	if limit == "" {
		lim = defaultLimit
	} else {
		lim, err = strconv.Atoi(limit)
		if err != nil {
			return nil, err
		}
	}
	return &QueryParams{
		Offset: off,
		Limit:  lim,
	}, nil
}

func SortWord(word string) string {
	splitWord := strings.Split(word, "")
	sort.Strings(splitWord)
	return strings.Join(splitWord, "")
}

func GenerateHttpErrorResp(w http.ResponseWriter, err error) error {
	resp, err := GenerateHttpResp(false, err, nil)
	if err != nil {
		return err
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return err
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	return nil
}

func GenerateHttpResp(success bool, err error, data interface{}) ([]byte, error) {
	httpResp := HttpResp{
		Success: success,
		Data:    data,
	}
	if err != nil {
		httpResp.Error = err.Error()
	}
	return json.Marshal(httpResp)
}
