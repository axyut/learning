import random
from turtle import Turtle, Screen

is_race_on = False
screen = Screen()
y_positions = [125, 75, 25, -25, -75, -125]
colors = ["red", "yellow", "green", "purple", "orange", "black"]
all_turtles = []

screen.setup(500, 400)   # width, height
user_bet = screen.textinput("Make your Bet!", prompt=f"Which turtle will win the race? Enter a color. \n {colors} ")

for place in range(0, 6):
    tim = Turtle("turtle")
    tim.color(colors[place])
    tim.penup()         # go without making line
    tim.goto(-230, y_positions[place])   # x axis and y axis

    all_turtles.append(tim)

if user_bet:
    is_race_on = True

while is_race_on:

    for turtle in all_turtles:
        rand_distance = random.randint(-2, 20)
        turtle.forward(rand_distance)

        if turtle.xcor() > 210:
            is_race_on = False
            winning_color = turtle.pencolor()

            if winning_color == user_bet:
                turtle.clear()
                turtle.backward((turtle.getscreen().window_width()) - 50)
                message = f"You've WON! The {winning_color} turtle is the winner."
                turtle.write(message, move=False, font=('monaco', 14, 'bold'), align='left')
            else:
                turtle.clear()
                turtle.backward((turtle.getscreen().window_width()) - 50)
                message = f"You've LOST! The {winning_color} turtle won the race."
                turtle.write(message, move=False, font=('monaco', 14, 'bold'), align='left')

screen.exitonclick()
