package models

type Formula struct {
	Targets []struct {
		Field string `json:"field"`
		Prop  string `json:"prop"`
	} `json:"targets"`
	Template string `json:"template"`
	Form     string `json:"form"`
	Model    string `json:"model"`
}
type FormItem struct {
	Model       string      `json:"model"`
	Title       string      `json:"title"`
	DbType      string      `json:"dbType"`
	Table       string      `json:"table,omitempty"`
	Key         string      `json:"key"`
	Extra       string      `json:"extra,omitempty"`
	ID          string      `json:"id"`
	Type        string      `json:"type"`
	Identity    string      `json:"identity"`
	Label       string      `json:"label"`
	PlaceHolder string      `json:"placeHolder"`
	Hidden      bool        `json:"hidden"`
	Disabled    bool        `json:"disabled"`
	Default     interface{} `json:"default"`
	Prefix      string      `json:"prefix"`
	Ifshowhide  string      `json:"ifshowhide"`
	Rules       []struct {
		Type string `json:"type"`
		Msg  string `json:"msg"`
	} `json:"rules"`
	HasTranslation bool   `json:"hasTranslation"`
	HasUserID      bool   `json:"hasUserId"`
	HasEquation    bool   `json:"hasEquation"`
	Equations      string `json:"equations"`
	IsGridSearch   bool   `json:"isGridSearch"`
	GridSearch     struct {
		Grid     interface{} `json:"grid"`
		Key      interface{} `json:"key"`
		Labels   interface{} `json:"labels"`
		Multiple bool        `json:"multiple"`
	} `json:"gridSearch"`
	IsFkey   bool `json:"isFkey"`
	Relation struct {
		Table              interface{}   `json:"table"`
		Key                interface{}   `json:"key"`
		Fields             []interface{} `json:"fields"`
		FilterWithUser     []interface{} `json:"filterWithUser"`
		SortField          interface{}   `json:"sortField"`
		SortOrder          string        `json:"sortOrder"`
		Multiple           bool          `json:"multiple"`
		Filter             string        `json:"filter"`
		ParentFieldOfForm  string        `json:"parentFieldOfForm"`
		ParentFieldOfTable string        `json:"parentFieldOfTable"`
	} `json:"relation,omitempty"`
	Span struct {
		Xs int `json:"xs"`
		Sm int `json:"sm"`
		Md int `json:"md"`
		Lg int `json:"lg"`
	} `json:"span"`
	Trigger        string `json:"trigger"`
	TriggerTimeout int    `json:"triggerTimeout"`
	File           struct {
		IsMultiple bool   `json:"isMultiple"`
		Count      int    `json:"count"`
		MaxSize    int    `json:"maxSize"`
		Type       string `json:"type"`
	} `json:"file,omitempty"`
	Options          []interface{} `json:"options"`
	PasswordOption   interface{}   `json:"passwordOption"`
	GeographicOption interface{}   `json:"GeographicOption"`
	EditorType       interface{}   `json:"editorType"`
	SchemaID         string        `json:"schemaID,omitempty"`

	//subForm data
	Name            string     `json:"name"`
	SubType         string     `json:"subtype"`
	Parent          string     `json:"parent"`
	FormId          uint64     `json:"formId"`
	FormType        string     `json:"formType"`
	MinHeight       string     `json:"min_height"`
	DisableDelete   bool       `json:"disableDelete"`
	DisableCreate   bool       `json:"disableCreate"`
	ShowRowNumber   bool       `json:"showRowNumber"`
	UseTableType    bool       `json:"useTableType"`
	TableTypeColumn string     `json:"tableTypeColumn"`
	TableTypeValue  string     `json:"tableTypeValue"`
	Schema          []FormItem `json:"schema"`
}
type SCHEMA struct {
	Model         string      `json:"model"`
	Identity      string      `json:"identity"`
	Timestamp     bool        `json:"timestamp"`
	LabelPosition string      `json:"labelPosition"`
	LabelWidth    interface{} `json:"labelWidth"`
	Width         string      `json:"width"`
	Padding       int         `json:"padding"`
	Schema        []FormItem  `json:"schema"`
	UI            interface{} `json:"ui"`
	Formula       []Formula   `json:"formula"`
	Triggers      struct {
		Namespace string `json:"namespace"`
		Insert    struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"insert"`
		Update struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"update"`
	} `json:"triggers"`
	SortField string `json:"sortField"`
	SordOrder string `json:"sordOrder"`
}