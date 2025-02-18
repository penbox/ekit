// Copyright 2021 gotomicro
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package slice

// Diff 差集，只支持 comparable 类型
func Diff[T comparable](src, dst []T) []T {
	if src == nil {
		return nil
	}
	if dst == nil {
		return src
	}
	mp := make(map[T]struct{}, len(dst))
	for _, d := range dst {
		mp[d] = struct{}{}
	}

	diff := make([]T, 0)
	for i := 0; i < len(src); i++ {
		if _, ok := mp[src[i]]; !ok {
			diff = append(diff, src[i])
			mp[src[i]] = struct{}{}
		}
	}

	return diff
}

// DiffFunc 差集
// 你应该优先使用 Diff
func DiffFunc[T any](src, dst []T, equal EqualFunc[T]) []T {
	if src == nil {
		return nil
	}
	if dst == nil {
		return src
	}

	diff := make([]T, 0)
	for i := 0; i < len(src); i++ {
		if !ContainsFunc(dst, src[i], equal) && !ContainsFunc(diff, src[i], equal) {
			diff = append(diff, src[i])
		}
	}

	return diff
}

// SymmetricDiff 对称差集
func SymmetricDiff[T comparable](src, dst []T) []T {
	diff := Diff(src, dst)
	diff = append(diff, Diff(dst, src)...)
	return diff
}

// SymmetricDiffFunc 对称差集
// 你应该优先使用 SymmetricDiff
func SymmetricDiffFunc[T any](src, dst []T, equal EqualFunc[T]) []T {
	diff := DiffFunc(src, dst, equal)
	diff = append(diff, DiffFunc(dst, src, equal)...)
	return diff
}
