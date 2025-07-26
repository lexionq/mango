# mango 🥭
Password manager for CLI, written in Golang.  

<p align="center">
  <img alt="GitHub License" src="https://img.shields.io/github/license/lexionq/mango?style=for-the-badge&logoColor=blue&color=blue">
  <img src="https://img.shields.io/github/languages/top/lexionq/mango?style=for-the-badge&color=cyan">
  <img src="https://img.shields.io/github/v/release/lexionq/mango?style=for-the-badge&color=purple">  
  <img width="90" height="90" alt="image" src="https://github.com/user-attachments/assets/5e5a2689-3f9f-4e80-a56c-f040f54e6073" />
</p>
<p align="center">
  <img src="https://img.shields.io/github/go-mod/go-version/lexionq/mango?style=for-the-badge&color=darkblue">
  <img alt="GitHub Actions Workflow Status" src="https://img.shields.io/github/actions/workflow/status/lexionq/mango/go.yml?style=for-the-badge&color=darkgreen">
</p>
  
## Disclaimer ⚠️
>[!CAUTION]
> ***This project is still in beta. The program owner, lexionq, says they're not responsible if you lose or leak any passwords you've saved using the mango password manager.***
>
> ***When you use Mango, it will ask you for a master password. Save this password and don't forget it! Otherwise, all your passwords will be irretrievable!***

## Platform 💻
This project works in Linux systems(if system have `nano`)

## Feautures 🌠
When you use mango, you encrypt registers in the format `[name,website,username,password,note]` using a master password of your choice, and then use this password to access them.
- Add registers
- List registers
- Edit registers

## How to use?

### Add Register
`./mango add`
### List Registers
`./mango list`
### Edit Registers
`./mango edit`
### Search a word in Registers
`./mango search <word>`

## How to install?

#### For compiled version
- Enter [mango GitHub Releases](https://github.com/mango/releases).
- Download the latest version.
- And run `mango` with use `./mango`.

#### Use on source code
```bash
git clone https://github.com/lexionq/mango mango
cd mango
go run main.go
```
