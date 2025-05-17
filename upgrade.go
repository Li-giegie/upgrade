package upgrade

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// Upgrade 把一个运行中的可执行文件（源文件），替换（升级）为目标文件，如果不在运行中则不要用本函数，srcPath如果不传递参数则使用os.Args[0]作为srcPath
func Upgrade(dstPath string, srcPath ...string) error {
	return upgrade(dstPath, _default(os.Args[0], srcPath...), nil)
}

// UpgradeWithOutFilename 同Upgrade类似，outFileName 参数获取源文件重命名后的文件，在windows平台下有效
func UpgradeWithOutFilename(dstPath string, outFileName *string, srcPath ...string) error {
	return upgrade(dstPath, _default(os.Args[0], srcPath...), outFileName)
}

func upgrade(dst, src string, oldFileName *string) (err error) {
	if !isFileExist(src) {
		return errors.New("newFile is not exist")
	}
	if !isFileExist(dst) {
		return errors.New("oldFilename is not exist")
	}
	if runtime.GOOS == "windows" {
		sDir, sFile := filepath.Split(src)
		tempFilename := filepath.Join(sDir, fmt.Sprintf("%d.%d.%s", time.Now().UnixNano(), rand.Uint32(), sFile))
		if oldFileName != nil {
			*oldFileName = tempFilename
		}
		if err = os.Rename(src, tempFilename); err != nil {
			return err
		}
		err = os.Rename(dst, src)
	} else {
		err = os.Rename(dst, src)
	}
	return
}

func isFileExist(filename string) bool {
	info, err := os.Stat(filename)
	if err == nil {
		return !info.IsDir()
	}
	return os.IsExist(err)
}

func _default(defaultVal string, args ...string) string {
	if len(args) == 0 {
		return defaultVal
	}
	return args[0]
}
