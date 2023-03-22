package dao_model

import (
	"time"
)

type RelactionSiiSession struct {
	Id            int       `xorm:"'id' not null pk autoincr comment('唯一值') INT(11)"`
	SessionInfoId int       `xorm:"'session_info_id' not null default 0 comment('課程編號') INT(11)"`
	Style         int       `xorm:"'style' not null default 0 comment('風格') INT(11)"`
	Item          int       `xorm:"'item' not null default 0 comment('品項') INT(11)"`
	Tag           int       `xorm:"'tag' not null default 0 comment('標籤') INT(11)"`
	Valid         int       `xorm:"'valid' not null default 1 comment('0:停用 1:啟用') TINYINT(2)"`
	IsDel         int       `xorm:"'is_del' not null default 0 comment('默认0：未删除的有效记录 1 : 逻辑删除的无效记录') TINYINT(2)"`
	CreatorId     int       `xorm:"'creator_id' not null default 0 comment('建立人员') INT(11)"`
	CreatedAt     time.Time `xorm:"'created_at' not null default CURRENT_TIMESTAMP comment('系统时间 默认(now()或 sysdate())') DATETIME <-"`
	ModifierId    int       `xorm:"'modifier_id' not null default 0 comment('修改人员') INT(11)"`
	ModifiedAt    time.Time `xorm:"'modified_at' not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME <-"`
	RefNo         *string   `xorm:"'ref_no' comment('關聯號(直播課:ShopLab LiveSessionId)') VARCHAR(45)"`
}

func (m *RelactionSiiSession) TableName() string {
	return "relaction_sii_session"
}

/*
// FindEasyToSQL : simple query
func (m *RelactionSiiSession) FindEasyToSQL(cond *RelactionSiiSessionCond, columns ...string) (sql string, args []interface{}, err error) {
    return GenSql(m.TableName(), columns, cond, []string{}, nil, nil, []string{})
}

// FindToSQL : complex query
func (m *RelactionSiiSession) FindToSQL(cond *RelactionSiiSessionCond, page *Page, order map[string]string, groups []string, columns ...string) (sql string, args []interface{}, err error) {
	return GenSql(m.TableName(), columns, cond, []string{}, page, order, groups)
}
*/

var RelactionSiiSession_RefNo_DEFAULT string

func (m *RelactionSiiSession) IsSetId() bool {
	return true
}

func (m *RelactionSiiSession) GetId() int {
	return m.Id
}

func (m *RelactionSiiSession) IsSetSessionInfoId() bool {
	return true
}

func (m *RelactionSiiSession) GetSessionInfoId() int {
	return m.SessionInfoId
}

func (m *RelactionSiiSession) IsSetStyle() bool {
	return true
}

func (m *RelactionSiiSession) GetStyle() int {
	return m.Style
}

func (m *RelactionSiiSession) IsSetItem() bool {
	return true
}

func (m *RelactionSiiSession) GetItem() int {
	return m.Item
}

func (m *RelactionSiiSession) IsSetTag() bool {
	return true
}

func (m *RelactionSiiSession) GetTag() int {
	return m.Tag
}

func (m *RelactionSiiSession) IsSetValid() bool {
	return true
}

func (m *RelactionSiiSession) GetValid() int {
	return m.Valid
}

func (m *RelactionSiiSession) IsSetIsDel() bool {
	return true
}

func (m *RelactionSiiSession) GetIsDel() int {
	return m.IsDel
}

func (m *RelactionSiiSession) IsSetCreatorId() bool {
	return true
}

func (m *RelactionSiiSession) GetCreatorId() int {
	return m.CreatorId
}

func (m *RelactionSiiSession) IsSetCreatedAt() bool {
	return true
}

func (m *RelactionSiiSession) GetCreatedAt() time.Time {
	return m.CreatedAt
}

func (m *RelactionSiiSession) IsSetModifierId() bool {
	return true
}

func (m *RelactionSiiSession) GetModifierId() int {
	return m.ModifierId
}

func (m *RelactionSiiSession) IsSetModifiedAt() bool {
	return true
}

func (m *RelactionSiiSession) GetModifiedAt() time.Time {
	return m.ModifiedAt
}

func (m *RelactionSiiSession) IsSetRefNo() bool {
	return m.RefNo != nil
}

func (m *RelactionSiiSession) GetRefNo() string {
	if !m.IsSetRefNo() {
		return RelactionSiiSession_RefNo_DEFAULT
	}
	return *m.RefNo
}
