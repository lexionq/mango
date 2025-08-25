package password

import (
	"fmt"
	"regexp"
	"strings"
)

type PasswordStrength struct {
	Score       int
	Level       string
	Issues      []string
	Suggestions []string
}

func CheckPasswordStrength(password string) PasswordStrength {
	var strength PasswordStrength
	score := 0
	var issues []string
	var suggestions []string

	// Length check
	length := len(password)
	if length >= 12 {
		score += 25
	} else if length >= 8 {
		score += 15
	} else if length >= 6 {
		score += 5
		issues = append(issues, "Password is too short")
		suggestions = append(suggestions, "Use at least 8 characters (12+ recommended)")
	} else {
		issues = append(issues, "Password is very short")
		suggestions = append(suggestions, "Use at least 8 characters (12+ recommended)")
	}

	// Lowercase letters
	if matched, _ := regexp.MatchString(`[a-z]`, password); matched {
		score += 5
	} else {
		issues = append(issues, "No lowercase letters")
		suggestions = append(suggestions, "Add lowercase letters (a-z)")
	}

	// Uppercase letters
	if matched, _ := regexp.MatchString(`[A-Z]`, password); matched {
		score += 5
	} else {
		issues = append(issues, "No uppercase letters")
		suggestions = append(suggestions, "Add uppercase letters (A-Z)")
	}

	// Numbers
	if matched, _ := regexp.MatchString(`[0-9]`, password); matched {
		score += 5
	} else {
		issues = append(issues, "No numbers")
		suggestions = append(suggestions, "Add numbers (0-9)")
	}

	// Special characters
	if matched, _ := regexp.MatchString(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`, password); matched {
		score += 10
	} else {
		issues = append(issues, "No special characters")
		suggestions = append(suggestions, "Add special characters (!@#$%^&*)")
	}

	// Variety bonus
	uniqueChars := make(map[rune]bool)
	for _, char := range password {
		uniqueChars[char] = true
	}

	if len(uniqueChars) >= length/2 {
		score += 10
	}

	// Common patterns penalty
	commonPatterns := []string{
		"123", "abc", "qwe", "asd", "zxc", "password", "admin", "user",
		"000", "111", "222", "333", "444", "555", "666", "777", "888", "999",
	}

	lowerPassword := strings.ToLower(password)
	for _, pattern := range commonPatterns {
		if strings.Contains(lowerPassword, pattern) {
			score -= 10
			issues = append(issues, fmt.Sprintf("Contains common pattern: %s", pattern))
			suggestions = append(suggestions, "Avoid common patterns and sequences")
			break
		}
	}

	// Repeated characters penalty
	for i := 0; i < len(password)-2; i++ {
		if password[i] == password[i+1] && password[i+1] == password[i+2] {
			score -= 15
			issues = append(issues, "Contains repeated characters")
			suggestions = append(suggestions, "Avoid repeating the same character")
			break
		}
	}

	// Determine strength level
	if score >= 70 {
		strength.Level = "Very Strong 🟢"
	} else if score >= 50 {
		strength.Level = "Strong 🟡"
	} else if score >= 30 {
		strength.Level = "Moderate 🟠"
	} else if score >= 15 {
		strength.Level = "Weak 🔴"
	} else {
		strength.Level = "Very Weak ⚠️"
	}

	strength.Score = score
	strength.Issues = issues
	strength.Suggestions = suggestions

	return strength
}

func DisplayPasswordStrength(password string) {
	strength := CheckPasswordStrength(password)

	fmt.Printf("\n🔒 Password Strength Analysis:\n")
	fmt.Printf("Score: %d/100\n", strength.Score)
	fmt.Printf("Level: %s\n", strength.Level)

	if len(strength.Issues) > 0 {
		fmt.Println("\n⚠️  Issues found:")
		for _, issue := range strength.Issues {
			fmt.Printf("  • %s\n", issue)
		}
	}

	if len(strength.Suggestions) > 0 {
		fmt.Println("\n💡 Suggestions:")
		for _, suggestion := range strength.Suggestions {
			fmt.Printf("  • %s\n", suggestion)
		}
	}

	if strength.Score >= 70 {
		fmt.Println("\n✅ Great! Your password is strong.")
	} else if strength.Score >= 50 {
		fmt.Println("\n👍 Good password, but could be stronger.")
	} else {
		fmt.Println("\n⚠️  Consider improving your password strength.")
	}
}
