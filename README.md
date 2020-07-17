# News_API
## This project is a part of xsolla summer school entry test.

In this vary basic API you can work with:
- Article - simple model with title, tag and array of users, who liked it.
- Tag - just made for filtering purposes.
- User - just here to add like and remove it.
- Like - simple article and user pair.

What can you do with it?
1. Article
  - Get all articles in storage with filters and sort options
  - Get article by id
  - Post article
2. Tag
  - Get all tags in storage
  - Get tag by id
  - Post tag
3. User
  - Get all users
4. Like
  - Get all likes in storage
  - Get like by id
  - Post like
  - Delete like by id
5. View swagger docs at /docs

Things to know about:
  - There are *some* tests in main_test.go that by all means don't cover everything
  - Open API is generated with swagger package https://github.com/go-swagger/go-swagger using:
  ```
    swagger generate spec -o ./swagger.yaml --scan-models
  ```
  - You can find published Open API specification on swaggerhub at https://app.swaggerhub.com/apis/Wardenclock1759/of-news_api/1.0.0
