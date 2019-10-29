# digest
golang http digest client

// The digest package provides an implementation of http.RoundTripper that takes
// care of HTTP Digest Authentication (http://www.ietf.org/rfc/rfc2617.txt).
// This only implements the MD5 and "auth" portions of the RFC, but that covers
// the majority of avalible server side implementations including apache web
// server.
//
// Example usage:
//
```go
func main() {
	resp, err := Get("http://admin:admin@192.168.0.108/cgi-bin/magicBox.cgi?action=getSystemInfo")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll((resp.Body))
	fmt.Println(string(body))
}
```
//	t := NewTransport("myUserName", "myP@55w0rd")
//	req, err := http.NewRequest("GET", "http://notreal.com/path?arg=1", nil)
//	if err != nil {
//		return err
//	}
//	resp, err := t.RoundTrip(req)
//	if err != nil {
//		return err
//	}
//
// OR it can be used as a client:
//
//	c, err := t.Client()
//	if err != nil {
//		return err
//	}
//	resp, err := c.Get("http://notreal.com/path?arg=1")
//	if err != nil {
//		return err
//	}
//