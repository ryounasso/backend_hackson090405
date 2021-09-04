echo
gorm


### run

`
  DATABASE_URL=postgres://[username]:[password]@127.0.0.1:5432/[dbname]?sslmode=disable make run
`

### heroku endpoint
https://whispering-bayou-86182.herokuapp.com/

### curl でリクエストを送る
GET: curl http://localhost:8000/todos/1

POST: curl -F "title=ebata" -F "description=mailmail" -F "todoId=1" http://localhost:8000/todos/add

### やること
editTodoやaddTodoでリクエストを受け付けることができたが、gormがよくわかっていないので、正しいところが編集できずにid=0が編集されているっぽい?