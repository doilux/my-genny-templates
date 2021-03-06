// Code generated by "my-genny-templates"; DO NOT EDIT.

package slice

import "github.com/cheekybits/genny/generic"

// ElementType - genny用変数
type ElementType generic.Type

// ElementTypeList - genny用変数
type ElementTypeList []ElementType

// Where - fnがtrueを返す要素で構成されたスライスを返す。
func (rcv ElementTypeList) Where(fn func(n ElementType) bool) ElementTypeList {
	result := make([]ElementType, 0, len(rcv))
	for _, v := range rcv {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Count - fnがtrueを返す要素で構成された要素の数を返す
func (rcv ElementTypeList) Count(fn func(n ElementType) bool) int {
	return len(rcv.Where(fn))
}

// Any - fnがtrueを返す要素が存在するときにtrueを返す。
func (rcv ElementTypeList) Any(fn func(n ElementType) bool) bool {
	for _, v := range rcv {
		if fn(v) {
			return true
		}
	}
	return false
}

// FoldLeftInt64 - defを初期値として、Sliceの各エレメントに対してopを実行した結果を返す
func (rcv ElementTypeList) FoldLeftInt64(def int64, op func(acc int64, x ElementType) int64) int64 {
	result := def
	for _, v := range rcv {
		result = op(result, v)
	}
	return result
}

// FoldLeftInt32 - defを初期値として、Sliceの各エレメントに対してopを実行した結果を返す
func (rcv ElementTypeList) FoldLeftInt32(def int32, op func(acc int32, x ElementType) int32) int32 {
	result := def
	for _, v := range rcv {
		result = op(result, v)
	}
	return result
}

// First - 最初の要素のポインタを返す。配列が空の時はnilを返す
func (rcv ElementTypeList) First() *ElementType {
	if len(rcv) < 1 {
		return nil
	}
	return &rcv[0]
}

// First - 最初の要素のポインタを返す。配列が空の時はdefValを返す
func (rcv ElementTypeList) FirstOrElse(defVal ElementType) ElementType {
	if len(rcv) < 1 {
		return defVal
	}
	return rcv[0]
}