package replace

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/pkg/errors"
)

// Do replaces substring in file by pattern
// used as replacement for tools like "sed" as they some of them are not POSIX-compatible on different platforms
func Do(file, pattern, value string) error {
	path, err := filepath.Abs(file)
	if err != nil {
		return errors.Wrap(err, "file path")
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.Wrap(err, "not found")
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return errors.Wrap(err, "regexp")
	}
	res := re.ReplaceAllString(string(content), value)
	return ioutil.WriteFile(path, []byte(res), os.ModeAppend)
}
