package articles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArticle(t *testing.T) {
	testData := `input: |-
  This is where input prompts should be provided.
  New line here.
content: |-
  Some test content
output: |-
  Some text to instruct AI what to output
`

	article, err := NewArticle([]byte(testData))
	assert.Nil(t, err)

	assert.EqualValues(
		t,
		`This is where input prompts should be provided.
New line here.`,
		article.InputPrompt,
	)
	assert.EqualValues(t, "Some test content", article.Content)
	assert.EqualValues(t, "Some text to instruct AI what to output", article.OutputPrompt)
}
