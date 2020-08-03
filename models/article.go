package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Model

	TagId int `json:"tag_id" gorm:"index"`  // gorm:index，用于声明这个字段为索引，如果你使用了自动迁移功能则会有所影响，在不使用则无影响
	Tag Tag `json:"tag"`    // Tag字段，实际是一个嵌套的struct，它利用TagID与Tag模型相互关联，在执行查询的时候，能够达到Article、Tag关联查询的功能

	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}
/*
Article是如何关联到Tag?

func GetArticle(id int) (article Article) {
    db.Where("id = ?", id).First(&article)
    db.Model(&article).Related(&article.Tag)

    return
}
能够达到关联，首先是gorm本身做了大量的约定俗成

Article有一个结构体成员是TagID，就是外键。gorm会通过类名+ID 的方式去找到这两个类之间的关联关系
Article有一个结构体成员是Tag，就是我们嵌套在Article里的Tag结构体，我们可以通过Related进行关联查询
2、 Preload是什么东西，为什么查询可以得出每一项的关联Tag？

func GetArticles(pageNum int, pageSize int, maps interface {}) (articles []Article) {
    db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

    return
}
Preload就是一个预加载器，它会执行两条 SQL，分别是SELECT * FROM blog_articles;和SELECT * FROM blog_tag WHERE id IN (1,2,3,4);，那么在查询出结构后，gorm内部处理对应的映射逻辑，将其填充到Article的Tag中，会特别方便，并且避免了循环查询

那么有没有别的办法呢，大致是两种

gorm的Join
循环Related
综合之下，还是Preload更好，如果你有更优的方案，欢迎说一下 :)

 */

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	
	if article.ID > 0 {
		return true
	}
	return false
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	
	return 
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (article []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&article)	
	
	return 
}

func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)

	return
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Update(data)
	
	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article {
		TagId: data["tag_id"].(int),
		Title : data["title"].(string),
		Desc : data["desc"].(string),
		Content : data["content"].(string),
		CreatedBy : data["created_by"].(string),
		State : data["state"].(int),
	})
	
	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})
	
	return true
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("CreatedOn", time.Now().Unix())
	
	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}