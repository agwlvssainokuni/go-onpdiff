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

package wudiff

import (
	"fmt"
	"testing"
)

func idiff(a []interface{}, b []interface{}) *DiffInfo {
	return WuDiff(a, b, func(ia interface{}, ib interface{}) bool {
		return ia.(int) == ib.(int)
	})
}

// 同一
func Test同一シーケンス(t *testing.T) {
	a := []interface{}{0, 1, 2, 3, 4, 5, 6}
	b := []interface{}{0, 1, 2, 3, 4, 5, 6}
	result := idiff(a, b)
	assertInfo(t, "Test同一シーケンス", result, //
		0, // edist
		[]int{0, 1, 2, 3, 4, 5, 6},                                                  //lcs
		[]int{DiffSame, DiffSame, DiffSame, DiffSame, DiffSame, DiffSame, DiffSame}) // ses
}

// 追加
// ・先頭
// ・末尾
// ・途中

func Test追加_先頭_1要素(t *testing.T) {
	a := []interface{}{0, 1, 2, 3, 4, 5, 6}
	b := []interface{}{7, 0, 1, 2, 3, 4, 5, 6}
	result := idiff(a, b)
	assertInfo(t, "Test追加_先頭_1要素", result, //
		1, // edist
		[]int{0, 1, 2, 3, 4, 5, 6},                                                           //lcs
		[]int{DiffAdd, DiffSame, DiffSame, DiffSame, DiffSame, DiffSame, DiffSame, DiffSame}) // ses
}

func Test追加_先頭_3要素(t *testing.T) {
	a := []interface{}{0, 1, 2, 3, 4, 5, 6}
	b := []interface{}{9, 8, 7, 0, 1, 2, 3, 4, 5, 6}
	result := idiff(a, b)
	assertInfo(t, "Test追加_先頭_3要素", result, //
		3, // edist
		[]int{0, 1, 2, 3, 4, 5, 6},                                                                             //lcs
		[]int{DiffAdd, DiffAdd, DiffAdd, DiffSame, DiffSame, DiffSame, DiffSame, DiffSame, DiffSame, DiffSame}) // ses
}

func Test追加_末尾_1要素(t *testing.T) {
	a := []interface{}{0, 1, 2, 3, 4, 5, 6}
	b := []interface{}{0, 1, 2, 3, 4, 5, 6, 7}
	result := idiff(a, b)
	assertInfo(t, "Test追加_末尾_1要素", result, //
		1, // edist
		[]int{0, 1, 2, 3, 4, 5, 6},                                                           //lcs
		[]int{DiffSame, DiffSame, DiffSame, DiffSame, DiffSame, DiffSame, DiffSame, DiffAdd}) // ses
}

func Test追加_末尾_3要素(t *testing.T) {
	a := []interface{}{0, 1, 2, 3, 4, 5, 6}
	b := []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := idiff(a, b)
	assertInfo(t, "Test追加_末尾_3要素", result, //
		3, // edist
		[]int{0, 1, 2, 3, 4, 5, 6},                                                                             //lcs
		[]int{DiffSame, DiffSame, DiffSame, DiffSame, DiffSame, DiffSame, DiffSame, DiffAdd, DiffAdd, DiffAdd}) // ses
}

func Test追加_途中_1要素(t *testing.T) {
	a := []interface{}{0, 1, 2, 3, 4, 5, 6}
	b := []interface{}{0, 1, 2, 7, 3, 4, 5, 6}
	result := idiff(a, b)
	assertInfo(t, "Test追加_途中_1要素", result, //
		1, // edist
		[]int{0, 1, 2, 3, 4, 5, 6},                                                           //lcs
		[]int{DiffSame, DiffSame, DiffSame, DiffAdd, DiffSame, DiffSame, DiffSame, DiffSame}) // ses
}

func Test追加_途中_3要素(t *testing.T) {
	a := []interface{}{0, 1, 2, 3, 4, 5, 6}
	b := []interface{}{0, 1, 2, 7, 8, 9, 3, 4, 5, 6}
	result := idiff(a, b)
	assertInfo(t, "Test追加_途中_3要素", result, //
		3, // edist
		[]int{0, 1, 2, 3, 4, 5, 6},                                                                             //lcs
		[]int{DiffSame, DiffSame, DiffSame, DiffAdd, DiffAdd, DiffAdd, DiffSame, DiffSame, DiffSame, DiffSame}) // ses
}

