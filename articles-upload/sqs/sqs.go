package sqs

type MessageBody struct {
	Records []Record `json:"Records"`
}

type Record struct {
	EventName string    `json:"eventName"`
	S3        S3Message `json:"s3"`
}

type S3Message struct {
	Bucket S3Bucket `json:"bucket"`
	Object S3Object `json:"object"`
}

type S3Bucket struct {
	Name string `json:"name"`
}

type S3Object struct {
	Key string `json:"key"`
}
