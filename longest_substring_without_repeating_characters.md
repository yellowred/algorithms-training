# 3. Longest Substring Without Repeating Characters
Given a string s, find the length of the longest substring without repeating characters.

Example 1:
```
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
```

## Solution

```python
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        chars = [0] * 128
        
        left = right = 0
        
        res = 0
        
        while right < len(s):
            chars[ord(s[right])] += 1
            
            while chars[ord(s[right])] > 1:
                chars[ord(s[left])] -= 1
                left += 1
                
            res = max(res, right - left + 1)
            right += 1
        
        return res
```
