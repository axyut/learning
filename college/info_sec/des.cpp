#include <iostream>
#include <bitset>

using namespace std;

// Initial permutation table
int IP[] = {58, 50, 42, 34, 26, 18, 10, 2,
             60, 52, 44, 36, 28, 20, 12, 4,
             62, 54, 46, 38, 30, 22, 14, 6,
             64, 56, 48, 40, 32, 24, 16, 8,
             57, 49, 41, 33, 25, 17, 9, 1,
             59, 51, 43, 35, 27, 19, 11, 3,
             61, 53, 45, 37, 29, 21, 13, 5,
             63, 55, 47, 39, 31, 23, 15, 7};

// Final permutation table
int IP_inverse[] = {40, 8, 48, 16, 56, 24, 64, 32,
                    39, 7, 47, 15, 55, 23, 63, 31,
                    38, 6, 46, 14, 54, 22, 62, 30,
                    37, 5, 45, 13, 53, 21, 61, 29,
                    36, 4, 44, 12, 52, 20, 60, 28,
                    35, 3, 43, 11, 51, 19, 59, 27,
                    34, 2, 42, 10, 50, 18, 58, 26,
                    33, 1, 41, 9, 49, 17, 57, 25};

// Expansion permutation table
int E[] = {32, 1, 2, 3, 4, 5,
           4, 5, 6, 7, 8, 9,
           8, 9, 10, 11, 12, 13,
           12, 13, 14, 15, 16, 17,
           16, 17, 18, 19, 20, 21,
           20, 21, 22, 23, 24, 25,
           24, 25, 26, 27, 28, 29,
           28, 29, 30, 31, 32, 1};

// Permutation table
int P[] = {16, 7, 20, 21,
           29, 12, 28, 17,
           1, 15, 23, 26,
           5, 18, 31, 10,
           2, 8, 24, 14,
           32, 27, 3, 9,
           19, 13, 30, 6,
           22, 11, 4, 25};

// S-boxes
int S[8][4][16] = {
    {
        // S1
        {14, 4, 13, 1, 2, 15, 11, 8, 3, 10, 6, 12, 5, 9, 0, 7},
        {0, 15, 7, 4, 14, 2, 13, 1, 10, 6, 12, 11, 9, 5, 3, 8},
        {4, 1, 14, 8, 13, 6, 2, 11, 15, 12, 9, 7, 3, 10, 5, 0},
        {15, 12, 8, 2, 4, 9, 1, 7, 5, 11, 3, 14, 10, 0, 6, 13}
    },
    {
        // S2
        {15, 1, 8, 14, 6, 11, 3, 4, 9, 7, 2, 13, 12, 0, 5, 10},
        {3, 13, 4, 7, 15, 2, 8, 14, 12, 0, 1, 10, 6, 9, 11, 5},
        {0, 14, 7, 11, 10, 4, 13, 1, 5, 8, 12, 6, 9, 3, 2, 15},
        {13, 8, 10, 1, 3, 15, 4, 2, 11, 6, 7, 12, 0, 5, 14, 9}
    },
    {
        // S3
        {10, 0, 9, 14, 6, 3, 15, 5, 1, 13, 12, 7, 11, 4, 2, 8},
        {13, 7, 0, 9, 3, 4, 6, 10, 2, 8, 5, 14, 12, 11, 15, 1},
        {13, 6, 4, 9, 8, 15, 3, 0, 11, 1, 2, 12, 5, 10, 14, 7},
        {1, 10, 13, 0, 6, 9, 8, 7, 4, 15, 14, 3, 11, 5, 2, 12}
    },
    {
        // S4
        {7, 13, 14, 3, 0, 6, 9, 10, 1, 2, 8, 5, 11, 12, 4, 15},
        {13, 8, 11, 5, 6, 15, 0, 3, 4, 7, 2, 12, 1, 10, 14, 9},
        {10, 6, 9, 0, 12, 11, 7, 13, 15, 1, 3, 14, 5, 2, 8, 4},
        {3, 15, 0, 6, 10, 1, 13, 8, 9, 4, 5, 11, 12, 7, 2, 14}
    },
    {
        // S5
        {2, 12, 4, 1, 7, 10, 11, 6, 8, 5, 3, 15, 13, 0, 14, 9},
        {14, 11, 2, 12, 4, 7, 13, 1, 5, 0, 15, 10, 3, 9, 8, 6},
        {4, 2, 1, 11, 10, 13, 7, 8, 15, 9, 12, 5, 6, 3, 0, 14},
        {11, 8, 12, 7, 1, 14, 2, 13, 6, 15, 0, 9, 10, 4, 5, 3}
    },
    {
        // S6
        {12, 1, 10, 15, 9, 2, 6, 8, 0, 13, 3, 4, 14, 7, 5, 11},
        {10, 15, 4, 2, 7, 12, 9, 5, 6, 1, 13, 14, 0, 11, 3, 8},
        {9, 14, 15, 5, 2, 8, 12, 3, 7, 0, 4, 10, 1, 13, 11, 6},
        {4, 3, 2, 12, 9, 5, 15, 10, 11, 14, 1, 7, 6, 0, 8, 13}
    },
    {
        // S7
        {4, 11, 2, 14, 15, 0, 8, 13, 3, 12, 9, 7, 5, 10, 6, 1},
        {13, 0, 11, 7, 4, 9, 1, 10, 14, 3, 5, 12, 2, 15, 8, 6},
        {1, 4, 11, 13, 12, 3, 7, 14, 10, 15, 6, 8, 0, 5, 9, 2},
        {6, 11, 13, 8, 1, 4, 10, 7, 9, 5, 0, 15, 14, 2, 3, 12}
    },
    {
        // S8
        {13, 2, 8, 4, 6, 15, 11, 1, 10, 9, 3, 14, 5, 0, 12, 7},
        {1, 15, 13, 8, 10, 3, 7, 4, 12, 5, 6, 11, 0, 14, 9, 2},
        {7, 11, 4, 1, 9, 12, 14, 2, 0, 6, 10, 13, 15, 3, 5, 8},
        {2, 1, 14, 7, 4, 10, 8, 13, 15, 12, 9, 0, 3, 5, 6, 11}
    }
};

