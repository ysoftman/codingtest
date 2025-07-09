// https://www.geeksforgeeks.org/gcd-two-array-numbers/
/*
배열내의 모든 값들의 gcd(최대공약수) 구하기
*/
#include <algorithm>
#include <iostream>
using namespace std;

int gcd(int a, int b) {
    if (b == 0) {
        return a;
    }
    return gcd(b, a % b);
}

int gcd_array(int *arr, int len) {
    int max_gcd = 1;
    for (int i = 0; i < len; ++i) {
        cout << arr[i] << " ";
    }
    cout << endl;
    max_gcd = arr[0];
    for (int i = 1; i < len; ++i) {
        max_gcd = gcd(max_gcd, arr[i]);
    }
    return max_gcd;
}

int main() {
    int arr[] = {88, 80, 14, 64, 32};
    int len   = 5;
    cout << gcd_array(arr, len) << endl;
}
