Request:
```
curl --location 'localhost:8080/calculate' \
--header 'User-Access: superuser' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "2+2"
}'
```

Result:
```
{
    "result": "4"
}
```

