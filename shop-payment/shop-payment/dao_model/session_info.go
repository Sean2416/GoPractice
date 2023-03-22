package dao_model

import (
	"time"
)

type SessionInfo struct {
	Id               int        `xorm:"'id' not null pk autoincr comment('唯一值') INT(11)"`
	SchoolInfoId     int        `xorm:"'school_info_id' not null default 0 comment('學院代號') INT(11)"`
	MemberInfoId     int        `xorm:"'member_info_id' not null default 0 comment('講師會員代號') INT(11)"`
	GiveCourseId     int        `xorm:"'give_course_id' not null default 0 comment('開課人員會員代號') INT(11)"`
	ChannelId        int        `xorm:"'channel_id' not null default 0 comment('頻道編號') INT(11)"`
	RefNo1           string     `xorm:"'ref_no1' not null default '' comment('關連代號1(TutorABC LobbySn)') VARCHAR(50)"`
	RefNo2           string     `xorm:"'ref_no2' not null default '' comment('關連代號2(TutorABC FreeLobbySn)') VARCHAR(50)"`
	RefNo3           string     `xorm:"'ref_no3' not null default '' comment('關連代號3(TutorABC SessionSn)') VARCHAR(50)"`
	SessionStartTime *time.Time `xorm:"'session_start_time' comment('課程開始時間') DATETIME"`
	SessionEndTime   *time.Time `xorm:"'session_end_time' comment('課程結束時間') DATETIME"`
	Period           int        `xorm:"'period' not null default 45 comment('課程時長') INT(11)"`
	Topic            string     `xorm:"'topic' not null default '' comment('主題') VARCHAR(400)"`
	Intro            string     `xorm:"'intro' not null default '' comment('內文') VARCHAR(12000)"`
	UrlMaterial      string     `xorm:"'url_material' not null default '' comment('教材圖片') VARCHAR(200)"`
	UrlVideo         string     `xorm:"'url_video' not null default '' VARCHAR(200)"`
	UrlVideoType     int        `xorm:"'url_video_type' not null default 0 comment('影片類型') INT(11)"`
	Valid            int        `xorm:"'valid' not null default 1 comment('0:停用 1:啟用') TINYINT(2)"`
	IsLocked         int        `xorm:"'is_locked' not null default 0 comment('是否被鎖定 (不被排程更新)') TINYINT(2)"`
	AlreadyTime      *time.Time `xorm:"'already_time' comment('發佈日期') DATETIME"`
	IsTest           int        `xorm:"'is_test' not null default 0 comment('是否為測試') TINYINT(2)"`
	IsPrivate        int        `xorm:"'is_private' not null default 0 comment('是否為私有(是則不露出)') TINYINT(2)"`
	IsDel            int        `xorm:"'is_del' not null default 0 comment('默认0：未删除的有效记录 1 : 逻辑删除的无效记录') TINYINT(2)"`
	CreatorId        int        `xorm:"'creator_id' not null default 0 comment('建立人员') INT(11)"`
	CreatedAt        time.Time  `xorm:"'created_at' not null default CURRENT_TIMESTAMP comment('系统时间 默认(now()或 sysdate())') DATETIME <-"`
	ModifierId       int        `xorm:"'modifier_id' not null default 0 comment('修改人员') INT(11)"`
	ModifiedAt       time.Time  `xorm:"'modified_at' not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME <-"`
	AttachUrl        *string    `xorm:"'attach_url' comment('教材附件路徑') VARCHAR(400)"`
	RoomType         int        `xorm:"'room_type' not null default 0 comment('1:拓課
2:TMC5
3:ZOOM') INT(11)"`
}

func (m *SessionInfo) TableName() string {
	return "session_info"
}

/*
// FindEasyToSQL : simple query
func (m *SessionInfo) FindEasyToSQL(cond *SessionInfoCond, columns ...string) (sql string, args []interface{}, err error) {
    return GenSql(m.TableName(), columns, cond, []string{}, nil, nil, []string{})
}

// FindToSQL : complex query
func (m *SessionInfo) FindToSQL(cond *SessionInfoCond, page *Page, order map[string]string, groups []string, columns ...string) (sql string, args []interface{}, err error) {
	return GenSql(m.TableName(), columns, cond, []string{}, page, order, groups)
}
*/

var SessionInfo_SessionStartTime_DEFAULT time.Time
var SessionInfo_SessionEndTime_DEFAULT time.Time

var SessionInfo_AlreadyTime_DEFAULT time.Time

var SessionInfo_AttachUrl_DEFAULT string

func (m *SessionInfo) IsSetId() bool {
	return true
}

func (m *SessionInfo) GetId() int {
	return m.Id
}

func (m *SessionInfo) IsSetSchoolInfoId() bool {
	return true
}

func (m *SessionInfo) GetSchoolInfoId() int {
	return m.SchoolInfoId
}

