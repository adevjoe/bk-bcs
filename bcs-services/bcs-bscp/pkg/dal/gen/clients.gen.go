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

func newClient(db *gorm.DB, opts ...gen.DOOption) client {
	_client := client{}

	_client.clientDo.UseDB(db, opts...)
	_client.clientDo.UseModel(&table.Client{})

	tableName := _client.clientDo.TableName()
	_client.ALL = field.NewAsterisk(tableName)
	_client.ID = field.NewUint32(tableName, "id")
	_client.UID = field.NewString(tableName, "uid")
	_client.BizID = field.NewUint32(tableName, "biz_id")
	_client.AppID = field.NewUint32(tableName, "app_id")
	_client.ClientVersion = field.NewString(tableName, "client_version")
	_client.ClientType = field.NewString(tableName, "client_type")
	_client.Ip = field.NewString(tableName, "ip")
	_client.Labels = field.NewString(tableName, "labels")
	_client.Annotations = field.NewString(tableName, "annotations")
	_client.FirstConnectTime = field.NewTime(tableName, "first_connect_time")
	_client.LastHeartbeatTime = field.NewTime(tableName, "last_heartbeat_time")
	_client.OnlineStatus = field.NewString(tableName, "online_status")
	_client.CpuUsage = field.NewFloat64(tableName, "cpu_usage")
	_client.CpuMaxUsage = field.NewFloat64(tableName, "cpu_max_usage")
	_client.CpuMinUsage = field.NewFloat64(tableName, "cpu_min_usage")
	_client.CpuAvgUsage = field.NewFloat64(tableName, "cpu_avg_usage")
	_client.MemoryUsage = field.NewUint64(tableName, "memory_usage")
	_client.MemoryMaxUsage = field.NewUint64(tableName, "memory_max_usage")
	_client.MemoryMinUsage = field.NewUint64(tableName, "memory_min_usage")
	_client.MemoryAvgUsage = field.NewUint64(tableName, "memory_avg_usage")
	_client.CurrentReleaseID = field.NewUint32(tableName, "current_release_id")
	_client.TargetReleaseID = field.NewUint32(tableName, "target_release_id")
	_client.ReleaseChangeStatus = field.NewString(tableName, "release_change_status")
	_client.ReleaseChangeFailedReason = field.NewString(tableName, "release_change_failed_reason")
	_client.SpecificFailedReason = field.NewString(tableName, "specific_failed_reason")
	_client.FailedDetailReason = field.NewString(tableName, "failed_detail_reason")
	_client.TotalSeconds = field.NewFloat64(tableName, "total_seconds")

	_client.fillFieldMap()

	return _client
}

type client struct {
	clientDo clientDo

	ALL                       field.Asterisk
	ID                        field.Uint32
	UID                       field.String
	BizID                     field.Uint32
	AppID                     field.Uint32
	ClientVersion             field.String
	ClientType                field.String
	Ip                        field.String
	Labels                    field.String
	Annotations               field.String
	FirstConnectTime          field.Time
	LastHeartbeatTime         field.Time
	OnlineStatus              field.String
	CpuUsage                  field.Float64
	CpuMaxUsage               field.Float64
	CpuMinUsage               field.Float64
	CpuAvgUsage               field.Float64
	MemoryUsage               field.Uint64
	MemoryMaxUsage            field.Uint64
	MemoryMinUsage            field.Uint64
	MemoryAvgUsage            field.Uint64
	CurrentReleaseID          field.Uint32
	TargetReleaseID           field.Uint32
	ReleaseChangeStatus       field.String
	ReleaseChangeFailedReason field.String
	SpecificFailedReason      field.String
	FailedDetailReason        field.String
	TotalSeconds              field.Float64

	fieldMap map[string]field.Expr
}

