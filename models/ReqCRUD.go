package models

type ReqCRUD struct {
	//general
	Operation  string   `xml:"operation" json:"operation"`
	Table      string   `xml:"table" json:"table"`
	Columns    []string `xml:"columns" json:"columns"`
	Conditions string   `xml:"conditions" json:"conditions"`
	TTL        int      `xml:"ttl" json:"ttl"`
	Timestamp  int      `xml:"timestamp" json:"timestamp"`
	//select
	Join           string `xml:"join" json:"join"`
	JoinCondition  string `xml:"joincondition" json:"joinCondition"`
	Orderby        string `xml:"orderby" json:"orderby"`
	Limit          int    `xml:"limit" json:"limit"`
	AllowFiltering bool   `xml:"allowfiltering" json:"allowFiltering"`
	//insert
	ColumnsList  []string   `xml:"columnslist" json:"columnsList"`
	InsertValues [][]string `xml:"insertvalues" json:"insertValues"`
	IfNotExists  bool       `xml:"ifnotexists" json:"ifNotExists"`
	//update
	ExistsCondition string `xml:"existscondition" json:"existsCondition"`
	Set             string `xml:"set" json:"set"`
}
