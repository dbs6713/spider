package crawler

import (
	"io"
	"io/ioutil"
	"regexp"
)

func ParseHTML(r io.ReadCloser) (body string, urls []string, err error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", nil, err
	}
	bs := string(b)
	re := regexp.MustCompile(`(http|https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`)
	u := re.FindAllString(bs, -1)

	return bs, u, nil
}
