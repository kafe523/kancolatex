package access

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type ErrInvalidIndex struct {
	dataIdx int
}

func (e *ErrInvalidIndex) Error() string {
	if e.dataIdx < 0 {
		return fmt.Sprintf("dataIdx %d should greater than 0", e.dataIdx)
	}
	return fmt.Sprintf("dataPtr %d out of range", e.dataIdx)
}

func accessReflex(dataPtr *any, acStr string) (any, error) {
	var dataElem reflect.Value
	dataElem = reflect.ValueOf(dataPtr)

	if dataElem.Kind() == reflect.Pointer {
		dataElem = reflect.ValueOf(*dataPtr)
		if dataElem.Kind() == reflect.Interface || dataElem.Kind() == reflect.Pointer {
			dataElem = dataElem.Elem()
		}
	}

	switch dataElem.Kind() {
	case reflect.Struct:
		for i := 0; i < dataElem.NumField(); i++ {
			dataElemTypeField := dataElem.Type().Field(i)
			if dataElemTypeField.Name == acStr {
				result := dataElem.FieldByName(acStr)
				if result.IsValid() {
					return result.Interface(), nil
				}
			}
		}

	case reflect.Map:
		result := dataElem.MapIndex(reflect.ValueOf(acStr))
		if result.IsValid() {
			return result.Interface(), nil
		}

	case reflect.Array:
		fallthrough
	case reflect.Slice:
		idx, err := strconv.Atoi(acStr)
		if err != nil {
			return nil, err
		}

		if idx < 0 {
			return nil, &ErrInvalidIndex{dataIdx: idx}
		}

		if idx >= dataElem.Len() {
			return nil, &ErrInvalidIndex{dataIdx: idx}
		}

		result := dataElem.Index(idx)
		if result.IsValid() {
			return result.Interface(), nil
		}
	}

	return nil, nil
}

func Access(data any, pat []string) any {
	var result any
	var err error

	var cur any = data

	for _, k := range pat {
		result, err = accessReflex(&cur, k)

		if err != nil {
			var errIdx *ErrInvalidIndex
			if errors.As(err, &errIdx) {
				return ""
			}
			panic(err)
		}

		if result == nil {
			return nil
		}

		switch reflect.TypeOf(result).Kind() {
		case reflect.Array:
			fallthrough
		case reflect.Slice:
			fallthrough
		case reflect.Map:
			fallthrough
		case reflect.Struct:
			cur = result
			continue
		}

		break
	}

	return result
}

type tokenizerState struct {
	value string
	start int
	cur   int
}

func (s *tokenizerState) IsEnd() bool { return s.cur >= len(s.value) }

func (s *tokenizerState) Advance() { s.cur += 1 }

func (s *tokenizerState) Consume() { s.start = s.cur }

// return segment string, boolean indicate the end of lookup, please break when true.
func (s *tokenizerState) Next() (string, bool) {
	var segment string
	var stop bool

	if s.start >= len(s.value) {
		stop = true
		return "", stop
	}

	for !s.IsEnd() {
		if s.value[s.cur] == '.' {
			segment = s.value[s.start:s.cur]
			s.Consume()
			s.Advance()
			s.Consume()
			return segment, false
		}

		s.Advance()
	}

	if s.IsEnd() {
		segment = s.value[s.start:s.cur]
		s.Consume()
		stop = true
	}

	return segment, stop
}

func PatternTokenizer(patStr string) []string {
	var result []string

	state := tokenizerState{patStr, 0, 0}

	for {
		r, stop := state.Next()
		result = append(result, strings.TrimSpace(r))

		if stop {
			break
		}
	}

	return result
}
