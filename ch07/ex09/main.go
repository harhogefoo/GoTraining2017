package main

import (
	"sort"
	"net/http"
	"strings"
	"fmt"
)

type LessFunc func(p1, p2 interface{}) bool

type MultiKeysSorter struct {
	lessFuncs []LessFunc
}

func (m *MultiKeysSorter) AddSortKey(key LessFunc) {
	m.lessFuncs = append(m.lessFuncs, key)
}

func (m *MultiKeysSorter) LessWithMultiKeys(p, q interface{}) bool {
	if len(m.lessFuncs) == 0 {
		panic("Not Key is added as LessFunc")
	}

	var k int
	for k = 0; k < len(m.lessFuncs)-1; k++ {
		less := m.lessFuncs[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return m.lessFuncs[k](p, q)
}

type TrackMultiKeysSorter struct {
	tracks []*Track
	MultiKeysSorter
}

func (t *TrackMultiKeysSorter) Len() int {
	return len(t.tracks)
}

func (t *TrackMultiKeysSorter) Swap(i, j int) {
	t.tracks[i], t.tracks[j] = t.tracks[j], t.tracks[i]
}

func (t *TrackMultiKeysSorter) Less(i, j int) bool {
	return t.LessWithMultiKeys(t.tracks[i], t.tracks[j])
}


func main() {
	sortNormal()
}

var titleKey = func(p, q interface{}) bool {
	tp := p.(*Track)
	tq := q.(*Track)
	return tp.Title < tq.Title
}

var artistKey = func(p, q interface{}) bool {
	tp := p.(*Track)
	tq := q.(*Track)
	return tp.Artist < tq.Artist
}

var albumKey = func(p, q interface{}) bool {
	tp := p.(*Track)
	tq := q.(*Track)
	return tp.Album < tq.Album
}

var yearKey = func(p, q interface{}) bool {
	tp := p.(*Track)
	tq := q.(*Track)
	return tp.Year < tq.Year
}

var lengthKey = func(p, q interface{}) bool {
	tp := p.(*Track)
	tq := q.(*Track)
	return tp.Length < tq.Length
}

var sortKeyFuncs = map[string]func(p, q interface{}) bool{
	"TITLE":  titleKey,
	"ARTIST": artistKey,
	"ALBUM":  albumKey,
	"YEAR":   yearKey,
	"LENGTH": lengthKey,
}

func sortNormal() {
	d := make([]*Track, len(tracks))
	copy(d, tracks)

	handler := func(w http.ResponseWriter, r *http.Request) {
		value := r.URL.Query().Get("sort")
		f, keys := sortOrders(strings.Split(value, ","))
		if len(f) == 0 {
			printTracksHTML(w, d, keys)
			return
		}

		table := TrackMultiKeysSorter{tracks: d}
		for _, sk := range f {
			table.AddSortKey(sk)
		}
		sort.Sort(&table)
		printTracksHTML(w, d, keys)
	}

	uri := "localhost:8001"
	fmt.Printf("[OPEN] %s\n", uri)
	http.HandleFunc("/", handler)
	http.ListenAndServe(uri, nil)
}

func sortOrders(keys []string) ([]func(p, q interface{}) bool, []string) {
	updatedKeys := make([]string, 0, len(keys))
	sortOrderFuncs := make([]func(p, q interface{}) bool, 0, len(keys))

	for _, key := range keys {
		key := strings.TrimSpace(key)
		f, ok := sortKeyFuncs[strings.ToUpper(key)]
		if ok {
			updatedKeys = append(updatedKeys, key)
			sortOrderFuncs = append(sortOrderFuncs, f)
		}
	}
	return sortOrderFuncs, updatedKeys
}
