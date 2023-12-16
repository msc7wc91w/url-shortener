curl --request POST \
--data '{
    "long_url": "https://www.google.com/",
    "user_id" : "123456789"
}' \
  http://localhost:9808/create-short-url
