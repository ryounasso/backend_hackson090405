# hackathon090405 チーム不退転 サーバーサイド

### 開発環境
1. `docker-compose build`
2. `docker-compose up`

### dbに接続
1. docker exec -it postgres /bin/bash 
2. su - postgres
3. psql --username=hackson --dbname=mydb
### 使用技術
- golang
  - echo
  - gorm
- heroku 



### heroku endpoint
https://whispering-bayou-86182.herokuapp.com/

GET /todos/:userId ユーザーのTODOを所得 
例 `curl http://localhost:8000/todos/user`

POST /todos/edit/:todoId TODOの編集 例 `curl -F "title=titledited" -F "description=descriptionedited" http://localhost:8000/todos/edit/1`

POST /todos/add TODOの追加 例 `curl -F "title=test" -F "description=test-description" -F "userId=user" http://localhost:8000/todos/add`

POST /todos/toggle_todo/:todoId TODOが完了したかしないか 例 `curl http://localhost:8000/todos/toggle_todo/1`

DELETE /todos/delete/:todoId TODOの削除 例 `curl http://localhost:8000/todos/delete/1`

### heroku がアプリケーションエラーとなったら
`heroku restart web.1 --app whispering-bayou-86182`

### やること
- []
