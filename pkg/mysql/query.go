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

// 認証コードと更新時刻を取得
var GetAuthCode = "SELECT auth_code, updated_at FROM users WHERE email = ?"

// ######################
// ### memoパッケージ用 ###
// ######################
// メモ一覧を取得
var GetMemoAll = "SELECT id, amount, partner, memo, date, period, type FROM memos WHERE uid = ?"

// メモを1件取得
var GetMemoOne = "SELECT id, amount, partner, memo, date, period, type FROM memos WHERE id = ? AND uid = ?"

// メモを登録
var CreateMemo = "INSERT INTO memos (uid, amount, partner, memo, date, period, type, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"

// メモを更新
var UpdateMemo = "UPDATE memos SET amount = ?, partner = ?, memo = ?, date = ?, period = ?, type = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ? AND uid = ?"

// メモを削除
var DeleteMemo = "DELETE FROM memos WHERE id = ? AND uid = ?"
