def substringWithConcatationOfAllWords(s, words)
  return [] if words.empty?
  return [] if s.length < words.map(&:length).inject(:+)
  words.sort! { |a, b| a.length <=> b.length }
  words.each_with_object([]) do |word, result|
    if s.include?(word)
      result << word
      s.sub!(word, '')
      substringWithConcatationOfAllWords(s, words)
    end
  end
end

//test
p substringWithConcatationOfAllWords('barfoothefoobarman', ['foo', 'bar'])
p substringWithConcatationOfAllWords('wordgoodgoodgoodbestword', ['word', 'good', 'best', 'word'])
p substringWithConcatationOfAllWords('barfoofoobarthefoobarman', ['bar', 'foo', 'the'])
p substringWithConcatationOfAllWords('wordgoodgoodgoodbestword', ['word', 'good', 'best', 'good'])