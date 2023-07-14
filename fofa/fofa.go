package fofa

import (
	"encoding/base64"
	"errors"
	"github.com/astaxie/beego/httplib"
	"github.com/tidwall/gjson"
	"strings"
)

func Search(email string, key string, query string) (result []string, err error) {
	query = base64.StdEncoding.EncodeToString([]byte(query))
	sResult, err := httplib.Get("https://fofa.info/api/v1/search/all?email=" + email + "&key=" + key + "&size=10000&qbase64=" + query).
		String()
	if err != nil {
		return
	}
	for _, v := range gjson.Parse(sResult).Get("results").Array() {
		if len(v.Array()) == 0 {
			continue
		}
		vv := v.Array()[0].Str
		if strings.HasPrefix(vv, "https://") {
			result = append(result, vv)
		} else {
			result = append(result, "http://"+vv)
		}
	}
	if len(result) == 0 {
		err = errors.New(sResult)
	}
	return
}
