# https://leetcode.com/problems/valid-phone-numbers/
# 193. Valid Phone Numbers
# Easy
# Given a text file file.txt that contains a list of phone numbers (one per line), write a one-liner bash script to print all valid phone numbers.

# You may assume that a valid phone number must appear in one of the following two formats: (xxx) xxx-xxxx or xxx-xxx-xxxx. (x means a digit)

# You may also assume each line in the text file must not contain leading or trailing white spaces.

# Example:

# Assume that file.txt has the following content:

# 987-123-4567
# 123 456 7890
# (123) 456-7890

# Your script should output the following valid phone numbers:

# 987-123-4567
# (123) 456-7890

# Read from the file file.txt and output all valid phone numbers to stdout.

cat <<zzz >file.txt
987-123-4567
123 456 7890
(123) 456-7890
zzz

regx1="^[0-9]{3}-[0-9]{3}-[0-9]{4}$"
regx2="^\([0-9]{3}\) [0-9]{3}-[0-9]{4}$"
while read line; do
    #echo -n "${line} ==> "
    # =~ 정규식 매칭
    if [[ ${line} =~ $regx1 ]] || [[ ${line} =~ $regx2 ]]; then
        echo "${line}"
    fi
done <file.txt

rm -rf file.txt
