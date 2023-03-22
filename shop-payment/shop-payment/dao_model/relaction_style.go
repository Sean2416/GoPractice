package dao_model

import (
	"time"
)

type RelactionStyle struct {
	Id          int       `xorm:"'id' not null pk autoincr comment('唯一值') INT(11)"`
	Key         string    `xorm:"'key' not null default '' comment('對應值') VARCHAR(45)"`
	Title       string    `xorm:"'title' not null comment('標題') VARCHAR(100)"`
	Description string    `xorm:"'description' not null comment('描述') VARCHAR(100)"`
	Sort        int       `xorm:"'sort' not null default 0 comment('Sort') TINYINT(4)"`
	Valid       int       `xorm:"'valid' not null default 1 comment('0:停用 1:啟用') TINYINT(2)"`
	IsDel       int       `xorm:"'is_del' not null default 0 comment('默认0：未删除的有效记录 1 : 逻辑删除的无效记录') TINYINT(2)"`
	CreatorId   int       `xorm:"'creator_id' not null default 0 comment('建立人员') INT(11)"`
	CreatedAt   time.Time `xorm:"'created_at' not null default CURRENT_TIMESTAMP comment('系统时间 默认(now()或 sysdate())') DATETIME <-"`
	ModifierId  int       `xorm:"'modifier_id' not null default 0 comment('修改人员') INT(11)"`
	ModifiedAt  time.Time `xorm:"'modified_at' not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME <-"`
}

func (m *RelactionStyle) TableName() string {
	return "relaction_style"
}

/*
// FindEasyToSQL : simple query
func (m *RelactionStyle) FindEasyToSQL(cond *RelactionStyleCond, columns ...string) (sql string, args []interface{}, err error) {
    return GenSql(m.TableName(), columns, cond, []string{}, nil, nil, []string{})
}

// FindToSQL : complex query
func (m *RelactionStyle) FindToSQL(cond *RelactionStyleCond, page *Page, order map[string]string, groups []string, columns ...string) (sql string, args []interface{}, err error) {
	return GenSql(m.TableName(), columns, cond, []string{}, page, order, groups)
}
*/

func (m *RelactionStyle) IsSetId() bool {
	return true
}

func (m *RelactionStyle) GetId() int {
	return m.Id
}

func (m *RelactionStyle) IsSetKey() bool {
	return true
}

func (m *RelactionStyle) GetKey() string {
	return m.Key
}

func (m *RelactionStyle) IsSetTitle() bool {
	return true
}

func (m *RelactionStyle) GetTitle() string {
	return m.Title
}

func (m *RelactionStyle) IsSetDescription() bool {
	return true
}

func (m *RelactionStyle) GetDescription() string {
	return m.Description
}

func (m *RelactionStyle) IsSetSort() bool {
	return true
}

func (m *RelactionStyle) GetSort() int {
	return m.Sort
}

func (m *RelactionStyle) IsSetValid() bool {
	return true
}

func (m *RelactionStyle) GetValid() int {
	return m.Valid
}

func (m *RelactionStyle) IsSetIsDel() bool {
	return true
}

func (m *RelactionStyle) GetIsDel() int {
	return m.IsDel
}

func (m *RelactionStyle) IsSetCreatorId() bool {
	return true
}

func (m *RelactionStyle) GetCreatorId() int {
	return m.CreatorId
}

func (m *RelactionStyle) IsSetCreatedAt() bool {
	return true
}

func (m *RelactionStyle) GetCreatedAt() time.Time {
	return m.CreatedAt
}

func (m *RelactionStyle) IsSetModifierId() bool {
	return true
}

func (m *RelactionStyle) GetModifierId() int {
	return m.ModifierId
}

func (m *RelactionStyle) IsSetModifiedAt() bool {
	return true
}

func (m *RelactionStyle) GetModifiedAt() time.Time {
	return m.ModifiedAt
}