func (c client) Table(newTableName string) *client {
	c.clientDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c client) As(alias string) *client {
	c.clientDo.DO = *(c.clientDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *client) updateTableName(table string) *client {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewUint32(table, "id")
	c.UID = field.NewString(table, "uid")
	c.BizID = field.NewUint32(table, "biz_id")
	c.AppID = field.NewUint32(table, "app_id")
	c.ClientVersion = field.NewString(table, "client_version")
	c.ClientType = field.NewString(table, "client_type")
	c.Ip = field.NewString(table, "ip")
	c.Labels = field.NewString(table, "labels")
	c.Annotations = field.NewString(table, "annotations")
	c.FirstConnectTime = field.NewTime(table, "first_connect_time")
	c.LastHeartbeatTime = field.NewTime(table, "last_heartbeat_time")
	c.OnlineStatus = field.NewString(table, "online_status")
	c.CpuUsage = field.NewFloat64(table, "cpu_usage")
	c.CpuMaxUsage = field.NewFloat64(table, "cpu_max_usage")
	c.CpuMinUsage = field.NewFloat64(table, "cpu_min_usage")
	c.CpuAvgUsage = field.NewFloat64(table, "cpu_avg_usage")
	c.MemoryUsage = field.NewUint64(table, "memory_usage")
	c.MemoryMaxUsage = field.NewUint64(table, "memory_max_usage")
	c.MemoryMinUsage = field.NewUint64(table, "memory_min_usage")
	c.MemoryAvgUsage = field.NewUint64(table, "memory_avg_usage")
	c.CurrentReleaseID = field.NewUint32(table, "current_release_id")
	c.TargetReleaseID = field.NewUint32(table, "target_release_id")
	c.ReleaseChangeStatus = field.NewString(table, "release_change_status")
	c.ReleaseChangeFailedReason = field.NewString(table, "release_change_failed_reason")
	c.SpecificFailedReason = field.NewString(table, "specific_failed_reason")
	c.FailedDetailReason = field.NewString(table, "failed_detail_reason")
	c.TotalSeconds = field.NewFloat64(table, "total_seconds")

	c.fillFieldMap()

	return c
}

func (c *client) WithContext(ctx context.Context) IClientDo { return c.clientDo.WithContext(ctx) }

func (c client) TableName() string { return c.clientDo.TableName() }

func (c client) Alias() string { return c.clientDo.Alias() }

func (c client) Columns(cols ...field.Expr) gen.Columns { return c.clientDo.Columns(cols...) }

func (c *client) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *client) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 27)
	c.fieldMap["id"] = c.ID
	c.fieldMap["uid"] = c.UID
	c.fieldMap["biz_id"] = c.BizID
	c.fieldMap["app_id"] = c.AppID
	c.fieldMap["client_version"] = c.ClientVersion
	c.fieldMap["client_type"] = c.ClientType
	c.fieldMap["ip"] = c.Ip
	c.fieldMap["labels"] = c.Labels
	c.fieldMap["annotations"] = c.Annotations
	c.fieldMap["first_connect_time"] = c.FirstConnectTime
	c.fieldMap["last_heartbeat_time"] = c.LastHeartbeatTime
	c.fieldMap["online_status"] = c.OnlineStatus
	c.fieldMap["cpu_usage"] = c.CpuUsage
	c.fieldMap["cpu_max_usage"] = c.CpuMaxUsage
	c.fieldMap["cpu_min_usage"] = c.CpuMinUsage
	c.fieldMap["cpu_avg_usage"] = c.CpuAvgUsage
	c.fieldMap["memory_usage"] = c.MemoryUsage
	c.fieldMap["memory_max_usage"] = c.MemoryMaxUsage
	c.fieldMap["memory_min_usage"] = c.MemoryMinUsage
	c.fieldMap["memory_avg_usage"] = c.MemoryAvgUsage
	c.fieldMap["current_release_id"] = c.CurrentReleaseID
	c.fieldMap["target_release_id"] = c.TargetReleaseID
	c.fieldMap["release_change_status"] = c.ReleaseChangeStatus
	c.fieldMap["release_change_failed_reason"] = c.ReleaseChangeFailedReason
	c.fieldMap["specific_failed_reason"] = c.SpecificFailedReason
	c.fieldMap["failed_detail_reason"] = c.FailedDetailReason
	c.fieldMap["total_seconds"] = c.TotalSeconds
}

func (c client) clone(db *gorm.DB) client {
	c.clientDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c client) replaceDB(db *gorm.DB) client {
	c.clientDo.ReplaceDB(db)
	return c
}

type clientDo struct{ gen.DO }

