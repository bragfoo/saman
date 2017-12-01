package model

type PlatformType struct {
	Ids         string `json:"Ids"`
	Name        string `json:"Name"`
	NameChinese string `json:"NameChinese"`
}

type PlatformFans struct {
	Ids        string `json:"Ids"`
	CreateTime int64  `json:"CreateTime"`
	Sum        int    `json:"Sum"`
	Decrease   int    `json:"Decrease"`
	Increase   int    `json:"Increase"`
	Type       string `json:"Type"`
	PlatIds    string `json:"PlatIds"`
}

type Video struct {
	Ids        string `json:"Ids"`
	VideoIds   string `json:"VideoIds"`
	PlatIds    string `json:"PlatIds"`
	Title      string `json:"Title"`
	Link       string `json:"Link"`
	CreateTime int64  `json:"CreateTime"`
}

type VideoPlayAmount struct {
	Ids         string `json:"Ids"`
	VideoIds    string `json:"VideoIds"`
	PlatIds     string `json:"PlatIds"`
	Title       string `json:"Title"`
	Link        string `json:"Link"`
	CreateTime  int64  `json:"CreateTime"`
	Sum         int    `json:"Sum"`
	NameChinese string `json:"NameChinese"`
}

type PlayAmount struct {
	Ids        string `json:"Ids"`
	VideoIds   string `json:"VideoIds"`
	CreateTime int64  `json:"CreateTime"`
	Sum        int    `json:"Sum" `
}

type MobileData struct {
	Ids        string `json:"Ids"`
	CreateTime int64  `json:"CreateTime"`
	Active     int    `json:"Active"`
	Launch     int    `json:"Launch"`
	Channel    string `json:"Channel"`
	SystemType int    `json:"SystemType"`
	ChannelIds string `json:"ChannelIds"`
}

type Channel struct {
	Ids  string `json:"Ids"`
	Name string `json:"Name"`
}

type AppUGC struct {
	Ids        string `json:"Ids"`
	CreateTime int64  `json:"CreateTime"`
	Like       int64  `json:"Like"`
	CommentSum int64  `json:"CommentSum"`
	ShareSum   int64  `json:"ShareSum"`
	PicSum     int64  `json:"PicSum"`
	VideoSum   int64  `json:"VideoSum"`
	VideoStay  int64  `json:"VideoStay"`
}

type Event struct {
	Ids          string `json:"Ids"`
	Name         string `json:"Name"`
	StartDate    int64  `json:"StartDate"`
	EndDate      int64  `json:"EndDate"`
	TotalPeople  int    `json:"TotalPeople"`
	TotalWork    int    `json:"TotalWork"`
	UploadPeople int    `json:"UploadPeople"`
}

type AppData struct {
	Ids         string `json:"Ids"`
	CreateTime  int64  `json:"CreateTime"`
	PicUpload   int    `json:"PicUpload"`
	VideoUpload int    `json:"VideoUpload"`
	TalentSum   int    `json:"TalentSum"`
	ActiveUser  int    `json:"ActiveUser"`
}

type Talent struct {
	Ids        string `json:"Ids"`
	User       string `json:"User"`
	Type       string `json:"Type"`
	CreateTime int64  `json:"CreateTime"`
	Submitted  bool   `json:"Submitted"`
	SkillName  string `json:"SkillName"`
}

type Skill struct {
	Ids  string `json:"Ids"`
	Name string `json:"Name"`
}

type AppStatistics struct {
	Ids        string `json:"Ids"`
	CreateTime int64  `json:"CreateTime"`
	video      int    `json:"Video"`
	Photo      int    `json:"Photo"`
	Article    int    `json:"Article"`
	Sqsp       int    `json:"Sqsp"`
}

type DayStatistics struct {
	Type       int
	CreateTime string
	Total      int
}

type PlayExcel struct {
	IType      string
	Title      string
	Link       string
	PlayAmount string
	CreateTime int64
}

type SumExcel struct {
	Type       string
	Total      int
	Grow       int
	CreateDate int64
}

type FansExcel struct {
	Increase    int
	Cancel      int
	NetIncrease int
	Total       int
	CreateTime  int64
	Type        string
}

type PlayAmountLiner struct {
	Sum        int64  `json:"Sum"`
	CreateTime int64  `json:"CreateTime"`
	PlatIds    string `json:"PlatIds"`
}

type PlatPlayAmount struct {
	Ids        string `json:"Ids"`
	CreateTime int64  `json:"CreateTime"`
	PlatType   string `json:"PlatType"`
	Sum        int    `json:"Sum"`
	Grow       int    `json:"Grow"`
}
type WeekGrow struct {
	Ids        string `json:"Ids"`
	PlatIds    string `json:"PlatIds"`
	Grow       int    `json:"Grow"`
	CreateTime int64  `json:"CreateTime"`
}
