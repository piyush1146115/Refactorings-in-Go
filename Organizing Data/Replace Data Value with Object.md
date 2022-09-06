# Replace Data Value with Object

# Ref: https://refactoring.guru/replace-data-value-with-object

## Problem
A class (or group of classes) contains a data field. The field has its own behavior and associated data.

## Solution
Create a new class, place the old field and its behavior in the class, and store the object of the class in the original class.

## Why Refactor

With replacement of a data value with an object, we have a primitive field (number, string, etc.) that’s no longer so simple due to growth of the program and now has associated data and behaviors. On the one hand, there’s nothing scary about these fields in and of themselves. However, this fields-and-behaviors family can be present in several classes simultaneously, creating duplicate code. Therefore, for all this we create a new class and move both the field and the related data and behaviors to it.

