package dao_model

import "errors"

var (
	ErrNotFoundData = errors.New("Not Found Data")
)

type OrderBy map[string]string

// 分頁
type PageSet struct {
	No   int64
	Size int64
}

// mysql分頁
type Page struct {
	Limit  int
	Offset int
}

//默認 1000筆
func NewDefPage() *Page {
	pset := &PageSet{
		No:   1,
		Size: 1000,
	}
	return NewPage(pset)
}

func NewPageSize(no int64, size int64) *Page {
	pset := &PageSet{
		No:   no,
		Size: size,
	}
	return NewPage(pset)
}

//會轉成mysql的 limit / offset
func NewPage(pageSet *PageSet) *Page {
	if pageSet == nil {
		return &Page{
			Limit:  1000,
			Offset: 0,
		}
	}

	// 不作分頁
	if pageSet.Size < 0 {
		return &Page{
			Limit:  -1,
			Offset: -1,
		}
	}

	limit := int(pageSet.Size)
	offset := int((pageSet.No - 1)) * limit
	return &Page{
		Limit:  limit,
		Offset: offset,
	}
}
