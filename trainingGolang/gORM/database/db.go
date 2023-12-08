package database
import (
"fmt"
"learning=gorm/models"
"Log"
"gorm.io/driver/postgres"
"gorm.io/gorm"
var (
host
= "localhost"
user
= "sofyan"
password = "postgres"
dbPort
= "5432"
dbname
= "learning-gorm"
db
*gorm. DB
err
error
func StartDB() {
config := ft. Sprintf (host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user,
password, dbname,
dbPort)
db, err = gorm. Open (postgres. Open (config), &gorm.Config{})
if err I= nil {
log. Fatal"error connecting to database :", err)]
db. Debug() . AutoMigrate (models.User {}, models. Product{})