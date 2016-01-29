package controllers

import (
    "github.com/revel/revel"
    "rucos/app/models"  // modelディレクトリを読み込み
    "rucos/app/routes"  // routesを読み込み
)

type App struct {
	*revel.Controller
}

// indexページ
func (c App) Index() revel.Result {
  rows, _ := DbMap.Select(models.User{}, "select * from user")

  var users []*models.User

  // rowにrowsの中身が入ってくる
  // rangeを使用すると、for文でindexを取って、rows[index]などやらなくて良いので楽
  for _, row := range rows {
    user := row.(*models.User)
    users = append(users, user)
  }
    return c.Render(users)
}


// 登録
// 登録後はindexページにリダイレクトする
func (c App) UserPost(name string) revel.Result {

  // 第一引数がIDでkeyに指定しているので、0を指定すると、自動でインクリメントして登録される
  DbMap.Insert(&models.User{0, name})
  return c.Redirect(routes.App.Index())
}

// 削除
// 登録後はindexページにリダイレクトする
func (c App) Delete(id int) revel.Result {

  // id値を受け取って削除する
  DbMap.Exec("delete from user where Id=?", id)
  return c.Redirect(routes.App.Index())
}
