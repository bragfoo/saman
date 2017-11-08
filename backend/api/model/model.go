package model

type PlatformType struct {
	Ids         string
	Name        string
	NameChinese string
}

type PlatformFans struct {
	Ids        string
	CreateTime int64
	Sum        int
	Decrease   int
	Increase   int
	Type       string
	PlatIds    string
}

type Video struct {
	Ids      string
	VideoIds string
	PlatIds  string
	Title    string
	Link     string
}

type VideoPlayAmount struct {
	Ids         string
	VideoIds    string
	PlatIds     string
	Title       string
	Link        string
	CreateTime  int64
	Sum         int
	NameChinese string
}

type PlayAmount struct {
	Ids        string
	VideoIds   string
	CreateTime int64
	Sum        int
}

type MobileData struct {
	Ids        string
	CreateTime int64
	Active     int
	Launch     int
	Channel    string
	SystemType int
	ChannelIds string
}

type Channel struct {
	Ids  string
	Name string
}

type AppUGC struct {
	Ids        string
	CreateTime int64
	Like       int
	CommentSum int
	ShareSum   int
	PicSum     int
	VideoSum   int
}

type Event struct {
	Ids          string
	Name         string
	StartDate    int64
	EndDate      int64
	TotalPeople  int
	TotalWork    int
	UploadPeople int
}

type AppData struct {
	Ids         string
	CreateTime  int64
	PicUpload   int
	VideoUpload int
	TalentSum   int
	ActiveUser  int
}
