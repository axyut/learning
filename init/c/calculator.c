#include <stdio.h>
#include <math.h>
#include <conio.h>
int main()
{
	printf("\n\t\t************Welcome to My MINOR project 1************");
	printf("\n");
	printf("\t\t\tThis is my calculator application");	
	printf("\n");
	printf("\t\t\tproject is done by SWATIMONALISA PATRA");
	printf("\n");
	printf("\t\t\tfrom oretes academy");
	printf("\n");
	int i=1;
	char operation;
	while(i)
	{
		options();
		printf("\t\t\t");
		operation = getch();
		
		switch(operation)
		{
			case '+': addition();
				break;
				
			case '-': subtraction();
				break;
			
			case '*':multiplication();
			    break;
			    
			case '/':division();
			    break;
			
			case 'r':
			case 'R':modulus();
			    break;	    
			    
			case '!':factorial();
			    break;
			
			case 's':
			case 'S':squareRoot();
			    break;   
			
			case '^': power();
				break;
			
			case 'q':
			case 'Q': exit(0);
				break;
				
			case 'D':
			case 'd':
				break;
				
			default: printf("\n\t\tYou entered unavailable option");
		}
	}
	
}

void options()
{
	printf("\n");
	printf("\t\t\tEnter your choice of operation");
	printf("\n");
	printf("\t\t\tEnter + for Addition");
	printf("\n");
	printf("\t\t\tEnter - for Subtraction");
	printf("\n");
	printf("\t\t\tEnter * for Multiplication");
	printf("\n");
	printf("\t\t\tEnter / for Division");
	printf("\n");
	printf("\t\t\tEnter r for Modulus");
	printf("\n");
	printf("\t\t\tEnter ^ for Power");
	printf("\n");
	printf("\t\t\tEnter ! for Factorial");
	printf("\n");
	printf("\t\t\tEnter s for Squareroot");
	printf("\n");
	printf("\t\t\tEnter q to quit");
	printf("\n");
	printf("\t\t\tEnter d to display the menu again");
	printf("\n");
}

void addition()
{
	int a, b, sum;
	printf("\n\t\tEnter two numbers for addition: ");
	scanf("%d %d", &a, &b);
	sum = a+b;
	printf("\n\t\tSum of %d and %d is %d\n", a, b, sum);
}

void subtraction()
{
	int a, b, diff;
	printf("\n\t\tEnter two numbers for subtraction: ");
	scanf("%d %d", &a, &b);
	diff = a-b;
	printf("\n\t\tSubtracting %d from %d, we get %d\n", a, b, diff);
}
void multiplication()
{
	int a, b, mul;
	printf("\n\t\tEnter two numbers for multiplication: ");
	scanf("%d %d", &a, &b);
	mul = a*b;
	printf("\n\t\tmultiplying %d with %d, we get %d\n", a, b, mul);
}
void division()
{
	int a, b, div;
	printf("\n\t\tEnter two numbers for division: ");
	scanf("%d %d", &a, &b);
	div = a/b;
	printf("\n\t\tdivision %d with %d, we get %d\n", a, b, div);
}
void modulus()
{
	double a, b, mod_res;
	printf("\n\t\tEnter two numbers for modulus: ");
	scanf("%lf %lf", &a, &b);
	mod_res= fmod(a,b);
	printf("\n\t\tdivision %lf from %lf, we get %lf\n", a, b, mod_res);
}
void power()
{
	double a, b, power;
	printf("\n\t\tEnter number to find power: ");
	scanf("%lf", &a);
	printf("\n\t\tEnter power: ");
	scanf("%lf", &b);
	power = pow(a, b);
	printf("\n\t\t%.2lf to the power %.2lf is %.2lf\n", a, b, power);
}
void factorial()
{
	int i,fact=1,num;
	printf("\n\t\tEnter a number to find factorial: ");
	scanf("%d", &num);
	for(i=1;i<=num;i++)
	{
		fact=fact*i;
	}
	printf("\n\t\tfactorial of %d is %d\n",num, fact);
}

void squareRoot()
{
	double num,squareRoot;
	printf("\n\t\tEnter a number to find squareRoot: ");
	scanf("%lf", &num);
	squareRoot=sqrt(num);
	printf("\n\t\tsquareRoot of %.2lf is %.2lf\n",num, squareRoot);
}



