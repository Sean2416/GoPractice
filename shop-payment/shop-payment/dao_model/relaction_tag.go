package dao_model

import (
	"time"
)

type RelactionTag struct {
	Id         int       `xorm:"'id' not null pk autoincr comment('唯一值') INT(11)"`
	Content    string    `xorm:"'content' not null default '' comment('標註內容') unique VARCHAR(50)"`
	Valid      int       `xorm:"'valid' not null default 1 comment('0:停用 1:啟用') TINYINT(2)"`
	IsDel      int       `xorm:"'is_del' not null default 0 comment('默认0：未删除的有效记录 1 : 逻辑删除的无效记录') TINYINT(2)"`
	CreatorId  int       `xorm:"'creator_id' not null default 0 comment('建立人员') INT(11)"`
	CreatedAt  time.Time `xorm:"'created_at' not null default CURRENT_TIMESTAMP comment('系统时间 默认(now()或 sysdate())') DATETIME <-"`
	ModifierId int       `xorm:"'modifier_id' not null default 0 comment('修改人员') INT(11)"`
	ModifiedAt time.Time `xorm:"'modified_at' not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME <-"`
}

func (m *RelactionTag) TableName() string {
	return "relaction_tag"
}

/*
// FindEasyToSQL : simple query
func (m *RelactionTag) FindEasyToSQL(cond *RelactionTagCond, columns ...string) (sql string, args []interface{}, err error) {
    return GenSql(m.TableName(), columns, cond, []string{}, nil, nil, []string{})
}

// FindToSQL : complex query
func (m *RelactionTag) FindToSQL(cond *RelactionTagCond, page *Page, order map[string]string, groups []string, columns ...string) (sql string, args []interface{}, err error) {
	return GenSql(m.TableName(), columns, cond, []string{}, page, order, groups)
}
*/

func (m *RelactionTag) IsSetId() bool {
	return true
}

func (m *RelactionTag) GetId() int {
	return m.Id
}

func (m *RelactionTag) IsSetContent() bool {
	return true
}

func (m *RelactionTag) GetContent() string {
	return m.Content
}

func (m *RelactionTag) IsSetValid() bool {
	return true
}

func (m *RelactionTag) GetValid() int {
	return m.Valid
}

func (m *RelactionTag) IsSetIsDel() bool {
	return true
}

func (m *RelactionTag) GetIsDel() int {
	return m.IsDel
}

func (m *RelactionTag) IsSetCreatorId() bool {
	return true
}

func (m *RelactionTag) GetCreatorId() int {
	return m.CreatorId
}

func (m *RelactionTag) IsSetCreatedAt() bool {
	return true
}

func (m *RelactionTag) GetCreatedAt() time.Time {
	return m.CreatedAt
}

func (m *RelactionTag) IsSetModifierId() bool {
	return true
}

func (m *RelactionTag) GetModifierId() int {
	return m.ModifierId
}

func (m *RelactionTag) IsSetModifiedAt() bool {
	return true
}

func (m *RelactionTag) GetModifiedAt() time.Time {
	return m.ModifiedAt
}