// 削除
// ・先頭
// ・末尾
// ・途中

func Test削除_先頭_1要素(t *testing.T) {
	a := []interface{}{0, 1, 2, 3, 4, 5, 6}
	b := []interface{}{1, 2, 3, 4, 5, 6, 9}
	result := idiff(a, b)
	assertInfo(t, "Test削除_先頭_1要素", result, //
		1+1, // edist
		[]int{1, 2, 3, 4, 5, 6},                                                             //lcs
		[]int{DiffDel, DiffSame, DiffSame, DiffSame, DiffSame, DiffSame, DiffSame, DiffAdd}) // ses
}

func Test削除_先頭_3要素(t *testing.T) {
	a := []interface{}{0, 1, 2, 3, 4, 5, 6}
	b := []interface{}{3, 4, 5, 6, 9, 9, 9}
	result := idiff(a, b)
	assertInfo(t, "Test削除_先頭_3要素", result, //
		3+3,               // edist
		[]int{3, 4, 5, 6}, //lcs
		[]int{DiffDel, DiffDel, DiffDel, DiffSame, DiffSame, DiffSame, DiffSame, DiffAdd, DiffAdd, DiffAdd}) // ses
}

func Test削除_末尾_1要素(t *testing.T) {
	a := []interface{}{0, 1, 2, 3, 4, 5, 6}
	b := []interface{}{9, 0, 1, 2, 3, 4, 5}
	result := idiff(a, b)
	assertInfo(t, "Test削除_末尾_1要素", result, //
		1+1, // edist
		[]int{0, 1, 2, 3, 4, 5},                                                             //lcs
		[]int{DiffAdd, DiffSame, DiffSame, DiffSame, DiffSame, DiffSame, DiffSame, DiffDel}) // ses
}

func Test削除_末尾_3要素(t *testing.T) {
	a := []interface{}{0, 1, 2, 3, 4, 5, 6}
	b := []interface{}{9, 9, 9, 0, 1, 2, 3}
	result := idiff(a, b)
	assertInfo(t, "Test削除_末尾_3要素", result, //
		3+3,               // edist
		[]int{0, 1, 2, 3}, //lcs
		[]int{DiffAdd, DiffAdd, DiffAdd, DiffSame, DiffSame, DiffSame, DiffSame, DiffDel, DiffDel, DiffDel}) // ses
}

func Test削除_途中_1要素(t *testing.T) {
	a := []interface{}{0, 1, 2, 3, 4, 5, 6}
	b := []interface{}{0, 1, 2, 4, 5, 6, 9}
	result := idiff(a, b)
	assertInfo(t, "Test削除_途中_1要素", result, //
		1+1, // edist
		[]int{0, 1, 2, 4, 5, 6},                                                             //lcs
		[]int{DiffSame, DiffSame, DiffSame, DiffDel, DiffSame, DiffSame, DiffSame, DiffAdd}) // ses
}

func Test削除_途中_3要素(t *testing.T) {
	a := []interface{}{0, 1, 2, 3, 4, 5, 6}
	b := []interface{}{0, 1, 5, 6, 9, 9, 9}
	result := idiff(a, b)
	assertInfo(t, "Test削除_途中_3要素", result, //
		3+3,               // edist
		[]int{0, 1, 5, 6}, //lcs
		[]int{DiffSame, DiffSame, DiffDel, DiffDel, DiffDel, DiffSame, DiffSame, DiffAdd, DiffAdd, DiffAdd}) // ses
}

// 変更
// ・先頭
// ・末尾
// ・途中

func Test変更_先頭_1要素(t *testing.T) {
	a := []interface{}{0, 1, 2, 3, 4, 5, 6}
	b := []interface{}{7, 1, 2, 3, 4, 5, 6}
	result := idiff(a, b)
	assertInfo(t, "Test変更_先頭_1要素", result, //
		2, // edist
		[]int{1, 2, 3, 4, 5, 6},                                                             //lcs
		[]int{DiffDel, DiffAdd, DiffSame, DiffSame, DiffSame, DiffSame, DiffSame, DiffSame}) // ses
}

