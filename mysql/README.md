# How to Write and Run a MySQL Server in Docker

## 1. Mindset

### 1.1. First Step: Understand What You Need
Ask yourself:

- What service do you want? (e.g., PostgreSQL, MongoDB, Redis, etc.)
- Do you want to use an existing image (official Docker image) or build your own?
- Do you need custom configuration, data persistence (storage), or initialization scripts?

**Example:**
> "I want to run PostgreSQL database with Docker, use version 15, and set a custom password."

### 1.2. Search for an Official Image on Docker Hub
- Go to [Docker Hub](https://hub.docker.com/)
- Search for the service (e.g., PostgreSQL, MySQL)
- Find and use the official image (they are very well maintained!)

---

## 2. Common Docker Compose Fields for Database Services

When writing `docker-compose.yml` files for any database service (MySQL, MongoDB, Redis, etc.), you usually define these fields:

```yaml
services:
  database-service:
    image: <database-image>
    container_name: <your-container-name>
    restart: always
    environment:
      - TZ=Asia/Ho_Chi_Minh
      - <DB_USER_VARIABLE>=<your-user>
      - <DB_PASSWORD_VARIABLE>=<your-password>
      - <DB_DATABASE_VARIABLE>=<your-database>
    ports:
      - "<host-port>:<container-port>"
    volumes:
      - ./data:/var/lib/<database-folder>
    networks:
      - backend-network

networks:
  backend-network:
    driver: bridge
```

**Notes:**
- Always set `restart: always` for auto-restart if container crashes.
- Set timezone `TZ` properly to your server's timezone.
- Expose ports properly (e.g., 3306 for MySQL, 27017 for MongoDB, 6379 for Redis).
- Use volumes to persist database data even when the container is deleted.

Examples for Environment Variables:
- **MySQL**: `MYSQL_ROOT_PASSWORD`, `MYSQL_DATABASE`, `MYSQL_USER`, `MYSQL_PASSWORD`
- **PostgreSQL**: `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB`
- **MongoDB**: `MONGO_INITDB_ROOT_USERNAME`, `MONGO_INITDB_ROOT_PASSWORD`

---

## 3. Dockerfile and Entrypoint

### 3.1. Dockerfile for MySQL

#### 3.1.1. What is a Dockerfile?
- A Dockerfile is a blueprint that Docker uses to build a container image.

#### 3.1.2. How Does It Work?
- It specifies the base image, copies files, sets environment variables, installs dependencies, and defines the default behavior of the container.

### 3.2. Entrypoint

- The **Entrypoint** defines the **main command** that always runs when the container starts.
- It specifies the **executable** (a set of commands) and defines how the container should behave.

#### 3.2.1. Why Customize the Entrypoint?
- To prepare the environment before the main process starts (e.g., set permissions, initialize settings).
- To override or extend the default behavior provided by the base image.

**Some examples of Dockerfiles based on different MySQL versions:**
- [Supported tags and respective Dockerfile links](https://github.com/docker-library/mysql/blob/8ade9b2c9a32a79fbaa44b564d09a40744f1d105/8.4/Dockerfile.oracle)

---

## 4. Custom MySQL Configuration

- The default configuration directory for MySQL Docker containers is `/etc/mysql/conf.d`.
- Mongo allow mount the FOLDER contains CONFIG file
- Set ro (read-only) to avoid modify the config

```yaml
volumes:
  - ./data:/var/lib/mysql
  - ./config:/etc/mysql/conf.d
```

### 4.1. Important Notes

1. Store your custom configuration file inside a folder and mount it **read-only (ro)** to `/etc/mysql/conf.d`.
2. Assign permission `644` to your configuration file.
   - **Reason:** MySQL will ignore configuration files that are too permissive (e.g., with `777` permission).

### 4.2. Useful Commands Inside the Container

- `whoami`: Check the current running user inside the Docker container.
- `echo $ENV_VAR`: Check the value of an environment variable.

```yaml
environment:
  TZ: Asia/Ho_Chi_Minh
  MYSQL_ROOT_PASSWORD: ${ROOT_PASSWORD}
  MYSQL_DATABASE: chuong
```

- `printenv`: Print all environment variables inside the Docker container.

---

## 5. Overriding Command to Prepare and Start Database

When you need to fix file permissions or run extra preparation steps before starting the database server, you can override the default `command:` in `docker-compose.yml`.

### 5.1. Purpose
- Fix file permissions inside container.
- Start the database server properly with correct permissions.

### 5.2. Example

```yaml
command:
  - /bin/bash
  - -c
  - |
    chmod 644 /etc/mysql/conf.d/custom.cnf
    exec mysqld
```

**Explanation:**
- `/bin/bash -c` means execute the next block as a shell script.
- `chmod 644` fixes permission on your mounted configuration file.
- `exec mysqld` starts MySQL server cleanly, replacing the bash process.

This pattern is clean, safe, and works across MySQL, PostgreSQL, MongoDB, etc.

---

## 6. Backup and Restore MySQL Data

### 6.1. Backup

```bash
docker exec mysql-db sh -c 'exec mysqldump -uroot -p"${MYSQL_ROOT_PASSWORD}" ${MYSQL_DATABASE}' > backup.sql
```

---

### 6.2. Restore

Restore a database from a SQL file into a running container:
```bash
docker exec -i <container-name> sh -c 'exec mysql -u<username> -p"<password>" <database-name> < <input-file>.sql'
```

---

# Summary Checklist

- [x] Understand the service requirements.
- [x] Search and use an official Docker image.
- [x] Use a Dockerfile if heavy customization is needed.
- [x] Define volumes correctly for data persistence.
- [x] Customize `command:` or `entrypoint:` if preparation is needed before running the database.
- [x] Properly backup and restore databases with Docker commands.

---
