package web

import (
	"imagine-flow/vars"
	"io/ioutil"
	"os"
	"path"
)

func getPublicImage(name string) []byte {
	p := path.Join(vars.PublicImagesFolder, name)
	f, e := os.Open(p)
	if e != nil {
		return nil
	}
	inBuf, _ := ioutil.ReadAll(f)
	return inBuf
}
