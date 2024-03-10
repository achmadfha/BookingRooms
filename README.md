# Booking Room Application

## Deskripsi Aplikasi

Aplikasi Booking Room dirancang untuk Perusahaan Enigma Camp dengan tujuan menggantikan sistem pemesanan ruangan yang saat ini masih dilakukan secara lisan dengan aplikasi berbasis digital. Aplikasi ini bertujuan untuk mengelola dan menyediakan informasi ketersediaan ruangan bagi semua karyawan.

## Kebutuhan Aplikasi

- Admin mempersiapkan data master yang mencakup informasi karyawan (divisi, jabatan, kontak), data ruangan (jenis ruangan, kapasitas, fasilitas), serta data login pengguna.
- Karyawan dapat mengakses sistem untuk memeriksa ketersediaan ruangan melalui antarmuka aplikasi.
- Karyawan yang ingin memesan ruangan harus masuk ke dalam aplikasi menggunakan kredensial login.
- GA memiliki akses untuk melihat daftar permintaan pemesanan dan memberikan persetujuan (accept) atau menolak (decline) dengan mencatat alasan penolakan jika diperlukan.
- Jika diperlukan, karyawan yang memesan ruangan dapat menyatakan kebutuhan peralatan tambahan seperti proyektor, mikrofon, notebook, dan sejenisnya.
- GA akan memperbarui status ketersediaan ruangan dalam sistem, menandai ruangan sebagai 'available' atau 'booked' berdasarkan persetujuan yang diberikan.
- GA, manajer, dan direktur memiliki hak akses untuk melihat seluruh proses pemesanan ruangan.
- Kemampuan untuk membuat laporan harian, bulanan, dan tahunan dari aktivitas pemesanan ruangan.
- Pengguna dapat mengunduh laporan dalam format CSV untuk analisis lebih lanjut.

## Actor (role)

- Admin
- Employee
- GA

## JSON API Collections

This repository contains a collection of API endpoints in JSON format, organized by functionality. Below are instructions on how to use each endpoint along with example requests and responses.

### Base Path

All endpoints in this collection share a common base path: `/api/v1/`. Ensure that you prepend this base path to all endpoint URLs when making requests.

## Authentication

All endpoints in this API require authentication using JWT (JSON Web Tokens). To access any endpoint, you must include a valid JWT token in the Authorization header of your requests.

## Endpoints

### Authentication Login Employee

#### Endpoint: `/auth/login`

- **Description**: This endpoint is used to log in to the system using basic authentication.
- **Methods**:
  - `POST`: Authenticate employee credentials and retrieve an access token.
- **Header**:
  - `Authorization`: `Basic Auth admin:admin`
- **Request**:
  - **Status Code**: 200 OK
  - **Body**:

```json
{
  "username": "string",
  "password": "string"
}
```

- **Response**:
  - **Status Code**: 200 OK
  - **Body**:

```json
{
  "responseCode": "string",
  "responseMessage": "string",
  "data": {
    "access_token": "string"
  }
}
```

#### Endpoint: `/auth/password`

- **Description**: This endpoint is used to update the password of a user.
- **Methods**:
  - `POST`: Change user password.
- **Header**:
  - `Authorization`: `Bearer <your JWT token>`
- **Request**:
  - **Status Code**: 200 OK
  - **Body**:

```json
{
  "username": "string",
  "oldPassword": "string",
  "newPassword": "string",
  "confirmPassword": "string"
}
```

- **Response**:
  - **Status Code**: 200 OK
  - **Body**:

```json
{
  "responseCode": "string",
  "responseMessage": "string"
}
```

### Employee Management

#### Endpoint: `/employee`

- **Description**: Retrieve all employee data using this endpoint.
- **Methods**:
  - `GET`: Retrieve a list of employee.
- **Header**:
  - `Authorization`: `Bearer <your JWT token>`
- Optional Parameters:
  - **page**: (Optional) The page number for paginated results.
  - **size**: (Optional) The number of transactions per page.
