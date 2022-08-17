# Move Field

## Ref: https://refactoring.guru/move-field

## Problem:
A field is used more in another class than in its own class.

## Solution:
Create a field in a new class and redirect all users of the old field to it.

## Why Refactor:
Here is our rule of thumb: put a field in the same place as the methods that use it (or else where most of these methods are)