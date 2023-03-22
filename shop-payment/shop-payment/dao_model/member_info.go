package dao_model

import (
	"time"
)

type MemberInfo struct {
	Id                  int        `xorm:"'id' not null pk autoincr comment('唯一值') INT(11)"`
	ClientSn            int        `xorm:"'client_sn' not null default 0 INT(11)"`
	ClientSnStep        int        `xorm:"'client_sn_step' not null default 0 comment('ClientSn資料同步為步驟幾') INT(11)"`
	Email               string     `xorm:"'email' not null default '' comment('信箱') VARCHAR(100)"`
	Phone               string     `xorm:"'phone' not null default '' comment('電話') VARCHAR(20)"`
	PhoneCode           string     `xorm:"'phone_code' not null default '886' VARCHAR(10)"`
	ChnName             string     `xorm:"'chn_name' not null default '' comment('中文姓名') VARCHAR(45)"`
	EnFname             string     `xorm:"'en_fname' not null default '' comment('英文名(前)') VARCHAR(45)"`
	EnMname             string     `xorm:"'en_mname' not null default '' comment('英文名(中)') VARCHAR(45)"`
	EnLname             string     `xorm:"'en_lname' not null default '' comment('英文名(後)') VARCHAR(45)"`
	NickName            string     `xorm:"'nick_name' not null default '' comment('綽號') VARCHAR(45)"`
	Birthday            *time.Time `xorm:"'birthday' comment('生日') DATETIME"`
	Gender              int        `xorm:"'gender' not null default 1 comment('性別 1:男 2:女') TINYINT(2)"`
	AddressCountry      *string    `xorm:"'address_country' comment('地區') VARCHAR(45)"`
	AddressCity         *string    `xorm:"'address_city' comment('城市') VARCHAR(45)"`
	Address             *string    `xorm:"'address' comment('地址') VARCHAR(200)"`
	AddressCode         *string    `xorm:"'address_code' comment('地址區碼') VARCHAR(45)"`
	Type                int        `xorm:"'type' not null default 0 comment('會員類型\n1:一般會員 2:學院講師 3:學院講師(自動建立)') INT(11)"`
	UrlPicture          string     `xorm:"'url_picture' not null default '' comment('頭像圖片連結') VARCHAR(400)"`
	UrlBanner           string     `xorm:"'url_banner' not null default '' comment('主頁Banner圖片連結') VARCHAR(400)"`
	AboutMe             string     `xorm:"'about_me' not null default '' comment('關於我') VARCHAR(4000)"`
	MoreInformation     *string    `xorm:"'more_information' comment('更多資訊') VARCHAR(4000)"`
	Professional        string     `xorm:"'professional' not null default '' comment('專業領域') VARCHAR(100)"`
	UrlMyPage           *string    `xorm:"'url_my_page' comment('個人網址') VARCHAR(45)"`
	ShareOpenRecord     *int       `xorm:"'share_open_record' default 0 comment('公開分享我開的課') TINYINT(2)"`
	ShareArticle        *int       `xorm:"'share_article' default 0 comment('公開分享我的文章') TINYINT(2)"`
	SharePodcast        *int       `xorm:"'share_podcast' default 0 comment('公開分享我的Podcast') TINYINT(2)"`
	ShareReservedRecord *int       `xorm:"'share_reserved_record' default 0 comment('公開分享我修的課') TINYINT(2)"`
	SharePortfolio      *int       `xorm:"'share_portfolio' default 0 comment('公開分享我的課程作品') TINYINT(2)"`
	ShareFollower       *int       `xorm:"'share_follower' default 0 comment('公開分享我的追蹤') TINYINT(2)"`
	ShareCollection     *int       `xorm:"'share_collection' default 0 comment('公開分享我的收藏') TINYINT(2)"`
	IsGiftFreeContract  int        `xorm:"'is_gift_free_contract' not null default 0 comment('是否贈送過免費合約(8776)') TINYINT(2)"`
	Valid               int        `xorm:"'valid' not null default 1 comment('0:停用 1:啟用') TINYINT(2)"`
	IsDel               int        `xorm:"'is_del' not null default 0 comment('默认0：未删除的有效记录 1 : 逻辑删除的无效记录') TINYINT(2)"`
	CreatorId           int        `xorm:"'creator_id' not null default 0 comment('建立人员') INT(11)"`
	CreatedAt           time.Time  `xorm:"'created_at' not null default CURRENT_TIMESTAMP comment('系统时间 默认(now()或 sysdate())') DATETIME <-"`
	ModifierId          int        `xorm:"'modifier_id' not null default 0 comment('修改人员') INT(11)"`
	Updated             time.Time  `xorm:"'updated' not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
}

func (m *MemberInfo) TableName() string {
	return "member_info"
}


// FindEasyToSQL : simple query
func (m *MemberInfo) FindEasyToSQL(cond *MemberInfoCond, columns ...string) (sql string, args []interface{}, err error) {
    return GenSql(m.TableName(), columns, cond, []string{}, nil, nil, []string{})
}

// FindToSQL : complex query
func (m *MemberInfo) FindToSQL(cond *MemberInfoCond, page *Page, order map[string]string, groups []string, columns ...string) (sql string, args []interface{}, err error) {
	return GenSql(m.TableName(), columns, cond, []string{}, page, order, groups)
}


var MemberInfo_Birthday_DEFAULT time.Time

var MemberInfo_AddressCountry_DEFAULT string
var MemberInfo_AddressCity_DEFAULT string
var MemberInfo_Address_DEFAULT string
var MemberInfo_AddressCode_DEFAULT string

var MemberInfo_MoreInformation_DEFAULT string

var MemberInfo_UrlMyPage_DEFAULT string
var MemberInfo_ShareOpenRecord_DEFAULT int
var MemberInfo_ShareArticle_DEFAULT int
var MemberInfo_SharePodcast_DEFAULT int
var MemberInfo_ShareReservedRecord_DEFAULT int
var MemberInfo_SharePortfolio_DEFAULT int
var MemberInfo_ShareFollower_DEFAULT int
var MemberInfo_ShareCollection_DEFAULT int

func (m *MemberInfo) IsSetId() bool {
	return true
}

func (m *MemberInfo) GetId() int {
	return m.Id
}

func (m *MemberInfo) IsSetClientSn() bool {
	return true
}

func (m *MemberInfo) GetClientSn() int {
	return m.ClientSn
}

func (m *MemberInfo) IsSetClientSnStep() bool {
	return true
}

func (m *MemberInfo) GetClientSnStep() int {
	return m.ClientSnStep
}

func (m *MemberInfo) IsSetEmail() bool {
	return true
}

func (m *MemberInfo) GetEmail() string {
	return m.Email
}

func (m *MemberInfo) IsSetPhone() bool {
	return true
}

func (m *MemberInfo) GetPhone() string {
	return m.Phone
}

func (m *MemberInfo) IsSetPhoneCode() bool {
	return true
}

func (m *MemberInfo) GetPhoneCode() string {
	return m.PhoneCode
}

func (m *MemberInfo) IsSetChnName() bool {
	return true
}

func (m *MemberInfo) GetChnName() string {
	return m.ChnName
}

func (m *MemberInfo) IsSetEnFname() bool {
	return true
}

func (m *MemberInfo) GetEnFname() string {
	return m.EnFname
}

func (m *MemberInfo) IsSetEnMname() bool {
	return true
}

func (m *MemberInfo) GetEnMname() string {
	return m.EnMname
}

func (m *MemberInfo) IsSetEnLname() bool {
	return true
}

func (m *MemberInfo) GetEnLname() string {
	return m.EnLname
}

func (m *MemberInfo) IsSetNickName() bool {
	return true
}

func (m *MemberInfo) GetNickName() string {
	return m.NickName
}

func (m *MemberInfo) IsSetBirthday() bool {
	return m.Birthday != nil
}

func (m *MemberInfo) GetBirthday() time.Time {
	if !m.IsSetBirthday() {
		return MemberInfo_Birthday_DEFAULT
	}
	return *m.Birthday
}

func (m *MemberInfo) IsSetGender() bool {
	return true
}

func (m *MemberInfo) GetGender() int {
	return m.Gender
}

func (m *MemberInfo) IsSetAddressCountry() bool {
	return m.AddressCountry != nil
}

func (m *MemberInfo) GetAddressCountry() string {
	if !m.IsSetAddressCountry() {
		return MemberInfo_AddressCountry_DEFAULT
	}
	return *m.AddressCountry
}

func (m *MemberInfo) IsSetAddressCity() bool {
	return m.AddressCity != nil
}

func (m *MemberInfo) GetAddressCity() string {
	if !m.IsSetAddressCity() {
		return MemberInfo_AddressCity_DEFAULT
	}
	return *m.AddressCity
}

func (m *MemberInfo) IsSetAddress() bool {
	return m.Address != nil
}

func (m *MemberInfo) GetAddress() string {
	if !m.IsSetAddress() {
		return MemberInfo_Address_DEFAULT
	}
	return *m.Address
}

func (m *MemberInfo) IsSetAddressCode() bool {
	return m.AddressCode != nil
}

func (m *MemberInfo) GetAddressCode() string {
	if !m.IsSetAddressCode() {
		return MemberInfo_AddressCode_DEFAULT
	}
	return *m.AddressCode
}

func (m *MemberInfo) IsSetType() bool {
	return true
}

func (m *MemberInfo) GetType() int {
	return m.Type
}

func (m *MemberInfo) IsSetUrlPicture() bool {
	return true
}

func (m *MemberInfo) GetUrlPicture() string {
	return m.UrlPicture
}

func (m *MemberInfo) IsSetUrlBanner() bool {
	return true
}

func (m *MemberInfo) GetUrlBanner() string {
	return m.UrlBanner
}

func (m *MemberInfo) IsSetAboutMe() bool {
	return true
}

func (m *MemberInfo) GetAboutMe() string {
	return m.AboutMe
}

func (m *MemberInfo) IsSetMoreInformation() bool {
	return m.MoreInformation != nil
}

func (m *MemberInfo) GetMoreInformation() string {
	if !m.IsSetMoreInformation() {
		return MemberInfo_MoreInformation_DEFAULT
	}
	return *m.MoreInformation
}

func (m *MemberInfo) IsSetProfessional() bool {
	return true
}

func (m *MemberInfo) GetProfessional() string {
	return m.Professional
}

func (m *MemberInfo) IsSetUrlMyPage() bool {
	return m.UrlMyPage != nil
}

func (m *MemberInfo) GetUrlMyPage() string {
	if !m.IsSetUrlMyPage() {
		return MemberInfo_UrlMyPage_DEFAULT
	}
	return *m.UrlMyPage
}

func (m *MemberInfo) IsSetShareOpenRecord() bool {
	return m.ShareOpenRecord != nil
}

func (m *MemberInfo) GetShareOpenRecord() int {
	if !m.IsSetShareOpenRecord() {
		return MemberInfo_ShareOpenRecord_DEFAULT
	}
	return *m.ShareOpenRecord
}

func (m *MemberInfo) IsSetShareArticle() bool {
	return m.ShareArticle != nil
}

func (m *MemberInfo) GetShareArticle() int {
	if !m.IsSetShareArticle() {
		return MemberInfo_ShareArticle_DEFAULT
	}
	return *m.ShareArticle
}

func (m *MemberInfo) IsSetSharePodcast() bool {
	return m.SharePodcast != nil
}

func (m *MemberInfo) GetSharePodcast() int {
	if !m.IsSetSharePodcast() {
		return MemberInfo_SharePodcast_DEFAULT
	}
	return *m.SharePodcast
}

func (m *MemberInfo) IsSetShareReservedRecord() bool {
	return m.ShareReservedRecord != nil
}

func (m *MemberInfo) GetShareReservedRecord() int {
	if !m.IsSetShareReservedRecord() {
		return MemberInfo_ShareReservedRecord_DEFAULT
	}
	return *m.ShareReservedRecord
}

func (m *MemberInfo) IsSetSharePortfolio() bool {
	return m.SharePortfolio != nil
}

func (m *MemberInfo) GetSharePortfolio() int {
	if !m.IsSetSharePortfolio() {
		return MemberInfo_SharePortfolio_DEFAULT
	}
	return *m.SharePortfolio
}

func (m *MemberInfo) IsSetShareFollower() bool {
	return m.ShareFollower != nil
}

func (m *MemberInfo) GetShareFollower() int {
	if !m.IsSetShareFollower() {
		return MemberInfo_ShareFollower_DEFAULT
	}
	return *m.ShareFollower
}

func (m *MemberInfo) IsSetShareCollection() bool {
	return m.ShareCollection != nil
}

func (m *MemberInfo) GetShareCollection() int {
	if !m.IsSetShareCollection() {
		return MemberInfo_ShareCollection_DEFAULT
	}
	return *m.ShareCollection
}

func (m *MemberInfo) IsSetIsGiftFreeContract() bool {
	return true
}

func (m *MemberInfo) GetIsGiftFreeContract() int {
	return m.IsGiftFreeContract
}

func (m *MemberInfo) IsSetValid() bool {
	return true
}

func (m *MemberInfo) GetValid() int {
	return m.Valid
}

func (m *MemberInfo) IsSetIsDel() bool {
	return true
}

func (m *MemberInfo) GetIsDel() int {
	return m.IsDel
}

func (m *MemberInfo) IsSetCreatorId() bool {
	return true
}

func (m *MemberInfo) GetCreatorId() int {
	return m.CreatorId
}

func (m *MemberInfo) IsSetCreatedAt() bool {
	return true
}

func (m *MemberInfo) GetCreatedAt() time.Time {
	return m.CreatedAt
}

func (m *MemberInfo) IsSetModifierId() bool {
	return true
}

func (m *MemberInfo) GetModifierId() int {
	return m.ModifierId
}

func (m *MemberInfo) IsSetUpdated() bool {
	return true
}

func (m *MemberInfo) GetUpdated() time.Time {
	return m.Updated
}
