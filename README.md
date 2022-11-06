# shorturl-GO
POST
URL => http://localhost:8080/v1/short-url/
Request => {
    "long_url": "https://github.com/amdtaufiq/TodoList-GO"
}
Response => {
    "message": "Short URL createad",
    "sort_url": "http://localhost:8080/v1/short-url/KuETbg"
}

GET 
URL => http://localhost:8080/v1/short-url/:url
