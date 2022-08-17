# Extract class

## Ref: https://refactoring.guru/extract-class

## Problem:
When one class does the work of two, awkwardness results.

## Solution:
Instead, create a new class and place the fields and methods responsible for the relevant functionality in it.

## Why Refactor
Classes always start out clear and easy to understand. They do their jon and mind their own business as it were, without butting into the work of other classes. But as the program expands, a method is added and then a field...and eventually some classes are performing more responsibilities than ever envisioned. 
