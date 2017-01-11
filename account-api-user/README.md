# account-api-user
micro service for user domain 

# API route 
 * GET /users/{userId} 
    * Input contract userId int
    * Output contract 
        {
        "id": int,
        "first_name": string,
        "last_name": string,
        "gender": string,
        "nick_name": string,
        "email": string
        }

* PUT /users/{userId} 
    * Input contract userId int +
        {
        "email": string
        }
    * Output contract 
        {
        "id": int,
        "first_name": string,
        "last_name": string,
        "gender": string,
        "nick_name": string,
        "email": string
        }
