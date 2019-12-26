// ysoftman
// matrix 회전하기

/*
https://www.geeksforgeeks.org/inplace-rotate-square-matrix-by-90-degrees/

Given an square matrix, turn it by 90 degrees in anti-clockwise direction without using any extra space.

Input
 1  2  3
 4  5  6
 7  8  9

Output:
 3  6  9 
 2  5  8 
 1  4  7 

Input:
 1  2  3  4 
 5  6  7  8 
 9 10 11 12 
13 14 15 16 

Output:
 4  8 12 16 
 3  7 11 15 
 2  6 10 14 
 1  5  9 13
*/

#include <iostream>

using namespace std;

void printMatrix(int matrix[], int m, int n)
{
    for (int i = 0; i < m; i++)
    {
        for (int j = 0; j < n; j++)
        {
            printf("%2d ", matrix[m * i + j]);
        }
        cout << endl;
    }

    cout << endl;
}

// 반시계 방향으로 90도 회전
int *rotateMatrix(int matrix[], int m, int n)
{
    int *newMatrix = new int[3 * 3];

    int a = m - 1, b = 0;
    for (int i = 0; i < m; i++)
    {
        for (int j = 0; j < n; j++)
        {
            a = m - 1 - j;
            b = i;
            newMatrix[m * a + b] = matrix[m * i + j];
        }
    }
    return newMatrix;
}

int main()
{
    int m1[][3] = {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}};
    printMatrix(&m1[0][0], 3, 3);

    int *nM1 = rotateMatrix(&m1[0][0], 3, 3);
    printMatrix(nM1, 3, 3);

    int m2[][4] = {{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}};
    printMatrix(&m2[0][0], 4, 4);

    int *nM2 = rotateMatrix(&m2[0][0], 4, 4);
    printMatrix(nM2, 4, 4);
    return 0;
}
