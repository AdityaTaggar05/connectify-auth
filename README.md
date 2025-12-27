# Annora Auth Service

A **production-ready authentication microservice** built in **Go**, providing secure JWT-based authentication with **refresh token rotation**, **PostgreSQL persistence**, and a **clean layered architecture**.

This service is designed to be **independently deployable**, **frontend-agnostic**, and **microservice-ready**, with a strong focus on correctness, security, and maintainability.

---

## ✨ Features

- JWT **access tokens** (short-lived)
- **Refresh tokens** with rotation & revocation
- Secure password hashing (bcrypt)
- PostgreSQL-backed session persistence
- Logout & session expiry handling
- Clean **Handler → Service → Repository** architecture

---

## 🧱 Architecture

This service follows a **Layered Architecture (Service–Repository pattern)**:

- **Handlers**: HTTP concerns only
- **Services**: Authentication & session logic
- **Repositories**: Database access only

The service is designed to be deployed as **one microservice** in a larger system.

---

## 🔐 Authentication Model (Tokens)

| Token              | Purpose              | Lifetime            | Stored                     |
| ------------------ | -------------------- | ------------------- | -------------------------- |
| Access Token (JWT) | API authorization    | Short (e.g. 15 min) | Client memory              |
| Refresh Token      | Session continuation | Long (e.g. 7 days)  | DB + secure client storage |

- Access tokens are **stateless**
- Refresh tokens are **stored, rotated, and revocable**
- Logout revokes refresh tokens
- Expired tokens are cleaned up asynchronously

---
