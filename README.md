
# CRUD written in ð±GO with docker and ngnix 

---

## How to run ð¤:

    docker-compose up -d

---

## How to use ð¿: 

    # Get users list ð
    curl -X GET localhost/api/users

    # Get user ð
    curl -X GET localhost/api/users/{id}
    
    # Create ð
    curl -X POST localhost/api/users -d "{\"nickname\": \"nickname\", \"email\": \"example@gmail.com\", \"password\": \"password\"}"

    # Update ð£
    curl -X PUT localhost/api/users/{id} -d "{\"nickname\": \"meh\", \"email\": \"niceass@gmail.com\", \"password\": \"password\"}"
    
    # Delete user â
    curl -X DELETE localhost/api/users/{id}
