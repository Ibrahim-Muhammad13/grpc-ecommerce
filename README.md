# small e-commerce app using golang
this an implementation of grpc server, using evans as client

## what i used
- grpc
- gorm
- mysql
- jwt

## example of requests


_Create user:_

![alt text](https://github.com/ibrahim-muhammad13/grpc-ecommerce/blob/main/img/createuser.PNG?raw=true)

_Login:_

![alt text](https://github.com/ibrahim-muhammad13/grpc-ecommerce/blob/main/img/login.PNG?raw=true)


_Create product:_
**_Notice:_** only admin can create product, after providing token and checking role you can create product

![alt text](https://github.com/ibrahim-muhammad13/grpc-ecommerce/blob/main/img/unauth.PNG?raw=true)

![alt text](https://github.com/ibrahim-muhammad13/grpc-ecommerce/blob/main/img/createproduct.PNG?raw=true)

_Get all products:_

![alt text](https://github.com/ibrahim-muhammad13/grpc-ecommerce/blob/main/img/allproducts.PNG?raw=true)


_Search product by name:_

![alt text](https://github.com/ibrahim-muhammad13/grpc-ecommerce/blob/main/img/search.PNG?raw=true)

_Update product:_

![alt text](https://github.com/ibrahim-muhammad13/grpc-ecommerce/blob/main/img/updated.PNG?raw=true)

_Get Product by id :_

![alt text](https://github.com/ibrahim-muhammad13/grpc-ecommerce/blob/main/img/byid.PNG?raw=true)



_Delete:_

![alt text](https://github.com/ibrahim-muhammad13/grpc-ecommerce/blob/main/img/delete.PNG?raw=true)


