# NewsManagementSystem-Go
Written by Jonathan Richard for home takeaway project in recruitment process
Packages:
### 1. github.com/go-sql-driver/mysql 
### 2. github.com/gorilla/mux 

# APIs
## /news
#### GET : Retrieve all news 
#### POST : Add a news (Will reject duplicates)
Example POST /news (Content-type: JSON)
```
{
    "title": "Hello Bareksa",
    "topic": "Investation",
    "tags" : ["Economy", "Business"],
    "status" : "Draft"
}
```
## /news/{title}
#### GET : Retrieve news with the corresponding title 
Example: GET /news/Hello%20Bareksa
#### DELETE : Remove the news with the corresponding title(will be added into the deleted table)
Example: DELETE /news/Hello%20Bareksa
#### PATCH : Modify Title, topic, tags (old tags would be removed), or status, unmodified field shall be empty
Example  PATCH /news/Hello%20Bareksa (Content-type: JSON)
Payload:
```
{
    "title": "Healthy Investation for Everyone",
    "tags" : ["Investation", "Money"],
    "topic": "Money",
    "status" : "Published"
}
``` 
or if you only want to modify , lets say, the status field, you can leave everything else empty
```
{
    "status" : "Published"
}
```

## /news/{title}/tags
#### POST: Add new tags into the existing news, will ignore duplicates
Example:  POST /news/Hello%20Bareksa/tags (Content-type: JSON)
Payload:
```
{
    "tags" : ["Bank", "Further Use"],
}
```
## /news/{title}/tags/{tags}
#### Delete: Remove tags of news if exist
Example: DELETE /news/Hello%20Bareksa/tags/money

## /news/topic/ {topic}
#### GET : Retrieve news with the corresponding topic
Example: GET/news/topic/investation

## /news/status/{status}
#### GET : Retrieve news with the corresponding status
Example: GET /news/status/deleted 
or
GET /news/status/draft


# Database
I provided databaseinit.sql that contains queries that creates the required table 
## News
Contains id,title,topic,and status of all news. 
## Tags
Contains tag name, and id of the corresponding news that references News.id 
## Deleted
Contains id and title of a deleted news

