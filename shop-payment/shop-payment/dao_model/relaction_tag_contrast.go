package dao_model

import (
	"time"
)

type RelactionTagContrast struct {
	Id         int       `xorm:"'id' not null pk autoincr comment('唯一值') INT(11)"`
	Style      int       `xorm:"'style' not null default 0 comment('風格編號') INT(11)"`
	Item       int       `xorm:"'item' not null default 0 comment('品項編號') INT(11)"`
	Tag        int       `xorm:"'tag' not null default 0 comment('tag') INT(11)"`
	Contrast   string    `xorm:"'contrast' not null default '' comment('對照詞') VARCHAR(200)"`
	Valid      int       `xorm:"'valid' not null default 1 comment('0:停用 1:啟用') TINYINT(2)"`
	IsDel      int       `xorm:"'is_del' not null default 0 comment('默认0：未删除的有效记录 1 : 逻辑删除的无效记录') TINYINT(2)"`
	CreatorId  int       `xorm:"'creator_id' not null default 0 comment('建立人员') INT(11)"`
	CreatedAt  time.Time `xorm:"'created_at' not null default CURRENT_TIMESTAMP comment('系统时间 默认(now()或 sysdate())') DATETIME <-"`
	ModifierId int       `xorm:"'modifier_id' not null default 0 comment('修改人员') INT(11)"`
	ModifiedAt time.Time `xorm:"'modified_at' not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME <-"`
}

func (m *RelactionTagContrast) TableName() string {
	return "relaction_tag_contrast"
}

/*
// FindEasyToSQL : simple query
func (m *RelactionTagContrast) FindEasyToSQL(cond *RelactionTagContrastCond, columns ...string) (sql string, args []interface{}, err error) {
    return GenSql(m.TableName(), columns, cond, []string{}, nil, nil, []string{})
}

// FindToSQL : complex query
func (m *RelactionTagContrast) FindToSQL(cond *RelactionTagContrastCond, page *Page, order map[string]string, groups []string, columns ...string) (sql string, args []interface{}, err error) {
	return GenSql(m.TableName(), columns, cond, []string{}, page, order, groups)
}
*/

func (m *RelactionTagContrast) IsSetId() bool {
	return true
}

func (m *RelactionTagContrast) GetId() int {
	return m.Id
}

func (m *RelactionTagContrast) IsSetStyle() bool {
	return true
}

func (m *RelactionTagContrast) GetStyle() int {
	return m.Style
}

func (m *RelactionTagContrast) IsSetItem() bool {
	return true
}

func (m *RelactionTagContrast) GetItem() int {
	return m.Item
}

func (m *RelactionTagContrast) IsSetTag() bool {
	return true
}

func (m *RelactionTagContrast) GetTag() int {
	return m.Tag
}

func (m *RelactionTagContrast) IsSetContrast() bool {
	return true
}

func (m *RelactionTagContrast) GetContrast() string {
	return m.Contrast
}

func (m *RelactionTagContrast) IsSetValid() bool {
	return true
}

func (m *RelactionTagContrast) GetValid() int {
	return m.Valid
}

func (m *RelactionTagContrast) IsSetIsDel() bool {
	return true
}

func (m *RelactionTagContrast) GetIsDel() int {
	return m.IsDel
}

func (m *RelactionTagContrast) IsSetCreatorId() bool {
	return true
}

func (m *RelactionTagContrast) GetCreatorId() int {
	return m.CreatorId
}

func (m *RelactionTagContrast) IsSetCreatedAt() bool {
	return true
}

func (m *RelactionTagContrast) GetCreatedAt() time.Time {
	return m.CreatedAt
}

func (m *RelactionTagContrast) IsSetModifierId() bool {
	return true
}

func (m *RelactionTagContrast) GetModifierId() int {
	return m.ModifierId
}

func (m *RelactionTagContrast) IsSetModifiedAt() bool {
	return true
}

func (m *RelactionTagContrast) GetModifiedAt() time.Time {
	return m.ModifiedAt
}
