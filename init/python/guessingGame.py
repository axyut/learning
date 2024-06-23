import random
print("********** Welcome to Number Guessing Game! *********")
print("\nWait, I'm thinking number between 1 and 100")
answer = random.randint(1,100)
#print(f"testing: {answer}")

easy_level = 10
hard_level = 5

def set_difficulty():
    level = input("Choose a difficulty . Type 'easy' for 10 lives or 'hard' for 5 lives.\n ==> ")
    if level == "easy":
        return easy_level
    else:
        return hard_level

def checker(guess, answer, turns):
    if guess > answer:
        print("Too High!")
        return turns - 1
    elif guess < answer:
        print('Too Low!')
        return turns - 1
    else:
        input(f" You got it. The answer is {answer}")
        return turns - 11    # now turns is always negative(less than zero)

def game():
    turns = set_difficulty()
    guess =0

    while guess != answer:

        guess = int(input("\nMake a guess: "))
        turns =checker(guess, answer, turns)

        if turns ==0:
            input("You've run out of guesses. YOU LOSE!!!!")
            return
        elif turns <0 :
            return
        else:
            print(f"You have {turns} remaining attempts to guess the number.")
game()