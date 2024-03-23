// @ Memory Limit(MB): 512, Execution Timeout(ms): 2000
// N persons from different countries have gathered at a party. You are the organizer of the party. You
// have distributed name tags to the participants with only their initials on them. Though the purpose of
// the party is interaction, you have found that participants are not actively taking part in conversation.
// You have the idea to pair people up who have opposite (reversed) sets of initials.
// For instance， a participant with the initials Ai., Bi can be paired with a participant with the initials Bi,
// Ai. One participant cannot be paired with multiple other participants аt the same time﹒
// Please calculate the maximum possible number of pairs.
// As the participants are from different countries, there are some names that cannot be expressed using
// the Roman alphabet. Thus， the initials of participant ﹔ are expressed as a pair of integers (Ai, Bi).
// The participant will then 66 paired with someone that has the initials ( Bi，Ai﹚， where (1 < i < N).
// Implementation details
// CLI
// Please implement a CLI application that takes the input value as an argument and outputs the result to
// standard output.
// For details, 566 the “CLI application template“ section а{ the bottom of this page.

// Input rule
// The program is executed as follows:
// ./myApp ﹤ input.in
// The standard input format is as follows:
// N
// A1 B1
// Ay By
// The conditions are as follows:
// 1 <=N < 10^5, integer
// 1 <= Ai,Bi <10^9，(1 < i < N), integer
// Output rule
// Output the maximum possible number of pairs.
// MO examples
// Example 1
// Standard input ( g9_samplel.in ﹚
// Copy
// 3
// 1 2
// 2 3
// 2 1
// Standard output
// 1
// Participant 1 and participant 3 can be paired.
// Example 2
// Standard input ( ee_sample2.in ﹚
// 3
// 1 2
// 2 1
// 2 1
// Standard output
// 1
// Participant 1 can be paired with either participant 2 or participant 3. As participant 1 cannot be paired with both participant 2 and participant 3 at the same time, the maximum number of
// pairs at one time is 1.
// Note that the expected O 15 defined in test/basic_testcases.json.
// Please refer 10 this е for implementation.

using System;
using System.IO;
using System.Collections.Generic;

public class MainApp
{

    static public void Main(string[] args)
    {
        // このコードは標準入力と標準出力を用いたサンプルコードです。
        // このコードは好きなように編集・削除してもらって構いません。
        // ---
        // This is a sample code to use stdin and stdout.
        // Edit and remove this code as you like.

        var lines = GetStdin();
        var n = int.Parse(lines[0]);

        var dict = new Dictionary<Tuple<long, long>, int>();
        for (var i = 1; i <= n; i++)
        {
            var line = lines[i];
            var tokens = line.Split(' ');
            var a = long.Parse(tokens[0]);
            var b = long.Parse(tokens[1]);
            var key = new Tuple<long, long>(a, b);
            if (dict.ContainsKey(key))
            {
                dict[key] += 1;
            }
            else
            {
                dict[key] = 1;
            }
        }
        var max = 0;
        foreach (var pair in dict)
        {
            var key = pair.Key;
            var value = pair.Value;
            var reversedKey = new Tuple<long, long>(key.Item2, key.Item1);
            if (dict.ContainsKey(reversedKey))
            {
                var reversedValue = dict[reversedKey];
                var sum = value + reversedValue;
                if (sum > max)
                {
                    max = sum;
                }
            }
        }
        Console.WriteLine(max);
    }


    static private string[] GetStdin()
    {
        var list = new List<string>();
        string line;
        while ((line = Console.ReadLine()) != null)
        {
            list.Add(line);
        }
        return list.ToArray();
    }
}