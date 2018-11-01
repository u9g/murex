package httpclient

import (
	"errors"
	"io"
	"strconv"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/readall"
)

func cmdPost(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Json)

	if p.Parameters.Len() == 0 {
		return errors.New("URL required.")
	}

	var jHttp jsonHttp

	url, err := p.Parameters.String(0)
	if err != nil {
		return err
	}
	validateURL(&url, p.Config)

	var body io.Reader
	//var contentType string
	if p.IsMethod {
		body = p.Stdin

		//contentType, err = p.Parameters.String(1)
		//if err != nil {
		//	return err
		//}
	} else {
		body = nil
	}

	resp, err := Request(p.Context, "POST", url, body, p.Config, enableTimeout)
	if err != nil {
		return err
	}

	jHttp.Status.Code, _ = strconv.Atoi(resp.Status[:3])
	jHttp.Status.Message = resp.Status[4:]

	jHttp.Headers = resp.Header
	b, err := readall.ReadAll(p.Context, resp.Body)
	resp.Body.Close()
	jHttp.Body = string(b)
	if err != nil {
		return err
	}

	b, err = json.Marshal(jHttp, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)

	return err
}