type IClientDo interface {
	gen.SubQuery
	Debug() IClientDo
	WithContext(ctx context.Context) IClientDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IClientDo
	WriteDB() IClientDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IClientDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IClientDo
	Not(conds ...gen.Condition) IClientDo
	Or(conds ...gen.Condition) IClientDo
	Select(conds ...field.Expr) IClientDo
	Where(conds ...gen.Condition) IClientDo
	Order(conds ...field.Expr) IClientDo
	Distinct(cols ...field.Expr) IClientDo
	Omit(cols ...field.Expr) IClientDo
	Join(table schema.Tabler, on ...field.Expr) IClientDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IClientDo
	RightJoin(table schema.Tabler, on ...field.Expr) IClientDo
	Group(cols ...field.Expr) IClientDo
	Having(conds ...gen.Condition) IClientDo
	Limit(limit int) IClientDo
	Offset(offset int) IClientDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IClientDo
	Unscoped() IClientDo
	Create(values ...*table.Client) error
	CreateInBatches(values []*table.Client, batchSize int) error
	Save(values ...*table.Client) error
	First() (*table.Client, error)
	Take() (*table.Client, error)
	Last() (*table.Client, error)
	Find() ([]*table.Client, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.Client, err error)
	FindInBatches(result *[]*table.Client, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.Client) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IClientDo
	Assign(attrs ...field.AssignExpr) IClientDo
	Joins(fields ...field.RelationField) IClientDo
	Preload(fields ...field.RelationField) IClientDo
	FirstOrInit() (*table.Client, error)
	FirstOrCreate() (*table.Client, error)
	FindByPage(offset int, limit int) (result []*table.Client, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IClientDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c clientDo) Debug() IClientDo {
	return c.withDO(c.DO.Debug())
}

func (c clientDo) WithContext(ctx context.Context) IClientDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c clientDo) ReadDB() IClientDo {
	return c.Clauses(dbresolver.Read)
}

func (c clientDo) WriteDB() IClientDo {
	return c.Clauses(dbresolver.Write)
}

func (c clientDo) Session(config *gorm.Session) IClientDo {
	return c.withDO(c.DO.Session(config))
}

func (c clientDo) Clauses(conds ...clause.Expression) IClientDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c clientDo) Returning(value interface{}, columns ...string) IClientDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c clientDo) Not(conds ...gen.Condition) IClientDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c clientDo) Or(conds ...gen.Condition) IClientDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c clientDo) Select(conds ...field.Expr) IClientDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c clientDo) Where(conds ...gen.Condition) IClientDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c clientDo) Order(conds ...field.Expr) IClientDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c clientDo) Distinct(cols ...field.Expr) IClientDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c clientDo) Omit(cols ...field.Expr) IClientDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c clientDo) Join(table schema.Tabler, on ...field.Expr) IClientDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c clientDo) LeftJoin(table schema.Tabler, on ...field.Expr) IClientDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c clientDo) RightJoin(table schema.Tabler, on ...field.Expr) IClientDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c clientDo) Group(cols ...field.Expr) IClientDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c clientDo) Having(conds ...gen.Condition) IClientDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c clientDo) Limit(limit int) IClientDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c clientDo) Offset(offset int) IClientDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c clientDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IClientDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c clientDo) Unscoped() IClientDo {
	return c.withDO(c.DO.Unscoped())
}

func (c clientDo) Create(values ...*table.Client) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c clientDo) CreateInBatches(values []*table.Client, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c clientDo) Save(values ...*table.Client) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c clientDo) First() (*table.Client, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.Client), nil
	}
}

func (c clientDo) Take() (*table.Client, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.Client), nil
	}
}

func (c clientDo) Last() (*table.Client, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.Client), nil
	}
}

func (c clientDo) Find() ([]*table.Client, error) {
	result, err := c.DO.Find()
	return result.([]*table.Client), err
}

func (c clientDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.Client, err error) {
	buf := make([]*table.Client, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c clientDo) FindInBatches(result *[]*table.Client, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c clientDo) Attrs(attrs ...field.AssignExpr) IClientDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c clientDo) Assign(attrs ...field.AssignExpr) IClientDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c clientDo) Joins(fields ...field.RelationField) IClientDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c clientDo) Preload(fields ...field.RelationField) IClientDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c clientDo) FirstOrInit() (*table.Client, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.Client), nil
	}
}

func (c clientDo) FirstOrCreate() (*table.Client, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.Client), nil
	}
}

func (c clientDo) FindByPage(offset int, limit int) (result []*table.Client, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c clientDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c clientDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c clientDo) Delete(models ...*table.Client) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *clientDo) withDO(do gen.Dao) *clientDo {
	c.DO = *do.(*gen.DO)
	return c
}
