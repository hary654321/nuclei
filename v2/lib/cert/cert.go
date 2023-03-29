package cert

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
)

type certInfo struct {
	CN string //公用名
	O  string //组织
	OU string //组织单位
}

func HttpGetCert(seedUrl string) *x509.Certificate {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	//seedUrl := "https://www.baidu.com"
	resp, err := client.Get(seedUrl)
	if err != nil {
		//slog.Println(slog.DEBUG, "请求失败", err.Error())
		return nil
	}
	if resp == nil {
		return nil
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.TLS == nil {
		return nil
	}
	if resp.TLS.PeerCertificates == nil {
		return nil
	}
	//fmt.Println(resp.TLS.PeerCertificates[0])
	certInfo := resp.TLS.PeerCertificates[0]
	//fmt.Println("颁发时间:", certInfo.NotBefore)
	//fmt.Println("过期时间:", certInfo.NotAfter)
	//fmt.Println("组织信息:", certInfo.Subject)
	//fmt.Println("颁发者:", certInfo.Issuer)
	//fmt.Println("版本:", certInfo.Version)
	//fmt.Println("PublicKeyAlgorithm", certInfo.PublicKeyAlgorithm)
	//fmt.Println("SignatureAlgorithm", certInfo.SignatureAlgorithm)
	//signStr := fmt.SDEBUG("%x", certInfo.Signature)
	//fmt.Println("Signature:", signStr)
	//fmt.Println("KeyUsage:", certInfo.KeyUsage)
	//fmt.DEBUG("%+v", reflect.TypeOf(certInfo.PublicKey))
	//PublicKey := certInfo.PublicKey
	//fmt.DEBUG("%+v", PublicKey)
	//fmt.Println("KeyUsage:", PublicKey.E)
	//fmt.Println("KeyUsage:", PublicKey.N)

	return certInfo
}
