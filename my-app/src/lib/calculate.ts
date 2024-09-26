// CalculateCutoff takes in the scheme and score and calculate the score I would need to win
function CalculateCutoff(herScore: number, scheme: number[][]): number {
    let res = 0

    // go through all the schemes 
    for (let i = 0; i < scheme.length; i++) {
        let s = scheme[i]
        // in the case where herScore runs is less than scheme afford
        if (herScore <= s[0]) {
            res += herScore * s[1]
            herScore = 0
            break
        }
        if (i === 0) {
            res += s[0] * s[1]
            herScore -= s[0]
        } else {
            res += (s[0] - scheme[i-1][0]) * s[1]
            herScore -= (s[0] - scheme[i-1][0])
        }
    }

    // if there is leftover just add
    res += herScore
    return res
}

// FOR TESTING
interface testCase {
    herScore: number,
    scheme: number[][] // [distance, multiplier]
    expRes: number,
}

const testCases: testCase[] = [
    {
        herScore: 2,
        scheme: [[1,2]],
        expRes: 3,
    },
    {
        herScore: 2,
        scheme: [[1,2], [2,3]],
        expRes: 5,
    },
    {
        herScore: 19,
        scheme: [[5,2],[10,1.5]],
        expRes: 26.5,
    }
]

function testCalculateCutoff() {
    for (let tc of testCases) {
        let res = CalculateCutoff(tc.herScore, tc.scheme)
        console.log("output:", res, "expected:", tc.expRes)
    }
}

testCalculateCutoff()