func Test変更_先頭_3要素(t *testing.T) {
	a := []interface{}{0, 1, 2, 3, 4, 5, 6}
	b := []interface{}{7, 8, 9, 3, 4, 5, 6}
	result := idiff(a, b)
	assertInfo(t, "Test変更_先頭_3要素", result, //
		6,                 // edist
		[]int{3, 4, 5, 6}, //lcs
		[]int{DiffDel, DiffDel, DiffDel, DiffAdd, DiffAdd, DiffAdd, DiffSame, DiffSame, DiffSame, DiffSame}) // ses
}

func Test変更_末尾_1要素(t *testing.T) {
	a := []interface{}{0, 1, 2, 3, 4, 5, 6}
	b := []interface{}{0, 1, 2, 3, 4, 5, 7}
	result := idiff(a, b)
	assertInfo(t, "Test変更_末尾_1要素", result, //
		2, // edist
		[]int{0, 1, 2, 3, 4, 5},                                                             //lcs
		[]int{DiffSame, DiffSame, DiffSame, DiffSame, DiffSame, DiffSame, DiffDel, DiffAdd}) // ses
}

func Test変更_末尾_3要素(t *testing.T) {
	a := []interface{}{0, 1, 2, 3, 4, 5, 6}
	b := []interface{}{0, 1, 2, 3, 7, 8, 9}
	result := idiff(a, b)
	assertInfo(t, "Test変更_末尾_3要素", result, //
		6,                 // edist
		[]int{0, 1, 2, 3}, //lcs
		[]int{DiffSame, DiffSame, DiffSame, DiffSame, DiffDel, DiffDel, DiffDel, DiffAdd, DiffAdd, DiffAdd}) // ses
}

func Test変更_途中_1要素(t *testing.T) {
	a := []interface{}{0, 1, 2, 3, 4, 5, 6}
	b := []interface{}{0, 1, 2, 7, 4, 5, 6}
	result := idiff(a, b)
	assertInfo(t, "Test変更_途中_1要素", result, //
		2, // edist
		[]int{0, 1, 2, 4, 5, 6},                                                             //lcs
		[]int{DiffSame, DiffSame, DiffSame, DiffDel, DiffAdd, DiffSame, DiffSame, DiffSame}) // ses
}

func Test変更_途中_3要素(t *testing.T) {
	a := []interface{}{0, 1, 2, 3, 4, 5, 6}
	b := []interface{}{0, 1, 7, 8, 9, 5, 6}
	result := idiff(a, b)
	assertInfo(t, "Test変更_途中_3要素", result, //
		6,                 // edist
		[]int{0, 1, 5, 6}, //lcs
		[]int{DiffSame, DiffSame, DiffDel, DiffDel, DiffDel, DiffAdd, DiffAdd, DiffAdd, DiffSame, DiffSame}) // ses
}

// AB入れ替え実行の場合

func TestAB入れ替え実行の場合(t *testing.T) {
	a := []interface{}{0, 1, 2, 3, 4, 5, 6}
	b := []interface{}{0, 3, 7, 4, 8, 6}
	result := idiff(a, b)
	assertInfo(t, "TestAB入れ替え実行の場合", result, //
		5,                 // edist
		[]int{0, 3, 4, 6}, //lcs
		[]int{DiffSame, DiffDel, DiffDel, DiffSame, DiffAdd, DiffSame, DiffDel, DiffAdd, DiffSame}) // ses
}

func assertInfo(t *testing.T, name string, result *DiffInfo, edist int, lcs []int, ses []int) {
	if edist != result.Edist {
		t.Error(name)
	}
	if len(lcs) != len(result.Lcs) {
		t.Error(name)
	}
	for i := 0; i < len(lcs); i++ {
		if lcs[i] != result.Lcs[i].(int) {
			t.Error(name)
		}
	}
	if len(ses) != len(result.Ses) {
		t.Error(name)
	}
	for i := 0; i < len(ses); i++ {
		if ses[i] != result.Ses[i].Type {
			t.Error(name)
		}
	}
}

func printSes(result *DiffInfo) {
	for i := 0; i < len(result.Ses); i++ {
		s := result.Ses[i]
		if s.Type == DiffSame {
			fmt.Print("=")
		} else if s.Type == DiffAdd {
			fmt.Print("+")
		} else if s.Type == DiffDel {
			fmt.Print("-")
		} else {
			fmt.Print("?")
		}
		fmt.Println(s.Value)
	}
}
