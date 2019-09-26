package lib
import (
	"crypto/md5"
	"io"
	"fmt"
	"strings"
)

func Md5(data string) string{
	t := md5.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
func strFirstToUpper(str string) string {
    if len(str) < 1 {
        return ""
    }
    strArry := []rune(str)
    if strArry[0] >= 97 && strArry[0] <= 122  {
        strArry[0] -=  32
    }
    return string(strArry)
}
func ParseInput(methods []string,line string) (string,[]string){
	args := strings.Split(strings.TrimSpace(line), " ")
	args[0] = strings.ToLower(args[0])
	args[0] = strFirstToUpper(args[0])
	if Contain(methods,args[0]){
		return args[0],args[1:]
	}else{
		return "",[]string{}
	}
}

func Contain(methods []string,method string) bool{
	for  _,m := range methods {
		if m == method{
			return true
		}
	}
	return false
}