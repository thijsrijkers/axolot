# Axolot

**Axolot** is a lightweight, secure, in-memory user account manager written in Go. It allows for the creation of user accounts with usernames and passwords, with all data encrypted and stored in memory. Axolot ensures user data is tightly bound to the host machine, making it inaccessible to anything other than the application itself.

## ✨ Features

- 👤 **User Account Creation**  
  Easily create user accounts using a `username` and `password`.

- 🧠 **Encrypted In-Memory Storage**  
  All user data is encrypted and stored in memory — never written to disk unless explicitly exported.

- 🔐 **Machine-Tied Encryption**  
  Encryption is uniquely derived from host-specific identifiers, making data unreadable outside the original host.

- 📦 **Encrypted Export**  
  Export the user store in a secure, encrypted format for backup or migration.

## 🔁 Portable Encryption Rules

- Export the encryption rules separately to re-import and decrypt data on another host machine securely.

> ⚠️ **One-Time Export Rule**  
> The encryption rules can only be exported **once** during the **first launch** of the application. After that, they are locked and cannot be retrieved again.
> 


## 🛡️ Security Principles

- In-memory only storage to prevent disk-level access.
- Encryption keys are never exposed externally.
- Data cannot be accessed by any process or application other than Axolot itself.
- All exported data must be decrypted using matching host-derived rules.
