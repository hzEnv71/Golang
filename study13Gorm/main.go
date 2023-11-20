package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id       int
	Username string `gorm:"default:lhz"`
	Address  string `gorm:"default:河南""`
	//gorm.Model
}

func (User) TableName() string {
	return "tb_user"
}

// 创建钩子
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	//rand.Seed(time.Now().UnixNano())
	//u.Id = rand.Intn(1000)
	var count int64
	tx.Table("tb_user").Count(&count)
	//fmt.Println(count)
	u.Id = int(count + 1)
	if u.Id == 0 {
		return errors.New("invalid id")
	}
	return
}

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	//if u.Id == 0 {
	//	return errors.New("admin user not allowed to delete")
	//}
	return
}
func main() {
	dsn := "root:Li20031202@tcp(127.0.0.1:3306)/data1?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//MySQL 驱动程序提供了 一些高级配置 可以在初始化过程中使用，例如：
	//db, err := gorm.Open(mysql.New(mysql.Config{
	//	DSN:                       "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
	//	DefaultStringSize:         256,                                                                        // string 类型字段的默认长度
	//	DisableDatetimePrecision:  true,                                                                       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	//	DontSupportRenameIndex:    true,                                                                       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	//	DontSupportRenameColumn:   true,                                                                       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	//	SkipInitializeWithVersion: false,                                                                      // 根据当前 MySQL 版本自动配置
	//}), &gorm.Config{})
	//现有的数据库连接
	//sqlDB, err := sql.Open("mysql", dsn)
	//if err != nil {
	//	fmt.Println("err:", err)
	//}
	//gormDB, err := gorm.Open(mysql.New(mysql.Config{
	//	Conn: sqlDB,
	//}), &gorm.Config{})
	//fmt.Println(gormDB)
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("err:", err)
	}
	//fmt.Println(gormDB)
	//Create(gormDB)
	Select(gormDB)
	//Update(gormDB)
	//Delete(gormDB)
}
func Create(gormDB *gorm.DB) {
	//user := User{Username: "蒙田", Address: "广州"}
	//fmt.Println(user.username, user.address)

	//result := gormDB.Create(&user) // 通过数据的指针来创建
	//创建记录并更新给出的字段。
	//result := gormDB.Select("Username", "Address").Create(&user)
	//创建一个记录且一同忽略传递给略去的字段值。
	//result := gormDB.Omit("Username").Create(&user)

	//批量插入
	users := []User{{Username: "jinzhu1"}, {Username: "jinzhu2"}, {Username: "jinzhu3"}}
	//gormDB.Create(&users)
	//使用 CreateInBatches 分批创建时，你可以指定每批的数量
	gormDB.CreateInBatches(users, 10)
	//fmt.Println(result.RowsAffected)
	//更据map创建
	gormDB.Model(&User{}).Create([]map[string]interface{}{
		{"Username": "lhz111", "Address": "深圳"},
		{"Username": "lhz222", "Address": "广州"},
	})

}
func Select(gormDB *gorm.DB) {
	var user1 = User{}
	var user2 = User{}
	//var user3 = User{}
	//result1 := gormDB.First(&user1)
	result2 := gormDB.Take(&user2)
	//result3 := gormDB.Last(&user3)
	//fmt.Println(user1.Id, user1.Username, user1.Address, result1.RowsAffected)
	fmt.Println(user2.Id, user2.Username, user2.Address, result2.RowsAffected)
	//fmt.Println(user3.Id, user3.Username, user3.Address, result3.RowsAffected)

	result3 := map[string]interface{}{}
	//gormDB.Model(&User{}).First(&result3)
	gormDB.Model(&User{}).Last(&result3)
	fmt.Println(result3)
	//根据主键检索
	users := []User{}
	gormDB.Find(&users, []int{1, 2, 3, 4, 5})

	fmt.Println(users)
	//检索全部对象
	allUsers := []User{}
	gormDB.Find(&allUsers)
	fmt.Println(allUsers)
	//where
	gormDB.Where("username = ?", "吕布").Find(&user1)
	fmt.Println(user1.Username, user1.Address)
	users1 := []User{}
	gormDB.Where("username LIKE ?", "%jin%").Find(&users1)
	fmt.Println(users1)
	users2 := []User{}
	gormDB.Where("username LIKE ? AND id >= ?", "%jin%", "250").Find(&users2)
	fmt.Println(users2)
	users3 := []User{}
	gormDB.Where([]int64{1, 3, 5}).Find(&users3)
	fmt.Println(users3)
	//当使用struct查询时，GORM只对非零字段进行查询，也就是说如果你的字段的值是0，''，false或其他零值，它将不会被用来建立查询条件
	//not
	users4 := []User{}
	gormDB.Not("username", []string{"jinzhu2", "jinzhu1"}).Find(&users4)
	fmt.Println("users4:", users4)
	//不在主键 slice 中
	user4 := User{}
	gormDB.Not([]int64{1, 2, 3}).First(&user4)
	fmt.Println("user4", user4)
	// 原生 SQL
	user5 := User{}
	gormDB.Not("username = ?", "柳岩").First(&user5)
	fmt.Println(user5)
	// Struct
	user6 := User{}
	gormDB.Not(User{Username: "柳岩"}).First(&user6)
	fmt.Println(user6)
	//or
	users7 := User{}
	gormDB.Where("username = ?", "zhi").Or("username = ?", "吕布").Find(&users7)
	fmt.Println(users7)
	//指定列
	userss := []User{}
	gormDB.Select("username, address").Find(&userss)
	fmt.Println("userss:", userss)
	userss1 := []User{}
	gormDB.Select([]string{"username", "address"}).Find(&userss1)
	fmt.Println("userss1:", userss1)
	//// SELECT name, age FROM users;
	var u = User{}
	//可以返回参数中的第一个非空表达式
	rows, err := gormDB.Table("tb_user").Select("COALESCE(address,?,?)", "河南", "河北").Rows()
	if err != nil {
		fmt.Println(err)
	}
	i := 0
	for rows.Next() {
		i++
		rows.Scan(u)
		fmt.Println("u:", i, u)
	}
	//order
	//使用 Order 从数据库查询记录时，当第二个参数设置为 true 时，将会覆盖之前的定义条件
	us := []User{}
	gormDB.Order("id asc, username").Find(&us)  //升序
	gormDB.Order("id desc, username").Find(&us) //降序
	fmt.Println(us)
	u1, u2 := []User{}, []User{}
	gormDB.Order("id desc").Find(&u1).Order("username desc").Find(&u2)
	fmt.Println(u1, "===", u2)
	//Limit指定要检索的最大记录数。 Offset指定在开始返回记录前要跳过的记录数。
	us1 := []User{}
	gormDB.Limit(5).Find(&us1)
	us2 := []User{}
	gormDB.Limit(10).Offset(5).Find(&us2)
	fmt.Println(us1, "==========", us2)
	//count 获取模型记录数
	var count int64
	gormDB.Where("username = ?", "jinzhu").Or("username = ?", "jinzhu2").Find(&users).Count(&count)
	fmt.Println(count)
	gormDB.Model(&User{}).Where("address = ?", "河南").Count(&count)
	fmt.Println(count)

	gormDB.Table("tb_user").Count(&count)
	fmt.Println(count)
	//. Group 和 Having
	//Joins
	//Pluck
}
func Update(gormDB *gorm.DB) {
	user := User{}
	//gormDB.First(&user)

	//user.Username = "jinzhu9"
	//user.Address = "湖北"
	//gormDB.Save(&user)
	//gormDB.Model(&user).Update("username", "hello")
	//gormDB.Model(&user).Where("username = ?", "蒙田").Update("address", "上海")
	//map
	//gormDB.Model(&user).Where("id=?", 20).Updates(map[string]interface{}{"username": "hello11", "address": "天津"})
	//struct
	//gormDB.Model(&user).Where("id=?", 2).Updates(User{Username: "hello0"})
	//更改选中的字段
	//gormDB.Model(&user).Where("username = ?", "蒙田").Select("address").Update("address", "南京")
	//更改非选中的字段
	//gormDB.Model(&user).Where("username = ?", "蒙田").Omit("address").Update("username", "蒙田")
	//批量更新
	gormDB.Table("tb_user").Where("id IN (?)", []int{10, 11}).Updates(map[string]interface{}{"address": "北京"})
	//带有表达式的 SQL 更新
	gormDB.Model(&user).Where("username=?", "吕布").Update("username", gorm.Expr("?+?", "率", "fen"))
}
func Delete(gormDB *gorm.DB) {
	var user User
	user.Id = 10
	gormDB.Delete(&user)
	//var u User
	//gormDB.Where("username = ?", "jinzhu8").Delete(&u)
	////根据主键删除
	//gormDB.Delete(&User{}, 10)
	//gormDB.Delete(&User{}, "10")
	//gormDB.Delete(&User{}, []int{1, 2, 3})
	////如果指定的值不包括主属性，那么 GORM 会执行批量删除，它将删除所有匹配的记录
	//gormDB.Where("username LIKE ?", "%6%").Delete(&User{})
	//gormDB.Delete(&User{}, "username LIKE ?", "%jinzhu%")
	//返回被删除的数据，仅适用于支持 Returning 的数据库
	//var users []User
	//gormDB.Clauses(clause.Returning{}).Where("username = ?", "马超").Delete(&users)
	//fmt.Println(users)
	//软删除
	//gormDB.Unscoped().Find(&user)

}
