package articles

import "gopkg.in/yaml.v3"

type Article struct {
	InputPrompt  string `yaml:"input"`
	Content      string `yaml:"content"`
	OutputPrompt string `yaml:"output"`
}

func NewArticle(data []byte) (*Article, error) {
	var article Article
	if err := yaml.Unmarshal(data, &article); err != nil {
		return nil, err
	}
	return &article, nil
}
