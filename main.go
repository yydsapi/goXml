package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"
	"tools/xmlquery"
)

var HiddenItem []string
var IsContains bool
var IsDetailContains bool
var MainDoc map[string]string
var DetailDoc map[string]string
var MainDocs []Pair
var DetailDocs []Pair
var MainView []Pair
var DetailView []Pair
var MainViewFinish []Pair
var DetailViewSub map[string]string
var XmlDir string
//var ShowItem []string
func main() {
	start := time.Now()
	loadConfig()
	XmlDir = "./Document"
	ds, _ := os.Open(XmlDir)
	names, _ := ds.Readdirnames(-1)
	ds.Close()
	//title = strings.Join(names, "")
	v := 0
	unid := ""
	MainDoc = make(map[string]string)
	DetailDoc = make(map[string]string)
	DetailViewSub = make(map[string]string)
	for _, name := range names {
		if path.Ext(strings.ToLower(name)) == ".xml" && strings.Index(name, "Document_") == 0 {
			v++
			n := strings.Split(name, ".")
			if strings.Index(n[0], "Detail_") > 0 {
				unid = strings.ReplaceAll(n[0], MConfigs.DetailPrefix, "")
				DetailDoc[unid] = getTime(XmlDir+"/"+name, 1)
				pid := getPid(XmlDir+"/"+name, 2)
				if pid != "" {
					DetailViewSub[unid] = pid
				}
				parseFile(unid, "2")
			} else {
				unid = strings.ReplaceAll(n[0], MConfigs.Prefix, "")
				MainDoc[unid] = getTime(XmlDir+"/"+name, 1)
				//parseFile(unid, "1")
			}
		}

	}
	MainDocs = Sort(MainDoc)
	DetailDocs = Sort(DetailDoc)
	tbview := StrHtmlTop("../tbview.css")
	MainView = getMainView("./View/Views_MainView.xml")
	DetailView = getDetailView("./View/Views_(MainDetail).xml")
	MainViewFinish = getMainView("./View/Views_MainViewFinish.xml")
	tbview += "<p>" + strings.ReplaceAll(strings.ReplaceAll(MConfigs.Prefix, "Document", ""), "_", "") + "_View</p><br><div style='width:90%;margin:0 auto;'><table><tr>"
	for _, name := range names {
		if path.Ext(strings.ToLower(name)) == ".xml" && strings.Index(name, "Document_") == 0 {
			v++
			n := strings.Split(name, ".")
			if strings.Index(n[0], "Detail_") > 0 {
			} else {
				unid = strings.ReplaceAll(n[0], MConfigs.Prefix, "")
				parseFile(unid, "1")
			}
		}

	}
	for _, b := range MainView {
		//fmt.Println(b)
		tbview += "<td style='background-color:#F0F0F0;padding-bottom: 10px;vertical-align: bottom;'>" + b.Key + "</td>"
	}
	tbview += "</tr>"
	for i := 0; i < len(MainDocs); i++ {
		f, err := os.Open(XmlDir + "/" + MConfigs.Prefix + MainDocs[i].Key + ".xml")
		if err != nil {
			panic(err)
		}
		doc, err := xmlquery.Parse(f)
		if err != nil {
			panic(err)
		}
		tbview += "<tr>"
		for _, b := range MainView {
			tbview += `<td onmouseover="this.style.cursor='pointer'" onClick=window.open('Form_Payment_html/` + MConfigs.Prefix + MainDocs[i].Key + ".html')>" + getItemValue(doc, b.Value) + "</td>"
		}
		tbview += "</tr>"
	}
	tbview += "</table></div></body></html>"
	err := os.WriteFile("html/"+"Views_MainView.html", []byte(tbview), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	tc := time.Since(start)
	fmt.Printf("main"+" : time cost = %v\n", tc)
}

func getMainView(name string) []Pair {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	doc, err := xmlquery.Parse(f)
	if err != nil {
		panic(err)
	}
	nodes := xmlquery.Find(doc, "//Elements/Field")
	pairs := make([]Pair, len(nodes))
	i := 0
	for _, n := range nodes {
		pairs[i] = Pair{n.Attr[2].Value, n.Attr[3].Value}
		i++
	}
	pairs[4] = Pair{"步骤(Steps)", "gtStatusFlag"}
	pairs[5] = Pair{MConfigs.Sort, MConfigs.Sort}
	return pairs
}
func getDetailView(name string) []Pair {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	doc, err := xmlquery.Parse(f)
	if err != nil {
		panic(err)
	}
	nodes := xmlquery.Find(doc, "//Elements/Field")
	pairs := make([]Pair, len(nodes))
	i := 0
	for _, n := range nodes {
		pairs[i] = Pair{n.Attr[2].Value, n.Attr[3].Value}
		i++
	}
	//	pairs[4] = Pair{"步骤(Steps)", "gtStatusFlag"}
	//	pairs[5] = Pair{MConfigs.Sort, MConfigs.Sort}
	return pairs
}
func getTime(name string, pos int) string {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	doc, err := xmlquery.Parse(f)
	if err != nil {
		panic(err)
	}
	for _, n := range xmlquery.Find(doc, "//Elements/Field") {
		//fmt.Println(i, n)
		if strings.Contains(strings.ToLower(n.Attr[pos].Value), strings.ToLower(MConfigs.Sort)) {
			return n.Attr[pos+1].Value
		}
	}
	return ""
}
func getPid(name string, pos int) string {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	doc, err := xmlquery.Parse(f)
	if err != nil {
		panic(err)
	}
	for _, n := range xmlquery.Find(doc, "//Elements/Field") {
		//fmt.Println(i, n)
		if len(n.Attr) == 4 {
			if strings.Contains(strings.ToLower(n.Attr[pos].Value), strings.ToLower("ParentNo")) {
				return n.Attr[pos+1].Value
			}
		}
	}
	return ""
}
func readFile(path string) string {
	fileObj, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(fileObj)
}
func getTitleItemValue(doc *xmlquery.Node, h string) (string, string) {
	for _, n := range xmlquery.Find(doc, "//Elements/Field") {
		//fmt.Println(len(n.Attr))
		if len(n.Attr) == 3 {
			if strings.Contains(strings.ToLower(n.Attr[1].Value), strings.ToLower(h)) {
				return replaceItemValue(n.Attr[1].Value), replaceItemValue(n.Attr[2].Value)
			}
		} else {
			if strings.Contains(strings.ToLower(n.Attr[2].Value), strings.ToLower(h)) {
				return replaceItemValue(n.Attr[2].Value), replaceItemValue(n.Attr[3].Value)
			}
		}
	}
	return "", ""
}
func getItemValue(doc *xmlquery.Node, h string) string {
	for _, n := range xmlquery.Find(doc, "//Elements/Field") {
		//fmt.Println(len(n.Attr))
		if len(n.Attr) == 3 {
			if strings.Contains(strings.ToLower(n.Attr[1].Value), strings.ToLower(h)) {
				return replaceItemValue(n.Attr[2].Value)
			}
		} else {
			if strings.Contains(strings.ToLower(n.Attr[2].Value), strings.ToLower(h)) {
				return replaceItemValue(n.Attr[3].Value)
			}
		}
	}
	return ""
}
func replaceItemValue(o string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(o, "];[", "];<br>["), "-------", ""), "--;", "<br>")
}
