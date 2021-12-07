package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	// 단어 못찾은 에러
	errNotFound = errors.New("단어가 없는데요?")
	// 이미 있는 단어를 등록하려할 때 에러
	errWordExists = errors.New("이미 등록된 단어에요")
	// 등록되지 않은 단어를 업데이트 하려 할 때 에러
	errCantUpdate = errors.New("없는 단어라 업데이트 할 수 없어요")
	// 등록되지 않은 단어를 삭제하려 할 때 에러
	errCantDelete = errors.New("삭제할 단어가 없는데영")
)

// 단어 검색
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// 단어 추가하기
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	// if err == errNotFound {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExists
	// }
	// return nil
	// 스위치로 바꿔 쓸 경우
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// 기존 단어 업데이트
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// 단어 삭제
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}
	return nil
}
