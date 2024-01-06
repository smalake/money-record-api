package mysql

// ######################
// ### authパッケージ用 ###
// ######################
// メールアドレスによるログイン
var LoginMail = "SELECT id, password FROM users WHERE email = ? AND register_type = 1 "

// Googleアカウントによるログイン
var LoginGoogle = "SELECT id FROM users WHERE email = ? AND register_type = 2"

// メールアドレスに紐づいている認証コードを取得
var CheckEmail = "SELECT auth_code FROM users WHERE email = ?"

// ユーザ情報を再登録
var ResendAuthCode = "UPDATE users SET auth_code = ?, password = ?, name = ?, updated_at = CURRENT_TIMESTAMP WHERE email = ?"

// 認証コードを更新
var UpdateAuthCode = "UPDATE users SET auth_code = ?, updated_at = CURRENT_TIMESTAMP WHERE email = ?"

// ユーザ情報を登録
var CreateUser = "INSERT INTO users (email, password, name, register_type, auth_code) VALUES (?, ?, ?, ?, ?)"

// グループを作成
var CreateGroup = "INSERT INTO `groups` (manage_user) VALUES (?)"

// グループIDをユーザに紐付ける
var UpdateGroup = "UPDATE users SET group_id = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?"

// ユーザが親か子かを判定
var LoginCheck = "SELECT COUNT(id) FROM groups WHERE manage_user = ?"
