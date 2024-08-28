package main

import (
	"fmt"
	"os"
	"strings"
	"tools/xmlquery"
)

func parseFile(unid string, tp string) {
	filename := ""
	hasDetail := false
	if tp == "1" {
		filename = "./Document/" + MConfigs.Prefix + unid + ".xml"
	} else {
		filename = "./Document/" + MConfigs.DetailPrefix + unid + ".xml"
	}
	fmt.Println(filename, tp)
	f, err := os.Open(filename)
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
	strHtmlTop := "<!DOCTYPE html><html lang='zh-CN'><head><link rel='stylesheet' type='text/css' href='../../tb.css'></head>"
	d, _ := os.Open("./Document/" + MConfigs.Prefix + unid)
	names, _ := d.Readdirnames(-1)
	d.Close()
	//title = strings.Join(names, "")
	//title = ""
	if tp == "1" {
		for _, name := range names {
			title += "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href ='../../Document/" + MConfigs.Prefix + unid + "/" + name + "' target=_blank>" + name + "</a>"
		}
	}
	if tp == "1" {
		strHtmlTop += "<p>" + strings.ReplaceAll(strings.ReplaceAll(MConfigs.Prefix, "Document", ""), "_", "") + "</p><br><body><div style='width:90%;margin:0 auto;'><table><tbody><tr>"
	} else {
		strHtmlTop += "<a href=# onclick='window.close()' style='margin-left:200px'>close</a><p>" + strings.ReplaceAll(strings.ReplaceAll(MConfigs.DetailPrefix, "Document", ""), "_", "") + "</p><br><body><div style='width:90%;margin:0 auto;'><table><tbody><tr>"
	}
	pid := ""
	for i, n := range xmlquery.Find(doc, "//Elements/Field") {

		if i > 67 {
			//break
		}
		//fmt.Printf("#%d %s\n", i, n.Data)
		//str = ""
		if tp == "1" {
			IsContains = false
			for _, h := range ShowItem {
				if len(n.Attr) == 3 {
					if strings.Contains(strings.ToLower(n.Attr[1].Value), strings.ToLower(h)) {
						IsContains = true
					}
				} else {
					if strings.Contains(strings.ToLower(n.Attr[2].Value), strings.ToLower(h)) {
						IsContains = true
					}
				}
			}
		} else {
			IsDetailContains = false
			for _, h := range ShowDetail {
				if len(n.Attr) == 3 {
					if strings.Contains(strings.ToLower(n.Attr[1].Value), strings.ToLower(h)) {
						IsDetailContains = true
					}
				} else {
					if strings.Contains(strings.ToLower(n.Attr[2].Value), strings.ToLower(h)) {
						IsDetailContains = true
					}
				}
			}
		}
		if IsContains {

		} else {
			if len(n.Attr) == 3 {
				trHide += 2
				strHtmlHide += "<td width='20%'>" + replaceItemValue(n.Attr[1].Value) + "&nbsp;&nbsp;</td><td >" + replaceItemValue(n.Attr[2].Value) + "&nbsp;&nbsp;</td>"
			}
			if len(n.Attr) == 4 {
				trHide += 2
				strHtmlHide += "<td>" + replaceItemValue(n.Attr[2].Value) + "&nbsp;&nbsp;</td><td >" + replaceItemValue(n.Attr[3].Value) + "&nbsp;&nbsp;</td>"
			}
		}
		if tp != "1" {
			if IsDetailContains {
				if len(n.Attr) == 3 {
					tr += 2
					strHtml += "<td width='20%'>" + replaceItemValue(n.Attr[1].Value) + "&nbsp;&nbsp;</td><td >" + replaceItemValue(n.Attr[2].Value) + "&nbsp;&nbsp;</td>"
				}
				if len(n.Attr) == 4 {
					tr += 2
					strHtml += "<td>" + replaceItemValue(n.Attr[2].Value) + "&nbsp;&nbsp;</td><td >" + replaceItemValue(n.Attr[3].Value) + "&nbsp;&nbsp;</td>"
				}
			}
		}
		if trHide == 4 {
			strHtmlHide += "</tr><tr>"
			trHide = 0
		}
		if tr == 4 {
			strHtml += "</tr><tr>"
			tr = 0
		}

	}
	if tp == "1" {
		tr = 0
		for _, h := range ShowItem {
			tr += 2
			a, b := getTitleItemValue(doc, h)
			if h == "No" {
				pid = b
			}
			strHtml += "<td width='20%'>" + strings.ReplaceAll(strings.ReplaceAll(a, "$quot;", "'"), ";", "<br><br>") + "&nbsp;&nbsp;</td>"
			strHtml += "<td >" + strings.ReplaceAll(strings.ReplaceAll(b, "$quot;", "'"), ";", "<br><br>") + "&nbsp;&nbsp;</td>"
			if tr == 4 {
				strHtml += "</tr><tr>"
				tr = 0
			}
		}
	} else {

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
	//detail
	hasDetail = false
	tbview := `<tr> <td colpan=4 ><font onclick='showHideDetail()' color= 'blue' onmouseover="this.style.cursor='pointer'">Detail======></font><br><br><div id='hideContentDetail' style='width:100%;margin:0 auto;'><table><tr>`
	for _, b := range DetailView {
		//fmt.Println(b)
		tbview += "<td style='padding: 7px 0;!important'>" + b.Key + "</td>"
	}
	tbview += "</tr>"
	for a, c := range DetailViewSub {
		if pid == c {
			hasDetail = true
			f, err := os.Open(XmlDir + "/" + MConfigs.DetailPrefix + a + ".xml")
			if err != nil {
				panic(err)
			}
			doc, err := xmlquery.Parse(f)
			if err != nil {
				panic(err)
			}
			tbview += "<tr>"
			for _, b := range DetailView {
				tbview += `<td style='padding: 7px 0;!important' onmouseover="this.style.cursor='pointer'" onClick=window.open('../Form_PaymentDetail_html/` + MConfigs.DetailPrefix + a + ".html')>" + getItemValue(doc, b.Value) + "</td>"
			}
			tbview += "</tr>"
		}
	}
	tbview += "</table></div></td></tr>"
	strHtmlHide = "<table><tbody><tr>" + strHtmlHide + "</tr></tbody></table></div>"
	//println(strHtml)
	//println(strHtmlTop + title + strHtml)
	zd := `<br>` + `<br>` + title + `<br><font  style="display: none;" onclick="showhide()" color= "blue" onmouseover="this.style.cursor='pointer'">查看隐藏字段>>>>>></font><br><br>
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
function showHideDetail(){
	if (document.getElementById("hideContentDetail").style.display=='none'){
	document.getElementById("hideContentDetail").style.display='block'
	}else{
	document.getElementById("hideContentDetail").style.display='none'
	}
}
</script>`
	if hasDetail {
		zd = `<br>` + tbview + zd
	}
	if tp == "2" {
		zd = ""
	}
	strHtml += "</tr></tbody></table>" + zd + "</div></body></html>"
	PathExists("./html/Form_Payment_html")
	PathExists("./html/Form_PaymentDetail_html")
	if tp == "1" {
		err = os.WriteFile("./html/Form_Payment_html/"+MConfigs.Prefix+unid+".html", []byte(strHtmlTop+strHtml), 0644)
	} else {
		err = os.WriteFile("./html/Form_PaymentDetail_html/"+MConfigs.DetailPrefix+unid+".html", []byte(strHtmlTop+strHtml), 0644)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
}
func PathExists(path string) {
	_, err := os.Stat(path)
	if err != nil {
		if err != nil {
			fmt.Println(err)
		}
	}
	if os.IsNotExist(err) {
		err = os.Mkdir(path, 0777)
		if err != nil {
			fmt.Println(err)
		}
	}
}
