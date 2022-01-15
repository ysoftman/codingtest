/*
https://www.educative.io/m/reverse-words-in-a-sentence

Reverse Words in a Sentence
Problem Statement
Reverse the order of words in a given sentence (an array of characters). Take the “Hello World” string for example:

*/

#include <iostream>
#include <vector>
using std::cout;
using std::endl;
using std::string;
using std::vector;
void reverse_words(char *sentence)
{
    //TODO: Write - Your - Code
    vector<string> vec;
    string temp;
    char *head = sentence;
    while (*sentence != '\0')
    {
        if (*sentence == ' ')
        {
            vec.push_back(temp);
            temp = "";
        }
        else
        {
            temp += *sentence;
        }

        if (*(sentence + 1) == '\0')
        {
            vec.push_back(temp);
        }
        sentence++;
    }

    string result;
    // cout << "for debugging" << endl;
    vector<string>::reverse_iterator rever_iter;
    for (rever_iter = vec.rbegin(); rever_iter != vec.rend(); rever_iter++)
    {
        // cout << *rever_iter << endl;
        result += *rever_iter + " ";
    }
    result[result.length() - 1] = '\0';

    strcpy(head, result.c_str());
}

int main()
{
    string str = "Hello World!";
    char *a = const_cast<char *>(str.c_str());
    cout << a << endl;
    reverse_words(a);
    cout << a << endl;
}