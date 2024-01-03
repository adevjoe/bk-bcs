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

func newAudit(db *gorm.DB, opts ...gen.DOOption) audit {
	_audit := audit{}

	_audit.auditDo.UseDB(db, opts...)
	_audit.auditDo.UseModel(&table.Audit{})

	tableName := _audit.auditDo.TableName()
	_audit.ALL = field.NewAsterisk(tableName)
	_audit.ID = field.NewUint32(tableName, "id")
	_audit.BizID = field.NewUint32(tableName, "biz_id")
	_audit.AppID = field.NewUint32(tableName, "app_id")
	_audit.ResourceType = field.NewString(tableName, "res_type")
	_audit.ResourceID = field.NewUint32(tableName, "res_id")
	_audit.Action = field.NewString(tableName, "action")
	_audit.Rid = field.NewString(tableName, "rid")
	_audit.AppCode = field.NewString(tableName, "app_code")
	_audit.Operator = field.NewString(tableName, "operator")
	_audit.CreatedAt = field.NewTime(tableName, "created_at")
	_audit.Detail = field.NewString(tableName, "detail")

	_audit.fillFieldMap()

	return _audit
}

type audit struct {
	auditDo auditDo

	ALL          field.Asterisk
	ID           field.Uint32
	BizID        field.Uint32
	AppID        field.Uint32
	ResourceType field.String
	ResourceID   field.Uint32
	Action       field.String
	Rid          field.String
	AppCode      field.String
	Operator     field.String
	CreatedAt    field.Time
	Detail       field.String

	fieldMap map[string]field.Expr
}

func (a audit) Table(newTableName string) *audit {
	a.auditDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a audit) As(alias string) *audit {
	a.auditDo.DO = *(a.auditDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *audit) updateTableName(table string) *audit {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint32(table, "id")
	a.BizID = field.NewUint32(table, "biz_id")
	a.AppID = field.NewUint32(table, "app_id")
	a.ResourceType = field.NewString(table, "res_type")
	a.ResourceID = field.NewUint32(table, "res_id")
	a.Action = field.NewString(table, "action")
	a.Rid = field.NewString(table, "rid")
	a.AppCode = field.NewString(table, "app_code")
	a.Operator = field.NewString(table, "operator")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.Detail = field.NewString(table, "detail")

	a.fillFieldMap()

	return a
}

func (a *audit) WithContext(ctx context.Context) IAuditDo { return a.auditDo.WithContext(ctx) }

func (a audit) TableName() string { return a.auditDo.TableName() }

func (a audit) Alias() string { return a.auditDo.Alias() }

func (a audit) Columns(cols ...field.Expr) gen.Columns { return a.auditDo.Columns(cols...) }

func (a *audit) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *audit) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 11)
	a.fieldMap["id"] = a.ID
	a.fieldMap["biz_id"] = a.BizID
	a.fieldMap["app_id"] = a.AppID
	a.fieldMap["res_type"] = a.ResourceType
	a.fieldMap["res_id"] = a.ResourceID
	a.fieldMap["action"] = a.Action
	a.fieldMap["rid"] = a.Rid
	a.fieldMap["app_code"] = a.AppCode
	a.fieldMap["operator"] = a.Operator
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["detail"] = a.Detail
}

func (a audit) clone(db *gorm.DB) audit {
	a.auditDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a audit) replaceDB(db *gorm.DB) audit {
	a.auditDo.ReplaceDB(db)
	return a
}

type auditDo struct{ gen.DO }

