## user
- /user
    - POST
    - Create User
- /user/login
    - GET
    - Logs user into the system
- /user/logout
    - GET
    - Logs out current logged in user session
- /user/{username}
    - GET
    - Get user by user name
- /user/{username}
    - PUT
    - Update user
- /user/{username}
    - DELETE
    - Delete user

## task
- /task
    - POST
    - Create task
- /task
    - GET
    - Get all tasks
- /task/{task_id}
    - GET
    - Get task description by task id
- /task/{task_id}
    - PUT
    - Update task
- /task/{task_id}
    - DELETE
    - Delete task

## memory
- /memory
    - GET 
    - Get all memories
- /memory/{memory_id}
    - GET
    - Get memory description
- /memory/{memory_id}
    - DELETE
    - Delete memory
- /memory/picture/{memory_id}
    - POST
    - Store picture
- /memory/picture/{memory_id}
    - DELETE
    - Delete picture

## tag
- /tag/{memory_id}
    - POST
    - Store tag
- /tag/{memory_id}
    - DELETE
    - Delete tag

## status
 - /status
    - PUT
    - Update task status
