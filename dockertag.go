package docker

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type DockerTag struct {
	Major  int64
	Minor  int64
	Patch  int64
	Prefix string
	Suffix string
}

func New(s string) (*DockerTag, error) {
	v := DockerTag{}

	if err := v.Set(s); err != nil {
		return nil, err
	}

	return &v, nil
}

func (d *DockerTag) Set(s string) error {
	//re := regexp.MustCompile(`-\d.\d.\d-`)
	//if re.MatchString(tagTemplate) {
	//	fmt.Println(strings.Split(tagTemplate, "-"))
	//}
	//fmt.Println(re.MatchString("a1c"))
	//fmt.Println(re.FindString(tagTemplate))
	//re = regexp.MustCompile("-")
	//fmt.Println(re.Split(tagTemplate,3))
	re := regexp.MustCompile(`\d.\d.\d`)
	if re.MatchString(s) {
		return errors.New("Don't match tag template!")
	}

	a := make([]string, 3, 3)
	a = strings.SplitN(s, "-", 3)
	parsedString := a[1]
	dotParts := strings.SplitN(parsedString, ".", 3)

	parsed := make([]int64, 3, 3)
	for i, v := range dotParts[:3] {
		val, err := strconv.ParseInt(v, 10, 64)
		parsed[i] = val
		if err != nil {
			return err
		}
	}

	d.Prefix = "a"
	d.Suffix = "b"
	d.Major = parsed[0]
	d.Minor = parsed[1]
	d.Patch = parsed[2]

	return nil
}
