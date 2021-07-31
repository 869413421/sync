package river

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/siddontang/go/ioutil2"
	"os"
	"path"
	"sync"
	"sync/pkg/logger"
	"time"
)

type MasterInfo struct {
	sync.RWMutex

	Name string `toml:"bin_name"`

	Pos uint32 `toml:"bin_pos"`

	filePath string

	lastSaveTime time.Time
}

// LoadMasterInfo 加载本机记录游标文件
func LoadMasterInfo(dir string) (*MasterInfo, error) {
	//1.检查路径
	var m MasterInfo
	if len(dir) == 0 {
		return &m, nil
	}

	//2.加载路径
	m.filePath = path.Join(dir, "master.info")
	m.lastSaveTime = time.Now()

	//3.创建文件夹
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	//4.判断文件是否存在
	f, err := os.Open(m.filePath)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	} else if os.IsNotExist(err) {
		return &m, nil
	}
	defer f.Close()

	//5.返回游标对象
	_, err = toml.DecodeReader(f, &m)
	return &m, err
}

// Position 获取binlog名称和游标
func (m *MasterInfo) Position() mysql.Position {
	m.RLock()
	defer m.RUnlock()

	return mysql.Position{
		Name: m.Name,
		Pos:  m.Pos,
	}
}

// Close 关闭保存binlog名称和游标
func (m *MasterInfo) Close() error {
	pos := m.Position()
	return m.Save(pos)
}

// Save 保存binlog名称和游标
func (m *MasterInfo) Save(pos mysql.Position) error {
	//1.加锁
	m.Lock()
	defer m.Unlock()

	//2.保存binlog名称游标
	m.Name = pos.Name
	m.Pos = pos.Pos
	if len(m.filePath) == 0 {
		return nil
	}

	//3.如果上次保存时间小于一秒，不处理
	n := time.Now()
	if n.Sub(m.lastSaveTime) < time.Second {
		return nil
	}

	m.lastSaveTime = n
	var buf bytes.Buffer
	e := toml.NewEncoder(&buf)

	e.Encode(m)

	var err error
	if err = ioutil2.WriteFileAtomic(m.filePath, buf.Bytes(), 0644); err != nil {
		logger.Danger("canal save master info to file %s err %v", m.filePath, err)
	}

	return err
}
