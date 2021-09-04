# hackathon090405 チーム不退転 サーバーサイド

### 使用技術
- golang
  - echo
  - gorm
- heroku 
### run

`
  DATABASE_URL=postgres://[username]:[password]@127.0.0.1:5432/[dbname]?sslmode=disable make run
`

### heroku endpoint
https://whispering-bayou-86182.herokuapp.com/

GET /todos/:userId ユーザーのTODOを所得

POST /todos/edit/:todoId TODOの編集

POST /todos/add TODOの追加

POST /todos/toggle_todo/:todoId TODOが完了したかしないか

DELETE /todos/delete/:todoId TODOの削除
### curl でリクエストを送る
GET: curl http://localhost:8000/todos/1

POST: curl -F "title=test" -F "description=test-description" -F "todoId=1" http://localhost:8000/todos/add

### やること
- []
