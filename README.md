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

### Cart Routes

---

### **Base URL**

`http://localhost:8080/api/carts`

---

**Header Example**:  
`Authorization: Bearer <your_token>`

---

### **Endpoints**

#### 1. **Get All Cart Items**

**GET** `/`  
Fetch all cart items.

**Request Example**:  
Headers:

```json
{
  "Authorization": "Bearer <your_token>"
}
```

**Response Example**:

```json
{
  "code": 200,
  "message": "Cart items fetched successfully",
  "data": [
    {
      "user_data": {
        "id": 3,
        "name": "Eva Doyle",
        "email": "jwt@gmail.com"
      },
      "cart_items": [
        {
          "id": 6,
          "name": "Oriental Fresh Sausages",
          "description": "Autem stabilis vinum ambulo ipsa. Causa molestiae surgo reprehenderit. Corroboro pax demitto.",
          "price": 1000000,
          "user": {
            "id": 3,
            "name": "Eva Doyle",
            "email": "eva@gmail.com"
          },
          "created_at": "2024-12-29T13:30:32.608467Z",
          "updated_at": "2024-12-29T13:30:32.608467Z",
          "quantity": 2
        }
      ]
    }
  ]
}
```

---

#### 2. **Add Item to Cart**

**POST** `/`  
Add an item to the cart.

**Request Body**:

```json
{
  "user_id": 3,
  "product_id": 6,
  "quantity": 2
}
```

**Response Example**:

```json
{
  "code": 200,
  "message": "Cart item created successfully",
  "data": {
    "user_data": {
      "id": 3,
      "name": "Eva Doyle",
      "email": "jwt@gmail.com"
    },
    "cart_items": [
      {
        "id": 6,
        "name": "Oriental Fresh Sausages",
        "description": "Autem stabilis vinum ambulo ipsa. Causa molestiae surgo reprehenderit. Corroboro pax demitto.",
        "price": 1000000,
        "user": {
          "id": 3,
          "name": "Eva Doyle",
          "email": "jwt@gmail.com"
        },
        "created_at": "2024-12-29T13:30:32.608467Z",
        "updated_at": "2024-12-29T13:30:32.608467Z",
        "quantity": 2
      }
    ]
  }
}
```

---

#### 3. **Update Cart Item**

**PUT** `/{cartId}`  
Update a specific cart item by its `cartId`.

**Request Parameters**:

- `cartId` (Path Parameter): The ID of the cart item to update.

**Request Body**:

```json
{
  "quantity": 5
}
```

**Response Example**:

```json
{
  "code": 200,
  "message": "Cart item updated successfully",
  "data": {
    "user_data": {
      "id": 3,
      "name": "Eva Doyle",
      "email": "jwt@gmail.com"
    },
    "cart_items": [
      {
        "id": 6,
        "name": "Oriental Fresh Sausages",
        "description": "Autem stabilis vinum ambulo ipsa. Causa molestiae surgo reprehenderit. Corroboro pax demitto.",
        "price": 1000000,
        "user": {
          "id": 3,
          "name": "Eva Doyle",
          "email": "jwt@gmail.com"
        },
        "created_at": "2024-12-29T13:30:32.608467Z",
        "updated_at": "2024-12-29T13:30:32.608467Z",
        "quantity": 5
      }
    ]
  }
}
```

---

#### 4. **Delete Cart Item**

**DELETE** `/{cartId}`  
Delete a specific cart item by its `cartId`.

**Request Parameters**:

- `cartId` (Path Parameter): The ID of the cart item to delete.

**Response Example**:

```json
{
  "code": 200,
  "message": "Cart item deleted successfully",
  "data": null
}
```

---

#### 5. **Get Cart Items by User ID**

**GET** `/user/{userId}`  
Fetch all cart items for a specific user by their `userId`.

**Request Parameters**:

- `userId` (Path Parameter): The ID of the user.

**Response Example**:

```json
{
  "code": 200,
  "message": "Cart items for user fetched successfully",
  "data": {
    "user_data": {
      "id": 3,
      "name": "Eva Doyle",
      "email": "jwt@gmail.com"
    },
    "cart_items": [
      {
        "id": 6,
        "name": "Oriental Fresh Sausages",
        "description": "Autem stabilis vinum ambulo ipsa. Causa molestiae surgo reprehenderit. Corroboro pax demitto.",
        "price": 1000000,
        "user": {
          "id": 3,
          "name": "Eva Doyle",
          "email": "jwt@gmail.com"
        },
        "created_at": "2024-12-29T13:30:32.608467Z",
        "updated_at": "2024-12-29T13:30:32.608467Z",
        "quantity": 2
      }
    ]
  }
}
```

