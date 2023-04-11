package ejercicios

import (
	"guia6/dictionary"

	"github.com/agrison/go-commons-lang/stringUtils"
)

func Frecuencia(s string) *dictionary.Dictionary[string, int] {
	dict := dictionary.NewDictionary[string, int]()
	for _, letra := range s {
		if stringUtils.IsAlpha(string(letra)) {

			var frecuencia int = 1
			if dict.Contains(string(letra)) {
				frecuencia = dict.Get(string(letra))
				frecuencia++
			}
			dict.Put(string(letra), frecuencia)
		}
	}
	return &dict
}
