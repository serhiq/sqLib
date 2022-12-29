# Home screenshot gallery

- api -  golang + iris
- frontend - Vue.js  + Quasar
- database - MySQL
- telegram bot - golang



### Featured app

* loading/deleting images
* tag assignment
* image filtering



### Start the app 

```bash
docker-compose up --build
```


**phpMyAdmin**

[http://localhost:8887/](http://localhost:8887/)

**Fronted**

[http://localhost:8080/](http://localhost:8080/)

**Backend**

[http://localhost:5555/](http://localhost:7777/)


**An Bot service**

telegram bot written in Golang to check the health of the service.

**Environment Variables**

| Name              | Description                           | Default     |
|-------------------|---------------------------------------|-------------|
| TELEGRAM_TOKEN    | Token create for Bot                  |   -         |
| CHAT_ID           | Chat id as receiver for our messages  |   -         |
