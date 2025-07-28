# mango 🥭
🔐🔑 Password manager for CLI, written in Golang.  

<p align="center">
  <img width="90" height="90" alt="image" src="https://github.com/user-attachments/assets/cc354bcc-0c60-49c6-b6ba-7837af62e9b8" />
  <img width="90" height="90" alt="image" src="https://github.com/user-attachments/assets/5e5a2689-3f9f-4e80-a56c-f040f54e6073" />
  <br>
  <img alt="GitHub License" src="https://img.shields.io/github/license/lexionq/mango?style=for-the-badge&logoColor=blue&color=blue">
  <img src="https://img.shields.io/github/languages/top/lexionq/mango?style=for-the-badge&color=cyan">
  <img src="https://img.shields.io/github/v/release/lexionq/mango?style=for-the-badge&color=purple">  
</p>

<p align="center">
  <img src="https://img.shields.io/github/go-mod/go-version/lexionq/mango?style=for-the-badge&color=darkblue">
  <img alt="GitHub Actions Workflow Status" src="https://img.shields.io/github/actions/workflow/status/lexionq/mango/go.yml?style=for-the-badge&color=darkgreen">
  <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/lexionq/mango?style=for-the-badge">
</p>
<div align="center"

[🥭 About Mango](#-about-mango) • [⚠️ Disclaimer](#%EF%B8%8F-disclaimer) • [Platform 💻](#platform-) • [Features 🌠](#features-) • [🔧 Available Commands](#-available-commands) • [📦 How to install?](#-how-to-install) • [🤝 Contributing](#-contributing) 

• [🧠 Future Plans](#-future-plans)

</div>

## 🥭 About Mango
Mango is a a lightweight and local password manager for the terminal. All your credentials are encrypted and stored **locally,** accessible **only with your master password.** 🔐🔓


🐧 Currently supports only Linux systems(requires `nano`)
🚀 Cross-platform support is coming soon! Stay tuned... 🥳

## ⚠️ Disclaimer 
>[!CAUTION]
>**🚨 Beta Warning**
>
> ***This project is still in beta. The program owner, lexionq, says they're not responsible if you lose or leak any passwords you've saved using the mango password manager.***
>
>**🧠 Don't Forget your Master Password**
>
> ***When you use Mango, it will ask you for a master password. Save this password and don't forget it! Otherwise, all your passwords will be irretrievable!***

## Platform 💻
🟢 **Linux *only*** (for now)

📝 Requires `nano` for editing registers

## Features 🌠
With Mango, you can securely manage your credentials in the format:

`[name, website, username, password, note]`

🔐 All data is AES encrypted using your master password.

## 🔧 Available Commands

- ➕ Add Register
`mango add`
- 📎 List Registers
`mango list`
- 📝 Edit Registers
`mango edit`
- 🔍 Search among Registers
`mango search <word>`

## 📦 How to install?
- 📥 Download the [`setup.sh`](https://github.com/lexionq/mango)
- 🛡️ Make it executable:
- ```bash
  chmod +x setup.sh
  ```
- ▶️ Run the script:
- ```bash
  ./setup.sh
  ```
- ✅ Mango is now installed and ready to use!

## 🔽 Or use source code
```bash
git clone https://github.com/lexionq/mango mango
cd mango
go run main.go
```

## 🤝 Contributing
PRs, issues, and suggestions are welcome!
Star ⭐ the repo if you like it — it motivates the developer 😄

## 🧠 Future Plans
- [x] 🔐 AES Encryption
- [x] 🧂 Salting
- [ ] 🪟 Windows & 🍏 macOS sypport
- [ ] 🌐 GUI version (maybe?)
- [ ] 🔁 Import/Export options