- **Response**:
  - **Status Code**: 200 OK
  - **Body**:

```json
{
    "responseCode": "string",
    "responseMessage": "string",
    "data": [
        {
            "employeeId": "uuid",
            "fullName": "string",
            "division": "string",
            "phoneNumber": "string",
            "position" : "string"
        },
        {
            "employeeId": "uuid",
            "fullName": "string",
            "division": "string",
            "phoneNumber": "string",
            "position" : "string"
        }
    ],
    "paging": {
        "page": int,
        "totalPages": int,
        "totalData": int
    }
}
```

#### Endpoint: `/employee`

- **Description**: This endpoint allows you to create a new employee.
- **Methods**: `POST`
- **Header**:
  - `Authorization`: `Bearer <your JWT token>`
- **Body** :

```json
{
  "fullName": "string",
  "division": "string",
  "phoneNumber": "string",
  "position": "string",
  "username": "string"
}
```

- **Response**:
  - **Status Code**: 200 OK
  - **Body**:

```json
{
  "responseCode": "string",
  "responseMessage": "string",
  "data": {
    "employeeId": "uuid",
    "fullName": "string",
    "division": "string",
    "phoneNumber": "string",
    "position": "string",
    "username": "string"
  }
}
```

#### Endpoint: `/employee/:id`

- **Description**: Retrieve a specific employee by its ID.
- **Methods**: `GET`
- **Header**:
  - `Authorization`: `Bearer <your JWT token>`
- **Response**:
  - **Status Code**: 200 OK
  - **Body**:

```json
{
  "responseCode": "string",
  "responseMessage": "string",
  "data": {
    "employeeId": "string",
    "fullName": "string",
    "division": "string",
    "phoneNumber": "string",
    "position": "string"
  }
}
```

#### Endpoint: `/employee/:id`

- **Description**: This endpoint allows you to update employee.
- **Methods**: `PUT`
- **Header**:
  - `Authorization`: `Bearer <your JWT token>`
- **Body** :
  - field body are optional, update fields u want

```json
{
  "fullName": "string",
  "division": "string",
  "phoneNumber": "string",
  "position": "string"
}
```

- **Response**:
  - **Status Code**: 200 OK
  - **Body**:

```json
{
  "responseCode": "string",
  "responseMessage": "string",
  "data": {
    "employeeId": "uuid",
    "fullName": "string",
    "division": "string",
    "phoneNumber": "string",
    "position": "string"
  }
}
```

#### Endpoint: `/employee/:id`

- **Description**: Delete a specific employee by its ID.
- **Methods**: `DELETE`
- **Header**:
  - `Authorization`: `Bearer <your JWT token>`
- **Response**:
  - **Status Code**: 200 OK
  - **Body**:

```json
{
  "responseCode": "string",
  "responseMessage": "string",
  "data": "string"
}
```

### Room Management

#### Endpoint: `/room`

- **Description**: Retrieve all room data using this endpoint.
- **Methods**:
  - `GET`: Retrieve a list of room.
- **Header**:
  - `Authorization`: `Bearer <your JWT token>`
- Optional Parameters:
  - **page**: (Optional) The page number for paginated results.
  - **size**: (Optional) The number of rooms per page.
- **Response**:
  - **Status Code**: 200 OK
  - **Body**:

```json
{
    "responseCode": "string",
    "responseMessage": "string",
    "data": [
        {
            "room_id": "uuid",
            "room_details_id": "uuid",
            "name": "string",
            "status": "string"
        },
        {
          "room_id": "uuid",
          "room_details_id": "uuid",
          "name": "string",
          "status": "string"
        }
    ],
    "paging": {
        "page": int,
        "totalPages": int,
        "totalData": int
    }
}
```

#### Endpoint: `/room/:id`

- **Description**: Retrieve a specific room by its ID.
- **Methods**: `GET`
- **Header**:
  - `Authorization`: `Bearer <your JWT token>`
- **Response**:
  - **Status Code**: 200 OK
  - **Body**:

