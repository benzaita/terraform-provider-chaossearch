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

type CreateObjectGroupRequest struct {
	Name             string
	SourceBucket     string
	Format           ObjectGroupFormat
	RetentionDays    int
	Filter           ObjectGroupFilter
	Options          ObjectGroupOptions
	DailyInterval    bool
	LiveEventsSqsArn string
	PartitionBy      string
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
