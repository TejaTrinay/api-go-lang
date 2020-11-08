# Development of basic version API using Go language 
Development of a HTTP JSON API using Go language by using MongoDB for storage

### Tech Stack used
- Go language to develop the API
- MongoDB for storage

### Endpoints and explanation
1. Creating an article

```
/api/articles
```
**Testing in the postman**
- Creating article
![Creating article](https://github.com/TejaTrinay/api-go-lang/blob/main/images/create_article.JPG)
- Getting the response ID
![Getting the response ID](https://github.com/TejaTrinay/api-go-lang/blob/main/images/create_article_response.JPG)

2. Getting an article using ID

```
/api/articles/{id}
```
**Testing in the postman**
- API request from getting article using ID
![Getting article](https://github.com/TejaTrinay/api-go-lang/blob/main/images/get_article_using_id.JPG)
- Getting the data
![Getting the data](https://github.com/TejaTrinay/api-go-lang/blob/main/images/get_article_using_id_response.JPG)

3. Listing all articles

```
/api/articles
```
**Testing in the postman**
- API request from getting all articles
![Getting article](https://github.com/TejaTrinay/api-go-lang/blob/main/images/get_articles.JPG)
- Getting the data
![Getting the data](https://github.com/TejaTrinay/api-go-lang/blob/main/images/get_articles_result.JPG)


 
