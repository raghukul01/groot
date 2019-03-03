package groot

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"io/ioutil"
)

func RequestHandler() {
	reader := bufio.NewReader(os.Stdin)
	for {
		// read the complete text
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\n", "", -1)
		words := strings.Fields(text)

		workerCount := viper.GetInt("NODE_COUNT")
		workerAddressMap := viper.GetStringMapString("NODES")

		var addressArray []string
		nodeName := "node"

		for i := 0; i < workerCount; i++ {
			key := nodeName + strconv.Itoa(i)
			addressArray = append(addressArray, workerAddressMap[key])
		}

		if words[0] == "grep" {
			grepHandler(words, addressArray)
		} else if words[0] == "live" {
			livelogHandler(words, addressArray)
		} else if words[0] == "between" {
			timeStampHandler(words, addressArray)
		} else {
			logrus.Fatal("Invalid argument passed!")
		}
	}
}

func grepHandler(words []string, addrArray []string) {
	key := words[1]
	var wg sync.WaitGroup

	endpoint := "/api/grep/" + key

	wg.Add(len(addrArray))
	respMap := make(map[int]string)

	for idx, addr := range addrArray {
		go func(idx int, addr string) {
			defer wg.Done()
			url := addr + endpoint
			resp, err := http.Get(url)
			if err != nil {
				logrus.Info("failed to make get request")
			}
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			respMap[idx] = string(body)
		}(idx, addr)
	}
	wg.Wait()
	logrus.Info(respMap)
}

func livelogHandler(words []string, addrArray []string) {

}

func timeStampHandler(words []string, addrArray []string) {

}
