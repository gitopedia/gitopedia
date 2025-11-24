package frontmatter

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type Data struct {
	ID       string                 `yaml:"id"`
	Title    string                 `yaml:"title"`
	Slug     string                 `yaml:"slug"`
	Tags     []string               `yaml:"tags"`
	People   []string               `yaml:"people"`
	Orgs     []string               `yaml:"orgs"`
	Places   []string               `yaml:"places"`
	DateRefs []string               `yaml:"date_refs"`
	Rest     map[string]interface{} `yaml:",inline"`
}

func Parse(content []byte) (Data, error) {
	s := string(content)
	if !strings.HasPrefix(s, "---") {
		return Data{}, fmt.Errorf("missing front matter")
	}
	parts := strings.SplitN(s, "---", 3)
	if len(parts) < 3 {
		return Data{}, fmt.Errorf("malformed front matter")
	}

	var fm Data
	if err := yaml.Unmarshal([]byte(parts[1]), &fm); err != nil {
		return Data{}, err
	}
	return fm, nil
}

