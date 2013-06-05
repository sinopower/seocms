/* 分类列表页(默认列表页) */

package controllers

import (
    "fmt"
    //"github.com/astaxie/beedb"
    "github.com/astaxie/beego"
)

type ListController struct {
    beego.Controller
}

const (
    SiteName = "SEOCMS"    // 网站名称
    ItemsPerPage = 10    // 列表页上每页显示文章数量
)

func (this *ListController) Get() {
    this.Layout = "layout.tpl"
    this.Data["SiteName"] = SiteName
    categoryNameEn := this.Ctx.Params[":category"]
    //Debug("Current category is `%s`.", categoryNameEn)

    // 获取分类列表，用于导航栏
    this.Data["Categories"] = GetCategories()

    if categoryNameEn == "" {    // 首页
        orm = InitDb()
        articles := []Article{}
        err = orm.OrderBy("-pubdate").Limit(ItemsPerPage).FindAll(&articles)
        Check(err)
        this.Data["Articles"] = articles

        // 设置页面标题
        this.Data["PageTitle"] = beego.AppConfig.String("appname")

        this.TplNames = "index.tpl"
    } else {    // 分类列表页
        // 获取当前分类
        orm = InitDb()
        category := Category{}
        err = orm.Where("name_en=?", categoryNameEn).Find(&category)
        Check(err)
        this.Data["Category"] = category

        // 获取当前分类文章列表
        orm = InitDb()
        articles := []Article{}
        err = orm.Where("category=?", category.Id).OrderBy("-pubdate").Limit(ItemsPerPage).FindAll(&articles)
        Check(err)
        this.Data["Articles"] = articles

        // 设置页面标题
        this.Data["PageTitle"] = fmt.Sprintf("%s相关文章_%s", category.Name, beego.AppConfig.String("appname"))

        this.TplNames = "list/category_list.tpl"
    }
}
