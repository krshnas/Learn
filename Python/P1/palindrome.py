# function which return reverse of a string

def is_palindrome(s):
    return s == s[::-1]


# Driver code
s = "malayalam"
ans = is_palindrome(s)

if ans:
    print("Yes")
else:
    print("No")