```json
{
    "responseCode": "string",
    "responseMessage": "string",
    "data": {
        "room_id": "uuid",
        "room_details": {
            "room_details_id": "uuid",
            "room_type": "string",
            "capacity": int,
            "facility": [
                "string",
                "string"
            ]
        },
        "name": "string",
        "status": "string"
    }
}
```

#### Endpoint: `/room`

- **Description**: This endpoint allows you to create a new room.
- **Methods**: `POST`
- **Header**:
  - `Authorization`: `Bearer <your JWT token>`
- **Body** :

```json
{
    "name": "string",
    "status": "string",
    "room_type": "string",
    "capacity": int,
    "facility": [
        "string",
        "string"
    ]
}
```

- **Response**:
  - **Status Code**: 201 Created
  - **Body**:

```json
{
  "responseCode": "string",
  "responseMessage": "string",
  "data": {
    "room_id": "uuid",
    "room_details": {
      "room_details_id": "uuid",
      "room_type": "string",
      "capacity": int,
      "facility": [
        "string",
        "string"
      ]
    },
    "name": "string",
    "status": "string"
  }
}
```

#### Endpoint: `/room/:id`

- **Description**: This endpoint allows you to update room.
- **Methods**: `PUT`
- **Header**:
  - `Authorization`: `Bearer <your JWT token>`
- **Body** :
  - field body are optional, update fields u want

```json
{
    "name": "string",
    "status": "string",
    "room_type": "string",
    "capacity": int,
    "facility": [
        "string",
        "string"
    ]
}
```

- **Response**:
  - **Status Code**: 200 OK
  - **Body**:

```json
{
  "responseCode": "string",
  "responseMessage": "string"
}
```

### Transactions Management

#### Endpoint: `/transactions`

- **Description**: Retrieve all transaction data using this endpoint.
- **Methods**:
  - `GET`: Retrieve a list of transactions.
- **Header**:
  - `Authorization`: `Bearer <your JWT token>`
- Optional Parameters:
  - **page**: (Optional) The page number for paginated results.
  - **size**: (Optional) The number of transactions per page.
  - **startDate**: (Optional) The start date for filtering transactions.
  - **endDate**: (Optional) The end date for filtering transactions.
- **Response**:
  - **Status Code**: 200 OK
  - **Body**:

```json
{
  "responseCode": "string",
  "responseMessage": "string",
  "data": [
    {
      "transaction_id": "uuid ",
      "employee_id": "uuid",
      "room_id": "uuid",
      "start_date": "date",
      "end_date": "date",
      "description": "string",
      "status": "string",
      "created_at": "date",
      "updated_at": "date"
    },
    {
      "transaction_id": "uuid ",
      "employee_id": "uuid",
      "room_id": "uuid",
      "start_date": "date",
      "end_date": "date",
      "description": "string",
      "status": "string",
      "created_at": "date",
      "updated_at": "date"
    }
  ],
  "paging": {
    "page": int,
    "totalPages": int,
    "totalData": int
  }
}
```

#### Endpoint: `/transactions/:id`

- **Description**: Retrieve a specific transaction by its ID.
- **Methods**:
  - `GET`: Retrieve a transaction by its ID.
- **Header**:
  - `Authorization`: `Bearer <your JWT token>`
- **Response**:
  - **Status Code**: 200 OK
  - **Body**:

```json
{
    "responseCode": "string",
    "responseMessage": "string",
    "data": {
        "transaction_id": "uuid",
        "employee": {
            "employee_id": "uuid",
            "full_name": "string",
            "division": "string",
            "phone_number": "string",
            "position": "string"
        },
        "room": {
            "room_id": "uuid",
            "room_details": {
                "room_details_id": "uuid",
                "room_type": "string",
                "capacity": int,
                "facility": [
                    "string",
                    "string"
                ]
            },
            "name": "string",
            "status": "string"
        },
        "start_date": "date",
        "end_date": "date",
        "description": "string",
        "status": "string",
        "created_at": "date",
        "updated_at": "date"
    }
}
```

