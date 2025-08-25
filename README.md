# mango 🥭

🔐🔑 Password manager for CLI, written in Golang.

<p align="center">
  <img width="90" height="90" alt="image" src="https://github.com/user-attachments/assets/cc354bcc-0c60-49c6-b6ba-7837af62e9b8" />
  <img width="90" height="90" alt="image" src="https://github.com/user-attachments/assets/5e5a2689-3f9f-4e80-a56c-f040f54e6073" />
  <br>
  <img alt="GitHub License" src="https://img.shields.io/github/license/lexionq/mango?style=for-the-badge&logoColor=blue&color=blue">
  <img src="https://img.shields.io/github/languages/top/lexionq/mango?style=for-the-badge&color=cyan">
  <img alt="GitHub Release" src="https://img.shields.io/github/v/release/lexionq/mango?style=for-the-badge&color=purple"> 
</p>

<p align="center">
  <img src="https://img.shields.io/github/go-mod/go-version/lexionq/mango?style=for-the-badge&color=darkblue">
  <img alt="GitHub Actions Workflow Status" src="https://img.shields.io/github/actions/workflow/status/lexionq/mango/go.yml?style=for-the-badge&color=darkgreen">
  <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/lexionq/mango?style=for-the-badge">
  <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/lexionq/mango?style=for-the-badge">
</p>
<div align="center"

