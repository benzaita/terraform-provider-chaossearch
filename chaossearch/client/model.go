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
	ID string
}

type ReadObjectGroupResponse struct {
	Compression      string
	FilterJSON       string
	Format           string
	LiveEventsSqsArn string
	PartitionBy      string
	SourceBucket     string
}

type CreateObjectGroupRequest struct {
	Name             string
	Compression      string
	FilterJSON       string
	Format           string
	LiveEventsSqsArn string
	PartitionBy      string
	SourceBucket     string
	Pattern			 string
}

type DeleteObjectGroupRequest struct {
	Name string
}

type UpdateObjectGroupRequest struct {
	Name           string
	IndexRetention int
}
