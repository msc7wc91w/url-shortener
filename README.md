# URL shortener

## Consume API

```bash
curl --request POST \
--data '{
    "long_url": "https://www.google.com/",
    "user_id" : "123456789"
}' \
  http://localhost:9808/create-short-url
```

## Install Redis

### Docker

```bash
# Start
sudo docker run -d --name redis -p 6379:6379 redis:7.2-alpine

# Stop
sudo docker stop redis
```

### [macOS](https://redis.io/docs/install/install-redis/install-redis-on-mac-os/)

```bash
# Install
brew install redis

# Start
brew services info redis

# Stop
brew services stop redis
```

### [Windows WSL 2](https://redis.io/docs/install/install-redis/install-redis-on-windows/)

```bash
# Install
curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg

echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/redis.list

sudo apt-get update
sudo apt-get install redis

# Start
sudo service redis-server start

# Stop
sudo service redis-server stop
```
