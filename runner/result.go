package runner

import (
	"bufio"
	"context"
	"os"
	"strings"

	"github.com/Tw1ps/ksubdomain/core"
	"github.com/Tw1ps/ksubdomain/core/gologger"
)

func (r *runner) handleResult(ctx context.Context) {
	var isWrite bool = false
	var err error
	var windowsWidth int

	if r.options.Silent {
		windowsWidth = 0
	} else {
		windowsWidth = core.GetWindowWith()
	}

	if r.options.Output != "" {
		isWrite = true
	}
	var foutput *os.File
	if isWrite {
		foutput, err = os.OpenFile(r.options.Output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
		if err != nil {
			gologger.Errorf("写入结果文件失败：%s\n", err.Error())
		}
	}
	onlyDomain := r.options.OnlyDomain
	notPrint := r.options.NotPrint
	for result := range r.recver {
		var msg string

		if onlyDomain {
			msg = result.Subdomain
		} else {
			var content = []string{
				result.Subdomain,
			}
			content = append(content, result.Answers...)
			msg = strings.Join(content, " => ")
		}

		if !notPrint {
			screenWidth := windowsWidth - len(msg) - 1
			if !r.options.Silent {
				if windowsWidth > 0 && screenWidth > 0 {
					gologger.Silentf("\r%s% *s\n", msg, screenWidth, "")
				} else {
					gologger.Silentf("\r%s\n", msg)
				}
				// 打印一下结果,可以看得更直观
				r.PrintStatus()
			} else {
				gologger.Silentf("%s\n", msg)
			}
		}

		if isWrite {
			w := bufio.NewWriter(foutput)
			_, err = w.WriteString(msg + "\n")
			if err != nil {
				gologger.Errorf("写入结果文件失败.Err:%s\n", err.Error())
			}
			_ = w.Flush()
		}
	}
}
