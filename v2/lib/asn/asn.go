package asn

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hary654321/nuclei/v2/core/slog"
)

var BaseUrlRt = "https://tools.ipip.net/as.php?a=rt&ip="

func GetAsnRt(ip string) map[string]interface{} {

	req, _ := http.NewRequest(http.MethodGet, BaseUrlRt+ip, nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	//req.Header.Add("Content-Length", "26")
	//req.Header.Add("Host", "tools.ipip.net")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	//req.Header.Add("Origin", "https://tools.ipip.net")
	//req.Header.Add("Referer", "https://tools.ipip.net/as.php")
	//req.Header.Add("X-Requested-With", "XMLHttpRequest")
	//req.Header.Add("sec-ch-ua", "Not?A_Brand";v="8", "Chromium";v="108", "Google Chrome";v="108")
	//req.Header.Add("sec-ch-ua-platform", "Windows")
	//req.Header.Add("Cookie", " __root_domain_v=.ipip.net; _qddaz=QD.810970811957720; lastSE=baidu; _ga=GA1.2.35085502.1672133495; _gid=GA1.2.1971748185.1672133495; https_waf_cookie=6797258e-51e5-4ac291087444bdc46389597bef154a261e17; LOVEAPP_SESSID=a7930364022e2f9c77c2e87ac7d4724d4c245d5f; Hm_lvt_6b4a9140aed51e46402f36e099e37baf=1670811959,1672133495,1672189861; _gat=1; Hm_lpvt_6b4a9140aed51e46402f36e099e37baf=1672193761; _qdda=3-1.1; _qddab=3-5qg2et.lc70zndh")

	var ResponseJsonRt []interface{}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Println(slog.DEBUG, err.Error())
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	//[{"asn":"AS16509","prefix":"3.33.144.0\/20","DEBUG":"AMAZON-02 - Amazon.com, Inc., US","paths":[{"time":"2022-12-25","as_path":"57695 16509","path":[{"asn":"AS57695","DEBUG":"MISAKA - Misaka Network, Inc., US"},{"asn":"AS16509","DEBUG":"AMAZON-02 - Amazon.com, Inc., US"}]},{"time":"2022-12-25","as_path":"57695 138195 16509","path":[{"asn":"AS57695","DEBUG":"MISAKA - Misaka Network, Inc., US"},{"asn":"AS138195","DEBUG":"MOACKCOLTD-AS-AP - MOACK.Co.LTD, KR"},{"asn":"AS16509","DEBUG":"AMAZON-02 - Amazon.com, Inc., US"}]},{"time":"2022-12-28 04:55:32","as_path":"57695 60068 174 16509","path":[{"asn":"AS57695","DEBUG":"MISAKA - Misaka Network, Inc., US"},{"asn":"AS60068","DEBUG":"CDN77 - Datacamp Limited, GB"},{"asn":"AS174","DEBUG":"COGENT-174 - Cogent Communications, US"},{"asn":"AS16509","DEBUG":"AMAZON-02 - Amazon.com, Inc., US"}]},{"time":"2022-12-25","as_path":"57695 60068 12956 16509","path":[{"asn":"AS57695","DEBUG":"MISAKA - Misaka Network, Inc., US"},{"asn":"AS60068","DEBUG":"CDN77 - Datacamp Limited, GB"},{"asn":"AS12956","DEBUG":"TELXIUS - TELEFONICA GLOBAL SOLUTIONS SL, ES"},{"asn":"AS16509","DEBUG":"AMAZON-02 - Amazon.com, Inc., US"}]},{"time":"2022-12-25","as_path":"57695 36236 16509","path":[{"asn":"AS57695","DEBUG":"MISAKA - Misaka Network, Inc., US"},{"asn":"AS36236","DEBUG":"NETACTUATE - NetActuate, Inc, US"},{"asn":"AS16509","DEBUG":"AMAZON-02 - Amazon.com, Inc., US"}]},{"time":"2022-12-25","as_path":"57695 32097 16509","path":[{"asn":"AS57695","DEBUG":"MISAKA - Misaka Network, Inc., US"},{"asn":"AS32097","DEBUG":"WII - WholeSale Internet, Inc., US"},{"asn":"AS16509","DEBUG":"AMAZON-02 - Amazon.com, Inc., US"}]},{"time":"2022-12-25","as_path":"57695 60068 33891 16509","path":[{"asn":"AS57695","DEBUG":"MISAKA - Misaka Network, Inc., US"},{"asn":"AS60068","DEBUG":"CDN77 - Datacamp Limited, GB"},{"asn":"AS33891","DEBUG":"CORE-BACKBONE - Core-Backbone GmbH, DE"},{"asn":"AS16509","DEBUG":"AMAZON-02 - Amazon.com, Inc., US"}]},{"time":"2022-12-25","as_path":"57695 56630 16509","path":[{"asn":"AS57695","DEBUG":"MISAKA - Misaka Network, Inc., US"},{"asn":"AS56630","DEBUG":"MELBICOM-EU-AS - Melbikomas UAB, LT"},{"asn":"AS16509","DEBUG":"AMAZON-02 - Amazon.com, Inc., US"}]},{"time":"2022-12-25","as_path":"57695 9009 16509","path":[{"asn":"AS57695","DEBUG":"MISAKA - Misaka Network, Inc., US"},{"asn":"AS9009","DEBUG":"M247 - M247 Europe SRL, RO"},{"asn":"AS16509","DEBUG":"AMAZON-02 - Amazon.com, Inc., US"}]},{"time":"2022-12-25","as_path":"57695 56630 3257 16509","path":[{"asn":"AS57695","DEBUG":"MISAKA - Misaka Network, Inc., US"},{"asn":"AS56630","DEBUG":"MELBICOM-EU-AS - Melbikomas UAB, LT"},{"asn":"AS3257","DEBUG":"GTT-BACKBONE - GTT Communications Inc., US"},{"asn":"AS16509","DEBUG":"AMAZON-02 - Amazon.com, Inc., US"}]}]}]slog.Println(slog.DEBUG, "返回:%s", string(body))
	if err != nil {
		slog.Println(slog.DEBUG, err.Error())
		return nil
	}
	if err = json.Unmarshal(body, &ResponseJsonRt); err != nil {
		slog.Println(slog.DEBUG, err.Error())
		return nil
	}

	if len(ResponseJsonRt) < 1 {
		slog.Println(slog.DEBUG, "ResponseJsonRt空数据")
		return nil
	}

	ObjI := ResponseJsonRt[0]

	outObj := ObjI.(map[string]interface{})

	//utils.PrinfI("outObj",outObj)

	return outObj
}
