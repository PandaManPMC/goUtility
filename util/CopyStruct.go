package util

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

var structCopyCache sync.Map

type structCopyCacheKey struct {
	src reflect.Type
	dst reflect.Type
}

type structFieldMapping struct {
	srcIndex []int
	dstIndex []int
	convert  bool
}

// CopyStruct 根据字段名复制两个结构体中类型兼容的字段。
// dst 必须是非 nil 的 struct 指针。
// src 可以是 struct 或 *struct。
func CopyStruct(dst any, src any) error {
	if dst == nil {
		return errors.New("CopyStruct: dst is nil")
	}

	if src == nil {
		return errors.New("CopyStruct: src is nil")
	}

	dstValue := reflect.ValueOf(dst)

	if dstValue.Kind() != reflect.Ptr || dstValue.IsNil() {
		return errors.New("CopyStruct: dst must be a non-nil pointer")
	}

	dstValue = dstValue.Elem()

	if dstValue.Kind() != reflect.Struct {
		return errors.New("CopyStruct: dst must point to struct")
	}

	srcValue := reflect.ValueOf(src)

	// 解引用指针
	for srcValue.Kind() == reflect.Ptr {
		if srcValue.IsNil() {
			return errors.New("CopyStruct: src is nil pointer")
		}

		srcValue = srcValue.Elem()
	}

	if srcValue.Kind() != reflect.Struct {
		return errors.New("CopyStruct: src must be struct or pointer to struct")
	}

	srcType := srcValue.Type()
	dstType := dstValue.Type()

	key := structCopyCacheKey{
		src: srcType,
		dst: dstType,
	}

	var mappings []structFieldMapping

	if cached, ok := structCopyCache.Load(key); ok {
		mappings = cached.([]structFieldMapping)
	} else {
		mappings = buildStructFieldMapping(srcType, dstType)

		// 防止多个 goroutine 同时构建同一个缓存
		actual, _ := structCopyCache.LoadOrStore(key, mappings)
		mappings = actual.([]structFieldMapping)
	}

	for _, mapping := range mappings {
		srcField := srcValue.FieldByIndex(mapping.srcIndex)
		dstField := dstValue.FieldByIndex(mapping.dstIndex)

		if !srcField.IsValid() || !dstField.IsValid() {
			continue
		}

		if !dstField.CanSet() {
			continue
		}

		if mapping.convert {
			dstField.Set(srcField.Convert(dstField.Type()))
		} else {
			dstField.Set(srcField)
		}
	}

	return nil
}

func buildStructFieldMapping(
	srcType reflect.Type,
	dstType reflect.Type,
) []structFieldMapping {

	mappings := make([]structFieldMapping, 0, dstType.NumField())

	srcFields := make(map[string]reflect.StructField, srcType.NumField())

	// 建立 source 字段索引
	for i := 0; i < srcType.NumField(); i++ {
		field := srcType.Field(i)

		// PkgPath != "" 表示非导出字段
		if field.PkgPath != "" {
			continue
		}

		srcFields[field.Name] = field
	}

	for i := 0; i < dstType.NumField(); i++ {
		dstField := dstType.Field(i)

		if dstField.PkgPath != "" {
			continue
		}

		srcField, ok := srcFields[dstField.Name]
		if !ok {
			continue
		}

		switch {
		case srcField.Type.AssignableTo(dstField.Type):
			mappings = append(mappings, structFieldMapping{
				srcIndex: srcField.Index,
				dstIndex: dstField.Index,
			})

		case srcField.Type.ConvertibleTo(dstField.Type):
			mappings = append(mappings, structFieldMapping{
				srcIndex: srcField.Index,
				dstIndex: dstField.Index,
				convert:  true,
			})
		}
	}

	return mappings
}

// CopyStructSlice 拷贝 [] ，自动推导 []D
// result, err := CopyStructSlice[Pet, ](models)
func CopyStructSlice[D any, S any](src []S) ([]D, error) {
	if len(src) == 0 {
		return []D{}, nil
	}

	result := make([]D, len(src))

	for i := range src {
		if err := CopyStruct(&result[i], &src[i]); err != nil {
			return nil, fmt.Errorf(
				"CopyStructSlice: index %d: %w",
				i,
				err,
			)
		}
	}

	return result, nil
}
