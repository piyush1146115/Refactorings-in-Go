# Inline class

## Ref: https://refactoring.guru/inline-class

## Problem 
A class does almost nothing and isn’t responsible for anything,
and no additional responsibilities are planned for it.

## Solution
Move all features from the class to another one.

### Why Refactor
Often this technique is needed after the features of one class are transplanted to other classes, leaving that class with little to do.

### Benefits
Eliminating needless classes frees up operating memory on the computer - and bandwidth in your head.

