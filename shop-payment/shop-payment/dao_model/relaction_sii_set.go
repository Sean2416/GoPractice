package dao_model

import (
	"time"
)

type RelactionSiiSet struct {
	Id         int       `xorm:"'id' not null pk autoincr comment('唯一值') INT(11)"`
	Style      int       `xorm:"'style' not null default 0 comment('風格') unique(uniq_s_i_t) INT(11)"`
	Item       int       `xorm:"'item' not null default 0 comment('品項') unique(uniq_s_i_t) INT(11)"`
	Tag        int       `xorm:"'tag' not null default 0 comment('標籤') unique(uniq_s_i_t) INT(11)"`
	Sort       int       `xorm:"'sort' not null default 0 comment('顯示順序') INT(11)"`
	Valid      int       `xorm:"'valid' not null default 1 comment('0:停用 1:啟用') TINYINT(2)"`
	IsDel      int       `xorm:"'is_del' not null default 0 comment('默认0：未删除的有效记录 1 : 逻辑删除的无效记录') TINYINT(2)"`
	CreatorId  int       `xorm:"'creator_id' not null default 0 comment('建立人员') INT(11)"`
	CreatedAt  time.Time `xorm:"'created_at' not null default CURRENT_TIMESTAMP comment('系统时间 默认(now()或 sysdate())') DATETIME <-"`
	ModifierId int       `xorm:"'modifier_id' not null default 0 comment('修改人员') INT(11)"`
	ModifiedAt time.Time `xorm:"'modified_at' not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME <-"`
}

func (m *RelactionSiiSet) TableName() string {
	return "relaction_sii_set"
}

/*
// FindEasyToSQL : simple query
func (m *RelactionSiiSet) FindEasyToSQL(cond *RelactionSiiSetCond, columns ...string) (sql string, args []interface{}, err error) {
    return GenSql(m.TableName(), columns, cond, []string{}, nil, nil, []string{})
}

// FindToSQL : complex query
func (m *RelactionSiiSet) FindToSQL(cond *RelactionSiiSetCond, page *Page, order map[string]string, groups []string, columns ...string) (sql string, args []interface{}, err error) {
	return GenSql(m.TableName(), columns, cond, []string{}, page, order, groups)
}
*/

func (m *RelactionSiiSet) IsSetId() bool {
	return true
}

func (m *RelactionSiiSet) GetId() int {
	return m.Id
}

func (m *RelactionSiiSet) IsSetStyle() bool {
	return true
}

func (m *RelactionSiiSet) GetStyle() int {
	return m.Style
}

func (m *RelactionSiiSet) IsSetItem() bool {
	return true
}

func (m *RelactionSiiSet) GetItem() int {
	return m.Item
}

func (m *RelactionSiiSet) IsSetTag() bool {
	return true
}

func (m *RelactionSiiSet) GetTag() int {
	return m.Tag
}

func (m *RelactionSiiSet) IsSetSort() bool {
	return true
}

func (m *RelactionSiiSet) GetSort() int {
	return m.Sort
}

func (m *RelactionSiiSet) IsSetValid() bool {
	return true
}

func (m *RelactionSiiSet) GetValid() int {
	return m.Valid
}

func (m *RelactionSiiSet) IsSetIsDel() bool {
	return true
}

func (m *RelactionSiiSet) GetIsDel() int {
	return m.IsDel
}

func (m *RelactionSiiSet) IsSetCreatorId() bool {
	return true
}

func (m *RelactionSiiSet) GetCreatorId() int {
	return m.CreatorId
}

func (m *RelactionSiiSet) IsSetCreatedAt() bool {
	return true
}

func (m *RelactionSiiSet) GetCreatedAt() time.Time {
	return m.CreatedAt
}

func (m *RelactionSiiSet) IsSetModifierId() bool {
	return true
}

func (m *RelactionSiiSet) GetModifierId() int {
	return m.ModifierId
}

func (m *RelactionSiiSet) IsSetModifiedAt() bool {
	return true
}

func (m *RelactionSiiSet) GetModifiedAt() time.Time {
	return m.ModifiedAt
}
