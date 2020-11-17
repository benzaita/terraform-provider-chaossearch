package client

type Bucket struct {
	Name         string `xml:"Name"`
	CreationDate string `xml:"CreationDate"`
}

type BucketCollection struct {
	Buckets []Bucket `xml:"Bucket"`
}

type ListBucketsResponse struct {
	BucketsCollection BucketCollection `xml:"Buckets"`
}

type ReadObjectGroupRequest struct {
	Id string
}

type ReadObjectGroupResponse struct {
	SourceBucket string
	Format       struct {
		Type        string `json:"_type"`
		Horizontal  bool
		StripPrefix bool `json:"stripPrefix"`
	}
	IndexRetention   int
	Filter           ObjectGroupFilter
	Options          ObjectGroupOptions
	DailyInterval    bool
	LiveEventsSqsArn string
	PartitionBy      string
}

type CreateObjectGroupRequest struct {
	Name             string
	SourceBucket     string
	Format           ObjectGroupFormat
	IndexRetention   int
	Filter           ObjectGroupFilter
	Options          ObjectGroupOptions
	DailyInterval    bool
	LiveEventsSqsArn string
	PartitionBy      string
}

type DeleteObjectGroupRequest struct {
	Name string
}

type UpdateObjectGroupRequest struct {
	Name           string
	IndexRetention int
}

type ObjectGroupFormat struct {
	Type        string
	Horizontal  bool
	StripPrefix bool
}

type ObjectGroupFilter struct {
	Operator string
	Operands []ObjectGroupFieldFilter
}

type ObjectGroupFieldFilter struct {
	FieldName   string
	ValuePrefix string
	ValueRegex  string
}

type ObjectGroupOptions struct {
	IgnoreIrregular bool
	Compression     string
}