// Permutation function
bitset<32> permutation(const bitset<32>& block, const int* table, int tableSize) {
    bitset<32> result;
    for (int i = 0; i < tableSize; ++i) {
        result[i] = block[table[i] - 1];
    }
    return result;
}

// Function for DES round
bitset<32> round(const bitset<32>& right, const bitset<48>& subKey) {
    // Expansion permutation
    bitset<48> expandedRight = permutation(right, E, 48);

    // XOR with subkey
    expandedRight ^= subKey;

    // S-box substitution
    bitset<32> substituted;
    for (int i = 0; i < 8; ++i) {
        int row = (expandedRight[i * 6] << 1) + expandedRight[i * 6 + 5];
        int col = (expandedRight[i * 6 + 1] << 3) + (expandedRight[i * 6 + 2] << 2)
                  + (expandedRight[i * 6 + 3] << 1) + expandedRight[i * 6 + 4];
        int value = S[i][row][col];
        substituted[i * 4] = (value >> 3) & 1;
        substituted[i * 4 + 1] = (value >> 2) & 1;
        substituted[i * 4 + 2] = (value >> 1) & 1;
        substituted[i * 4 + 3] = value & 1;
    }

    // Permutation
    return permutation(substituted, P, 32);
}

// Generate subkeys from the key
void generateSubKeys(const bitset<64>& key, bitset<48> subKeys[]) {
    // PC-1 permutation
    bitset<56> permutedKey = permutation(key, PC1, 56);

    // Split into left and right halves
    bitset<28> left = permutedKey >> 28;
    bitset<28> right = permutedKey & bitset<56>(string("00000000000000000000000000001111"));

    // Generate 16 subkeys
    for (int i = 0; i < 16; ++i) {
        // Left circular shift
        left = (left << shift[i]) | (left >> (28 - shift[i]));
        right = (right << shift[i]) | (right >> (28 - shift[i]));

        // Combine left and right halves
        bitset<56> combined = (left.to_ullong() << 28) | right.to_ullong();

        // PC-2 permutation
        subKeys[i] = permutation(combined, PC2, 48);
    }
}

// DES encryption
bitset<64> encrypt(const bitset<64>& plainText, const bitset<64>& key) {
    bitset<48> subKeys[16];
    generateSubKeys(key, subKeys);

    // Initial permutation
    bitset<64> permutedText = permutation(plainText, IP, 64);

    // Split into left and right halves
    bitset<32> left = permutedText >> 32;
    bitset<32> right = permutedText & bitset<64>(string("0000000000000000000000000000000011111111111111111111111111111111"));

    // 16 rounds
    for (int i = 0; i < 16; ++i) {
        bitset<32> temp = right;
        right = left ^ round(right, subKeys[i]);
        left = temp;
    }

    // Combine left and right halves
    bitset<64> result = (right.to_ullong() << 32) | left.to_ullong();

    // Final permutation
    return permutation(result, IP_inverse, 64);
}

int main() {
    // Example plaintext and key
    bitset<64> plainText("0000000100100011010001010110011110001001101010111100110111101111");
    bitset<64> key("0001001100110100010101110111100110011011101111001101111111110001");

    // Encrypt
    bitset<64> cipherText = encrypt(plainText, key);

    // Display ciphertext
    cout << "Ciphertext: " << hex << cipherText.to_ullong() << endl;

    return 0;
}
