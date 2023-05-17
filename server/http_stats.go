package monitor

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"../api-monitor/models"
	"github.com/tcnksm/go-httpstat"
)

func HTTPStat(method, URL string, triggerID int64) {
	req, err := http.NewRequest(method, URL, nil)
	if err != nil {
		fmt.Printf("error_req:%v \n", err)
	}

	var result httpstat.Result
	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error_client:%v \n", err)
	}

	if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
		fmt.Printf("error_ioutil:%v \n", err)
	}

	defer res.Body.Close()
	result.End(time.Now())

	fmt.Println(result)

	httpStat := &models.HTTPStat{
		TriggerID:  triggerID,
		DNS:        int(result.DNSLookup / time.Millisecond),
		TCP:        int(result.TCPConnection / time.Millisecond),
		TLS:        int(result.TLSHandshake / time.Millisecond),
		Processing: int(result.ServerProcessing / time.Millisecond),
		Transfer:   int(result.ContentTransfer(time.Now()) / time.Millisecond),
		StatusCode: res.StatusCode,
	}

	if err := models.CreateHttpStat(httpStat); err != nil {
		fmt.Println("[HTTPStat] can't CreateHttpStat:", err)
	}

}
