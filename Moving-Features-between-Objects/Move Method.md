# Move Method

## Ref: https://refactoring.guru/move-method

## Problem
A method is used more in another class than in its own class.

## Solution
Create a new method in the class that uses the method the most, then move code from the old method to there.
Turn the code of the original method into a reference to the new method in the other class or else remove it entirely.

## Why Refactor
- You want to move a method to a class that contains most of the data used by the method. This makes classes more internally coherent.
- You want to move a method in order to reduce or eliminate the dependency of the class calling the method on the class in which it’s located. This can be useful if the calling class is already dependent on the class to which you’re planning to move the method. This reduces dependency between classes.  

