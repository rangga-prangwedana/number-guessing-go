# 🎮 Number Guessing Game (Go CLI)

A fun, CLI-based number guessing game written in **Go**.  
The computer picks a random number between **1 and 100**, and your task is to guess it within a limited number of chances depending on the difficulty you choose.

## 🚀 Features
- Welcome screen with rules.
- Random number generation between **1–100**.
- Difficulty levels:
  - **Easy** → 10 chances  
  - **Medium** → 5 chances  
  - **Hard** → 3 chances  
- Feedback on each guess (`greater` or `less`).
- Congratulatory message with number of attempts on success.
- Game ends when:
  - You guess correctly ✅  
  - You run out of chances ❌  

## 📖 Sample Gameplay

```plaintext
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)
Enter your choice: 2
Great! You have selected the Medium difficulty level.
Let's start the game!
Enter your guess: 50
Incorrect! The number is less than 50.
Enter your guess: 25
Incorrect! The number is greater than 25.
Enter your guess: 35
Incorrect! The number is less than 35.
Enter your guess: 30
🎉 Congratulations! You guessed the correct number in 4 attempts.
```

## 🛠️ Installation & Usage

1. **Clone the repository**

   ```bash
   git clone https://github.com/your-username/number-guessing-game-go.git
   cd number-guessing-game-go
   ```

2. **Run the game**

   ```bash
   go run main.go
   ```

3. **Build an executable (optional)**

   ```bash
   go build -o guessing-game
   ./guessing-game
   ```

---



## 🧑‍💻 Tech Stack

* **Language**: Go (Golang)
* **Type**: CLI application

---

## 📜 License

This project is open-source and available under the [MIT License](LICENSE).
