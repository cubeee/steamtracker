package paginator

import (
	"github.com/jinzhu/gorm"
)

type Paginator struct {
	options      Options
	db           *gorm.DB
	calculations Calculations
}

type Calculations struct {
	pages      int64
	totalItems int64
}

type Options struct {
	PageSize int64
}

func NewPaginator(options Options, db *gorm.DB, itemType interface{}) Paginator {
	return Paginator{options: options, db: db}.paginate(itemType)
}

func (p Paginator) paginate(itemType interface{}) Paginator {
	p.calculations = Calculations{
		totalItems: p.countItems(itemType),
	}
	return p
}

func (p Paginator) Page(page uint64, out interface{}) {
	p.PageCustom(page, out, func(db *gorm.DB, out interface{}) {
		db.Find(&out)
	})
}

func (p Paginator) PageCustom(page uint64, out interface{}, customFunc func(db *gorm.DB, out interface{})) {
	var offset uint64
	if page > 1 {
		offset = page + 1
	}

	customFunc(p.db.Offset(offset).Limit(p.options.PageSize), out)
}

func (p Paginator) countItems(itemType interface{}) int64 {
	var pages int64
	p.db.Model(itemType).Count(&pages)
	return pages
}

func (p Paginator) Pages() int64 {
	totalPages := p.calculations.totalItems / p.options.PageSize
	return int64(totalPages + 1)
}
