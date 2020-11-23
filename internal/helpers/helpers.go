package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

const (
	KeyContentType      = "Content-Type"
	JSONContentType     = "application/json"
	defaultLimit        = 50
	maxWordLength       = 100
	wordTooLengthFormat = "word: %s is too length. max length is %d"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type QueryParams struct {
	Offset int
	Limit  int
}

type HTTPErrorResponse struct {
	Error string `json:"error"`
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

func SortWord(word string) (string, error) {
	if len(word) > maxWordLength {
		return "", fmt.Errorf(wordTooLengthFormat, word, maxWordLength)
	}
	splitWord := strings.Split(word, "")
	sort.Strings(splitWord)
	return strings.Join(splitWord, ""), nil
}

func GenerateHTTPErrorResp(w http.ResponseWriter, err error) error {
	rsp := HTTPErrorResponse{
		Error: err.Error(),
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Add(KeyContentType, JSONContentType)
	err = json.NewEncoder(w).Encode(rsp)
	if err != nil {
		return err
	}
	return nil
}

func LogHTTPError(logger logrus.FieldLogger, r *http.Request, err error) {
	logger.WithFields(logrus.Fields{
		"uri":    r.URL.Path,
		"method": r.Method,
	}).Error(err.Error())
}
