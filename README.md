# API REST - SERVER FILES

- deploy heroku ( https://api-url-images.herokuapp.com/ )

- EndPoints
    - / [GET]
    - api/v1/images [POST]
    - api/v1/images/:IDImage [GET]
    
- Request [form-data]
    - content-type [multipart/form-data]
    - Key ("file") [file]
    - value ("name-file.png")
    
- Response
```json
{
    "message_type": "MESSAGE",
    "message": "the file was saved successfully",
    "error": false,
    "data": "https://api-url-images.herokuapp.com/api/v1/images/125ce0baff3271761ca61843eccf7985.jpg"
}
```