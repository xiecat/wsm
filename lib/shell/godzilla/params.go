package godzilla

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ExecParams struct {
	Template    string `json:"template"`
	Command     string `json:"command"`
	CurrPath    string `json:"currPath"`
	RealCommand string `json:"realCommand"`
}

func (e *ExecParams) SetDefaultAndCheckValue() error {
	if len(e.Template) == 0 {
		e.Template = `cmd /c "{command}"`
	}
	if len(e.Command) == 0 {
		return errors.New("command is empty")
	}
	if len(e.CurrPath) == 0 {
		e.CurrPath = "C:/"
	}
	c := fmt.Sprintf(`cd /d "%s"&%s`, e.CurrPath, e.Command)
	if len(e.RealCommand) == 0 {
		e.RealCommand = strings.ReplaceAll(e.Template, "{command}", c)
	}
	return nil
}

var fileNameEmptyError = errors.New("file name is empty")

type GetFiles struct {
	DirName string `json:"dirName"`
}

func (g GetFiles) SetDefaultAndCheckValue() error {
	if len(g.DirName) == 0 {
		return errors.New("dir name is empty")
	}
	return nil
}

type NewDir struct {
	DirName string `json:"dirName"`
}

func (n NewDir) SetDefaultAndCheckValue() error {
	if len(n.DirName) == 0 {
		return errors.New("dir name is empty")
	}
	return nil
}

type DownloadFile struct {
	FileName string `json:"fileName"`
}

func (d DownloadFile) SetDefaultAndCheckValue() error {
	if len(d.FileName) == 0 {
		return fileNameEmptyError
	}
	return nil
}

type UploadFile struct {
	FileName  string `json:"fileName"`
	FileValue []byte `json:"fileValue"`
}

func (u UploadFile) SetDefaultAndCheckValue() error {
	if len(u.FileName) == 0 {
		return fileNameEmptyError
	}
	if len(u.FileValue) == 0 {
		return errors.New("file content is empty")
	}
	return nil
}

type CopyFile struct {
	SrcFileName  string `json:"srcFileName"`
	DestFileName string `json:"destFileName"`
}

func (c CopyFile) SetDefaultAndCheckValue() error {
	if len(c.SrcFileName) == 0 {
		return errors.New("src file name is empty")
	}
	if len(c.DestFileName) == 0 {
		return errors.New("dest file name is empty")
	}
	return nil
}

type MoveFile struct {
	SrcFileName  string `json:"srcFileName"`
	DestFileName string `json:"destFileName"`
}

func (m MoveFile) SetDefaultAndCheckValue() error {
	if len(m.SrcFileName) == 0 {
		return errors.New("src file name is empty")
	}
	if len(m.DestFileName) == 0 {
		return errors.New("dest file name is empty")
	}
	return nil
}

type DeleteFile struct {
	FileName string `json:"fileName"`
}

func (d DeleteFile) SetDefaultAndCheckValue() error {
	if len(d.FileName) == 0 {
		return fileNameEmptyError
	}
	return nil
}

type NewFile struct {
	FileName string `json:"fileName"`
}

func (n NewFile) SetDefaultAndCheckValue() error {
	if len(n.FileName) == 0 {
		return fileNameEmptyError
	}
	return nil
}

type BigFileUpload struct {
	FileName     string `json:"fileName"`
	FileContents []byte `json:"fileContents"`
	Position     int    `json:"position"`
}

func (b BigFileUpload) SetDefaultAndCheckValue() error {
	if len(b.FileName) == 0 {
		return fileNameEmptyError
	}
	if len(b.FileContents) == 0 {
		return errors.New("file content is empty")
	}
	return nil
}

type BigFileDownload struct {
	FileName    string `json:"fileName"`
	Position    int    `json:"position"`
	ReadByteNum int    `json:"readByteNum"`
}

func (b BigFileDownload) SetDefaultAndCheckValue() error {
	if len(b.FileName) == 0 {
		return fileNameEmptyError
	}
	return nil
}

type GetFileSize struct {
	FileName string `json:"fileName"`
}

func (g *GetFileSize) SetDefaultAndCheckValue() error {
	if len(g.FileName) == 0 {
		return errors.New("file name is empty")
	}
	return nil
}

// FileRemoteDown 指定一个 url 让目标下载
type FileRemoteDown struct {
	// eg. https://github.com/xxx/1.exe
	Url string `json:"url"`
	// 文件在目标服务器上保存的路径
	SaveFile string `json:"saveFile"`
}

func (f *FileRemoteDown) SetDefaultAndCheckValue() error {
	if len(f.Url) == 0 {
		return errors.New("url is empty")
	}
	u, err := url.Parse(f.Url)
	if err != nil {
		return err
	}
	f.Url = u.String()
	if len(f.SaveFile) == 0 {
		return errors.New("save file path is empty")
	}
	return nil
}

type FileAttr string

const (
	FileBasicAttr FileAttr = "fileBasicAttr"
	FileTimeAttr  FileAttr = "fileTimeAttr"
)

// FixFileAttr 修改文件时间戳、权限
type FixFileAttr struct {
	// 要修改的文件路径
	FileName string `json:"fileName"`
	// 修改时间戳还是权限
	FileAttr FileAttr
	// 如果选择修改时间戳，那么格式为 2006-01-02 15:04:05
	// 如果选择修改权限，那么格式为 RWX
	Attr string `json:"attr"`
}

func (s *FixFileAttr) SetDefaultAndCheckValue() error {
	if len(s.FileName) == 0 {
		return errors.New("file name is empty")
	}
	if len(s.Attr) == 0 {
		return errors.New("attr name is empty")
	}
	if s.FileAttr == FileTimeAttr {
		ok, _ := regexp.MatchString(`^\d{4}-\d{1,2}-\d{1,2} \d{2}:\d{2}:\d{2}$`, s.Attr)
		if !ok {
			return errors.New(fmt.Sprintf("%s Incorrect time format,eg. 2006-01-02 15:04:05", s.Attr))
		}
		tt, err := time.Parse("2006-01-02 15:04:05", s.Attr)
		if err != nil {
			return err
		}
		s.Attr = strconv.FormatInt(tt.Unix(), 10)
	}
	return nil
}

type DBManagerParams struct {
	DBType     string
	DBHost     string
	DBPort     int
	DBUsername string
	DBPassword string
	ExecType   string
	ExecSql    string
	DBCharset  string
	CurrentDB  string
}

func (d DBManagerParams) SetDefaultAndCheckValue() error {
	if len(d.DBType) == 0 {
		return errors.New("db type is empty")
	}
	if len(d.DBHost) == 0 {
		return errors.New("db host is empty")
	}
	if d.DBPort == 0 {
		return errors.New("db port is error")
	}
	if len(d.ExecSql) == 0 {
		return errors.New("db sql is empty")
	}
	return nil
}
