# mango 🥭
🔐🔑 Password manager for CLI, written in Golang.  
<p align="right"
  <img width="90" height="90" alt="image" src="https://github.com/user-attachments/assets/5e5a2689-3f9f-4e80-a56c-f040f54e6073" />
</p>
<p align="center">
  <img alt="GitHub License" src="https://img.shields.io/github/license/lexionq/mango?style=for-the-badge&logoColor=blue&color=blue">
  <img src="https://img.shields.io/github/languages/top/lexionq/mango?style=for-the-badge&color=cyan">
  <img src="https://img.shields.io/github/v/release/lexionq/mango?style=for-the-badge&color=purple">  
</p>

<p align="center">
  <img src="https://img.shields.io/github/go-mod/go-version/lexionq/mango?style=for-the-badge&color=darkblue">
  <img alt="GitHub Actions Workflow Status" src="https://img.shields.io/github/actions/workflow/status/lexionq/mango/go.yml?style=for-the-badge&color=darkgreen">
  <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/lexionq/mango?style=for-the-badge">
</p>


🥭 Mango is a password manager that lets you store all your passwords locally on your system using a master password that you set via the terminal interface.
Currently, it only supports Linux. Cross-platform support will be available soon. 🥳

## Disclaimer ⚠️
>[!CAUTION]
> ***This project is still in beta. The program owner, lexionq, says they're not responsible if you lose or leak any passwords you've saved using the mango password manager.***
>
> ***When you use Mango, it will ask you for a master password. Save this password and don't forget it! Otherwise, all your passwords will be irretrievable!***

## Platform 💻
This project works in Linux systems(if system have `nano`)

## Feautures 🌠
When you use mango, you encrypt registers in the format `[name,website,username,password,note]` using a master password of your choice, and then use this password to access them.
- *Add registers*
- *List registers*
- *Edit registers*
- *Search word among registers*

## How to use?

### Add Register
`mango add`
### List Registers
`mango list`
### Edit Registers
`mango edit`
### Search a word among Registers
`mango search <word>`

## How to install?
- Enter [mango GitHub repository](https://github.com/lexionq/mango)
- Download the `setup.sh`
- Run `chmod +x setup.sh`(if required)
- And, `./setup.sh`
- Finally, mango installed on your computer.

#### Use on source code
```bash
git clone https://github.com/lexionq/mango mango
cd mango
go run main.go
```
