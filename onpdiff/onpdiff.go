/*
 * Copyright 2018 agwlvssainokuni
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package onpdiff

const (
	DiffSame = iota
	DiffAdd
	DiffDel
)

type DiffElem struct {
	Type  int
	Value interface{}
}

type DiffInfo struct {
	Edist int
	Ses   []*DiffElem
	Lcs   []interface{}
}

type path struct {
	k    int
	y    int
	prev *path
}

func OnpDiff(a []interface{}, b []interface{}, equals func(interface{}, interface{}) bool) *DiffInfo {
	m := len(a)
	n := len(b)
	if m <= n {
		return doOnpDiff(a, b, m, n, equals, true)
	} else {
		return doOnpDiff(b, a, n, m, equals, false)
	}
}

func doOnpDiff(a []interface{}, b []interface{}, m int, n int, equals func(interface{}, interface{}) bool, normal bool) *DiffInfo {
	maxAndSnake := funcMaxAndSnake(a, b, m, n, equals, normal)

	// ////////////////////////////////////////////////////////
	// ON(NP)アルゴリズム本体：ここから
	offset := m + 1
	fp := make([]*path, (m+1)+(n+1)+1)
	for k := -(m + 1); k <= (n + 1); k++ {
		fp[k+offset] = &path{k, -1, nil}
	}

	delta := n - m
	p := -1
	for {
		p += 1
		for k := -p; k < delta; k++ {
			fp[k+offset] = maxAndSnake(k, fp[k-1+offset], fp[k+1+offset])
		}
		for k := delta + p; k > delta; k-- {
			fp[k+offset] = maxAndSnake(k, fp[k-1+offset], fp[k+1+offset])
		}
		fp[delta+offset] = maxAndSnake(delta, fp[delta-1+offset], fp[delta+1+offset])
		if fp[delta+offset].y >= n {
			break
		}
	}
	edist := delta + 2*p
	// ON(NP)アルゴリズム本体：ここまで
	// ////////////////////////////////////////////////////////

	// リストの先頭から見られるよう並べ直す。
	l := make([]*path, 0)
	for pt := fp[delta+offset]; pt.y >= 0; pt = pt.prev {
		l = append(l, pt)
	}

	// 差分の算出結果として以下の2点を導出する。
	// ・Shortest Edit Script
	// ・Longest Common Sequence
	ses := make([]*DiffElem, 0)
	lcs := make([]interface{}, 0)
	x := 0
	y := 0
	for i := len(l) - 1; i >= 0; i-- {
		pt := l[i]
		if y-x < pt.k {
			// 差分の経路(k線)は現在位置よりも右側：追加
			t := DiffAdd
			if !normal {
				t = DiffDel
			}
			for y-x < pt.k {
				ses = append(ses, &DiffElem{t, b[y]})
				y += 1
			}
		}
		if y-x > pt.k {
			// 差分の経路(k線)は現在位置よりも左側：削除
			t := DiffDel
			if !normal {
				t = DiffAdd
			}
			for y-x > pt.k {
				ses = append(ses, &DiffElem{t, a[x]})
				x += 1
			}
		}
		// 差分の経路(k線)上で進めるだけ進む。
		for y < pt.y {
			ses = append(ses, &DiffElem{DiffSame, b[y]})
			lcs = append(lcs, b[y])
			x += 1
			y += 1
		}
	}

	return &DiffInfo{edist, ses, lcs}
}

func funcMaxAndSnake(a []interface{}, b []interface{}, m int, n int, equals func(interface{}, interface{}) bool, normal bool) func(int, *path, *path) *path {
	return func(k int, pt1 *path, pt2 *path) *path {

		// ON(NP)アルゴリズム：max
		var (
			y  int
			pt *path
		)
		if pt1.y+1 == pt2.y {
			// k-1とk+1の両方から合流する場合は「削除、追加」の順になるよう選択する。
			if normal {
				y = pt1.y + 1
				pt = pt1
			} else {
				y = pt2.y
				pt = pt2
			}
		} else if pt1.y+1 > pt2.y {
			// k-1
			y = pt1.y + 1
			pt = pt1
		} else {
			// k+1
			y = pt2.y
			pt = pt2
		}

		// ON(NP)アルゴリズム：snake
		x := y - k
		for x < m && y < n && equals(a[x], b[y]) {
			x += 1
			y += 1
		}
		return &path{k, y, pt}
	}
}
