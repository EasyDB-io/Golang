# Golang-Client
EasyDB Golang Client

This gem has not yet been implemented. If you'd like to write it, I'll pay you!

Just implement the following 4 functions with their respective cURLs:

Get: 
```
curl --request GET \ 
--url https://app.easydb.io/database/{database_uuid}/{key} \
-H "token:{database_token}"
```

Put: 
```
curl --request POST \  
  --url https://app.easydb.io/database/{database_uuid} \
  -H 'content-type: application/json' \
  -H 'token: {database_token}' \
  -d '{ 
      "key": "somekey",
      "value": "{asdf}"
  }'
```

Delete:
```
curl --request DELETE \
   --url https://app.easydb.io/database/{database_uuid} \
   -H 'content-type: application/json' \
   -H 'token: {database_token}' \
   -d '{
       "key": "somekey"
   }
```

List: 
```
curl --request GET \ 
--url https://app.easydb.io/database/{database_uuid} \
-H "token:{database_token}"
```

If you create a PR implementing this language, I'm happy to send you $5 via Stellar.
