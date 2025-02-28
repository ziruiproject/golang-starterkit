### Authentication Routes

---

### **Base URL**

`http://localhost:8080/api/auth`

---

### **Endpoints**

---

#### 1. **Register a New User**

**POST** `/register`

This endpoint allows users to register an account.

**Request Body**:

```json
{
  "name": "John",
  "email": "asa@hotmail.com",
  "password": "rahasia123"
}
```

**Response Example (Success)**:

```json
{
  "code": 201,
  "message": "User registered successfully",
  "data": {
    "id": 7,
    "name": "John",
    "email": "asa@hotmail.com"
  }
}
```

---

#### 2. **Login**

**POST** `/login`

This endpoint allows users to log in and receive a token for authentication.

**Request Body**:

```json
{
  "email": "Annamae_Steuber@hotmail.com",
  "password": "rahasia123"
}
```

**Response Example (Success)**:

```json
{
  "code": 200,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...."
  }
}
```

### Protected User Routes

---

### **Base URL**

`http://localhost:8080/api/users`

---

### **Endpoints**

---

#### 1. **Get All Users**

**GET** `/`

This endpoint retrieves a list of all users.

**Response Example**:

```json
{
  "code": 200,
  "message": "Users fetched successfully",
  "data": [
    {
      "id": 4,
      "name": "Gerard Considine",
      "email": "Annamae_Steuber@hotmail.com"
    },
    {
      "id": 6,
      "name": "MD",
      "email": "Lauriane_Robel17@gmail.com"
    },
    {
      "id": 3,
      "name": "Eva Doyle",
      "email": "jwt@gmail.com"
    },
    {
      "id": 7,
      "name": "John",
      "email": "asa@hotmail.com"
    }
  ]
}
```

---

#### 2. **Get User by Identifier**

**GET** `/{identifier}`

This endpoint allows you to fetch a user by their identifier, which could be either their `id` or `email`.

**Path Parameter**:

- `identifier`: The `id` or `email` of the user.

**Response Example**:

```json
{
  "code": 200,
  "message": "User fetched successfully",
  "data": {
    "id": 3,
    "name": "Eva Doyle",
    "email": "jwt@gmail.com"
  }
}

```

---

#### 3. **Delete User by ID**

**DELETE** `/{userId}`

This endpoint allows for deleting a user by their `userId`.

**Path Parameter**:

- `userId`: The ID of the user to be deleted.

**Response Example (Success)**:

```json
{
  "code": 200,
  "message": "User deleted successfully",
  "data": null
}
```

---

#### 4. **Update User by ID**

**PUT** `/{userId}`

This endpoint allows updating a user by their `userId`.

**Path Parameter**:

- `userId`: The ID of the user to be updated.

**Request Body**:

```json
{
  "name": "John",
  "email": "jhon@hotmail.com"
}
```

**Response Example (Success)**:

```json
{
  "code": 200,
  "message": "User updated successfully",
  "data": {
    "id": 4,
    "name": "John",
    "email": "jhon@hotmail.com"
  }
}
```

---

### **Error Responses**

1. **Bad Request**  
   **Status Code**: `400 Bad Request`  
   **Response Example**:
   ```json
   {
       "code": 400,
       "message": "Invalid input data",
       "data": null
   }
   ```

2. **Unauthorized**  
   **Status Code**: `401 Unauthorized`  
   **Response Example**:
   ```json
   {
       "code": 401,
       "message": "Missing Authorization header",
       "data": null
   }
   ```
3. **Bad Request**  
   **Status Code**: `404 Not Found`  
   **Response Example**:
   ```json
   {
       "code": 404,
       "message": "resource not found",
       "data": []
   }
   ```
4. **Internal Server Error**  
   **Status Code**: `500 Internal Server Error`  
   **Response Example**:
   ```json
   {
       "code": 500,
       "message": "Internal server error",
       "data": null
   }
   ```

---