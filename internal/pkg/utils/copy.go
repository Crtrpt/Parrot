package utils

import (
	"github.com/jinzhu/copier"
)

func CopyArray[D, S any](dst *[]*D, src []*S) (err error) {
	for _, item := range src {
		t := new(D)
		err := copier.Copy(t, item)

		if err != nil {
			return err
		}
		*dst = append(*dst, t)
	}
	return nil
}

func CopyObject[D, S any](dst *D, src *S) (err error) {
	return copier.Copy(dst, src)
}