#### Endpoint: `/transactions`

- **Description**: This endpoint allows you to create a new transaction.
- **Methods**: `POST`
- **Header**:
  - `Authorization`: `Bearer <your JWT token>`
- **Body**:

```json
{
  "employee_id": "uuid",
  "room_id": "uuid",
  "start_date": "date",
  "end_date": "date",
  "description": "string"
}
```

- **Response**:
  - **Status Code**: 201 Created
  - **Body**:

```json
{
  "responseCode": "string",
  "responseMessage": "string",
  "data": {
    "id": "uuid",
    "employee_id": "uuid",
    "room_id": "uuid",
    "start_date": "date",
    "end_date": "date",
    "description": "string",
    "status": "string"
  }
}
```

#### Endpoint: `/transactions/logs/:id`

- **Description**: This endpoint allows you to accept or decline a transaction.
- **Methods**:
  - `PUT` : `/:id` use ur transactions log id
- **Header**:
  - `Authorization`: `Bearer <your JWT token>`
- **Body**:

```json
{
  "approved_by": "uuid",
  "approval_status": "string",
  "description": "string"
}
```

- **Response**:
  - **Status Code**: 200 OK
  - **Body**:

```json
{
  "responseCode": "string",
  "responseMessage": "string",
  "data": {
    "approved_by": "string",
    "approval_status": "string",
    "description": "string"
  }
}
```

#### Endpoint: `/transactions/logs/:id`

- **Description**: This endpoint allows you to see transactions accept or decline and the reason for it.
- **Methods**:
  - `GET` : `/:id` use ur transactions log id
- **Header**:
  - `Authorization`: `Bearer <your JWT token>`
- **Response**:
  - **Status Code**: 200 OK
  - **Body**:

```json
{
  "responseCode": "string",
  "responseMessage": "success",
  "data": {
    "transaction_log_id": "uuid",
    "transaction": {
      "transaction_id": "uuid",
      "employee": {
        "employee_id": "uuid",
        "full_name": "string",
        "division": "string",
        "phone_number": "string",
        "position": "string"
      },
      "room": {
        "room_id": "uuid",
        "room_details": {
          "room_details_id": "uuid",
          "room_type": "string",
          "capacity": int,
          "facility": [
            "string",
            "string"
          ]
        },
        "name": "string",
        "status": "string"
      },
      "start_date": "date",
      "end_date": "date",
      "description": "string",
      "status": "string",
      "created_at": "date",
      "updated_at": "date"
    },
    "approved_by": {
      "employee_id": "uuid",
      "full_name": "string",
      "division": "string",
      "phone_number": "string",
      "position": "string"
    },
    "approval_status": "string",
    "description": "string",
    "created_at": "date",
    "updated_at": "date"
  }
}
```

#### Endpoint: `/transactions/logs`

- **Description**: Retrieve all log transaction data using this endpoint.
- **Methods**:
  - `GET`: Retrieve a list of transaction logs.
- **Header**:
  - `Authorization`: `Bearer <your JWT token>`
- Optional Parameters:
  - **page**: (Optional) The page number for paginated results.
  - **size**: (Optional) The number of transactions per page.
  - **startDate**: (Optional) The start date for filtering transactions.
  - **endDate**: (Optional) The end date for filtering transactions.
- **Response**:
  - **Status Code**: 200 OK
  - **Body**:

```json
{
  "responseCode": "2000306",
  "responseMessage": "success",
  "data": [
    {
      "transaction_log_id": "uuid",
      "transaction_id": "uuid",
      "approved_by": "uuid",
      "approval_status": "string",
      "description": "string",
      "created_at": "date",
      "updated_at": "date"
    },
    {
      "transaction_log_id": "uuid",
      "transaction_id": "uuid",
      "approved_by": "uuid",
      "approval_status": "string",
      "description": "string",
      "created_at": "date",
      "updated_at": "date"
    }
  ],
  "paging": {
    "page": int,
    "totalPages": int,
    "totalData": int
  }
}
```
