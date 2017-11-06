package model

type PlatformType struct {
	Ids         string
	Name        string
	NameChinese string
}

type PlatformFans struct {
	Ids        string
	CreateTime int
	Sum        int
	Decrease   int
	Increase   int
	Type       string
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
	CreateTime  string
	Sum         int
	NameChinese string
}

type PlayAmount struct {
	Ids        string
	VideoIds   string
	CreateTime int
	Sum        int
}

type MobileData struct {
	Ids        string
	CreateTime int
	Active     int
	Launch     int
	Channel    string
	SystemType int
}

type Channel struct {
	Ids  string
	Name string
}

type AppUGC struct {
	Ids        string
	CreateTime int
	Like       int
	CommentSum int
	ShareSum   int
	PicSum     int
	VideoSum   int
}

type Event struct {
	Ids          string
	Name         string
	StartDate    int
	EndDate      int
	TotalPeople  int
	TotalWork    int
	UploadPeople int
}
