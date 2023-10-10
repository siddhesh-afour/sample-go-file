# Import the necessary modules
import math
import random

# Define a class
class Person:
  def __init__(self, name, age):
    self.name = name
    self.age = age

  def greet(self):
    print("Hello, my name is {} and I am {} years old.".format(self.name, self.age))

# Define a function
def calculate_area(radius):
  return math.pi * radius * radius

# Define a main function
def main():
  # Create a new instance of the Person class
  person = Person("John Doe", 30)

  # Greet the person
  person.greet()

  # Calculate the area of a circle
  area = calculate_area(5)

  # Print the area
  print("The area of a circle with a radius of 5 units is {}".format(area))

# If this file is being run as the main program, execute the main() function
if __name__ == "__main__":
  main()
  
