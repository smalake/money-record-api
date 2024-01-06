package mysql

// authパッケージ用
var LoginMail = "SELECT id, password FROM users WHERE email = ? AND register_type = 1 "
var LoginGoogle = "SELECT id FROM users WHERE email = ? AND register_type = 2"
var CheckEmail = "SELECT auth_code FROM users WHERE email = ?"
var ResendAuthCode = "UPDATE users SET auth_code = ?, password = ?, name = ?, updated_at = CURRENT_TIMESTAMP WHERE email = ?"
var UpdateAuthCode = "UPDATE users SET auth_code = ?, updated_at = CURRENT_TIMESTAMP WHERE email = ?"
var CreateUser = "INSERT INTO users (email, password, name, register_type, auth_code) VALUES (?, ?, ?, ?, ?)"
var CreateGroup = "INSERT INTO `groups` (manage_user) VALUES (?)"
var UpdateGroup = "UPDATE users SET group_id = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?"
var LoginCheck = "SELECT COUNT(id) FROM groups WHERE manage_user = ?"
