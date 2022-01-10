package kwords

import (
	"runtime"
	"sync"

	"github.com/jfcg/sorty/v2"
)

type OccurringWord struct {
	Word string
	N    int
}

func KWords(k int, words []string) []OccurringWord {
	n := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(n)

	// split the arr and count in multiple go routines
	c := make(chan map[string]int)
	for i, l := 0, len(words); i < n; i++ {
		r := l / n
		begin, end := i*r, (i+1)*r
		if i == n-1 {
			end = l
		}
		go func() {
			m := map[string]int{}
			for i := begin; i < end; i++ {
				w := words[i]
				m[w] = m[w] + 1
			}
			c <- m
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(c)
	}()

	// merge the results from go routines
	m := map[string]int{}
	for mc := range c {
		for w, n := range mc {
			m[w] += n
		}
	}

	// build array from map
	a, i := make([]OccurringWord, len(m)), 0
	for w, n := range m {
		a[i] = OccurringWord{w, n}
		i++
	}
	if i < k {
		k = i
	}

	// sort the array descending
	sorty.Sort(i, func(i, k, r, s int) bool {
		if a[i].N > a[k].N {
			if r != s {
				a[r], a[s] = a[s], a[r]
			}
			return true
		}
		return false
	})

	// take the first k most occurring words
	return a[:k]
}
