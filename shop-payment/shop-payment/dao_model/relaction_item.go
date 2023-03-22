package dao_model

import (
	"time"
)

type RelactionItem struct {
	Id          int       `xorm:"'id' not null pk autoincr comment('唯一值') INT(11)"`
	Title       string    `xorm:"'title' not null comment('標題') index VARCHAR(100)"`
	Description string    `xorm:"'description' not null comment('描述') VARCHAR(100)"`
	Valid       int       `xorm:"'valid' not null default 1 comment('0:停用 1:啟用') TINYINT(2)"`
	IsDel       int       `xorm:"'is_del' not null default 0 comment('默认0：未删除的有效记录 1 : 逻辑删除的无效记录') TINYINT(2)"`
	CreatorId   int       `xorm:"'creator_id' not null default 0 comment('建立人员') INT(11)"`
	CreatedAt   time.Time `xorm:"'created_at' not null default CURRENT_TIMESTAMP comment('系统时间 默认(now()或 sysdate())') DATETIME <-"`
	ModifierId  int       `xorm:"'modifier_id' not null default 0 comment('修改人员') INT(11)"`
	ModifiedAt  time.Time `xorm:"'modified_at' not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME <-"`
}

func (m *RelactionItem) TableName() string {
	return "relaction_item"
}

/*
// FindEasyToSQL : simple query
func (m *RelactionItem) FindEasyToSQL(cond *RelactionItemCond, columns ...string) (sql string, args []interface{}, err error) {
    return GenSql(m.TableName(), columns, cond, []string{}, nil, nil, []string{})
}

// FindToSQL : complex query
func (m *RelactionItem) FindToSQL(cond *RelactionItemCond, page *Page, order map[string]string, groups []string, columns ...string) (sql string, args []interface{}, err error) {
	return GenSql(m.TableName(), columns, cond, []string{}, page, order, groups)
}
*/

func (m *RelactionItem) IsSetId() bool {
	return true
}

func (m *RelactionItem) GetId() int {
	return m.Id
}

func (m *RelactionItem) IsSetTitle() bool {
	return true
}

func (m *RelactionItem) GetTitle() string {
	return m.Title
}

func (m *RelactionItem) IsSetDescription() bool {
	return true
}

func (m *RelactionItem) GetDescription() string {
	return m.Description
}

func (m *RelactionItem) IsSetValid() bool {
	return true
}

func (m *RelactionItem) GetValid() int {
	return m.Valid
}

func (m *RelactionItem) IsSetIsDel() bool {
	return true
}

func (m *RelactionItem) GetIsDel() int {
	return m.IsDel
}

func (m *RelactionItem) IsSetCreatorId() bool {
	return true
}

func (m *RelactionItem) GetCreatorId() int {
	return m.CreatorId
}

func (m *RelactionItem) IsSetCreatedAt() bool {
	return true
}

func (m *RelactionItem) GetCreatedAt() time.Time {
	return m.CreatedAt
}

func (m *RelactionItem) IsSetModifierId() bool {
	return true
}

func (m *RelactionItem) GetModifierId() int {
	return m.ModifierId
}

func (m *RelactionItem) IsSetModifiedAt() bool {
	return true
}

func (m *RelactionItem) GetModifiedAt() time.Time {
	return m.ModifiedAt
}
