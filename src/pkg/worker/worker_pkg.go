package worker

import (
	"net/http"
	"net/url"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func LiveHandler() {
	numProcess := viper.GetInt("PROCESS_COUNT")
	nodePath := viper.GetString("NODE_DIR")
	pathMap := viper.GetStringMapString("PROCESSES")

	lineCount := make(map[int]int)

	for i := 0; i < numProcess; i++ {
		lineCount[i] = 0
	}
	for {
		// check for new updates every 100 miliseconds
		time.Sleep(100 * time.Millisecond)
		for i := 0; i < numProcess; i++ {
			filePath := nodePath + pathMap[strconv.Itoa(i)]
			lineCountStr, _ := exec.Command("wc", "-l", filePath).Output()
			tmpArr := strings.Fields(string(lineCountStr))
			lineCnt, _ := strconv.Atoi(tmpArr[0])
			if lineCnt > lineCount[i] {
				lastLines, _ := exec.Command("tail",
					strconv.Itoa(lineCount[i]-lineCnt),
					filePath).
					Output()
				link := viper.GetString("LIVE_ENDPOINT") + url.PathEscape(string(lastLines))
				client := http.Client{}
				request, _ := http.NewRequest("GET", link, nil)
				resp, _ := client.Do(request)
				defer resp.Body.Close()
				lineCount[i] = lineCnt
			}
		}
	}
}
