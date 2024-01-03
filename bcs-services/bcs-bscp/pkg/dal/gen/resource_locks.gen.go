// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/dal/table"
)

func newResourceLock(db *gorm.DB, opts ...gen.DOOption) resourceLock {
	_resourceLock := resourceLock{}

	_resourceLock.resourceLockDo.UseDB(db, opts...)
	_resourceLock.resourceLockDo.UseModel(&table.ResourceLock{})

	tableName := _resourceLock.resourceLockDo.TableName()
	_resourceLock.ALL = field.NewAsterisk(tableName)
	_resourceLock.ID = field.NewUint32(tableName, "id")
	_resourceLock.BizID = field.NewUint32(tableName, "biz_id")
	_resourceLock.ResType = field.NewString(tableName, "res_type")
	_resourceLock.ResKey = field.NewString(tableName, "res_key")
	_resourceLock.ResCount = field.NewUint32(tableName, "res_count")

	_resourceLock.fillFieldMap()

	return _resourceLock
}

type resourceLock struct {
	resourceLockDo resourceLockDo

	ALL      field.Asterisk
	ID       field.Uint32
	BizID    field.Uint32
	ResType  field.String
	ResKey   field.String
	ResCount field.Uint32

	fieldMap map[string]field.Expr
}

func (r resourceLock) Table(newTableName string) *resourceLock {
	r.resourceLockDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r resourceLock) As(alias string) *resourceLock {
	r.resourceLockDo.DO = *(r.resourceLockDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *resourceLock) updateTableName(table string) *resourceLock {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewUint32(table, "id")
	r.BizID = field.NewUint32(table, "biz_id")
	r.ResType = field.NewString(table, "res_type")
	r.ResKey = field.NewString(table, "res_key")
	r.ResCount = field.NewUint32(table, "res_count")

	r.fillFieldMap()

	return r
}

func (r *resourceLock) WithContext(ctx context.Context) IResourceLockDo {
	return r.resourceLockDo.WithContext(ctx)
}

func (r resourceLock) TableName() string { return r.resourceLockDo.TableName() }

func (r resourceLock) Alias() string { return r.resourceLockDo.Alias() }

func (r resourceLock) Columns(cols ...field.Expr) gen.Columns {
	return r.resourceLockDo.Columns(cols...)
}

func (r *resourceLock) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *resourceLock) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 5)
	r.fieldMap["id"] = r.ID
	r.fieldMap["biz_id"] = r.BizID
	r.fieldMap["res_type"] = r.ResType
	r.fieldMap["res_key"] = r.ResKey
	r.fieldMap["res_count"] = r.ResCount
}

func (r resourceLock) clone(db *gorm.DB) resourceLock {
	r.resourceLockDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r resourceLock) replaceDB(db *gorm.DB) resourceLock {
	r.resourceLockDo.ReplaceDB(db)
	return r
}

type resourceLockDo struct{ gen.DO }

type IResourceLockDo interface {
	gen.SubQuery
	Debug() IResourceLockDo
	WithContext(ctx context.Context) IResourceLockDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IResourceLockDo
	WriteDB() IResourceLockDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IResourceLockDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IResourceLockDo
	Not(conds ...gen.Condition) IResourceLockDo
	Or(conds ...gen.Condition) IResourceLockDo
	Select(conds ...field.Expr) IResourceLockDo
	Where(conds ...gen.Condition) IResourceLockDo
	Order(conds ...field.Expr) IResourceLockDo
	Distinct(cols ...field.Expr) IResourceLockDo
	Omit(cols ...field.Expr) IResourceLockDo
	Join(table schema.Tabler, on ...field.Expr) IResourceLockDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IResourceLockDo
	RightJoin(table schema.Tabler, on ...field.Expr) IResourceLockDo
	Group(cols ...field.Expr) IResourceLockDo
	Having(conds ...gen.Condition) IResourceLockDo
	Limit(limit int) IResourceLockDo
	Offset(offset int) IResourceLockDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IResourceLockDo
	Unscoped() IResourceLockDo
	Create(values ...*table.ResourceLock) error
	CreateInBatches(values []*table.ResourceLock, batchSize int) error
	Save(values ...*table.ResourceLock) error
	First() (*table.ResourceLock, error)
	Take() (*table.ResourceLock, error)
	Last() (*table.ResourceLock, error)
	Find() ([]*table.ResourceLock, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.ResourceLock, err error)
	FindInBatches(result *[]*table.ResourceLock, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.ResourceLock) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IResourceLockDo
	Assign(attrs ...field.AssignExpr) IResourceLockDo
	Joins(fields ...field.RelationField) IResourceLockDo
	Preload(fields ...field.RelationField) IResourceLockDo
	FirstOrInit() (*table.ResourceLock, error)
	FirstOrCreate() (*table.ResourceLock, error)
	FindByPage(offset int, limit int) (result []*table.ResourceLock, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IResourceLockDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r resourceLockDo) Debug() IResourceLockDo {
	return r.withDO(r.DO.Debug())
}

func (r resourceLockDo) WithContext(ctx context.Context) IResourceLockDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r resourceLockDo) ReadDB() IResourceLockDo {
	return r.Clauses(dbresolver.Read)
}

