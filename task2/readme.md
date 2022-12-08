## Task2 - Anagrams
### Programmers Notes
First of all it was very fun to solve this challange

The comparison takes a lot of time. O(n*n+1/2)
So i found it correct to use the map approach.

The tricky part was finding a unique key for words and it should be same for anagrams.
In my first try I solved this by converting strings to rune array and sorting by letters but sorting also takes time.
So I solved this by creating a slice that counts all letters for words.

Ex:
aabc
```text
 a b c d e f g h i j k l m n o p q r s t u v w x y z
[2,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
```

we can convert words to numbers. So main problem is finding equal key for same sets of numbers. I also thought I could solve it with the area of the polygons, but I couldn't find a constant function for the areas of the polygons. Maybe it can be solved more easily with another mathematical approach.