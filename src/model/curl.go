package model

import (
	"fmt"
	"reflect"
	"strings"
)

type Curl struct {
	// most used options
	// NOTE: ignore -O, -o, -C, -#, -D, -c
	// should not support -h, -M, -V
	Header      []string `short:"H" long:"header" description:"curl headers"`
	Method      string   `short:"X" long:"request" description:"request command"`
	Body        string   `short:"d" long:"data" description:"curl request body"`
	Form        []string `short:"F" long:"form" description:"Specify multipart MIME data"`
	FormString  []string `long:"form-string" description:"Specify multipart MIME data"`
	User        string   `short:"u" long:"user" description:"Server user and password"`
	Verbose     []bool   `short:"v" long:"verbose" description:"make the operation more talkative"`
	RemoteName  []bool   `short:"O" long:"remote-name" description:"Write output to a file named as the remote file"`
	Head        []bool   `short:"I" long:"head" description:"Show document info only"`
	Location    []bool   `short:"L" long:"location" description:"Follow redirects"`
	Get         []bool   `short:"G" long:"get" description:"Put the post data in the URL and use GET"`
	UserAgent   string   `short:"A" long:"user-agent" description:"Send User-Agent <name> to server"`
	Basic       []string `long:"basic" description:"Use HTTP Basic Authentication"`
	Proxy       string   `short:"x" long:"proxy" description:"Use this proxy"`
	ProxyUser   string   `short:"U" long:"proxy-user" description:"Proxy user and password"`
	Proxytunnel []bool   `short:"p" long:"proxytunnel" description:"Operate through an HTTP proxy tunnel (using CONNECT)"`
	Compressed  []bool   `long:"compressed" description:"Request compressed response"`
	// other useful options
	ContinueAt          string `short:"C" long:"continue-at" description:"Resumed transfer offset"`
	MaxTime             string `short:"m" long:"max-time" description:"Maximum time allowed for the transfer"`
	Insecure            []bool `short:"k" long:"insecure" description:"Allow insecure server connections when using SSL"`
	NoBuffer            []bool `short:"N" long:"no-buffer" description:"Display transfer progress as a bar"`
	DumpHeader          string `short:"D" long:"dump-header" description:"Write the received headers to <filename>"`
	ProcessBar          []bool `short:"#" long:"progress-bar" description:"Disable buffering of the output stream"`
	Output              string `short:"o" long:"output" description:"Write to file instead of stdout"` // need ignore
	Include             []bool `short:"i" long:"include" description:"Include protocol response headers in the output"`
	JunkSessionCookies  []bool `short:"j" long:"junk-session-cookies" description:"Ignore session cookies read from file"`
	ListOnly            []bool `short:"l" long:"list-only" description:"List only mode"`
	Ipv4                []bool `short:"4" long:"ipv4" description:"Resolve names to IPv4 addresses"`
	Ipv6                []bool `short:"6" long:"ipv6" description:"Resolve names to IPv6 addresses"`
	Http0d9             []bool `long:"http0.9" description:"Allow HTTP 0.9 responses"`
	Http1d0             []bool `short:"0" long:"http1.0" description:"Use HTTP 1.0"`
	Http1d1             []bool `long:"http1.1" description:"Use HTTP 1.1"`
	Http2               []bool `long:"http2" description:"Use HTTP 2"`
	Http2PriorKnowledge []bool `long:"http2-prior-knowledge" description:"Use HTTP 2 without HTTP/1.1 Upgrade"`
	Ssl                 []bool `long:"ssl" description:"Try SSL/TLS"`
	Sslv2               []bool `short:"2" long:"sslv2" description:"Use SSLv2"`
	Sslv3               []bool `short:"3" long:"sslv3" description:"Use SSLv3"`
	Tlsv1               []bool `short:"1" long:"tlsv1" description:"Use TLSv1.0 or greater"`
	Tlsv1d0             []bool `long:"tlsv1.0" description:"Use TLSv1.0 or greater"`
	Tlsv1d1             []bool `long:"tlsv1.1" description:"Use TLSv1.1 or greater"`
	Tlsv1d2             []bool `long:"tlsv1.2" description:"Use TLSv1.2 or greater"`
	Tlsv1d3             []bool `long:"tlsv1.3" description:"Use TLSv1.3 or greater"`
	IgnoreContentLength []bool `long:"ignore-content-length" description:"Ignore the size of the remote resource"`
	Quote               []bool `short:"Q" long:"quote" description:"Send command(s) to server before transfer"`
	UseAscii            []bool `short:"B" long:"use-ascii" description:"Use ASCII/text transfer"`
	Silent              []bool `short:"s" long:"silent" description:"Silent mode"`
	LimitRate           string `long:"limit-rate" description:"Limit transfer speed to RATE"`
	Url                 string `long:"url" description:"URL to work with"`
	Cookie              string `short:"b" long:"cookie" description:"Send cookies from string/file"`
	CookieJar           string `short:"c" long:"cookie-jar" description:"Write cookies to <filename> after operation"`
	Cert                string `short:"E" long:"Cert" description:"Client certificate file and password"`
	Rawcurl             string
}

func (c *Curl) GetHeader() []string {
	if c == nil {
		return make([]string, 0)
	}
	return c.Header
}

func (c *Curl) GetMethod() string {
	if c == nil {
		return ""
	}
	return c.Method
}

func (c *Curl) SetMethod(m string) {
	if c == nil {
		return
	}
	c.Method = m
}

func (c *Curl) GetUrl() string {
	if c == nil {
		return ""
	}
	return c.Url
}

func (c *Curl) GetBody() string {
	if c == nil {
		return ""
	}
	return c.Body
}

func (c Curl) BuildCurlCmd() string {
	cmdParts := make([]string, 0)
	boolParts := make([]string, 0)
	val := reflect.ValueOf(c)
	typ := reflect.TypeOf(c)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		shortTag := fieldType.Tag.Get("short")
		longTag := fieldType.Tag.Get("long")
		if longTag == "url" {
			continue
		}
		flag := fmt.Sprintf("--%s", longTag)
		if shortTag != "" {
			flag = fmt.Sprintf("-%s", shortTag)
		}

		fieldIntf := field.Interface()
		switch fieldIntf.(type) {
		case string:
			cmd := field.String()
			if cmd != "" {
				cmdParts = append(cmdParts, fmt.Sprintf("%s '%s'", flag, cmd))
			}
		case []string:
			cmds := fieldIntf.([]string)
			if len(cmds) > 0 {
				for _, cmd := range cmds {
					cmdParts = append(cmdParts, fmt.Sprintf("%s '%s'", flag, cmd))
				}
			}
		case []bool:
			cmds := fieldIntf.([]bool)
			if len(cmds) > 0 {
				boolParts = append(boolParts, fmt.Sprintf("%s", flag))
			}
		}
	}

	var res strings.Builder
	res.WriteString("curl '")
	res.WriteString(c.Url)
	res.WriteString("'")
	if len(cmdParts) > 0 {
		res.WriteString(" \\\n")
		res.WriteString(strings.Join(cmdParts, " \\\n"))
	}
	if len(boolParts) > 0 {
		res.WriteString(" \\\n")
		res.WriteString(strings.Join(boolParts, " "))
	}

	return res.String()
}