func (r resourceLockDo) WriteDB() IResourceLockDo {
	return r.Clauses(dbresolver.Write)
}

func (r resourceLockDo) Session(config *gorm.Session) IResourceLockDo {
	return r.withDO(r.DO.Session(config))
}

func (r resourceLockDo) Clauses(conds ...clause.Expression) IResourceLockDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r resourceLockDo) Returning(value interface{}, columns ...string) IResourceLockDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r resourceLockDo) Not(conds ...gen.Condition) IResourceLockDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r resourceLockDo) Or(conds ...gen.Condition) IResourceLockDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r resourceLockDo) Select(conds ...field.Expr) IResourceLockDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r resourceLockDo) Where(conds ...gen.Condition) IResourceLockDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r resourceLockDo) Order(conds ...field.Expr) IResourceLockDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r resourceLockDo) Distinct(cols ...field.Expr) IResourceLockDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r resourceLockDo) Omit(cols ...field.Expr) IResourceLockDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r resourceLockDo) Join(table schema.Tabler, on ...field.Expr) IResourceLockDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r resourceLockDo) LeftJoin(table schema.Tabler, on ...field.Expr) IResourceLockDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r resourceLockDo) RightJoin(table schema.Tabler, on ...field.Expr) IResourceLockDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r resourceLockDo) Group(cols ...field.Expr) IResourceLockDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r resourceLockDo) Having(conds ...gen.Condition) IResourceLockDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r resourceLockDo) Limit(limit int) IResourceLockDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r resourceLockDo) Offset(offset int) IResourceLockDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r resourceLockDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IResourceLockDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r resourceLockDo) Unscoped() IResourceLockDo {
	return r.withDO(r.DO.Unscoped())
}

func (r resourceLockDo) Create(values ...*table.ResourceLock) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r resourceLockDo) CreateInBatches(values []*table.ResourceLock, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r resourceLockDo) Save(values ...*table.ResourceLock) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r resourceLockDo) First() (*table.ResourceLock, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.ResourceLock), nil
	}
}

func (r resourceLockDo) Take() (*table.ResourceLock, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.ResourceLock), nil
	}
}

func (r resourceLockDo) Last() (*table.ResourceLock, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.ResourceLock), nil
	}
}

func (r resourceLockDo) Find() ([]*table.ResourceLock, error) {
	result, err := r.DO.Find()
	return result.([]*table.ResourceLock), err
}

func (r resourceLockDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.ResourceLock, err error) {
	buf := make([]*table.ResourceLock, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r resourceLockDo) FindInBatches(result *[]*table.ResourceLock, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r resourceLockDo) Attrs(attrs ...field.AssignExpr) IResourceLockDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r resourceLockDo) Assign(attrs ...field.AssignExpr) IResourceLockDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r resourceLockDo) Joins(fields ...field.RelationField) IResourceLockDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r resourceLockDo) Preload(fields ...field.RelationField) IResourceLockDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r resourceLockDo) FirstOrInit() (*table.ResourceLock, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.ResourceLock), nil
	}
}

func (r resourceLockDo) FirstOrCreate() (*table.ResourceLock, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.ResourceLock), nil
	}
}

func (r resourceLockDo) FindByPage(offset int, limit int) (result []*table.ResourceLock, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r resourceLockDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r resourceLockDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r resourceLockDo) Delete(models ...*table.ResourceLock) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *resourceLockDo) withDO(do gen.Dao) *resourceLockDo {
	r.DO = *do.(*gen.DO)
	return r
}
