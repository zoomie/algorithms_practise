function createCharCounter(str) {
    var hashmap = {}
    for (var i=0; i<str.length; i++) {
        char = str.charAt(i)
        if (!(char in hashmap)) {
            hashmap[char] = 0
        }
        hashmap[char]++
    }
    return hashmap
}


function compateStrings(str1, str2) {
    charCounter1 = createCharCounter(str1)
    charCounter2 = createCharCounter(str2)
    for (var key in charCounter1) {
        if (charCounter1[key] > charCounter2[key]) {
            return false
        }
    }
    return true
}

console.log(compateStrings("test", "testtt"))
