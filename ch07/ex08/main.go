package main

import (
	"fmt"
	"sort"
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


type StableTrackMultiKeysSorter struct {
	tracks    []*Track
	lessFuncs []LessFunc
	lessIndex int
}

func (t *StableTrackMultiKeysSorter) AddSortKey(key LessFunc) {
	t.lessFuncs = append(t.lessFuncs, key)
	t.lessIndex++
}

func (t *StableTrackMultiKeysSorter) Len() int {
	return len(t.tracks)
}

func (t *StableTrackMultiKeysSorter) Swap(i, j int) {
	t.tracks[i], t.tracks[j] = t.tracks[j], t.tracks[i]
}

func (t *StableTrackMultiKeysSorter) Less(i, j int) bool {
	if t.lessIndex < 0 {
		panic(fmt.Errorf("Out of Index: %d", t.lessIndex))
	}
	return t.lessFuncs[t.lessIndex](t.tracks[i], t.tracks[j])
}

func (t *StableTrackMultiKeysSorter) HasNext() bool {
	t.lessIndex--
	return t.lessIndex >= 0
}


func main() {
	fmt.Println("origin")
	printTracks(tracks)

	sortNormal()
	sortStable()
}

var title = func(p, q interface{}) bool {
	tp := p.(*Track)
	tq := q.(*Track)
	return tp.Title < tq.Title
}

var year = func(p, q interface{}) bool {
	tp := p.(*Track)
	tq := q.(*Track)
	return tp.Year < tq.Year
}

func sortNormal() {
	fmt.Println("\n=== sort.Sort ===")

	d := make([]*Track, len(tracks))
	copy(d, tracks)
	table := TrackMultiKeysSorter{tracks: d}

	table.AddSortKey(title)
	table.AddSortKey(year)
	sort.Sort(&table)
	printTracks(d)
}

func sortStable() {
	fmt.Println("\n=== sort.Stable ===")

	d := make([]*Track, len(tracks))
	copy(d, tracks)
	table := StableTrackMultiKeysSorter{tracks: d}

	table.AddSortKey(title)
	table.AddSortKey(year)

	for table.HasNext() {
		sort.Stable(&table)
	}
	printTracks(d)
}

