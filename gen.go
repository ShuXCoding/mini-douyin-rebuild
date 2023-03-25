package main

import (
	"gorm.io/gen"
	"mini-douyin-rebuild/db"
	"mini-douyin-rebuild/model"
)

type Queries interface {

	//  SELECT * FROM user WhERE id=@id
	GetUserInfo(id int) (gen.T, error)
	// SELECT password FROM user WHERE username=@username
	GetUserPassword(username string) (password string, err error)
	// SELECT id FROM user WHERE username=@username
	FindUserByUsername(username string) (id int, err error)
}

func GenInit() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dal",
		Mode:    gen.WithDefaultQuery | gen.WithoutContext,
	})
	g.UseDB(db.Conn)
	//g.ApplyBasic(
	//	g.GenerateModelAs("user", "User", gen.FieldIgnore("create_time")),
	//	g.GenerateModelAs("comment","Comment"),
	//	g.GenerateModelAs("favorite","Favorite",gen.FieldIgnore("create_time")),
	//	g.GenerateModelAs("message","Message"),
	//	g.GenerateModelAs("relation","Relation",gen.FieldIgnore("create_time")),
	//	g.GenerateModelAs("video","Video",gen.FieldIgnore("update_date")))
	g.ApplyInterface(func(queries Queries) {}, model.User{})
	g.Execute()
}
