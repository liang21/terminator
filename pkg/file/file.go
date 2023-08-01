package file

import (
	"github.com/liang21/terminator/pkg/config"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func LoadFile(path string) (bs *config.Bootstrap, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &bs)
	if err != nil {
		return nil, err
	}
	return bs, nil
}
