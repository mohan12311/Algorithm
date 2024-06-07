func solution(my_string string, overwrite_string string, s int) string {
    my_bytes := []byte(my_string)
    overwrite_bytes := []byte(overwrite_string)
    
    l := len(overwrite_bytes)
    i := 0
    
    for i < l {
        my_bytes[s] = overwrite_bytes[i]
        s++
        i++
    }
    
    return string(my_bytes)
}