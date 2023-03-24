package merge

func MergeStrings(s1, s2 string) string {
   
	
    for i := len(s1) - 1; i >= 0; i-- {
		
        if s1[i:] == s2[:len(s1)-i] {
            return s1[:i] + s2
			
        } 
    }
    return s1 + s2
}

