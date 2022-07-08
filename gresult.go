package wsm

import (
	"encoding/json"
	"errors"
	"strings"
)

type Mode string

const (
	Raw       Mode = "raw"
	BasicInfo Mode = "basicInfo"
	FileOpt   Mode = "fileOpt"
)

type gResult struct {
	mode   Mode
	Raw    []byte
	Body   map[string]string
	Status bool
}

func newGResult(raw []byte, mode Mode) *gResult {
	return &gResult{
		Raw:  raw,
		mode: mode,
	}
}

func (g *gResult) Parser() error {
	switch g.mode {
	case Raw:
		result := make(map[string]string, 1)
		result["raw"] = string(g.Raw)
		g.Body = result
	case BasicInfo:
		rawList := strings.Split(string(g.Raw), "\n")
		result := make(map[string]string, len(rawList))
		for _, r := range rawList {
			if len(r) == 0 {
				continue
			}
			rL := strings.SplitN(r, " : ", 2)
			key := rL[0]
			value := ""
			if len(rL) == 2 {
				value = rL[1]
			}
			result[key] = value
		}
		g.Body = result
	case FileOpt:
		jsonStr, err := parserFileOptToMap(string(g.Raw))
		if err != nil {
			return err
		}
		g.Body = jsonStr
	}

	return nil
}

func (g *gResult) ToMap() map[string]string {
	return g.Body
}

func (g *gResult) ToString() string {
	return string(g.Raw)
}

type fileInfo struct {
	Name         string `json:"name"`
	LastModified string `json:"lastModified"`
	FileType     string `json:"type"`
	FileSize     string `json:"size"`
	Perm         string `json:"perm"`
}

func parserFileOptToMap(raw string) (map[string]string, error) {
	var fileInfoList []string
	result := make(map[string]string, 2)

	rawList := strings.Split(raw, "\n")
	if rawList[0] != "ok" {
		return nil, errors.New("目标返回异常,无法正常格式化数据")
	}
	result["msg"] = rawList[0]
	result["currentDir"] = rawList[1]
	for i := 2; i < len(rawList); i++ {
		var fi fileInfo
		rawStr := strings.Split(rawList[i], "\t")
		if len(rawStr) == 5 {
			if rawStr[1] == "0" {
				fi.FileType = "dir"
			} else {
				fi.FileType = "file"
			}
			fi.Name = rawStr[0]
			fi.LastModified = rawStr[2]
			fi.FileSize = rawStr[3]
			fi.Perm = rawStr[4]
			jsons, err := json.Marshal(fi)
			if err != nil {
				return nil, err
			}
			fileInfoList = append(fileInfoList, string(jsons))
		}
	}
	result["fileList"] = strings.Join(fileInfoList, " , ")
	return result, nil
}
