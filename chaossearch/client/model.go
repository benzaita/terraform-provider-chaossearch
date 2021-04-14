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
	IndexRetention   int
	Active           bool
}

type CreateObjectGroupRequest struct {
	Name             string
	Compression      string
	FilterJSON       string
	Format           string
	LiveEventsSqsArn string
	PartitionBy      string
	SourceBucket     string
	Pattern          string
	IndexRetention   int
}

type SetActiveRequest struct {
	ObjectGroupName string
	Active          bool
}

type DeleteObjectGroupRequest struct {
	Name string
}

type UpdateObjectGroupRequest struct {
	Name           string
	IndexRetention int
}

type ReadIndexingStateRequest struct {
	ObjectGroupName string
}

type ReadBucketMetadataResponse struct {
	Metadata map[string]BucketMetadata `json:"Metadata"`
}

// There are other nested structures that we don't marshal here because we don't need to.
//  Feel free to add them in the future when needed
type BucketMetadata struct {
	Bucket string `json:"Bucket"`
	State  string `json:"State"`
}

type IndexingState struct {
	ObjectGroupName string
	Active          bool
}
