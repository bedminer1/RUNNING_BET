// CalculateCutoff takes in the scheme and score and calculate the score I would need to win
export function CalculateCutoff(herScore: number, scheme: number[][]): number {
    let res = 0

    // go through all the schemes 
    for (let i = 0; i < scheme.length; i++) {
        const [dist, mult] = scheme[i]
        const prevDist = i > 0 ? scheme[i-1][0] : 0
        const range = dist - prevDist
        // in the case where herScore is less than scheme's total coverage
        if (herScore <= range) {
            res += herScore * mult
            return res
        }
        
        res += range * mult
        herScore -= range
    }

    // if there is leftover just add with 1x multiplier
    res += herScore
    return res
}

// FOR TESTING
interface testCase {
    name: string,
    herScore: number,
    scheme: number[][] // [distance, multiplier]
    expRes: number,
}

const testCases: testCase[] = [
    {
        name: "herScoreLessThanCoverage",
        herScore: 5,
        scheme: [[10,2]],
        expRes: 10,
    },
    {
        name: "herScoreEqualToCoverage",
        herScore: 12,
        scheme: [[3, 2], [8, 1.5], [12, 3]],
        expRes: 25.5, // 3*2 + 5*1.5 + 4*3
    },
    {
        name: "herScoreMoreThanCoverage",
        herScore: 19,
        scheme: [[5,2],[10,1.5]],
        expRes: 26.5,
    },
    {
        name: "noSchemeOnlyBaseMultiplier",
        herScore: 7,
        scheme: [],
        expRes: 7, // No scheme, just 1x multiplier for herScore
    },
    {
        name: "herScoreZero",
        herScore: 0,
        scheme: [[5, 2], [10, 3]],
        expRes: 0, // herScore is 0
    },
    {
        name: "largeScoreMultipleSchemes",
        herScore: 50,
        scheme: [[10, 2], [20, 3], [30, 1.5]],
        expRes: 85, // 10*2 + 10*3 + 10*1.5 + remaining 20*1
    },
]

function testCalculateCutoff(verbose: boolean) {
    let testsPassed = 0
    for (let tc of testCases) {
        let res = CalculateCutoff(tc.herScore, tc.scheme)
        let testPassed = tc.expRes === res

        // if not verbose skip
        if (!verbose) {
            if (testPassed) testsPassed++
            continue
        }

        // if verbose, log it out
        console.log("Test:", tc.name)
        if (tc.expRes === res) {
            console.log("Test Status: Passed")
            continue
        }
        console.log("Test Status: Failed")
        console.log("Expected:", tc.expRes, "Output:", res)
    }

    console.log(testsPassed, "/", testCases.length, "tests passed")
}

testCalculateCutoff(false) // verbose = false