---

#### 6. **Checkout User's Cart**

**POST** `/user/{userId}/checkout`  
Checkout all cart items for a specific user.

**Request Parameters**:

- `userId` (Path Parameter): The ID of the user checking out their cart.

**Response Example**:

```json
{
  "code": 200,
  "message": "Checkout successful",
  "data": {
    "user_data": {
      "id": 4,
      "name": "Gerard Considine",
      "email": "Annamae_Steuber@hotmail.com"
    },
    "cart_items": []
  }
}
```

---

#### 7. **Get Cart Item by ID**

**GET** `/{cartId}`  
Fetch a specific cart item by its `cartId`.

**Request Parameters**:

- `cartId` (Path Parameter): The ID of the cart item to retrieve.

**Response Example**:

```json
{
  "code": 200,
  "message": "Cart item fetched successfully",
  "data": {
    "user_data": {
      "id": 3,
      "name": "Eva Doyle",
      "email": "jwt@gmail.com"
    },
    "cart_items": [
      {
        "id": 6,
        "name": "Oriental Fresh Sausages",
        "description": "Autem stabilis vinum ambulo ipsa. Causa molestiae surgo reprehenderit. Corroboro pax demitto.",
        "price": 1000000,
        "user": {
          "id": 3,
          "name": "Eva Doyle",
          "email": "jwt@gmail.com"
        },
        "created_at": "2024-12-29T13:30:32.608467Z",
        "updated_at": "2024-12-29T13:30:32.608467Z",
        "quantity": 2
      }
    ]
  }
}
```

---

### Bank Routes

---

### **Base URL**

`http://localhost:8080/api/banks`

---

### **Endpoints**

#### 1. **Create a New Bank Account**

**POST** `/`  
Create a new bank account.

**Request Body**:

```json
{
  "user_id": 4
}
```

**Response Example**:

```json
{
  "code": 201,
  "message": "Account created successfully",
  "data": {
    "id": 4,
    "user": {
      "id": 4,
      "name": "Gerard Considine",
      "email": "Annamae_Steuber@hotmail.com"
    },
    "balance": 0,
    "updated_at": "2025-01-01T11:04:13.94314095+07:00"
  }
}
```

---

#### 2. **Get All Bank Accounts**

**GET** `/`  
Retrieve all bank accounts.

**Request Example**:  
Headers:

```json
{
  "Authorization": "Bearer <your_token>"
}
```

**Response Example**:

```json
{
  "code": 200,
  "message": "Accounts fetched successfully",
  "data": [
    {
      "id": 4,
      "user": {
        "id": 4,
        "name": "Gerard Considine",
        "email": "Annamae_Steuber@hotmail.com"
      },
      "balance": 0,
      "updated_at": "2025-01-01T08:36:17.634273Z"
    },
    {
      "id": 1,
      "user": {
        "id": 3,
        "name": "Eva Doyle",
        "email": "jwt@gmail.com"
      },
      "balance": 1500000,
      "updated_at": "2025-01-01T06:01:26.91105Z"
    }
  ]
}
```

---

#### 3. **Get Bank Account by ID**

**GET** `/{id}`  
Retrieve a specific bank account by its `id`.

**Request Parameters**:

- `id` (Path Parameter): The ID of the bank account.

**Response Example**:

```json
{
  "code": 200,
  "message": "Account fetched successfully",
  "data": {
    "id": 4,
    "user": {
      "id": 4,
      "name": "Gerard Considine",
      "email": "Annamae_Steuber@hotmail.com"
    },
    "balance": 0,
    "updated_at": "2025-01-01T08:36:17.634273Z"
  }
}
```

---

#### 4. **Update Bank Account by ID**

**PUT** `/{id}`  
Update a specific bank account by its `id`.

**Request Parameters**:

- `id` (Path Parameter): The ID of the bank account.

**Request Body**:

```json
{
  "balance": 10000
}
```

**Response Example**:

```json
{
  "code": 200,
  "message": "Account updated successfully",
  "data": {
    "id": 4,
    "user": {
      "id": 4,
      "name": "Gerard Considine",
      "email": "Annamae_Steuber@hotmail.com"
    },
    "balance": 900000,
    "updated_at": "2025-01-01T10:59:33.619638Z"
  }
}
```