type IAuditDo interface {
	gen.SubQuery
	Debug() IAuditDo
	WithContext(ctx context.Context) IAuditDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAuditDo
	WriteDB() IAuditDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAuditDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAuditDo
	Not(conds ...gen.Condition) IAuditDo
	Or(conds ...gen.Condition) IAuditDo
	Select(conds ...field.Expr) IAuditDo
	Where(conds ...gen.Condition) IAuditDo
	Order(conds ...field.Expr) IAuditDo
	Distinct(cols ...field.Expr) IAuditDo
	Omit(cols ...field.Expr) IAuditDo
	Join(table schema.Tabler, on ...field.Expr) IAuditDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAuditDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAuditDo
	Group(cols ...field.Expr) IAuditDo
	Having(conds ...gen.Condition) IAuditDo
	Limit(limit int) IAuditDo
	Offset(offset int) IAuditDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAuditDo
	Unscoped() IAuditDo
	Create(values ...*table.Audit) error
	CreateInBatches(values []*table.Audit, batchSize int) error
	Save(values ...*table.Audit) error
	First() (*table.Audit, error)
	Take() (*table.Audit, error)
	Last() (*table.Audit, error)
	Find() ([]*table.Audit, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.Audit, err error)
	FindInBatches(result *[]*table.Audit, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.Audit) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAuditDo
	Assign(attrs ...field.AssignExpr) IAuditDo
	Joins(fields ...field.RelationField) IAuditDo
	Preload(fields ...field.RelationField) IAuditDo
	FirstOrInit() (*table.Audit, error)
	FirstOrCreate() (*table.Audit, error)
	FindByPage(offset int, limit int) (result []*table.Audit, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAuditDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a auditDo) Debug() IAuditDo {
	return a.withDO(a.DO.Debug())
}

func (a auditDo) WithContext(ctx context.Context) IAuditDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a auditDo) ReadDB() IAuditDo {
	return a.Clauses(dbresolver.Read)
}

func (a auditDo) WriteDB() IAuditDo {
	return a.Clauses(dbresolver.Write)
}

func (a auditDo) Session(config *gorm.Session) IAuditDo {
	return a.withDO(a.DO.Session(config))
}

func (a auditDo) Clauses(conds ...clause.Expression) IAuditDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a auditDo) Returning(value interface{}, columns ...string) IAuditDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a auditDo) Not(conds ...gen.Condition) IAuditDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a auditDo) Or(conds ...gen.Condition) IAuditDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a auditDo) Select(conds ...field.Expr) IAuditDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a auditDo) Where(conds ...gen.Condition) IAuditDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a auditDo) Order(conds ...field.Expr) IAuditDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a auditDo) Distinct(cols ...field.Expr) IAuditDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a auditDo) Omit(cols ...field.Expr) IAuditDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a auditDo) Join(table schema.Tabler, on ...field.Expr) IAuditDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a auditDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAuditDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a auditDo) RightJoin(table schema.Tabler, on ...field.Expr) IAuditDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a auditDo) Group(cols ...field.Expr) IAuditDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a auditDo) Having(conds ...gen.Condition) IAuditDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a auditDo) Limit(limit int) IAuditDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a auditDo) Offset(offset int) IAuditDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a auditDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAuditDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a auditDo) Unscoped() IAuditDo {
	return a.withDO(a.DO.Unscoped())
}

func (a auditDo) Create(values ...*table.Audit) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a auditDo) CreateInBatches(values []*table.Audit, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a auditDo) Save(values ...*table.Audit) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a auditDo) First() (*table.Audit, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.Audit), nil
	}
}

func (a auditDo) Take() (*table.Audit, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.Audit), nil
	}
}

func (a auditDo) Last() (*table.Audit, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.Audit), nil
	}
}

func (a auditDo) Find() ([]*table.Audit, error) {
	result, err := a.DO.Find()
	return result.([]*table.Audit), err
}

func (a auditDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.Audit, err error) {
	buf := make([]*table.Audit, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a auditDo) FindInBatches(result *[]*table.Audit, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a auditDo) Attrs(attrs ...field.AssignExpr) IAuditDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a auditDo) Assign(attrs ...field.AssignExpr) IAuditDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a auditDo) Joins(fields ...field.RelationField) IAuditDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a auditDo) Preload(fields ...field.RelationField) IAuditDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a auditDo) FirstOrInit() (*table.Audit, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.Audit), nil
	}
}

func (a auditDo) FirstOrCreate() (*table.Audit, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.Audit), nil
	}
}

func (a auditDo) FindByPage(offset int, limit int) (result []*table.Audit, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a auditDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a auditDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a auditDo) Delete(models ...*table.Audit) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *auditDo) withDO(do gen.Dao) *auditDo {
	a.DO = *do.(*gen.DO)
	return a
}
