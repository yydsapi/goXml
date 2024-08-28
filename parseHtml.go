package main

import (
	"fmt"
	"os"
	"strings"
	"tools/xmlquery"
)

func parseHtml(fn string) {
	fmt.Println(fn)
	f, err := os.Open(fn + ".xml")
	if err != nil {
		panic(err)
	}
	doc, err := xmlquery.Parse(f)
	if err != nil {
		panic(err)
	}
	//channel := xmlquery.FindOne(doc, "//Elements")
	//gtDocTitle
	title := ""
	//str := ""
	strHtml := ""
	strHtmlHide := ""
	tr := 0
	trHide := 0
	strHtmlTop := "<!DOCTYPE html><html lang='zh-CN'><head><link rel='stylesheet' type='text/css' href='tb.css'></head>"
	d, _ := os.Open(fn)
	names, _ := d.Readdirnames(-1)
	d.Close()
	//title = strings.Join(names, "")
	//title = ""
	for _, name := range names {
		title += "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href ='" + fn + "/" + name + "' target=_blank>" + name + "</a><br>"
	}
	strHtmlTop += "<br><body><div style='width:90%;margin:0 auto;'><table><tbody><tr>"
	for i, n := range xmlquery.Find(doc, "//Elements/Field") {

		if i > 67 {
			//break
		}
		//fmt.Printf("#%d %s\n", i, n.Data)
		//str = ""
		IsContains = false
		for _, h := range HiddenItem {
			if strings.Contains(strings.ToLower(n.Attr[2].Value), strings.ToLower(h)) {
				IsContains = true
			}
		}
		if !IsContains {
			print(len(n.Attr))
			if len(n.Attr) > 2 {
				tr += 1
				strHtml += "<td width='20%'>" + n.Attr[2].Value + " : &nbsp;&nbsp;</td>"
			}
			if len(n.Attr) > 3 {
				tr += 1
				strHtml += "<td>" + n.Attr[3].Value + "&nbsp;&nbsp;</td>"
			}

		} else {
			if len(n.Attr) > 2 {
				trHide += 1
				strHtmlHide += "<td width='20%'>" + n.Attr[2].Value + " : &nbsp;&nbsp;</td>"
			}
			if len(n.Attr) > 3 {
				trHide += 1
				strHtmlHide += "<td>" + n.Attr[3].Value + "&nbsp;&nbsp;</td>"
			}
		}
		/*		for j, attr := range n.Attr {
				if attr.Name.Local == "gtDocTitle" {
					title = attr.Value
				}

				if j == 2 || j == 3 {
					println("(" + Cstr(j) + "): " + attr.Name.Local + "|" + attr.Value + "------" + Cstr(Instr(attr.Value, "$")))
					if !strings.Contains(attr.Value, "$") {
						strHtml += "<td>" + attr.Value + "&nbsp;&nbsp;</td>"
						tr++
					}
					//strHtml += "<td>" + attr.Name.Local + "</td><td>" + attr.Value + "</td>"
				}
			}*/
		//println(i)
		if tr == 4 {
			strHtml += "</tr><tr>"
			tr = 0
		}
		if trHide == 4 {
			strHtmlHide += "</tr><tr>"
			trHide = 0
		}

	}
	if tr != 0 {
		for i := tr; i < 4; i++ {
			strHtml += "<td>&nbsp;&nbsp;</td>"
		}
	}
	if trHide != 0 {
		for i := trHide; i < 4; i++ {
			strHtmlHide += "<td>&nbsp;&nbsp;</td>"
		}
	}
	strHtmlHide = "<table><tbody><tr>" + strHtmlHide + "</tr></tbody></table>"
	//println(strHtml)
	//println(strHtmlTop + title + strHtml)
	zd := `<br><font onclick="showhide()" color= "blue" onmouseover="this.style.cursor='pointer';">查看隐藏字段>>>>>></font><br><br>
<div id='hideContent' style="display: none;">
    ` + strHtmlHide + `
 </div>
<script>
function showhide(){
	if (document.getElementById("hideContent").style.display=='none'){
	document.getElementById("hideContent").style.display='block'
	}else{
	document.getElementById("hideContent").style.display='none'
	}
}
</script>`
	strHtml += "</tr></tbody></table>" + zd + "</div></body></html>"
	err = os.WriteFile("html/"+fn+".html", []byte(strHtmlTop+title+strHtml), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