func (m *SessionInfo) IsSetMemberInfoId() bool {
	return true
}

func (m *SessionInfo) GetMemberInfoId() int {
	return m.MemberInfoId
}

func (m *SessionInfo) IsSetGiveCourseId() bool {
	return true
}

func (m *SessionInfo) GetGiveCourseId() int {
	return m.GiveCourseId
}

func (m *SessionInfo) IsSetChannelId() bool {
	return true
}

func (m *SessionInfo) GetChannelId() int {
	return m.ChannelId
}

func (m *SessionInfo) IsSetRefNo1() bool {
	return true
}

func (m *SessionInfo) GetRefNo1() string {
	return m.RefNo1
}

func (m *SessionInfo) IsSetRefNo2() bool {
	return true
}

func (m *SessionInfo) GetRefNo2() string {
	return m.RefNo2
}

func (m *SessionInfo) IsSetRefNo3() bool {
	return true
}

func (m *SessionInfo) GetRefNo3() string {
	return m.RefNo3
}

func (m *SessionInfo) IsSetSessionStartTime() bool {
	return m.SessionStartTime != nil
}

func (m *SessionInfo) GetSessionStartTime() time.Time {
	if !m.IsSetSessionStartTime() {
		return SessionInfo_SessionStartTime_DEFAULT
	}
	return *m.SessionStartTime
}

func (m *SessionInfo) IsSetSessionEndTime() bool {
	return m.SessionEndTime != nil
}

func (m *SessionInfo) GetSessionEndTime() time.Time {
	if !m.IsSetSessionEndTime() {
		return SessionInfo_SessionEndTime_DEFAULT
	}
	return *m.SessionEndTime
}

func (m *SessionInfo) IsSetPeriod() bool {
	return true
}

func (m *SessionInfo) GetPeriod() int {
	return m.Period
}

func (m *SessionInfo) IsSetTopic() bool {
	return true
}

func (m *SessionInfo) GetTopic() string {
	return m.Topic
}

func (m *SessionInfo) IsSetIntro() bool {
	return true
}

func (m *SessionInfo) GetIntro() string {
	return m.Intro
}

func (m *SessionInfo) IsSetUrlMaterial() bool {
	return true
}

func (m *SessionInfo) GetUrlMaterial() string {
	return m.UrlMaterial
}

func (m *SessionInfo) IsSetUrlVideo() bool {
	return true
}

func (m *SessionInfo) GetUrlVideo() string {
	return m.UrlVideo
}

func (m *SessionInfo) IsSetUrlVideoType() bool {
	return true
}

func (m *SessionInfo) GetUrlVideoType() int {
	return m.UrlVideoType
}

func (m *SessionInfo) IsSetValid() bool {
	return true
}

func (m *SessionInfo) GetValid() int {
	return m.Valid
}

func (m *SessionInfo) IsSetIsLocked() bool {
	return true
}

func (m *SessionInfo) GetIsLocked() int {
	return m.IsLocked
}

func (m *SessionInfo) IsSetAlreadyTime() bool {
	return m.AlreadyTime != nil
}

func (m *SessionInfo) GetAlreadyTime() time.Time {
	if !m.IsSetAlreadyTime() {
		return SessionInfo_AlreadyTime_DEFAULT
	}
	return *m.AlreadyTime
}

func (m *SessionInfo) IsSetIsTest() bool {
	return true
}

func (m *SessionInfo) GetIsTest() int {
	return m.IsTest
}

func (m *SessionInfo) IsSetIsPrivate() bool {
	return true
}

func (m *SessionInfo) GetIsPrivate() int {
	return m.IsPrivate
}

func (m *SessionInfo) IsSetIsDel() bool {
	return true
}

func (m *SessionInfo) GetIsDel() int {
	return m.IsDel
}

func (m *SessionInfo) IsSetCreatorId() bool {
	return true
}

func (m *SessionInfo) GetCreatorId() int {
	return m.CreatorId
}

func (m *SessionInfo) IsSetCreatedAt() bool {
	return true
}

func (m *SessionInfo) GetCreatedAt() time.Time {
	return m.CreatedAt
}

func (m *SessionInfo) IsSetModifierId() bool {
	return true
}

func (m *SessionInfo) GetModifierId() int {
	return m.ModifierId
}

func (m *SessionInfo) IsSetModifiedAt() bool {
	return true
}

func (m *SessionInfo) GetModifiedAt() time.Time {
	return m.ModifiedAt
}

func (m *SessionInfo) IsSetAttachUrl() bool {
	return m.AttachUrl != nil
}

func (m *SessionInfo) GetAttachUrl() string {
	if !m.IsSetAttachUrl() {
		return SessionInfo_AttachUrl_DEFAULT
	}
	return *m.AttachUrl
}

func (m *SessionInfo) IsSetRoomType() bool {
	return true
}

func (m *SessionInfo) GetRoomType() int {
	return m.RoomType
}
