package cert

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	"github.com/projectdiscovery/nuclei/v2/pkg/utils"
)

func TestCert(t *testing.T) {
	getCertInfo()
}

func getCertInfo() {
	domain := "www.baidu.com:443"
	conn, err := tls.DialWithDialer(&net.Dialer{
		Timeout:  time.Second * 5,
		Deadline: time.Now().Add(time.Second * 9),
	}, "tcp", domain, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	defer conn.Close()
	stats := conn.ConnectionState()
	certs := stats.PeerCertificates

	utils.PrinfI("c", certs)
	for i := range certs {
		utils.WriteJson("cert.json", certs[i])
		//fmt.Printf("%+v\n", certs[i])
		fmt.Println(certs[i].DNSNames)                        //所有域名
		fmt.Println(string(certs[i].RawSubject))              //单位
		fmt.Println(string(certs[i].Raw))                     //组织信息
		fmt.Println(string(certs[i].RawSubjectPublicKeyInfo)) //
		fmt.Println(string(certs[i].RawIssuer))               //单位
		fmt.Println(string(certs[i].RawTBSCertificate))       //单位
		//fmt.Println(certs[i].IsCA)
		//fmt.Println(certs[i].IsCA)
		//fmt.Println(certs[i].IPAddresses)
		//fmt.Println(certs[i].URIs)
		//fmt.Println(certs[i].EmailAddresses)
		//fmt.Println(certs[i].NotBefore)
		//fmt.Println(certs[i].NotAfter)
	}
}