---

#### 5. **Delete Bank Account by ID**

**DELETE** `/{id}`  
Delete a specific bank account by its `id`.

**Request Parameters**:

- `id` (Path Parameter): The ID of the bank account.

**Response Example**:

```json
{
  "code": 200,
  "message": "Account deleted successfully",
  "data": null
}
```

---

#### 6. **Transfer Between Accounts**

**POST** `/transfer`  
Transfer funds between two accounts.

**Request Body**:

```json
{
  "from_account_id": 1,
  "to_account_id": 2,
  "amount": 50000
}
```

**Response Example**:

```json
{
  "code": 200,
  "message": "Transfer completed successfully",
  "data": null
}
```

### Product Routes

---

### **Base URL**

`http://localhost:8080/api/products`

---

### **Middleware**

- **Public Routes**: Accessible without authentication.
- **Protected Routes**: Require authentication

**Header Example for Protected Routes**:

```json
{
  "Authorization": "Bearer <your_token>"
}
```

---

### **Endpoints**

---

### **Public Routes**

#### 1. **Find Product by ID**

**GET** `/api/products/{productId}`

Retrieve details of a specific product by its `productId`.

**Request Parameters**:

- `productId` (Path Parameter): The ID of the product to retrieve.

**Response Example**:

```json
{
  "code": 200,
  "message": "Product fetched successfully",
  "data": {
    "id": 5,
    "name": "Oriental Soft Gloves",
    "description": "Vinculum corpus aegrus angustus tempore condico canonicus vestigium. Abbas tunc alius colligo viridis conduco adflicto coepi acceptus absorbeo. Reprehenderit ulciscor surculus canis tui bene perspiciatis.",
    "price": 1000000,
    "user": {
      "id": 3,
      "name": "Eva Doyle",
      "email": "jwt@gmail.com"
    },
    "created_at": "2024-12-29T13:30:32.446838Z",
    "updated_at": "2024-12-29T13:30:32.446838Z"
  }
}
```

---

#### 2. **Search Products**

**GET** `/search`

Search for products using query parameters.

**Request Query Parameters**:

- `query` (Query Parameter): The search term or keyword.

**Request Example**:  
`GET /search?query=ori`

**Response Example**:

```json
{
  "code": 200,
  "message": "Products fetched successfully",
  "data": [
    {
      "id": 5,
      "name": "Oriental Soft Gloves",
      "description": "Vinculum corpus aegrus angustus tempore condico canonicus vestigium. Abbas tunc alius colligo viridis conduco adflicto coepi acceptus absorbeo. Reprehenderit ulciscor surculus canis tui bene perspiciatis.",
      "price": 1000000,
      "user": {
        "id": 3,
        "name": "Eva Doyle",
        "email": "jwt@gmail.com"
      },
      "created_at": "2024-12-29T13:30:32.446838Z",
      "updated_at": "2024-12-29T13:30:32.446838Z"
    },
    {
      "id": 6,
      "name": "Oriental Fresh Sausages",
      "description": "Autem stabilis vinum ambulo ipsa. Causa molestiae surgo reprehenderit. Corroboro pax demitto.",
      "price": 1000000,
      "user": {
        "id": 3,
        "name": "Eva Doyle",
        "email": "jwt@gmail.com"
      },
      "created_at": "2024-12-29T13:30:32.608467Z",
      "updated_at": "2024-12-29T13:30:32.608467Z"
    },
    {
      "id": 13,
      "name": "Oriental Steel Shirt",
      "description": "Atrox porro acer terga sumo volubilis cogo atrocitas suppellex. Tabella coepi charisma coerceo undique desparatus adduco. Vilicus demum quas arbor spiculum.",
      "price": 1000000,
      "user": {
        "id": 3,
        "name": "Eva Doyle",
        "email": "jwt@gmail.com"
      },
      "created_at": "2024-12-29T13:30:33.640919Z",
      "updated_at": "2024-12-29T13:30:33.640919Z"
    }
  ]
}
```

---

### **Protected Routes**

#### 3. **Get All Products**

**GET** `/`

Retrieve a list of all products.

**Response Example**:

```json
{
  "code": 200,
  "message": "Products fetched successfully",
  "data": [
    {
      "id": 5,
      "name": "Oriental Soft Gloves",
      "description": "Vinculum corpus aegrus angustus tempore condico canonicus vestigium. Abbas tunc alius colligo viridis conduco adflicto coepi acceptus absorbeo. Reprehenderit ulciscor surculus canis tui bene perspiciatis.",
      "price": 1000000,
      "user": {
        "id": 3,
        "name": "Eva Doyle",
        "email": "jwt@gmail.com"
      },
      "created_at": "2024-12-29T13:30:32.446838Z",
      "updated_at": "2024-12-29T13:30:32.446838Z"
    },
    {
      "id": 6,
      "name": "Oriental Fresh Sausages",
      "description": "Autem stabilis vinum ambulo ipsa. Causa molestiae surgo reprehenderit. Corroboro pax demitto.",
      "price": 1000000,
      "user": {
        "id": 3,
        "name": "Eva Doyle",
        "email": "jwt@gmail.com"
      },
      "created_at": "2024-12-29T13:30:32.608467Z",
      "updated_at": "2024-12-29T13:30:32.608467Z"
    },
    {
      "id": 13,
      "name": "Oriental Steel Shirt",
      "description": "Atrox porro acer terga sumo volubilis cogo atrocitas suppellex. Tabella coepi charisma coerceo undique desparatus adduco. Vilicus demum quas arbor spiculum.",
      "price": 1000000,
      "user": {
        "id": 3,
        "name": "Eva Doyle",
        "email": "jwt@gmail.com"
      },
      "created_at": "2024-12-29T13:30:33.640919Z",
      "updated_at": "2024-12-29T13:30:33.640919Z"
    }
  ]
}
```

---

#### 4. **Create a Product**

**POST** `/`

Create a new product.

**Request Body**:

```json
{
  "name": "Ergonomic Chair",
  "description": "A comfortable ergonomic chair.",
  "price": 150000
}
```

**Response Example**:

```json
{
  "code": 201,
  "message": "Product created successfully",
  "data": {
    "id": 32,
    "name": "Ergonomic Chair",
    "description": "A comfortable ergonomic chair.",
    "price": 150000,
    "user": {
      "id": 4,
      "name": "Gerard Considine",
      "email": "Annamae_Steuber@hotmail.com"
    },
    "created_at": "2025-01-01T11:19:38.613173152+07:00",
    "updated_at": "2025-01-01T11:19:38.613173201+07:00"
  }
}
```

---

#### 5. **Update a Product**

**PUT** `/{productId}`

Update an existing product by its `productId`.

**Request Parameters**:

- `productId` (Path Parameter): The ID of the product to update.

**Request Body**:

```json
{
  "name": "Ergonomic Chair",
  "description": "A comfortable ergonomic chair.",
  "price": 1700000
}
```

**Response Example**:

```json
{
  "code": 200,
  "message": "Product updated successfully",
  "data": {
    "id": 32,
    "name": "Ergonomic Chair",
    "description": "A comfortable ergonomic chair.",
    "price": 1700000,
    "user": {
      "id": 4,
      "name": "Gerard Considine",
      "email": "Annamae_Steuber@hotmail.com"
    },
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "2025-01-01T11:23:04.711809098+07:00"
  }
}
```

---

#### 6. **Delete a Product**

**DELETE** `/{productId}`

Delete a specific product by its `productId`.

**Request Parameters**:

- `productId` (Path Parameter): The ID of the product to delete.

**Response Example**:

```json
{
  "code": 200,
  "message": "Product deleted successfully",
  "data": null
}
```

#### 7. **Find Products by User ID**

**GET** `/user/{userId}`

This endpoint retrieves all products associated with a specific user by their `userId`.

**Path Parameter**:

- `userId`: The ID of the user to fetch products for.

**Response Example**:

```json
{
  "code": 200,
  "message": "Products fetched successfully",
  "data": [
    {
      "id": 27,
      "name": "Ergonomic Chair",
      "description": "A comfortable ergonomic chair.",
      "price": 150000,
      "user": {
        "id": 4,
        "name": "Gerard Considine",
        "email": "Annamae_Steuber@hotmail.com"
      },
      "created_at": "2025-01-01T11:14:32.948878Z",
      "updated_at": "2025-01-01T11:14:32.948878Z"
    },
    {
      "id": 33,
      "name": "Ergonomic Shoes",
      "description": "A comfortable ergonomic shoes.",
      "price": 190000,
      "user": {
        "id": 4,
        "name": "Gerard Considine",
        "email": "Annamae_Steuber@hotmail.com"
      },
      "created_at": "2025-01-01T11:37:46.928202986+07:00",
      "updated_at": "2025-01-01T11:37:46.928203038+07:00"
    }
  ]
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