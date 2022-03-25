package piglatin

import (
    "regexp"
    "strings"
)

func Sentence(sentence string) string {
	var pigLatin = make([]string, 0)
    for _, word := range strings.Fields(sentence) {
        pigLatin = append(pigLatin, translate(word))    
    }
    return strings.Join(pigLatin, " ")
}

func translate(word string) string {
    var result strings.Builder
    if match, _ := regexp.MatchString("^(yt|xr|[aeiou]|[aeiou]qu).*", word); match {
        result.WriteString(word)
    } else if match, _ := regexp.MatchString("^(thr|sch|[^aeiou]qu).*", word); match {
        result.WriteString(word[3:])
        result.WriteString(word[:3])
    } else if match, _ := regexp.MatchString("^(ch|qu|th|rh).*", word); match {
        result.WriteString(word[2:])
        result.WriteString(word[:2])
    } else {
        result.WriteString(word[1:])
        result.WriteString(word[:1])
    }
    result.WriteString("ay")
    return result.String() 
}
