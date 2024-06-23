
import random

stages = [''' 
  +---+|
  |   ||
  O   ||
 /|\  ||
 / \  ||
      ||
      ||
     /||\ 
============||
||||||||||||||
''', ''' 
   +---+|
  |   ||
  O   ||
 /|\  ||
 /    ||
      ||
      ||
     /||\ 
============||
||||||||||||||
''',''' 
  +---+|
  |   ||
  O   ||
 /|\  ||
      ||
      ||
      ||
     /||\ 
============||
||||||||||||||
''',''' 
  +---+|
  |   ||
  O   ||
 /|   ||
      ||
      ||
      ||
     /||\ 
============||
||||||||||||||
''',''' 
  +---+|
  |   ||
  O   ||
  |   ||
      ||
      ||
      ||
     /||\ 
============||
||||||||||||||
''',''' 
  +---+|
  |   ||
  O   ||
      ||
      ||
      ||
      ||
     /||\ 
============||
||||||||||||||
''']

print('''           Welcome to Hangman!!!
 _ 
| |
| |__   __ _ _ __   __ _ _ __ ___   __ _ _ __
| '_ \ / _' | '_ \ / _' | '_ ' '_\ / _' | '_ \ 
| | | | (_| | | | | (_| | | | | | | (_| | | | |
|_| |_|\__,_|_| |_|\__, |_| |_| |_|\__,_|_| |_|~~
                    __/ |
                   |___/' 

RULES: 1. You have 5 lives to guess the word before the man's body is completed and hanged.

1. To play with computer, type ==> '1'
2. To play with your friend, type ==> '2' ''')
play_with = input("==> ")

word_list = ["computers", "animals", "penguins", "rockets",'single', 'hippopotamus']

if int(play_with)==2:
    chosen_word = input("Type a secret word: ")
    print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n START")
elif int(play_with)==1:
    chosen_word = random.choice(word_list)
else:
    print("Please pick a valid number!")
    quit()

word_length = len(chosen_word)
display = []

for times in range(word_length):
    display += "_"

endOfGame= False
lives = 5

while not endOfGame:
    print(stages[lives])
    guess = input("Guess a letter: ").lower()
   

    for position in range(word_length):
        letter = chosen_word[position]
        if letter == guess:
            display[position] = letter
            print(f'''
            
            
Ahhhhyesss!!  ________ {letter} guessed.________''')
    
    if guess not in chosen_word:
        lives -= 1
        print(f'''
        
        
Oohuhh. limbs added.    ______Try again!______''')
        if lives == 0:
            endOfGame = True
            print("-------------You couldn't save him.------------------------------------------------------------------------------")
    
    print(f'''                                            
Remaining life = {lives}''')
    
    if "_" not in display:
        endOfGame= True
        print("**************** YOU SAVED THE MAN!!!!!!!***************************************************************************")
    
    
    print(f''' ==> {''.join(display)}.  It is a {word_length} letter word.''')