package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/prometheus/common/log"
)

const (
	DBUrl = "http://10.21.40.33:9200"
	//DBUrl = "http://10.72.17.30:29200"
	DocIndexOfBooklibBook = "avg_book_doc_index"
	DocTypeOfBooklibBook  = "avg_book_doc_type"
)

const (
	DocIndex = "ip_booklib_book"
	DocType  = "ip_booklib_book"
)
const DocMapping = `
{
	"mappings":{
		"book":{
			"properties":{
				"nick_name":{
					"type":"text",
					"store": true
				},
				"title":{
					"type":"text",
					"store": true
				},
				"desc":{
					"type":"text",
					"store": true
				}
			}
		}
	}
}`

var cli *elastic.Client

func Start() *elastic.Client {

	client, err := elastic.NewSimpleClient(elastic.SetURL(DBUrl))
	if err != nil {
		panic(err.Error())
	}

	ctx := context.Background()
	info, code, err := client.Ping(DBUrl).Do(ctx)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("32", info, code)

	return client
}

func main() {
	cli = Start()
	//IsIndexExist(DocIndex)
	//CreateIndex(DocIndex, DocType, DocMapping)
	//TypeExists("ip_booklib_book", "ip_booklib_book")
	data, err := SearchBook("时间不多", false, 1, 5)
	fmt.Println(data)
	fmt.Println(err)
}

func CreateIndex(docIndex, docType, mapping string) {
	ctx := context.Background()
	if _, err := cli.CreateIndex(docIndex).BodyString(mapping).Do(ctx); err != nil {
		log.Fatal("72", err.Error())
	}
}

func IsIndexExist(indexName string) {
	ctx := context.Background()
	yes, err := cli.IndexExists(indexName).Do(ctx)
	if err != nil {
		log.Fatal("80", err.Error())
	}
	fmt.Println("edwin 44", yes)
}

func TypeExists(indexName, typeName string) {
	ctx := context.Background()
	yes, err := cli.TypeExists().Type(typeName).Index(indexName).Do(ctx)
	if err != nil {
		log.Fatal("89", err.Error())
	}
	fmt.Println("edwin 44", yes)
}

func SearchBook(term string, highlight bool, pageNum, pageSize int32) (books []Book, err error) {
	if pageSize <= 0 {
		pageSize = 5
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	pageStart := int((pageNum - 1) * pageSize)
	ctx := context.Background()
	query := elastic.NewQueryStringQuery(term).Field("title").Field("nick_name")
	hl := elastic.NewHighlight()
	hl = hl.Fields(elastic.NewHighlighterField("title"))
	hl = hl.Fields(elastic.NewHighlighterField("nick_name"))
	var searchResult *elastic.SearchResult
	searchResult, err = cli.Search().
		Index(DocIndexOfBooklibBook).
		Type(DocTypeOfBooklibBook).
		Query(query).
		Highlight(hl).
		//Sort("collect_num", false). //不要排序，会破坏关联性
		From(pageStart).Size(int(pageSize)).
		Pretty(true).
		Do(ctx)
	if err != nil {
		panic(err.Error())
		return nil, err
	}
	var book Book

	for _, v := range searchResult.Hits.Hits {
		_ = json.Unmarshal(*v.Source, &book)
		if highlight {
			if hits, ok := v.Highlight["title"]; ok && len(hits) > 0 {
				book.Title = hits[0]
			}
			if hit, ok := v.Highlight["nick_name"]; ok && len(hit) > 0 {
				book.NickName = hit[0]
			}
		}
		books = append(books, book)
	}
	return books, nil
}

type Book struct {
	Id            int64  `xorm:"id"             json:"book_id"          part:"book_id"         editor:"book_id"`
	CateId        int64  `xorm:"cate_id"        json:"cate_id"          part:"cate_id"         editor:"cate_id"`
	UserId        int64  `xorm:"user_id"        json:"user_id"          part:"user_id"         editor:"user_id"`
	Title         string `xorm:"title"          json:"title"            part:"title"           editor:"name"`
	Intro         string `xorm:"intro"          json:"desc"             part:"-"               editor:"desc"`
	Image         string `xorm:"image"          json:"image"            part:"image"           editor:"image"`
	ImageHeight   int32  `xorm:"image_height"   json:"image_height"     part:"image_height"    editor:"image_height"`
	ImageWidth    int32  `xorm:"image_width"    json:"image_width"      part:"image_width"     editor:"image_width"`
	OnlineStatus  int32  `xorm:"online_status"  json:"online_status"    part:"online_status"   editor:"online_status"`
	WorkStatus    int32  `xorm:"work_status"    json:"work_status"      part:"work_status"     editor:"work_status"`
	Weight        int32  `xorm:"weight"         json:"weight"           part:"-"               editor:"-"`
	CollectNum    int32  `xorm:"collect_num"    json:"collect_num"      part:"-"               editor:"collect_num"`
	LikeNum       int32  `xorm:"like_num"       json:"like_num"         part:"-"               editor:"like_num"`
	ReadNum       int32  `xorm:"read_num"       json:"read_num"         part:"read_num"        editor:"read_num"`
	ShareNum      int32  `xorm:"share_num"      json:"share_num"        part:"-"               editor:"share_num"`
	CommentNum    int32  `xorm:"comment_num"    json:"comment_num"      part:"-"               editor:"comment_num"`
	Created       int32  `xorm:"created"        json:"created"          part:"-"               editor:"created"`
	Updated       int32  `xorm:"updated"        json:"updated"          part:"-"               editor:"updated"`
	BookURL       string `xorm:"-"              json:"book_url"         part:"book_url"        editor:"book_url"`
	PlayURL       string `xorm:"-"              json:"play_url"         part:"play_url"        editor:"play_url"`
	NickName      string `xorm:"-"              json:"nick_name"        part:"nick_name"       editor:"-"`
	Avatar        string `xorm:"-"              json:"avatar"           part:"avatar"          editor:"-"`
	PictureFrame  int32  `xorm:"-"              json:"picture_frame"    part:"picture_frame"   editor:"-"`
	CateName      string `xorm:"-"              json:"cate_name"        part:"cate_name"       editor:"cate_name"`
	PraiseStatus  int32  `xorm:"-"              json:"praise_status"    part:"-"               editor:"praise_status"`
	CollectStatus int32  `xorm:"-"              json:"collect_status"   part:"-"               editor:"collect_status"`
	WorkType      int32  `xorm:"-"              json:"work_type"        part:"work_type"       editor:"-"`
	Coauthors     string `xorm:"coauthors"      json:"coauthors"        part:"-"               editor:"-"`
	HImage        string `xorm:"h_image"        json:"h_image"          part:"h_image"         editor:"h_image"`
	HHeight       int32  `xorm:"h_height"       json:"h_height"         part:"h_height"        editor:"h_height"`
	HWidth        int32  `xorm:"h_width"        json:"h_width"          part:"h_width"         editor:"h_width"`
}
