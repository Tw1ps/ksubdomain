package core

import (
	"bufio"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

func RandomStr(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz1234567890")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

// LinesInFile 读取文件 返回每行的数组
func LinesInFile(fileName string) ([]string, error) {
	result := []string{}
	f, err := os.Open(fileName)
	if err != nil {
		return result, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			result = append(result, line)
		}
	}
	return result, nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func GetWindowWith() int {
	w, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0
	}
	return w
}

func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func SliceToString(items []string) string {
	ret := strings.Builder{}
	ret.WriteString("[")
	ret.WriteString(strings.Join(items, ","))
	ret.WriteString("]")
	return ret.String()
}

func Load_text(filename string) (string, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(f), err
}

type Decompose struct {
	Subdomain string `json:"subdomain"`
	Domain    string `json:"domain"`
	Suffix    string `json:"suffix"`
}

func Dismantl_domain(host string) Decompose {
	var ret Decompose
	public_suffix_list, err := Load_text("./data/public_suffix_list.bat")
	if err != nil {
		return ret
	}
	data := strings.Split(host, ".")
	for i, _ := range data {
		suffix := strings.Join(data[i:], ".")
		if find := strings.Contains(public_suffix_list, suffix); find {
			ret.Subdomain = strings.Join(data[:i-1], ".")
			ret.Domain = data[i-1]
			ret.Suffix = suffix
		}
	}
	return ret
}
