# Contributing Guide

Thank you for your interest in contributing to this project! 🥭  
Your help is highly appreciated. Please read the following guidelines before getting started.

---

## 🚀 How to Contribute

To contribute to this project, please follow these steps:

1. **Fork** this repository to your GitHub account.
2. **Clone** your fork to your local machine:

   ```bash
   git clone https://github.com/YOUR_USERNAME/REPO_NAME.git
   ```

3. Make your changes **directly on the `main` branch** of your fork.
4. Format your code using `gofmt`:

   ```bash
   gofmt -w .
   ```

5. Test your changes locally. If you're adding a new feature or fixing a bug, make sure it works as expected:

   ```bash
   go build
   go run main.go
   ```

6. Commit your changes with a **clear and descriptive message**:

   ```bash
   git commit -m "fix: corrected password decryption bug"
   ```

7. Push your changes to your fork:

   ```bash
   git push origin main
   ```

8. Open a **pull request** to the `main` branch of this repository.

> 📌 You don't need to create a new branch. All changes can be made directly on the `main` branch of your fork.

---

## ✅ Contribution Rules

- Keep your changes **focused and relevant**.
- Do **not** include unrelated changes in a single pull request.
- Write **clear and meaningful commit messages** (e.g., `add: new CLI command for listing entries`).
- Follow Go coding conventions and use tools like `gofmt`.
- Keep the code **simple, modular, and readable**.
- Avoid duplicating code – reuse existing functions when possible.

---

## 🧪 Testing (Optional but Recommended)

If the feature or fix can be tested:

- Write basic test functions under the appropriate package.
- Make sure all tests pass before submitting a pull request.

---

## 💬 Need Help?

If you have questions, suggestions, or if you're unsure about anything, feel free to:

- Open an issue on GitHub.
- Comment directly on an existing issue or pull request.
- Or contact me with e-mail : [my email](mailto:lexionq@proton.me)
    
---

Thanks again for helping improve this project! Your contributions make a difference. 🙌
