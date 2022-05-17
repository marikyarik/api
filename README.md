# api

go run main.go 

nginx proxy config:
```
server {
        listen 80;

        location / {
                proxy_pass http://127.0.0.1:8000/;
                proxy_set_header User "Test User";
                proxy_set_header Host $host;
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_pass_request_headers      on;
        }
}
```