# Login with Fiber Framework

This project is a simple login, registration, and user dashboard web application. It uses **Go Fiber** as the backend framework and serves static files (HTML, CSS, and JavaScript) for the frontend.

---

## **Project Structure**

```
├── configs/              # Configuration files
├── controllers/          # Controller logic for API
├── middlewares/          # Middleware functions
├── models/               # Data models
├── routes/               # API route definitions
├── tmp/                  # Temporary files (used by Air for hot reload)
├── UI/                   # Frontend files
│   ├── js/               # JavaScript logic
│   │   └── script.js
│   ├── index.html        # Login page
│   ├── register.html     # Registration page
│   └── user.html         # User dashboard
├── .air.toml             # Air configuration for hot reload
├── .env                  # Environment variables
├── curl-cmds.txt         # Sample curl commands for API testing
├── go.mod                # Go module dependencies
├── go.sum                # Checksums for dependencies
├── main.go               # Application entry point
```

---

## **Features**

1. **Login Page:** Accessible at `/index.html`.
2. **Register Page:** Accessible at `/register.html`.
3. **User Dashboard:** Shows user details and includes a logout button.
4. **Password Hashing:** Ensures secure storage of user passwords.
5. **Cookies for Auth Token:** Uses cookies to store authentication tokens securely.

---

## **Requirements**

- **Go 1.20+**
- **Air** (for hot reloading, optional but recommended)

---

## **Setup Instructions**

### 1. Clone the Repository

```bash
git clone <repository-url>
cd login-with-fiber-framework
```

### 2. Install Dependencies

Ensure you have all the required Go packages:

```bash
go mod tidy
```

### 3. Run the Project

#### Option 1: Using Air (Hot Reloading)

Install Air if you haven’t already:

```bash
go install github.com/cosmtrek/air@latest
```

Run the server with Air:

```bash
air
```

#### Option 2: Directly Run with Go

```bash
go run main.go
```

---

## **Usage**

1. Start the server (as described above).
2. Open your browser and navigate to `http://localhost:3000/`.
   - **Login Page:** `http://localhost:3000/index.html`
   - **Register Page:** `http://localhost:3000/register.html`
   - **User Dashboard:** `http://localhost:3000/user.html`
3. Use the following curl commands to test the API endpoints:

   **Login**
   ```bash
   curl -X POST "http://localhost:3000/api/v1/login" -d "email=jackson8000@gmail.com&password=pass1234"
   ```
   **Register**
   ```bash
   curl -X POST "http://localhost:3000/api/v1/register" -d "name=Jackson Miller&email=jackson8000@gmail.com&password=pass1234"
   ```
   **Logout**
   ```bash
   curl -X GET "http://localhost:3000/api/v1/logout"
   ```
   **Get User Info**
   ```bash
   curl -X GET "http://localhost:3000/api/v1/user"
   ```

---

## **Static File Hosting**

The project contains a basic User interface, `UI` directory contains the frontend files (HTML, CSS, JS) and is served as static files using the `fiber.Static()` function in `main.go`.\
You can access the following pages directly:

- `index.html`: Login page.
- `register.html`: Registration page.
- `user.html`: User dashboard.

---

## **Environment Variables**

- **MONGODB\_URI:**  Mongo DB connection string.
- **MONGO\_DATABASE:** Name of the database.
- **JWT\_SECRET:** Secret key for JWT encoding.

---

## **Contributing**

Feel free to fork the repository, create feature branches, and submit pull requests.

---

## **License**

This project is licensed under the MIT License.

