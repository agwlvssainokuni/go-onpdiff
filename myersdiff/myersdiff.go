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

package myersdiff

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
	x    int
	prev *path
}

func MyersDiff(a []interface{}, b []interface{}, equals func(interface{}, interface{}) bool) *DiffInfo {

	N := len(a)
	M := len(b)

	// ////////////////////////////////////////////////////////
	// Myersアルゴリズム本体：ここから
	MAX := N + M
	V := make([]*path, MAX+MAX+1)
	offset := MAX

	V[1+offset] = &path{1, 0, nil}
	edist := -1
LOOP:
	for D := 0; D <= MAX; D += 1 {
		for k := -D; k <= D; k += 2 {
			var (
				prev *path
				x    int
			)
			if k <= -D || (k < D && V[k-1+offset].x < V[k+1+offset].x) {
				// 左端またはk+1先行：k + 1から下へ
				prev = V[k+1+offset]
				x = prev.x
			} else {
				// 右端またはk-1先行：k - 1から右へ
				prev = V[k-1+offset]
				x = prev.x + 1
			}
			y := x - k
			for x < N && y < M && equals(a[x+1-1], b[y+1-1]) {
				x += 1
				y += 1
			}
			V[k+offset] = &path{k, x, prev}
			if x >= N && y >= M {
				edist = D
				break LOOP
			}
		}
	}
	// Myersアルゴリズム本体：ここまで
	// ////////////////////////////////////////////////////////

	// リストの先頭から見られるよう並べ直す。
	// ※Edit Graphの末端は「k=N-M」なので、そこから逆順に辿る。
	l := make([]*path, 0)
	for pt := V[N-M+offset]; pt.prev != nil; pt = pt.prev {
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
		if x-y < pt.k {
			// 差分の経路(k線)は現在位置よりも右側：削除
			t := DiffDel
			for x-y < pt.k {
				ses = append(ses, &DiffElem{t, a[x]})
				x += 1
			}
		}
		if x-y > pt.k {
			// 差分の経路(k線)は現在位置よりも左側：追加
			t := DiffAdd
			for x-y > pt.k {
				ses = append(ses, &DiffElem{t, b[y]})
				y += 1
			}
		}
		// 差分の経路(k線)上で進めるだけ進む。
		for x < pt.x {
			ses = append(ses, &DiffElem{DiffSame, b[y]})
			lcs = append(lcs, b[y])
			x += 1
			y += 1
		}
	}

	return &DiffInfo{edist, ses, lcs}
}
