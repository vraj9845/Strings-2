// Time Complexity: O(n^2)
// Space Complexity: O(1)
func strStr(haystack string, needle string) int {
    if len(needle)>len(haystack) {
        return -1
    }
    for i:=0;i< len(haystack)-len(needle)+1;i++ {
        if haystack[i]==needle[0]{
            j:=0
            for j<len(needle){
                if haystack[i+j]!=needle[j] || i+j == len(haystack){
                    break;
                }
                j++ 
            }
            if j==len(needle){
                    return i;
                }
            if i+j == len(haystack) {
                return -1;
            }
        }
    }
    return -1;
}
