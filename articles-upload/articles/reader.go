package articles

import (
	"io/ioutil"

	awsUtils "github.com/EdgeJay/psg-navi-bot/articles-upload/aws"
)

type Reader struct {
	S3Path awsUtils.S3Path
}

func NewReader(s3Path awsUtils.S3Path) *Reader {
	return &Reader{
		S3Path: s3Path,
	}
}

// Loads file from S3
func (r *Reader) LoadAndParseFile() (*Article, error) {
	s3Client := awsUtils.GetS3Client()
	obj, err := s3Client.GetObject(awsUtils.CreateGetObjectInput(r.S3Path))
	if err != nil {
		return nil, err
	}

	b, readErr := ioutil.ReadAll(obj.Body)
	if readErr != nil {
		return nil, readErr
	}

	return NewArticle(b)
}
