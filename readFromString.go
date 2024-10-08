package main

import (
	"fmt"
	"strconv"
	"strings"
	"tools/xmlquery"
)

func Cstr(s int) string {
	return strconv.Itoa(s)
}
func Instr(s string, subs string) int {
	return strings.Index(s, subs) //-1 表示字符串 s 不包含字符串 subs

}
func Cint(s string) int {
	b, _ := strconv.Atoi(s)
	return b
}
func Cint64(s string) int64 {
	b, _ := strconv.ParseInt(s, 10, 64)
	return b
}
func ReadFormString() {
	s := `<?xml version="1.0" encoding="UTF-8" ?>
<rss version="2.0">
<channel>
  <title>W3Schools Home Page</title>
  <link>https://www.w3schools.com</link>
  <description>Free web building tutorials</description>
  <item>
    <title>RSS Tutorial</title>
    <link>https://www.w3schools.com/xml/xml_rss.asp</link>
    <description>New RSS tutorial on W3Schools</description>
  </item>
  <item>
    <title>XML Tutorial</title>
    <link>https://www.w3schools.com/xml</link>
    <description>New XML tutorial on W3Schools</description>
  </item>
</channel>
</rss>`

	doc, err := xmlquery.Parse(strings.NewReader(s))
	if err != nil {
		panic(err)
	}
	channel := xmlquery.FindOne(doc, "//channel")
	if n := channel.SelectElement("title"); n != nil {
		fmt.Printf("title: %s\n", n.InnerText())
	}
	if n := channel.SelectElement("link"); n != nil {
		fmt.Printf("link: %s\n", n.InnerText())
	}
	for i, n := range xmlquery.Find(doc, "//item/title") {
		fmt.Printf("#%d %s\n", i, n.InnerText())
	}
}
