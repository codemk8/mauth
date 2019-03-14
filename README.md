## mauth

Authentication microservice written in go

## API Endpoints

* /register
* /login
* /auth
* /logout

## To run a micro service

```bash
export MAUTH_API_KEY="YOUR_API_KEY_HERE" 
./bin/mauth --alsologtostderr 
```


```bash
# test by curl
curl -H "Content-Type: application/json" -H "Authorization: token YOUR_API_KEY_HERE" http://localhost:55555/login
```