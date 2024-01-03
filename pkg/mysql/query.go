package mysql

// authパッケージ用
var LoginMail = "SELECT id, password FROM users WHERE email = ? AND register_type = 1 "
var LoginGoogle = "SELECT id FROM users WHERE email = ? AND register_type = 2"
var CreateUser = "INSERT INTO users (email, password, name, register_type) VALUES (?, ?, ?, ?)"
var CreateGroup = "INSERT INTO `groups` (manage_user) VALUES (?)"
var UpdateGroup = "UPDATE users SET group_id = ? WHERE id = ?"
var LoginCheck = "SELECT COUNT(id) FROM groups WHERE manage_user = ?"