[🥭 About](#-about-mango) • [✨ Features](#-key-features) • [📦 Install](#-installation) • [🔧 Commands](#-commands-reference) • [🆕 Updates](#-updates--maintenance) • [🤝 Contributing](#-contributing-and-contributors)

</div>

## 🥭 About Mango

Mango is a **lightweight, secure, and local** password manager for the terminal. All your credentials are encrypted and stored **locally** - no cloud, no servers, just you and your data. 🔐

### 🚀 Quick Start
```bash
# Install mango
curl -sSL https://raw.githubusercontent.com/lexionq/mango/main/setup.sh | bash

# First run - create your master password
mango add

# Add your first password
# Name: Gmail
# Website: gmail.com  
# Username: mysupersecure@gmail.com
# Password: [secure password]
# Notes: Personal email

# List all passwords
mango list

# Search for specific entries
mango search gmail

# Copy password to clipboard
mango copy

# Generate a strong password
mango generate 16

# Check password strength
mango strength mypassword123

# Create backup
mango backup
```

## ⚠️ Disclaimer

> [!CAUTION]
> **🚨 Warning**
>
> **_The program owner, lexionq, says they're not responsible if you lose or leak any passwords you've saved using the mango password manager._**
>
> **🧠 Don't Forget your Master Password**
>
> **_When you use Mango, it will ask you for a master password. Save this password and don't forget it! Otherwise, all your passwords will be irretrievable!_**

## Platform 💻

🟢 **Linux** ✅  
🍏 **macOS** ✅  
🪟 **Windows** ✅ (testing)

📝 Automatic editor detection (supports nano, vim, notepad, VS Code, and more)

## ✨ Key Features

### 🔒 Security First
- **AES-256 Encryption** - Military-grade encryption for your passwords
- **Salt & Hash** - Secure master password storage with bcrypt
- **Local Storage** - No cloud, no servers, your data stays on your device
- **Password Strength Analysis** - Real-time security assessment

### 🌍 Cross-Platform
- **Windows, macOS, Linux** - Works everywhere
- **Smart Editor Detection** - Automatically finds your preferred editor
- **Clipboard Integration** - Secure copy with auto-clear (30s)

### 📊 Data Management
- **Multiple Export Formats** - Encrypted backup & CSV export
- **Import/Export** - Easy migration between devices
- **Database Statistics** - Track your password health
- **Automatic Backups** - Never lose your data

### 💡 User Experience
- **Interactive CLI** - Clean, intuitive interface
- **Search & Filter** - Quickly find what you need
- **Batch Operations** - Manage multiple entries efficiently
- **Zero Dependencies** - Single binary, no installation hassles

**Data Format:** `[name, website, username, password, notes]`

## 🔧 Commands Reference

### 🔐 Basic Operations
| Command | Description | Example |
|---------|-------------|---------|
| `add` | Add new password entry | `mango add` |
| `list` | Show all saved passwords | `mango list` |
| `edit` | Edit existing entries | `mango edit` |
| `delete` | Remove a password entry | `mango delete` |
| `search` | Find entries by keyword | `mango search google` |

### �️ dPassword Tools
| Command | Description | Example |
|---------|-------------|---------|
| `generate` | Create secure password | `mango generate 16` |
| `strength` | Analyze password security | `mango strength mypass123` |
| `copy` | Copy password to clipboard | `mango copy` |
| `change` | Update master password | `mango change` |

### � Bacrkup & Data Management
| Command | Description | Example |
|---------|-------------|---------|
| `backup` | Create encrypted backup | `mango backup` |
| `restore` | Restore from backup | `mango restore` |
| `export` | Export encrypted database | `mango export` |
| `import` | Import encrypted database | `mango import` |
| `exportcsv` | Export to CSV format | `mango exportcsv` |
| `importcsv` | Import from CSV file | `mango importcsv` |

### 📊 Analytics
| Command | Description | Example |
|---------|-------------|---------|
| `stats` | View database statistics | `mango stats` |

### 🆘 Help
```bash
mango --help    # Show all commands
mango -h        # Short help
```

## 📦 Installation

### Method 1: Pre-built Binary (Recommended)
```bash
# Download and run the setup script
curl -sSL https://raw.githubusercontent.com/lexionq/mango/main/setup.sh | bash
```

### Method 2: Build from Source
```bash
# Clone the repository
git clone https://github.com/lexionq/mango.git
cd mango

# Using Makefile (recommended)
make build              # Standard build
make build-optimized    # Smaller binary
make install            # Build and install to PATH

# Or manual build
go build -o mango main.go
./mango --help
```

### Method 3: Go Install
```bash
go install github.com/lexionq/mango@latest
```

### Method 4: Development Mode
```bash
git clone https://github.com/lexionq/mango.git
cd mango
go run main.go --help
```

## 🤝 Contributing and Contributors

PRs, issues, and suggestions are welcome!
Star ⭐ the repo if you like it — it motivates the developer 😄

<a href="https://github.com/lexionq/mango/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=lexionq/mango" />
</a>

## 🎯 Features

- [x] 🔐 AES Encryption
- [x] 🧂 Salting
- [x] 🪟 Windows & 🍏 macOS support
- [x] 🔁 Import/Export options (encrypted & CSV)
- [x] 📋 Clipboard integration with auto-clear
- [x] 🔒 Password strength analysis
- [x] 💾 Automatic backup system
- [x] 📊 Database statistics
- [x] 🗑️ Secure password deletion
- [x] 🎯 Cross-platform editor support
- [ ] 🔐 Two-Factor Authentication (2FA)
- [ ] 📱 Password expiry notifications
- [ ] 🌐 GUI version (maybe?)
- [ ] 🏷️ Categories and tags system

## 🔄 Updates & Maintenance

### Update Mango
```bash
# If installed via setup script
curl -sSL https://raw.githubusercontent.com/lexionq/mango/main/setup.sh | bash

# If installed via go install
go install github.com/lexionq/mango@latest

# If built from source
git pull origin main && go build -o mango . 
```

### Build Options
```bash
# Standard build
go build -o mango .

# Optimized build (smaller binary)
go build -ldflags="-s -w" -o mango main.go

# Cross-compile for different platforms
GOOS=windows GOARCH=amd64 go build -o mango.exe main.go  # Windows
GOOS=darwin GOARCH=amd64 go build -o mango-mac main.go   # macOS
GOOS=linux GOARCH=amd64 go build -o mango-linux main.go  # Linux
```

### Performance Tips
- Use `mango backup` regularly to protect your data
- Check `mango stats` to monitor password health
- Run `mango strength` on important passwords periodically
