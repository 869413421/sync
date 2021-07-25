package river

import (
	"github.com/BurntSushi/toml"
	"os"
	"path"
	"sync"
	"time"
)

type MasterInfo struct {
	sync.RWMutex

	Name string

	Position uint32

	filePath string

	lastSaveTime time.Time
}

// LoadMasterInfo 加载本机记录游标文件
func LoadMasterInfo(dir string) (*MasterInfo, error) {
	var m MasterInfo

	if len(dir) == 0 {
		return &m, nil
	}

	m.filePath = path.Join(dir, "master.info")
	m.lastSaveTime = time.Now()

	file, err := os.Open(m.filePath)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	if os.IsNotExist(err) {
		return &m, nil
	}

	defer file.Close()

	_, err = toml.DecodeReader(file, &m)
	return &m, err
}
