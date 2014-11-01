package kdtgo

import (
	"fmt"
	"sort"
	"time"
)

func NowString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func TodayStartString() string {
	return time.Now().Format("2006-01-02") + " 00:00:00"
}

func TodayEndString() string {
	return time.Now().AddDate(0, 0, 1).Format("2006-01-02") + " 00:00:00"
}

type signMap struct {
	noSort     map[string]string
	sortedKeys []string
	isSorted   bool
}

func NewSign(_map map[string]string) *signMap {
	return &signMap{
		noSort:     _map,
		sortedKeys: make([]string, 0, len(_map)),
		isSorted:   false,
	}
}

func (sm *signMap) SortKeys() *signMap {
	if sm.isSorted {
		return sm
	}

	for k := range sm.noSort {
		sm.sortedKeys = append(sm.sortedKeys, k)
	}

	sort.Strings(sm.sortedKeys)

	sm.isSorted = true

	return sm
}

func (sm *signMap) ToString(separator ...string) string {
	if !sm.isSorted {
		sm.SortKeys()
	}

	var ret string

	_separator := ""
	if len(separator) >= 1 {
		_separator = separator[0]
	}

	for _, k := range sm.sortedKeys {
		ret += fmt.Sprintf("%s"+_separator+"%s", k, sm.noSort[k])
	}

	return ret
}

func (sm *signMap) BracketToString(bracket string, separator ...string) string {
	s := sm.ToString(separator...)
	return bracket + s + bracket
}
