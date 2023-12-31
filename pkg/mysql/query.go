package mysql

// authパッケージ用
var LoginMail = "SELECT id, password FROM users WHERE email = ?"
var RegisterUser = "INSERT INTO users (email, password, name) VALUES (?, ?, ?)"
