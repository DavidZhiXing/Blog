fn subStringWithConcatationOfAllWords(s: String, words: Vec<String>) -> Vec<i32> {
    let mut res = vec![];
    let mut i = 0;
    let mut j = 0;
    let mut k = 0;
    let mut word_len = words.len();
    let mut word_len_sum = 0;
    for word in words {
        word_len_sum += word.len();
    }
    while i + word_len_sum <= s.len() {
        let mut word_index = 0;
        let mut word_len = words[word_index].len();
        while j < word_len {
            if s[i + j] != words[word_index].as_bytes()[j] {
                break;
            }
            j += 1;
        }
        if j == word_len {
            word_index += 1;
            if word_index == word_len {
                res.push(i as i32);
                word_index = 0;
                word_len = words[word_index].len();
                i += word_len; 
                j = 0;
                k += 1;
                if k == word_len_sum {
                    break;
                }
            }
        } else {
            i += 1;
            j = 0;
            k = 0;
        }   
    }
    res
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_subStringWithConcatationOfAllWords() {
        assert_eq!(subStringWithConcatationOfAllWords("barfoothefoobarman".to_string(), vec!["foo".to_string(), "bar".to_string()]), vec![0, 9]);
        assert_eq!(subStringWithConcatationOfAllWords("wordgoodgoodgoodbestword".to_string(), vec!["word".to_string(), "good".to_string(), "best".to_string(), "word".to_string()]), vec![]);
        assert_eq!(subStringWithConcatationOfAllWords("barfoofoobarthefoobarman".to_string(), vec!["bar".to_string(), "foo".to_string(), "bar".to_string()]), vec![6, 9, 12]);
        assert_eq!(subStringWithConcatationOfAllWords("wordgoodgoodgoodbestword".to_string(), vec!["word".to_string(), "good".to_string(), "best".to_string(), "good".to_string()]), vec![8]);
        assert_eq!(subStringWithConcatationOfAllWords("barfoofoobarthefoobarman".to_string(), vec!["bar".to_string(), "foo".to_string(), "the".to_string()]), vec![6, 9, 12]);
        assert_eq!(subStringWithConcatationOfAllWords("barfoofoobarthefoobarman".to_string(), vec!["bar".to_string(), "foo".to_string(), "bar".to_string(), "the".to_string()]), vec![6, 9, 12]);
        assert_eq!(subStringWithConcatationOfAllWords("barfoofoobarthefoobarman".to_string(), vec!["bar".to_string(), "foo".to_string(), "bar".to_string(), "the".to_string(), "foo".to_string()]), vec![6, 9, 12]);
        assert_eq!(subStringWithConcatationOfAllWords("barfoofoobarthefoobarman".to_string(), vec!["bar".to_string(), "foo".to_string(), "bar".to_string(), "the".to_string(), "foo".to_string(), "bar".to_string()]), vec![6, 9, 12]);
    }
}