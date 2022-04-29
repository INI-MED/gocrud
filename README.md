
# CRUD written in 😱GO with docker and ngnix 

---

## How to run 🤔:

    docker-compose up -d

---

## How to use 🗿: 

    # Get users list 📝
    curl -X GET localhost/api/users

    # Get user 🙉
    curl -X GET localhost/api/users/{id}
    
    # Create 🎉
    curl -X POST localhost/api/users -d "{\"nickname\": \"nickname\", \"email\": \"example@gmail.com\", \"password\": \"password\"}"

    # Update 📣
    curl -X PUT localhost/api/users/{id} -d "{\"nickname\": \"meh\", \"email\": \"niceass@gmail.com\", \"password\": \"password\"}"
    
    # Delete user ❌
    curl -X DELETE localhost/api/users/{id}
