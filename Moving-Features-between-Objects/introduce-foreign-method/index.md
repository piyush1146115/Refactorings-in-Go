# Introduce Foreign Method

## Ref: https://refactoring.guru/introduce-foreign-method

## Problem

A utility class doesn’t contain the method that you need and you can’t add the method to the class.

## Solution

Add the method to a client class and pass an object of the utility class to it as an argument.

## Why Refactor

You have code that uses the data and methods of a certain class. You realize that the code will look and work much better inside a new method in the class. But you can’t add the method to the class because, for example, the class is located in a third-party library.

This refactoring has a big payoff when the code that you want to move to the method is repeated several times in different places in your program.

## Benefits

Removes code duplication. If your code is repeated in several places, you can replace these code fragments with a method call. This is better than duplication even considering that the foreign method is located in a suboptimal place.
